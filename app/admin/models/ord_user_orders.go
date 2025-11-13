package models

import (
	"time"

	"go-admin/common/models"
)

type OrdUserOrders struct {
	models.Model
	OrderNo      string    `json:"OrderNo" gorm:"type:varchar(255);comment:订单号"`
	UserId       int       `json:"userId" gorm:"type:bigint(20);comment:用户ID"`
	RegionId     int       `json:"regionId" gorm:"type:bigint(20);comment:地区ID"`
	CategoryId   int       `json:"categoryId" gorm:"type:bigint(20);comment:分类ID"`
	GiftcardId   int       `json:"giftcardId" gorm:"type:bigint(20);comment:礼品卡ID"`
	CardType     string    `json:"cardType" gorm:"type:enum('code','physical','image','horizontal','whiteboard');comment:卡片类型: code=兑换码, image=图片, physical=实体卡等"`
	GiftCardCode string    `json:"giftCardCode" gorm:"type:varchar(40);comment:礼品卡卡号或验证码"`
	Balance      string    `json:"balance" gorm:"type:decimal(20,8);comment:卡余额"`
	CurrencyCode string    `json:"currencyCode" gorm:"type:char(3);comment:币种，例如 USD, CNY"`
	DiscountRate string    `json:"discountRate" gorm:"type:decimal(5,2);comment:折扣"`
	Rate         string    `json:"rate" gorm:"type:decimal(20,8);comment:汇率"`
	VerifySpeed  string    `json:"verifySpeed" gorm:"type:varchar(32);comment:验证速度，如 instant, fast, normal, slow 或秒数描述"`
	BalanceType  int       `json:"balanceType" gorm:"type:varchar(255);comment:入账类型：1=法币余额，2=虚拟币余额"`
	Status       int       `json:"status" gorm:"type:tinyint(4);comment:订单状态: 0=待处理,1=已经接单,2=已完成,3=已取消,4=已经驳回"`
	CardExtra    string    `json:"cardExtra" gorm:"type:json;comment:卡片附加信息（如过期时间、面额说明等）"`
	CompletedAt  time.Time `json:"completedAt" gorm:"type:timestamp;comment:完成时间"`
	Remark       string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	AdminRemark  string    `json:"adminRemark" gorm:"type:varchar(120);comment:管理员备注"`

	AssignType           int       `json:"assignType" gorm:"type:tinyint(1);comment:接单类型：1=人工接单，2=系统派单"`
	AssignBy             int       `json:"assignBy" gorm:"type:bigint(20);comment:派单人ID（接单人ID）"`
	AssignName           string    `json:"assignName" gorm:"type:varchar(128);comment:派单人名称（冗余字段）"`
	ProcessingStartedEnd time.Time `json:"processingStartedEnd" gorm:"type:datetime(3);comment:管理员结束处理时间"`
	ProcessingStartedAt  time.Time `json:"processingStartedAt" gorm:"type:datetime(3);comment:管理员开始处理时间"`
	ProcessingDuration   int       `json:"processingDuration" gorm:"type:int(11);comment:处理耗时（秒）"`
	ProcessingStatus     int       `json:"processingStatus" gorm:"type:tinyint(1);comment:处理状态1正在处理2取消3完成"`
	CanceledAt           time.Time `json:"canceledAt" gorm:"type:timestamp;comment:取消时间"`

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
