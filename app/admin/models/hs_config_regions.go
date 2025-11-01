package models

import (

	"go-admin/common/models"

)

type HsConfigRegions struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(100);comment:地区名称"` 
    Code string `json:"code" gorm:"type:varchar(20);comment:地区代码，如 CN、US、JP 等"` 
    IsActive string `json:"isActive" gorm:"type:tinyint(4);comment:是否启用：1=启用，0=禁用"` 
    models.ModelTime
    models.ControlBy
}

func (HsConfigRegions) TableName() string {
    return "hs_config_regions"
}

func (e *HsConfigRegions) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsConfigRegions) GetId() interface{} {
	return e.Id
}