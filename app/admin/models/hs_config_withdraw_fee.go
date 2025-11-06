package models

import (

	"go-admin/common/models"

)

type HsConfigWithdrawFee struct {
    models.Model
    
    CurrencyCode string `json:"currencyCode" gorm:"type:char(3);comment:币种代码，如 USD/CNY/USDT"` 
    MinAmount string `json:"minAmount" gorm:"type:decimal(18,2);comment:最小提现金额"` 
    MaxAmount string `json:"maxAmount" gorm:"type:decimal(18,2);comment:最大提现金额"` 
    FeeRate string `json:"feeRate" gorm:"type:decimal(6,4);comment:手续费率（例如 0.015 表示 1.5%）"` 
    IsActive string `json:"isActive" gorm:"type:tinyint(1);comment:是否启用：1=启用，0=禁用"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigWithdrawFee) TableName() string {
    return "hs_config_withdraw_fee"
}

func (e *HsConfigWithdrawFee) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigWithdrawFee) GetId() interface{} {
	return e.Id
}