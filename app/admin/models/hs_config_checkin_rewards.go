package models

import (

	"go-admin/common/models"

)

type HsConfigCheckinRewards struct {
    models.Model
    
    ConsecutiveDays string `json:"consecutiveDays" gorm:"type:int(11);comment:连续签到天数"` 
    ExperienceReward string `json:"experienceReward" gorm:"type:int(11);comment:经验值奖励"` 
    ExtraRewards string `json:"extraRewards" gorm:"type:json;comment:额外奖励配置(JSON格式)"` 
    IsActive string `json:"isActive" gorm:"type:tinyint(1);comment:是否启用"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigCheckinRewards) TableName() string {
    return "hs_config_checkin_rewards"
}

func (e *HsConfigCheckinRewards) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigCheckinRewards) GetId() interface{} {
	return e.Id
}