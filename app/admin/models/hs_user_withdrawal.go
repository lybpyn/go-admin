package models

import (
    "time"

	"go-admin/common/models"

)

type HsUserWithdrawal struct {
    models.Model
    
    WithdrawNo string `json:"withdrawNo" gorm:"type:varchar(64);comment:提现单号，唯一"` 
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:用户ID"` 
    Currency string `json:"currency" gorm:"type:char(3);comment:ISO 4217币种代码，如 USD/CNY"` 
    Amount string `json:"amount" gorm:"type:decimal(18,2);comment:提现金额"` 
    Fee string `json:"fee" gorm:"type:decimal(18,2);comment:提现手续费"` 
    NetAmount string `json:"netAmount" gorm:"type:decimal(18,2);comment:实际出账金额（amount - fee）"` 
    Method string `json:"method" gorm:"type:varchar(32);comment:提现方式：bank/paypal/crypto/…"` 
    AccountInfo string `json:"accountInfo" gorm:"type:json;comment:提现账户信息（脱敏）"` 
    Status string `json:"status" gorm:"type:varchar(16);comment:状态：pending/review/processing/success/failed/canceled"` 
    Reason string `json:"reason" gorm:"type:varchar(255);comment:失败/取消原因"` 
    ChannelTxnId string `json:"channelTxnId" gorm:"type:varchar(128);comment:通道回执流水号"` 
    RequestedAt time.Time `json:"requestedAt" gorm:"type:datetime(3);comment:发起时间"` 
    ProcessedAt time.Time `json:"processedAt" gorm:"type:datetime(3);comment:处理完成时间"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserWithdrawal) TableName() string {
    return "hs_user_withdrawal"
}

func (e *HsUserWithdrawal) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserWithdrawal) GetId() interface{} {
	return e.Id
}