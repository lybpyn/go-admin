package models

import (
    "time"

	"go-admin/common/models"

)

type OrdConfigCurrencyRates struct {
    models.Model
    
    BaseCurrencyCode string `json:"baseCurrencyCode" gorm:"type:varchar(10);comment:基准货币代码 (ISO 4217)"` 
    QuoteCurrencyCode string `json:"quoteCurrencyCode" gorm:"type:varchar(10);comment:报价货币代码 (ISO 4217)"` 
    Rate string `json:"rate" gorm:"type:decimal(20,8);comment:汇率: 1 base_currency = rate quote_currency"` 
    RegionCode string `json:"regionCode" gorm:"type:varchar(10);comment:地区代码,为空表示全局汇率"` 
    RateType string `json:"rateType" gorm:"type:varchar(20);comment:汇率类型: standard=标准, buying=买入, selling=卖出"` 
    Source string `json:"source" gorm:"type:varchar(50);comment:汇率来源,如 manual, api, coingecko"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态: 1=启用, 0=禁用"` 
    ValidFrom time.Time `json:"validFrom" gorm:"type:datetime;comment:生效开始时间"` 
    ValidTo time.Time `json:"validTo" gorm:"type:datetime;comment:生效结束时间"` 
    models.ModelTime
    models.ControlBy
}

func (OrdConfigCurrencyRates) TableName() string {
    return "ord_config_currency_rates"
}

func (e *OrdConfigCurrencyRates) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdConfigCurrencyRates) GetId() interface{} {
	return e.Id
}