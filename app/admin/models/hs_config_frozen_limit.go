package models

import (

	"go-admin/common/models"

)

type HsConfigFrozenLimit struct {
    models.Model
    
    CurrencyCode string `json:"currencyCode" gorm:"type:char(3);comment:币种代码，如 USD/CNY/USDT"` 
    FrozenLimitAmount string `json:"frozenLimitAmount" gorm:"type:decimal(18,2);comment:可冻结金额上限 / 提现限制金额"` 
    IsActive string `json:"isActive" gorm:"type:tinyint(1);comment:是否启用：1=启用，0=禁用"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigFrozenLimit) TableName() string {
    return "hs_config_frozen_limit"
}

func (e *HsConfigFrozenLimit) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigFrozenLimit) GetId() interface{} {
	return e.Id
}