package models

import (

	"go-admin/common/models"

)

type OrdGiftCardSuppliers struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(200);comment:供应商名称"` 
    Code string `json:"code" gorm:"type:varchar(50);comment:供应商代码"` 
    ContactPerson string `json:"contactPerson" gorm:"type:varchar(100);comment:联系人"` 
    ContactEmail string `json:"contactEmail" gorm:"type:varchar(200);comment:联系邮箱"` 
    ContactPhone string `json:"contactPhone" gorm:"type:varchar(50);comment:联系电话"` 
    ApiEndpoint string `json:"apiEndpoint" gorm:"type:varchar(500);comment:API接口地址"` 
    ApiKey string `json:"apiKey" gorm:"type:varchar(200);comment:API密钥"` 
    ApiConfig string `json:"apiConfig" gorm:"type:json;comment:API配置信息"` 
    SettlementRate string `json:"settlementRate" gorm:"type:decimal(5,4);comment:结算费率"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态：0=禁用，1=启用"` 
    models.ModelTime
    models.ControlBy
}

func (OrdGiftCardSuppliers) TableName() string {
    return "ord_gift_card_suppliers"
}

func (e *OrdGiftCardSuppliers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdGiftCardSuppliers) GetId() interface{} {
	return e.Id
}