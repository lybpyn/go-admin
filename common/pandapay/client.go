package pandapay

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"go-admin/config"
)

// Client PandaPay客户端
type Client struct {
	ApiUrl     string
	MerchantId int
	AppSecret  string
	NotifyUrl  string
	Currency   string
}

// NewClient 创建PandaPay客户端
func NewClient() *Client {
	cfg := config.ExtConfig.PandaPay
	return &Client{
		ApiUrl:     cfg.ApiUrl,
		MerchantId: cfg.MerchantId,
		AppSecret:  cfg.AppSecret,
		NotifyUrl:  cfg.NotifyUrl,
		Currency:   cfg.Currency,
	}
}

// PayoutRequest 代付请求参数
type PayoutRequest struct {
	MerchantId      int         `json:"merchantId"`
	MerchantOrderId string      `json:"merchantOrderId"`
	Amount          float64     `json:"amount"`
	Currency        string      `json:"currency,omitempty"`
	Timestamp       int64       `json:"timestamp"`
	Sign            string      `json:"sign"`
	NotifyUrl       string      `json:"notifyUrl,omitempty"`
	FundAccount     FundAccount `json:"fundAccount"`
}

// FundAccount 收款账户信息
type FundAccount struct {
	AccountType string       `json:"accountType"` // bank_account 或 crypto
	Contact     *Contact     `json:"contact"`
	BankAccount *BankAccount `json:"bankAccount,omitempty"`
}

// Contact 联系人信息
type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Address string `json:"address,omitempty"`
}

// BankAccount 银行账户信息
type BankAccount struct {
	Name          string `json:"name"`
	BankCode      string `json:"bankCode"`
	AccountNumber string `json:"accountNumber"`
}

// PayoutResponse 代付响应
type PayoutResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderId         string `json:"orderId"`
		MerchantOrderId string `json:"merchantOrderId"`
		Status          int    `json:"status"` // 0=等待, 1=成功, 2=失败等
		ChannelTxnId    string `json:"channelTxnId,omitempty"`
	} `json:"data"`
}

// generateSign 生成签名
// 签名格式: merchantId=999&merchantOrderId=1000&amount=100.00&appSecret=商户appSecret
// 进行MD5加密并转小写
func (c *Client) generateSign(merchantOrderId string, amount float64) string {
	// 格式化金额为两位小数字符串（签名用）
	amountStr := fmt.Sprintf("%.2f", amount)

	// 拼接签名字符串
	signStr := fmt.Sprintf("merchantId=%d&merchantOrderId=%s&amount=%s&appSecret=%s",
		c.MerchantId, merchantOrderId, amountStr, c.AppSecret)

	// MD5加密并转小写
	hash := md5.Sum([]byte(signStr))
	return fmt.Sprintf("%x", hash)
}

// SubmitPayout 提交代付请求
func (c *Client) SubmitPayout(req *PayoutRequest) (*PayoutResponse, error) {
	// 设置默认值
	if req.MerchantId == 0 {
		req.MerchantId = c.MerchantId
	}
	if req.Currency == "" {
		req.Currency = c.Currency
	}
	if req.NotifyUrl == "" {
		req.NotifyUrl = c.NotifyUrl
	}
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().UnixMilli()
	}

	// 格式化金额为两位小数（确保签名和JSON中格式一致）
	amountStr := fmt.Sprintf("%.2f", req.Amount)

	// 生成签名
	req.Sign = c.generateSign(req.MerchantOrderId, req.Amount)

	// 日志：打印签名信息（用于调试）
	signStr := fmt.Sprintf("merchantId=%d&merchantOrderId=%s&amount=%s&appSecret=%s",
		req.MerchantId, req.MerchantOrderId, amountStr, c.AppSecret)
	fmt.Printf("[PandaPay] 签名字符串: %s\n", signStr)
	fmt.Printf("[PandaPay] 生成签名: %s\n", req.Sign)

	// 手动构建JSON，确保amount字段保持两位小数格式
	reqBodyMap := map[string]interface{}{
		"merchantId":      req.MerchantId,
		"merchantOrderId": req.MerchantOrderId,
		"amount":          json.Number(amountStr), // 使用json.Number确保格式
		"currency":        req.Currency,
		"timestamp":       req.Timestamp,
		"sign":            req.Sign,
		"notifyUrl":       req.NotifyUrl,
		"fundAccount":     req.FundAccount,
	}

	// 序列化请求
	reqBody, err := json.Marshal(reqBodyMap)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	// 日志：打印请求体（用于调试）
	fmt.Printf("[PandaPay] 请求URL: %s\n", c.ApiUrl)
	fmt.Printf("[PandaPay] 请求体: %s\n", string(reqBody))

	// 发送HTTP请求（配置中已包含完整URL，不再拼接）
	httpReq, err := http.NewRequest("POST", c.ApiUrl, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("创建HTTP请求失败: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: 30 * time.Second}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送HTTP请求失败: %w", err)
	}
	defer httpResp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 检查HTTP状态码
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d, 响应: %s", httpResp.StatusCode, string(respBody))
	}

	// 解析响应
	var resp PayoutResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w, 原始响应: %s", err, string(respBody))
	}

	// 检查业务状态码
	// 200/0: 成功
	// 1002: 已转账过，视为成功继续执行
	if resp.Code != 200 && resp.Code != 0 && resp.Code != 1002 {
		return &resp, fmt.Errorf("代付请求失败: %s (code: %d)", resp.Message, resp.Code)
	}

	// 如果是1002，记录日志
	if resp.Code == 1002 {
		fmt.Printf("[PandaPay] 订单已转账过 (code: 1002), 继续执行\n")
	}

	return &resp, nil
}

// BuildBankPayoutRequest 构建银行卡代付请求
func BuildBankPayoutRequest(
	withdrawNo string,
	amount float64,
	currency string,
	accountName string,
	bankCode string,
	accountNumber string,
	email string,
	mobile string,
	address string,
) *PayoutRequest {
	return &PayoutRequest{
		MerchantOrderId: withdrawNo,
		Amount:          amount,
		Currency:        currency,
		FundAccount: FundAccount{
			AccountType: "bank_account",
			Contact: &Contact{
				Name:    accountName,
				Email:   email,
				Mobile:  mobile,
				Address: address,
			},
			BankAccount: &BankAccount{
				Name:          accountName,
				BankCode:      bankCode,
				AccountNumber: accountNumber,
			},
		},
	}
}

// ParseAmount 解析字符串金额为float64
func ParseAmount(amountStr string) (float64, error) {
	if amountStr == "" {
		return 0, errors.New("金额不能为空")
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("金额格式错误: %w", err)
	}
	if amount <= 0 {
		return 0, errors.New("金额必须大于0")
	}
	return amount, nil
}
