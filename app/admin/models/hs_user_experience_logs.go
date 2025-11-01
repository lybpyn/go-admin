package models

import (

	"go-admin/common/models"

)

type HsUserExperienceLogs struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20);comment:用户ID"` 
    ExperienceChange string `json:"experienceChange" gorm:"type:int(11);comment:经验值变化(正数为增加，负数为减少)"` 
    ExperienceBefore string `json:"experienceBefore" gorm:"type:int(11);comment:变化前经验值"` 
    ExperienceAfter string `json:"experienceAfter" gorm:"type:int(11);comment:变化后经验值"` 
    SourceType string `json:"sourceType" gorm:"type:varchar(50);comment:经验来源类型(checkin,gift_card_exchange,order_complete等)"` 
    SourceId string `json:"sourceId" gorm:"type:varchar(100);comment:来源记录ID"` 
    Description string `json:"description" gorm:"type:varchar(255);comment:经验变化描述"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserExperienceLogs) TableName() string {
    return "hs_user_experience_logs"
}

func (e *HsUserExperienceLogs) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserExperienceLogs) GetId() interface{} {
	return e.Id
}