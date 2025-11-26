package models

import (

	"go-admin/common/models"

)

type HsBanks struct {
    models.Model

    BankCode      string `json:"bankCode" gorm:"type:varchar(64);comment:银行编码/银行行号，如 SWIFT/BIC 或自定义编码"`
    BankName      string `json:"bankName" gorm:"type:varchar(128);comment:银行名称"`
    BankNameShort string `json:"bankNameShort" gorm:"type:varchar(80);comment:银行名称缩写"`
    ChannelType   int64  `json:"channelType" gorm:"type:bigint(20);comment:通道类型"`
    Status        int    `json:"status" gorm:"type:tinyint(4);default:1;comment:状态: 0=禁用,1=启用"`
    models.ModelTime
    models.ControlBy
}

func (HsBanks) TableName() string {
    return "hs_banks"
}

func (e *HsBanks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsBanks) GetId() interface{} {
	return e.Id
}