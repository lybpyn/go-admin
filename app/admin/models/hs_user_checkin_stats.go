package models

import (

	"go-admin/common/models"

)

type HsUserCheckinStats struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20) unsigned;comment:用户ID"` 
    TotalCheckinDays string `json:"totalCheckinDays" gorm:"type:int(11);comment:总签到天数"` 
    CurrentConsecutiveDays string `json:"currentConsecutiveDays" gorm:"type:int(11);comment:当前连续签到天数"` 
    MaxConsecutiveDays string `json:"maxConsecutiveDays" gorm:"type:int(11);comment:最大连续签到天数"` 
    LastCheckinDate string `json:"lastCheckinDate" gorm:"type:date;comment:最后签到日期"` 
    ThisMonthCheckinDays string `json:"thisMonthCheckinDays" gorm:"type:int(11);comment:本月签到天数"` 
    models.ModelTime
    models.ControlBy
}

func (HsUserCheckinStats) TableName() string {
    return "hs_user_checkin_stats"
}

func (e *HsUserCheckinStats) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsUserCheckinStats) GetId() interface{} {
	return e.Id
}