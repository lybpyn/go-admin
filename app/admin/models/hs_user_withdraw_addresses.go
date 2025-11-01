package models

import (

	"go-admin/common/models"

)

type HsUserWithdrawAddresses struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:所属用户ID"` 
    ConfigSupportedChainsCoinId string `json:"configSupportedChainsCoinId" gorm:"type:bigint(20);comment:链类型，例如 TRC20、ERC20、BEP20"` 
    Address string `json:"address" gorm:"type:varchar(128);comment:钱包提现地址"` 
    Label string `json:"label" gorm:"type:varchar(64);comment:用户自定义备注标签，例如“我的TRC钱包”"` 
    IsDefault string `json:"isDefault" gorm:"type:tinyint(1);comment:是否默认提现地址：0=否，1=是"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:状态：1=启用，0=禁用，2=待审核，-1=拒绝"` 
    CreatedIp string `json:"createdIp" gorm:"type:varchar(64);comment:添加时的IP地址"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserWithdrawAddresses) TableName() string {
    return "hs_user_withdraw_addresses"
}

func (e *HsUserWithdrawAddresses) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserWithdrawAddresses) GetId() interface{} {
	return e.Id
}