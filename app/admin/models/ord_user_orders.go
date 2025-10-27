package models

import (
    "time"

	"go-admin/common/models"

)

type OrdUserOrders struct {
    models.Model
    
    UserId string `json:"userId" gorm:"type:bigint(20);comment:用户ID"` 
    RegionId string `json:"regionId" gorm:"type:bigint(20);comment:地区ID"` 
    CategoryId string `json:"categoryId" gorm:"type:bigint(20);comment:分类ID"` 
    OrderNo string `json:"orderNo" gorm:"type:varchar(64);comment:订单号"` 
    GiftcardId string `json:"giftcardId" gorm:"type:bigint(20);comment:礼品卡id"` 
    CardType string `json:"cardType" gorm:"type:enum('code','physical','horizontal','whiteboard');comment:卡片类型(Physical,Code,Receipy,Cash Receipt)"` 
    GiftCardCode string `json:"giftCardCode" gorm:"type:varchar(40);comment:礼品卡卡号验证码"` 
    Balance string `json:"balance" gorm:"type:decimal(20,8);comment:卡余额"` 
    Currency string `json:"currency" gorm:"type:char(3);comment:币种，例如 USD, CNY"` 
    DiscountRate string `json:"discountRate" gorm:"type:decimal(5,2);comment:折扣"` 
    Rate string `json:"rate" gorm:"type:decimal(20,8);comment:汇率"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:订单状态: 0=待支付,1=已支付,2=已发卡,3=已完成,4=已取消"` 
    CardExtra string `json:"cardExtra" gorm:"type:json;comment:CardExtra"` 
    CompletedAt time.Time `json:"completedAt" gorm:"type:timestamp;comment:完成时间"` 
    CanceledAt time.Time `json:"canceledAt" gorm:"type:timestamp;comment:取消时间"` 
    models.ModelTime
    models.ControlBy
}

func (OrdUserOrders) TableName() string {
    return "ord_user_orders"
}

func (e *OrdUserOrders) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *OrdUserOrders) GetId() interface{} {
	return e.Id
}