package models

import (

	"go-admin/common/models"

)

type HsConfigWithdrawLimit struct {
    models.Model
    
    CurrencyCode string `json:"currencyCode" gorm:"type:char(3);comment:币种代码，如 USD/CNY/USDT"` 
    SingleLimit string `json:"singleLimit" gorm:"type:decimal(18,2);comment:单笔提现限额"` 
    DailyLimitAmount string `json:"dailyLimitAmount" gorm:"type:decimal(18,2);comment:每日累计提现金额上限"` 
    DailyLimitCount string `json:"dailyLimitCount" gorm:"type:int(11);comment:每日提现次数上限"` 
    IsActive string `json:"isActive" gorm:"type:tinyint(1);comment:是否启用：1=启用，0=禁用"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigWithdrawLimit) TableName() string {
    return "hs_config_withdraw_limit"
}

func (e *HsConfigWithdrawLimit) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigWithdrawLimit) GetId() interface{} {
	return e.Id
}