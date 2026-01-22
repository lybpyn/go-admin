package dto

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FlexibleString 灵活字符串类型，可以接受字符串或数字
type FlexibleString string

// UnmarshalJSON 自定义 JSON 反序列化，支持字符串和数字
func (fs *FlexibleString) UnmarshalJSON(data []byte) error {
	// 尝试解析为字符串
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*fs = FlexibleString(str)
		return nil
	}

	// 尝试解析为数字
	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*fs = FlexibleString(fmt.Sprintf("%.2f", num))
		return nil
	}

	// 都失败则返回错误
	return fmt.Errorf("FlexibleString: cannot unmarshal %s", string(data))
}

// String 转换为字符串
func (fs FlexibleString) String() string {
	return string(fs)
}

// Float64 转换为浮点数
func (fs FlexibleString) Float64() (float64, error) {
	return strconv.ParseFloat(string(fs), 64)
}

// PandaPayCallbackReq PandaPay回调请求参数
type PandaPayCallbackReq struct {
	MerchantId           int32           `json:"merchantId" form:"merchantId" comment:"商户id"`
	MerchantOrderId      string          `json:"merchantOrderId" form:"merchantOrderId" comment:"商户订单id"`
	Amount               string          `json:"amount" form:"amount" comment:"订单金额（最多带两位小数）"`
	Timestamp            int64           `json:"timestamp" form:"timestamp" comment:"时间戳"`
	OrderId              int64           `json:"orderId" form:"orderId" comment:"平台订单id"`
	Status               int32           `json:"status" form:"status" comment:"订单状态 0 等待;1 成功;2 失败;3下单失败;4 处理中"`
	Sign                 string          `json:"sign" form:"sign" comment:"签名"`
	Msg                  string          `json:"msg" form:"msg" comment:"错误消息"`
	PayTime              string          `json:"payTime" form:"payTime" comment:"支付时间"`
	SessionId            string          `json:"sessionId" form:"sessionId" comment:"银行交易sessionId"`
	PayoutBankCode       string          `json:"payoutBankCode" form:"payoutBankCode" comment:"收款银行编码"`
	PayoutBankName       string          `json:"payoutBankName" form:"payoutBankName" comment:"收款银行名称"`
	PayoutCardName       string          `json:"payoutCardName" form:"payoutCardName" comment:"收款卡号持卡人姓名"`
	PayoutCardNumber     string          `json:"payoutCardNumber" form:"payoutCardNumber" comment:"收款卡号"`
	MerchantAmount       FlexibleString  `json:"merchantAmount" form:"merchantAmount" comment:"商户账号扣款金额"`
	MerchantRefundAmount FlexibleString  `json:"merchantRefundAmount" form:"merchantRefundAmount" comment:"代付交易退回商户账号金额"`
}

// PandaPayCallbackResp PandaPay回调响应
type PandaPayCallbackResp struct {
	Code    int    `json:"code" comment:"响应码 0=成功"`
	Message string `json:"message" comment:"响应消息"`
}
