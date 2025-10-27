package models

import (
    "time"

	"go-admin/common/models"

)

type HsUserCheckins struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:用户ID"` 
    CheckinDate string `json:"checkinDate" gorm:"type:date;comment:签到日期"` 
    CheckinTime time.Time `json:"checkinTime" gorm:"type:datetime;comment:签到时间"` 
    ConsecutiveDays string `json:"consecutiveDays" gorm:"type:int(11);comment:连续签到天数"` 
    RewardPoints string `json:"rewardPoints" gorm:"type:int(11);comment:签到奖励积分"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserCheckins) TableName() string {
    return "hs_user_checkins"
}

func (e *HsUserCheckins) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserCheckins) GetId() interface{} {
	return e.Id
}