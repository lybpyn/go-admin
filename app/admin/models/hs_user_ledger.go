package models

import (

	"go-admin/common/models"

)

type HsUserLedger struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:UserId"` 
    Currency string `json:"currency" gorm:"type:char(3);comment:ISO 4217，例如 USD/CNY"` 
    Direction string `json:"direction" gorm:"type:tinyint(4);comment:1=入账(credit), -1=出账(debit)"` 
    Amount string `json:"amount" gorm:"type:decimal(18,2);comment:本次发生额，>0"` 
    BalanceBefore string `json:"balanceBefore" gorm:"type:decimal(18,2);comment:记账前余额"` 
    BalanceAfter string `json:"balanceAfter" gorm:"type:decimal(18,2);comment:记账后余额"` 
    BizType string `json:"bizType" gorm:"type:varchar(32);comment:业务类型：order_settlement/withdraw/withdraw_fee/withdraw_reversal/manual_adjust_* 等"` 
    BizId string `json:"bizId" gorm:"type:varchar(64);comment:业务单号：例如订单号/提现单号"` 
    IdempotencyKey string `json:"idempotencyKey" gorm:"type:varchar(128);comment:用于幂等控制：如 ORDER_SETTLED:{order_no}"` 
    RefTable string `json:"refTable" gorm:"type:varchar(64);comment:可选：引用表名"` 
    RefId string `json:"refId" gorm:"type:varchar(64);comment:可选：引用ID"` 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:Remark"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:1=已入账，0=待入账，-1=冲正"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserLedger) TableName() string {
    return "hs_user_ledger"
}

func (e *HsUserLedger) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserLedger) GetId() interface{} {
	return e.Id
}