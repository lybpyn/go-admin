package models

import (

	"go-admin/common/models"

)

type OrdCurrencyRates struct {
    models.Model
    
    QuoteCurrency string `json:"quoteCurrency" gorm:"type:char(3);comment:计价货币 ISO4217，例如 CNY"` 
    Rate string `json:"rate" gorm:"type:decimal(20,8);comment:汇率，表示 1 base_currency = rate quote_currency"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态: 1=启用,0=禁用"` 
    models.ModelTime
    models.ControlBy
}

func (OrdCurrencyRates) TableName() string {
    return "ord_currency_rates"
}

func (e *OrdCurrencyRates) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdCurrencyRates) GetId() interface{} {
	return e.Id
}