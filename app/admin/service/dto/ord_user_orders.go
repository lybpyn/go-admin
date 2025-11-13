package dto

import (
    "strconv"
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdUserOrdersGetPageReq struct {
	dto.Pagination     `search:"-"`
	OrderNo string `form:"orderNo" search:"type:exact;column:order_no;table:ord_user_orders" comment:"订单号"`
	UserId string `form:"userId" search:"type:exact;column:user_id;table:ord_user_orders" comment:"用户ID"`
	Status int `form:"status" search:"type:exact;column:status;table:ord_user_orders" comment:"订单状态:5=待处理,1=已经接单,2=已完成,3=已取消,4=已经驳回"`
	ProcessingStatus int `form:"processingStatus" search:"type:exact;column:processing_status;table:ord_user_orders" comment:"管理员处理状态:1=正在处理,2=取消,3=完成"`
	BeginTime string `form:"beginTime" search:"type:gte;column:created_at;table:ord_user_orders" comment:"开始时间"`
	EndTime string `form:"endTime" search:"type:lte;column:created_at;table:ord_user_orders" comment:"结束时间"`
    OrdUserOrdersOrder
}

type OrdUserOrdersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_user_orders"`
    OrderNo string `form:"orderNoOrder"  search:"type:order;column:order_no;table:ord_user_orders"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:ord_user_orders"`
    RegionId string `form:"regionIdOrder"  search:"type:order;column:region_id;table:ord_user_orders"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:ord_user_orders"`
    GiftcardId string `form:"giftcardIdOrder"  search:"type:order;column:giftcard_id;table:ord_user_orders"`
    CardType string `form:"cardTypeOrder"  search:"type:order;column:card_type;table:ord_user_orders"`
    GiftCardCode string `form:"giftCardCodeOrder"  search:"type:order;column:gift_card_code;table:ord_user_orders"`
    Balance string `form:"balanceOrder"  search:"type:order;column:balance;table:ord_user_orders"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:ord_user_orders"`
    DiscountRate string `form:"discountRateOrder"  search:"type:order;column:discount_rate;table:ord_user_orders"`
    Rate string `form:"rateOrder"  search:"type:order;column:rate;table:ord_user_orders"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_user_orders"`
    CardExtra string `form:"cardExtraOrder"  search:"type:order;column:card_extra;table:ord_user_orders"`
    CompletedAt string `form:"completedAtOrder"  search:"type:order;column:completed_at;table:ord_user_orders"`
    CanceledAt string `form:"canceledAtOrder"  search:"type:order;column:canceled_at;table:ord_user_orders"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_user_orders"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_user_orders"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_user_orders"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_user_orders"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_user_orders"`

}

func (m *OrdUserOrdersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdUserOrdersInsertReq struct {
    Id int `json:"-" comment:"订单ID"` // 订单ID
    UserId string `json:"userId" comment:"用户ID"`
    RegionId string `json:"regionId" comment:"地区ID"`
    CategoryId string `json:"categoryId" comment:"分类ID"`
    GiftcardId string `json:"giftcardId" comment:"礼品卡id"`
    CardType string `json:"cardType" comment:"卡片类型(Physical,Code,Receipy,Cash Receipt)"`
    GiftCardCode string `json:"giftCardCode" comment:"礼品卡卡号验证码"`
    Balance string `json:"balance" comment:"卡余额"`
    CurrencyCode string `json:"currencyCode" comment:"币种，例如 USD, CNY"`
    DiscountRate string `json:"discountRate" comment:"折扣"`
    Rate string `json:"rate" comment:"汇率"`
    Status int `json:"status" comment:"订单状态:0=待处理,1=已经接单,2=已完成,3=已取消,4=已经驳回"`
    CardExtra string `json:"cardExtra" comment:""`
    CompletedAt time.Time `json:"completedAt" comment:"完成时间"`
    CanceledAt time.Time `json:"canceledAt" comment:"取消时间"`
    Remark string `json:"remark" comment:"备注"`
    common.ControlBy
}

func (s *OrdUserOrdersInsertReq) Generate(model *models.OrdUserOrders)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId, _ = strconv.Atoi(s.UserId)
    model.RegionId, _ = strconv.Atoi(s.RegionId)
    model.CategoryId, _ = strconv.Atoi(s.CategoryId)
    model.GiftcardId, _ = strconv.Atoi(s.GiftcardId)
    model.CardType = s.CardType
    model.GiftCardCode = s.GiftCardCode
    model.Balance = s.Balance
    model.CurrencyCode = s.CurrencyCode
    model.DiscountRate = s.DiscountRate
    model.Rate = s.Rate
    model.Status = s.Status
    model.CardExtra = s.CardExtra
    model.CompletedAt = s.CompletedAt
    model.CanceledAt = s.CanceledAt
    model.Remark = s.Remark
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdUserOrdersInsertReq) GetId() interface{} {
	return s.Id
}

