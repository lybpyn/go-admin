package models

import (

	"go-admin/common/models"

)

type OrdGiftcard struct {
    models.Model
    
    RegionId string `json:"regionId" gorm:"type:bigint(20);comment:地区ID"` 
    Name string `json:"name" gorm:"type:varchar(128);comment:卡名称，例如 Steam 50 USD"` 
    ValuesConfig string `json:"valuesConfig" gorm:"type:json;comment:面额配置，支持区间和固定值"` 
    DiscountRate string `json:"discountRate" gorm:"type:decimal(5,2);comment:折扣汇率，例如0.95 表示95折"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态: 1=启用, 0=禁用"` 
    Extra string `json:"extra" gorm:"type:json;comment:扩展信息，如购买说明、限制条件"` 
    models.ModelTime
    models.ControlBy
}

func (OrdGiftcard) TableName() string {
    return "ord_giftcard"
}

func (e *OrdGiftcard) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdGiftcard) GetId() interface{} {
	return e.Id
}