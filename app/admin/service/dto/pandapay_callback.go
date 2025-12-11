package dto

// PandaPayCallbackReq PandaPay回调请求参数
type PandaPayCallbackReq struct {
	MerchantId           int32  `json:"merchantId" comment:"商户id"`
	MerchantOrderId      string `json:"merchantOrderId" comment:"商户订单id"`
	Amount               string `json:"amount" comment:"订单金额（最多带两位小数）"`
	Timestamp            int64  `json:"timestamp" comment:"时间戳"`
	OrderId              int64  `json:"orderId" comment:"平台订单id"`
	Status               int32  `json:"status" comment:"订单状态 0 等待;1 成功;2 失败;3下单失败;4 处理中"`
	Sign                 string `json:"sign" comment:"签名"`
	Msg                  string `json:"msg" comment:"错误消息"`
	PayTime              string `json:"payTime" comment:"支付时间"`
	SessionId            string `json:"sessionId" comment:"银行交易sessionId"`
	PayoutBankCode       string `json:"payoutBankCode" comment:"收款银行编码"`
	PayoutBankName       string `json:"payoutBankName" comment:"收款银行名称"`
	PayoutCardName       string `json:"payoutCardName" comment:"收款卡号持卡人姓名"`
	PayoutCardNumber     string `json:"payoutCardNumber" comment:"收款卡号"`
	MerchantAmount       string `json:"merchantAmount" comment:"商户账号扣款金额"`
	MerchantRefundAmount string `json:"merchantRefundAmount" comment:"代付交易退回商户账号金额"`
}

// PandaPayCallbackResp PandaPay回调响应
type PandaPayCallbackResp struct {
	Code    int    `json:"code" comment:"响应码 0=成功"`
	Message string `json:"message" comment:"响应消息"`
}
