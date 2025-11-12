package models

import (
	"go-admin/common/models"
)

type OrdGiftcardWriteoffs struct {
	models.Model

	UserId     string `json:"userId" gorm:"type:bigint(20);comment:用户ID，表示提交/使用礼品卡的用户"`
	OrderId    string `json:"orderId" gorm:"type:bigint(20);comment:订单ID，关联 ord_user_orders.id，用于核销对应的订单"`
	GiftCardId string `json:"giftCardId" gorm:"type:bigint(20);comment:礼品卡ID，关联礼品卡主表（若有）"`
	Status     int    `json:"status" gorm:"type:tinyint(4);comment:核销状态：0=待核销，1=已核销，2=失败"`
	Remark     string `json:"remark" gorm:"type:varchar(255);comment:备注信息，例如失败原因、核销说明"`

	// 新增字段
	AdminRecognizedCode     string `json:"adminRecognizedCode" gorm:"type:varchar(64);comment:管理员识别的兑换码"`
	PlatformSaleRate        string `json:"platformSaleRate" gorm:"type:decimal(20,8);comment:平台售卡汇率"`
	RecognizedCardValue     string `json:"recognizedCardValue" gorm:"type:decimal(20,8);comment:识别的卡片面值"`
	FailureImageUrl         string `json:"failureImageUrl" gorm:"type:varchar(255);comment:失败时的截图URL"`
	SupplierId              string `json:"supplierId" gorm:"type:bigint(20);comment:收卡品牌商ID，关联 ord_gift_card_suppliers.id"`
	ConfigRate              string `json:"configRate" gorm:"type:decimal(20,8);comment:配置汇率（后台计算填写）"`
	UserLocalCurrencyAmount string `json:"userLocalCurrencyAmount" gorm:"type:decimal(20,8);comment:用户本地货币金额（后台计算填写）"`
	UserCurrencyCode        string `json:"userCurrencyCode" gorm:"type:varchar(10);comment:用户货币代码（从用户区域获取）"`

	// 平台入账相关字段
	PlatformSettlementAmount   string `json:"platformSettlementAmount" gorm:"type:decimal(20,8);comment:平台入账货币金额"`
	PlatformSettlementCurrency string `json:"platformSettlementCurrency" gorm:"type:varchar(10);comment:平台入账货币代码"`
	PlatformToUsdRate          string `json:"platformToUsdRate" gorm:"type:decimal(20,8);comment:平台入账货币对应美元汇率"`

	// 关联字段（不存储到数据库，仅用于返回）
	OrderNo      string `json:"orderNo" gorm:"column:order_no" comment:"订单号"`
	UserName     string `json:"userName" gorm:"column:user_name" comment:"用户名"`
	GiftCardName string `json:"giftCardName" gorm:"column:gift_card_name" comment:"礼品卡名称"`

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
