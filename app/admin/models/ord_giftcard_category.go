package models

import (

	"go-admin/common/models"

)

type OrdGiftcardCategory struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(128);comment:分类名称，如 Steam、eBay"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态: 1=启用, 0=禁用"` 
    DiscountRate string `json:"discountRate" gorm:"type:decimal(5,2);comment:汇率折扣展示用"` 
    SortOrder string `json:"sortOrder" gorm:"type:int(11);comment:SortOrder"` 
    models.ModelTime
    models.ControlBy
}

func (OrdGiftcardCategory) TableName() string {
    return "ord_giftcard_category"
}

func (e *OrdGiftcardCategory) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdGiftcardCategory) GetId() interface{} {
	return e.Id
}