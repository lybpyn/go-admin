package models

import (

	"go-admin/common/models"

)

type OrdGiftcardWriteoffs struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20);comment:用户ID，表示提交/使用礼品卡的用户"` 
    OrderId string `json:"orderId" gorm:"type:bigint(20);comment:订单ID，关联 ord_user_orders.id，用于核销对应的订单"` 
    GiftCardId string `json:"giftCardId" gorm:"type:bigint(20);comment:礼品卡ID，关联礼品卡主表（若有）"` 
    GiftCardCode string `json:"giftCardCode" gorm:"type:varchar(64);comment:礼品卡卡号/兑换码，保证唯一性"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:核销状态：0=待核销，1=已核销，2=失败"` 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:备注信息，例如失败原因、核销说明"` 
    models.ModelTime
    models.ControlBy
}

func (OrdGiftcardWriteoffs) TableName() string {
    return "ord_giftcard_writeoffs"
}

func (e *OrdGiftcardWriteoffs) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdGiftcardWriteoffs) GetId() interface{} {
	return e.Id
}