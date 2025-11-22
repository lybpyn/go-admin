package pandapay

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"go-admin/config"
)

// BankListRequest 银行列表请求参数
type BankListRequest struct {
	MerchantId int    `json:"merchantId"`
	Sign       string `json:"sign"`
}

// BankListResponse 银行列表响应
type BankListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		BankCode string `json:"bankCode"`
		BankName string `json:"bankName"`
	} `json:"data"`
}

// generateBankListSign 生成银行列表签名
// 签名格式: merchantId=999&appSecret=商户appSecret
// 进行MD5加密并转小写
func generateBankListSign(merchantId int, appSecret string) string {
	signStr := fmt.Sprintf("merchantId=%d&appSecret=%s", merchantId, appSecret)
	hash := md5.Sum([]byte(signStr))
	return fmt.Sprintf("%x", hash)
}

// TestGetBankList 测试获取银行列表
func TestGetBankList(t *testing.T) {
	// 加载配置（需要先初始化配置）
	// 如果配置未加载，使用默认测试值
	merchantId := 1099
	appSecret := "sUBI7ooN51T98T1TWtzarkMJpWvuAgXn"
	apiUrl := "https://winpay-oapi-dev.aicapay.com/api/payout/payout/bankList"

	// 如果配置已加载，使用配置值
	if config.ExtConfig.PandaPay.MerchantId != 0 {
		merchantId = config.ExtConfig.PandaPay.MerchantId
		appSecret = config.ExtConfig.PandaPay.AppSecret
	}

	// 生成签名
	sign := generateBankListSign(merchantId, appSecret)

	// 构建请求参数
	req := BankListRequest{
		MerchantId: merchantId,
		Sign:       sign,
	}

	// 序列化请求（格式化输出用）
	reqBodyPretty, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		t.Fatalf("序列化请求失败: %v", err)
	}

	// 实际发送的请求体（压缩格式）
	reqBody, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("序列化请求失败: %v", err)
	}

	// 打印调试信息
	t.Logf("请求URL: %s", apiUrl)
	t.Logf("商户ID: %d", merchantId)
	signStr := fmt.Sprintf("merchantId=%d&appSecret=%s", merchantId, appSecret)
	t.Logf("签名字符串: %s", signStr)
	t.Logf("生成签名: %s", sign)
	t.Logf("请求体JSON格式:\n%s", string(reqBodyPretty))

	// 发送HTTP请求
	httpReq, err := http.NewRequest("POST", apiUrl, bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("创建HTTP请求失败: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: 30 * time.Second}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		t.Fatalf("发送HTTP请求失败: %v", err)
	}
	defer httpResp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		t.Fatalf("读取响应失败: %v", err)
	}

	// 格式化响应体
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, respBody, "", "  "); err != nil {
		// 如果格式化失败，直接输出原始内容
		t.Logf("HTTP状态码: %d", httpResp.StatusCode)
		t.Logf("响应体: %s", string(respBody))
	} else {
		t.Logf("HTTP状态码: %d", httpResp.StatusCode)
		t.Logf("响应体JSON格式:\n%s", prettyJSON.String())
	}

	// 检查HTTP状态码
	if httpResp.StatusCode != http.StatusOK {
		t.Fatalf("HTTP请求失败，状态码: %d, 响应: %s", httpResp.StatusCode, string(respBody))
	}

	// 解析响应
	var resp BankListResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		t.Fatalf("解析响应失败: %v, 原始响应: %s", err, string(respBody))
	}

	// 检查业务状态码
	if resp.Code != 200 && resp.Code != 0 {
		t.Fatalf("获取银行列表失败: %s (code: %d)", resp.Message, resp.Code)
	}

	// 打印银行列表
	t.Logf("成功获取银行列表，共 %d 家银行:", len(resp.Data))
	for i, bank := range resp.Data {
		t.Logf("  [%d] %s - %s", i+1, bank.BankCode, bank.BankName)
	}
}

// TestPayoutSign 测试代付签名生成
func TestPayoutSign(t *testing.T) {
	// 测试参数
	merchantId := 1046
	merchantOrderId := "PX3323133213833222"
	amount := "1000.00"
	appSecret := "liEOD83b93G9kqiEqnxvWWR2Bn3Uc2NS"

	// 生成签名字符串
	signStr := fmt.Sprintf("merchantId=%d&merchantOrderId=%s&amount=%s&appSecret=%s",
		merchantId, merchantOrderId, amount, appSecret)

	// MD5加密并转小写
	hash := md5.Sum([]byte(signStr))
	sign := fmt.Sprintf("%x", hash)

	// 打印结果
	t.Logf("签名字符串: %s", signStr)
	t.Logf("生成签名: %s", sign)
	t.Logf("预期签名: d26e03109b447c59bf110cb4861e053f")

	// 验证签名是否匹配
	if sign == "d26e03109b447c59bf110cb4861e053f" {
		t.Logf("签名验证通过!")
	} else {
		t.Logf("签名不匹配!")
	}
}

// TestAmountFormat 测试金额格式是否保持一致
func TestAmountFormat(t *testing.T) {
	// 测试金额495.00
	amount := 495.0
	amountStr := fmt.Sprintf("%.2f", amount)

	// 构建 JSON
	reqBodyMap := map[string]interface{}{
		"merchantId":      1099,
		"merchantOrderId": "WD1763734583474162836",
		"amount":          json.Number(amountStr),
		"currency":        "NGN",
	}

	// 序列化
	reqBody, err := json.Marshal(reqBodyMap)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	// 打印结果
	t.Logf("金额: %.2f", amount)
	t.Logf("金额字符串: %s", amountStr)
	t.Logf("JSON输出:\n%s", string(reqBody))

	// 验证 JSON 中包含 "amount":495.00 (数字类型，不带引号)
	if bytes.Contains(reqBody, []byte(`"amount":495.00`)) {
		t.Logf("✓ JSON中金额格式正确: 495.00 (数字类型)")
	} else {
		t.Logf("✗ JSON中金额格式不正确")
	}
}
