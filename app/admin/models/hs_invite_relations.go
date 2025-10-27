package models

import (

	"go-admin/common/models"

)

type HsInviteRelations struct {
    models.Model
    
    Level1InviterId string `json:"level1InviterId" gorm:"type:bigint(20);comment:一级邀请人ID"` 
    Level2InviterId string `json:"level2InviterId" gorm:"type:bigint(20);comment:二级邀请人ID"` 
    models.ModelTime
    models.ControlBy
}

func (HsInviteRelations) TableName() string {
    return "hs_invite_relations"
}

func (e *HsInviteRelations) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsInviteRelations) GetId() interface{} {
	return e.Id
}