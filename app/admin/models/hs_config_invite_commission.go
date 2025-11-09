package models

import (

	"go-admin/common/models"

)

type HsConfigInviteCommission struct {
    models.Model
    
    ConfigName string `json:"configName" gorm:"type:varchar(100);comment:配置名称"` 
    FirstLevelRate string `json:"firstLevelRate" gorm:"type:decimal(8,4);comment:一级邀请分成比例（百分比，如5.0000表示5%）"` 
    SecondLevelRate string `json:"secondLevelRate" gorm:"type:decimal(8,4);comment:二级邀请分成比例（百分比，如3.0000表示3%）"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态：1=启用，0=禁用"` 
    Remark string `json:"remark" gorm:"type:varchar(500);comment:备注说明"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigInviteCommission) TableName() string {
    return "hs_config_invite_commission"
}

func (e *HsConfigInviteCommission) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigInviteCommission) GetId() interface{} {
	return e.Id
}