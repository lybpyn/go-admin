package models

import (

	"go-admin/common/models"

)

type HsBanks struct {
    models.Model
    
    BankCode string `json:"bankCode" gorm:"type:varchar(64);comment:银行编码/银行行号，如 SWIFT/BIC 或自定义编码"` 
    Name string `json:"name" gorm:"type:varchar(128);comment:银行名称"` 
    Country string `json:"country" gorm:"type:char(2);comment:国家 ISO2"` 
    SwiftCode string `json:"swiftCode" gorm:"type:varchar(32);comment:SWIFT/BIC (如适用)"` 
    RoutingNumber string `json:"routingNumber" gorm:"type:varchar(64);comment:路由/清算号(如 ABA)"` 
    SupportsInternational string `json:"supportsInternational" gorm:"type:tinyint(4);comment:是否支持国际汇款: 0否,1是"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态: 0=禁用,1=启用"` 
    Note string `json:"note" gorm:"type:text;comment:Note"` 
    Extra string `json:"extra" gorm:"type:json;comment:扩展字段，例如银行支付说明、证件要求等"` 
    models.ModelTime
    models.ControlBy
}

func (HsBanks) TableName() string {
    return "hs_banks"
}

func (e *HsBanks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsBanks) GetId() interface{} {
	return e.Id
}