type OrdUserOrdersUpdateReq struct {
    Id int `uri:"id" comment:"订单ID"` // 订单ID
    UserId string `json:"userId" comment:"用户ID"`
    RegionId string `json:"regionId" comment:"地区ID"`
    CategoryId string `json:"categoryId" comment:"分类ID"`
    GiftcardId string `json:"giftcardId" comment:"礼品卡id"`
    CardType string `json:"cardType" comment:"卡片类型(Physical,Code,Receipy,Cash Receipt)"`
    GiftCardCode string `json:"giftCardCode" comment:"礼品卡卡号验证码"`
    Balance string `json:"balance" comment:"卡余额"`
    CurrencyCode string `json:"currencyCode" comment:"币种，例如 USD, CNY"`
    DiscountRate string `json:"discountRate" comment:"折扣"`
    Rate string `json:"rate" comment:"汇率"`
    Status int `json:"status" comment:"订单状态:0=待处理,1=已经接单,2=已完成,3=已取消,4=已经驳回"`
    CardExtra string `json:"cardExtra" comment:""`
    CompletedAt time.Time `json:"completedAt" comment:"完成时间"`
    CanceledAt time.Time `json:"canceledAt" comment:"取消时间"`
    Remark string `json:"remark" comment:"备注"`
    common.ControlBy
}

func (s *OrdUserOrdersUpdateReq) Generate(model *models.OrdUserOrders)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId, _ = strconv.Atoi(s.UserId)
    model.RegionId, _ = strconv.Atoi(s.RegionId)
    model.CategoryId, _ = strconv.Atoi(s.CategoryId)
    model.GiftcardId, _ = strconv.Atoi(s.GiftcardId)
    model.CardType = s.CardType
    model.GiftCardCode = s.GiftCardCode
    model.Balance = s.Balance
    model.CurrencyCode = s.CurrencyCode
    model.DiscountRate = s.DiscountRate
    model.Rate = s.Rate
    model.Status = s.Status
    model.CardExtra = s.CardExtra
    model.CompletedAt = s.CompletedAt
    model.CanceledAt = s.CanceledAt
    model.Remark = s.Remark
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdUserOrdersUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdUserOrdersGetReq 功能获取请求参数
type OrdUserOrdersGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdUserOrdersGetReq) GetId() interface{} {
	return s.Id
}

// OrdUserOrdersDeleteReq 功能删除请求参数
type OrdUserOrdersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdUserOrdersDeleteReq) GetId() interface{} {
	return s.Ids
}

// OrdUserOrdersGetByAssignReq 根据接单人和状态查询订单请求参数
type OrdUserOrdersGetByAssignReq struct {
	dto.Pagination `search:"-"`
	OrderNo string `form:"orderNo" search:"type:exact;column:order_no;table:ord_user_orders" comment:"订单号"`
	Status int `form:"status" search:"type:exact;column:status;table:ord_user_orders" comment:"订单状态:5=待处理,1=已经接单,2=已完成,3=已取消,4=已经驳回"`
	BeginTime string `form:"beginTime" search:"type:gte;column:created_at;table:ord_user_orders" comment:"开始时间"`
	EndTime string `form:"endTime" search:"type:lte;column:created_at;table:ord_user_orders" comment:"结束时间"`
	AssignBy int `json:"-" search:"type:exact;column:assign_by;table:ord_user_orders" comment:"接单人ID（从context自动获取）"`
	OrdUserOrdersOrder
}

func (m *OrdUserOrdersGetByAssignReq) GetNeedSearch() interface{} {
	return *m
}

// OrdUserOrdersAcceptReq 管理员接单请求参数
type OrdUserOrdersAcceptReq struct {
	Id int `uri:"id" comment:"订单ID"`
}

func (s *OrdUserOrdersAcceptReq) GetId() interface{} {
	return s.Id
}

// OrdUserOrdersCancelAcceptReq 取消接单请求参数
type OrdUserOrdersCancelAcceptReq struct {
	Id int `uri:"id" comment:"订单ID"`
	AdminRemark string `json:"adminRemark" comment:"管理员备注"`
}

func (s *OrdUserOrdersCancelAcceptReq) GetId() interface{} {
	return s.Id
}
