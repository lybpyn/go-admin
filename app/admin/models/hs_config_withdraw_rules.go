package models

import (

	"go-admin/common/models"

)

type HsConfigWithdrawRules struct {
    models.Model
    
    CurrencyCode string `json:"currencyCode" gorm:"type:varchar(10);comment:币种代码，如 USD、CNY、USDT、BTC"` 
    CurrencyType string `json:"currencyType" gorm:"type:enum('fiat','crypto');comment:币种类型：fiat=法币，crypto=虚拟币"` 
    ChainType string `json:"chainType" gorm:"type:varchar(20);comment:链类型（仅虚拟币适用），如 ERC20/TRC20/BEP20"` 
    SingleMin string `json:"singleMin" gorm:"type:decimal(36,18);comment:单笔最小提现数量/金额"` 
    SingleMax string `json:"singleMax" gorm:"type:decimal(36,18);comment:单笔最大提现数量/金额"` 
    DailyLimitAmount string `json:"dailyLimitAmount" gorm:"type:decimal(36,18);comment:每日累计提现上限"` 
    DailyLimitCount string `json:"dailyLimitCount" gorm:"type:int(11);comment:每日提现次数上限"` 
    FeeType string `json:"feeType" gorm:"type:enum('fixed','rate','mixed','tiered');comment:手续费类型：fixed=固定，rate=按比例，mixed=固定+比例，tiered=阶梯收费"` 
    FeeFixed string `json:"feeFixed" gorm:"type:decimal(36,18);comment:固定手续费数量/金额"` 
    FeeRate string `json:"feeRate" gorm:"type:decimal(8,6);comment:手续费率（例如 0.015 表示 1.5%）"` 
    MinFee string `json:"minFee" gorm:"type:decimal(36,18);comment:最小手续费"` 
    MaxFee string `json:"maxFee" gorm:"type:decimal(36,18);comment:最高手续费"` 
    IsActive string `json:"isActive" gorm:"type:tinyint(1);comment:是否启用：1=启用，0=禁用"`
    FeeTiers []HsConfigWithdrawFeeTiers `json:"feeTiers" gorm:"foreignKey:RuleId;references:Id"`
    models.ModelTime
    models.ControlBy
}

func (HsConfigWithdrawRules) TableName() string {
    return "hs_config_withdraw_rules"
}

func (e *HsConfigWithdrawRules) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigWithdrawRules) GetId() interface{} {
	return e.Id
}