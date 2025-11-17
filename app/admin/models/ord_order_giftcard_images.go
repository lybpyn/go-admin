package models

import (

	"go-admin/common/models"

)

type OrdOrderGiftcardImages struct {
    models.Model
    
    OrderId string `json:"orderId" gorm:"type:bigint(20);comment:关联的订单ID"` 
    ImageUrl string `json:"imageUrl" gorm:"type:varchar(255);comment:礼品卡图片URL"` 
    SortOrder int `json:"sortOrder" gorm:"type:int(11);comment:排序顺序"` 
    models.ModelTime
    models.ControlBy
}

func (OrdOrderGiftcardImages) TableName() string {
    return "ord_order_giftcard_images"
}

func (e *OrdOrderGiftcardImages) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdOrderGiftcardImages) GetId() interface{} {
	return e.Id
}