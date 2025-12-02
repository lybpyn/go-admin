package models

import (

	"go-admin/common/models"

)

type OrdConfigGiftcardRegion struct {
    models.Model
    
    RegionName string `json:"regionName" gorm:"type:varchar(128);comment:区域名称"` 
    models.ModelTime
    models.ControlBy
}

func (OrdConfigGiftcardRegion) TableName() string {
    return "ord_config_giftcard_region"
}

func (e *OrdConfigGiftcardRegion) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdConfigGiftcardRegion) GetId() interface{} {
	return e.Id
}