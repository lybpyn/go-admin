package models

import (

	"go-admin/common/models"

)

type HsInvitesBind struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20);comment:被邀请用户ID"` 
    InviterId string `json:"inviterId" gorm:"type:bigint(20);comment:直接邀请人ID (一级)"` 
    Level string `json:"level" gorm:"type:tinyint(4);comment:层级: 1=一级, 2=二级"` 
    models.ModelTime
    models.ControlBy
}

func (HsInvitesBind) TableName() string {
    return "hs_invites_bind"
}

func (e *HsInvitesBind) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsInvitesBind) GetId() interface{} {
	return e.Id
}