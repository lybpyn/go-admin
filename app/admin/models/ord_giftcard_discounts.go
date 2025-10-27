package models

import (

	"go-admin/common/models"

)

type OrdGiftcardDiscounts struct {
    models.Model
    
    GiftcardId string `json:"giftcardId" gorm:"type:bigint(20);comment:礼品卡ID -> ord_giftcard.id"` 
    CardType string `json:"cardType" gorm:"type:enum('code','physical','horizontal','whiteboard');comment:卡类型"` 
    DiscountRate string `json:"discountRate" gorm:"type:decimal(5,2);comment:折扣汇率，例如0.95 表示95折"` 
    models.ModelTime
    models.ControlBy
}

func (OrdGiftcardDiscounts) TableName() string {
    return "ord_giftcard_discounts"
}

func (e *OrdGiftcardDiscounts) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdGiftcardDiscounts) GetId() interface{} {
	return e.Id
}