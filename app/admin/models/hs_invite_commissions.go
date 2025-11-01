package models

import (

	"go-admin/common/models"

)

type HsInviteCommissions struct {
    models.Model
    
    OrderId string `json:"orderId" gorm:"type:bigint(20);comment:订单ID"` 
    UserId string `json:"userId" gorm:"type:bigint(20);comment:分成获得者ID"` 
    CommissionLevel string `json:"commissionLevel" gorm:"type:tinyint(4);comment:分成层级: 1=一级 5%, 2=二级 3%"` 
    CommissionRate string `json:"commissionRate" gorm:"type:decimal(5,2);comment:分成比例，例如 0.05 表示5%"` 
    CommissionAmount string `json:"commissionAmount" gorm:"type:decimal(20,8);comment:分成金额"` 
    Status string `json:"status" gorm:"type:tinyint(4);comment:状态: 0=待结算,1=已结算,2=取消"` 
    models.ModelTime
    models.ControlBy
}

func (HsInviteCommissions) TableName() string {
    return "hs_invite_commissions"
}

func (e *HsInviteCommissions) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *HsInviteCommissions) GetId() interface{} {
	return e.Id
}