package models

import (

	"go-admin/common/models"

)

type HsUsers struct {
    models.Model
    
    Username string `json:"username" gorm:"type:varchar(50);comment:用户名（可选展示用）"` 
    PasswordHash string `json:"passwordHash" gorm:"type:varchar(255);comment:PasswordHash"` 
    Firstname string `json:"firstname" gorm:"type:varchar(20);comment:Firstname"` 
    Lastname string `json:"lastname" gorm:"type:varchar(20);comment:Lastname"` 
    Avatar string `json:"avatar" gorm:"type:varchar(255);comment:头像URL"` 
    Balance string `json:"balance" gorm:"type:decimal(12,2);comment:Balance"` 
    LevelId string `json:"levelId" gorm:"type:int(11);comment:用户等级ID"` 
    Experience string `json:"experience" gorm:"type:int(11);comment:当前经验"` 
    TotalExperience string `json:"totalExperience" gorm:"type:int(11);comment:累计经验"` 
    InviteCode string `json:"inviteCode" gorm:"type:varchar(8);comment:InviteCode"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态：1正常，0封禁"` 
    Version string `json:"version" gorm:"type:bigint(20);comment:Version"` 
    models.ModelTime
    models.ControlBy
}

func (HsUsers) TableName() string {
    return "hs_users"
}

func (e *HsUsers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUsers) GetId() interface{} {
	return e.Id
}