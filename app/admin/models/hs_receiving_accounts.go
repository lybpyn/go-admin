package models

import (

	"go-admin/common/models"

)

type HsReceivingAccounts struct {
    models.Model
    
    MerchantId string `json:"merchantId" gorm:"type:bigint(20);comment:关联 merchants.id"` 
    BankId string `json:"bankId" gorm:"type:bigint(20);comment:关联 banks.id (如果是银行账户)"` 
    AccountName string `json:"accountName" gorm:"type:varchar(128);comment:账户名/收款人名称"` 
    AccountNumber string `json:"accountNumber" gorm:"type:varchar(128);comment:账号（可能为卡号/IBAN/虚拟账号）"` 
    Currency string `json:"currency" gorm:"type:varchar(8);comment:货币代码 (USD/CNY/...)"` 
    MinAmount string `json:"minAmount" gorm:"type:decimal(20,4);comment:单次最小收款金额"` 
    MaxAmount string `json:"maxAmount" gorm:"type:decimal(20,4);comment:单次最大收款金额 (NULL 表示无限制)"` 
    DailyLimit string `json:"dailyLimit" gorm:"type:decimal(20,4);comment:每日上限 (平台/风控)"` 
    IsDefault string `json:"isDefault" gorm:"type:tinyint(4);comment:是否默认收款账号: 0否,1是"` 
    AutoTransferEnabled string `json:"autoTransferEnabled" gorm:"type:tinyint(4);comment:是否启用自动转出/自动结算"` 
    AutoTransferThreshold string `json:"autoTransferThreshold" gorm:"type:decimal(20,4);comment:自动转出阈值 (达到该金额触发)"` 
    RiskThresholdAmount string `json:"riskThresholdAmount" gorm:"type:decimal(20,4);comment:风控阈值(超过则人工审核) —— 可与汇率表联动计算"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态:0=禁用,1=启用,2=待审核"` 
    InterfaceType string `json:"interfaceType" gorm:"type:varchar(64);comment:接口/通道类型，如 bank_transfer, third_party_api 等"` 
    InterfaceConfig string `json:"interfaceConfig" gorm:"type:json;comment:接口配置信息（JSON）: 如 api_key, secret, callback_url, 接口码等"` 
    AllowedCountries string `json:"allowedCountries" gorm:"type:json;comment:允许接收国家数组，例如 ["CN","US"]"` 
    Note string `json:"note" gorm:"type:text;comment:Note"` 
    models.ModelTime
    models.ControlBy
}

func (HsReceivingAccounts) TableName() string {
    return "hs_receiving_accounts"
}

func (e *HsReceivingAccounts) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsReceivingAccounts) GetId() interface{} {
	return e.Id
}