package models

import (

	"go-admin/common/models"

)

type HsConfgiUserLevels struct {
    models.Model

    LevelName string `json:"levelName" gorm:"type:varchar(50);comment:等级名称"`
    UpExperience string `json:"upExperience" gorm:"type:int(11);comment:升级所需经验值"`
    LevelIcon string `json:"levelIcon" gorm:"type:varchar(255);comment:等级图标URL"`
    RebateRate string `json:"rebateRate" gorm:"type:decimal(5,2);comment:返利比例"`
    SortOrder int `json:"sortOrder" gorm:"type:int(11);comment:排序顺序"`
    IsActive string `json:"isActive" gorm:"type:tinyint(1);comment:是否启用"`
    models.ModelTime
    models.ControlBy
}

func (HsConfgiUserLevels) TableName() string {
    return "hs_confgi_user_levels"
}

func (e *HsConfgiUserLevels) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfgiUserLevels) GetId() interface{} {
	return e.Id
}