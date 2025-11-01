package models

import (

	"go-admin/common/models"

)

type OrdGiftcardRegion struct {
    models.Model
    
    CategoryId string `json:"categoryId" gorm:"type:bigint(20);comment:所属分类ID，外键 -> hs_giftcard_category.id"` 
    RegionCode string `json:"regionCode" gorm:"type:varchar(10);comment:地区代码，如 US、UK、CN"` 
    QuoteCurrency string `json:"quoteCurrency" gorm:"type:varchar(10);comment:货币代码，如 USD, GBP, CNY"` 
    QuoteCurrencySymbol string `json:"quoteCurrencySymbol" gorm:"type:varchar(10);comment:QuoteCurrencySymbol"` 
    Rate string `json:"rate" gorm:"type:decimal(10,4);comment:Rate"` 
    IsMain string `json:"isMain" gorm:"type:tinyint(1);comment:是否默认区域"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态: 1=启用, 0=禁用"` 
    models.ModelTime
    models.ControlBy
}

func (OrdGiftcardRegion) TableName() string {
    return "ord_giftcard_region"
}

func (e *OrdGiftcardRegion) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdGiftcardRegion) GetId() interface{} {
	return e.Id
}