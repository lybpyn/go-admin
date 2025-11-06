package models

import (

	"go-admin/common/models"

)

type HsUserFrozenLedger struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:UserId"` 
    CurrencyCode string `json:"currencyCode" gorm:"type:char(3);comment:币种代码，如 USD/CNY/USDT"` 
    Direction string `json:"direction" gorm:"type:tinyint(4);comment:1=冻结增加，-1=冻结减少(解冻)"` 
    Amount string `json:"amount" gorm:"type:decimal(18,2);comment:冻结或解冻金额"` 
    FrozenBefore string `json:"frozenBefore" gorm:"type:decimal(18,2);comment:变动前冻结余额"` 
    FrozenAfter string `json:"frozenAfter" gorm:"type:decimal(18,2);comment:变动后冻结余额"` 
    BizType string `json:"bizType" gorm:"type:varchar(32);comment:业务类型：invite_commissions/order_rebate等"` 
    BizId string `json:"bizId" gorm:"type:varchar(64);comment:业务单号"` 
    IdempotencyKey string `json:"idempotencyKey" gorm:"type:varchar(128);comment:幂等键"` 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:Remark"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:1=已冻结或解冻，0=待处理，-1=冲正"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserFrozenLedger) TableName() string {
    return "hs_user_frozen_ledger"
}

func (e *HsUserFrozenLedger) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserFrozenLedger) GetId() interface{} {
	return e.Id
}