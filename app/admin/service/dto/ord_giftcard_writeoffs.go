package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardWriteoffsGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrdGiftcardWriteoffsOrder
}

type OrdGiftcardWriteoffsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard_writeoffs"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:ord_giftcard_writeoffs"`
    OrderId string `form:"orderIdOrder"  search:"type:order;column:order_id;table:ord_giftcard_writeoffs"`
    GiftCardId string `form:"giftCardIdOrder"  search:"type:order;column:gift_card_id;table:ord_giftcard_writeoffs"`
    GiftCardCode string `form:"giftCardCodeOrder"  search:"type:order;column:gift_card_code;table:ord_giftcard_writeoffs"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_giftcard_writeoffs"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:ord_giftcard_writeoffs"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard_writeoffs"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard_writeoffs"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard_writeoffs"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard_writeoffs"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard_writeoffs"`
    
}

func (m *OrdGiftcardWriteoffsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardWriteoffsInsertReq struct {
    Id int `json:"-" comment:"主键ID，自增"` // 主键ID，自增
    UserId string `json:"userId" comment:"用户ID，表示提交/使用礼品卡的用户"`
    OrderId string `json:"orderId" comment:"订单ID，关联 ord_user_orders.id，用于核销对应的订单"`
    GiftCardId string `json:"giftCardId" comment:"礼品卡ID，关联礼品卡主表（若有）"`
    GiftCardCode string `json:"giftCardCode" comment:"礼品卡卡号/兑换码，保证唯一性"`
    Status string `json:"status" comment:"核销状态：0=待核销，1=已核销，2=失败"`
    Remark string `json:"remark" comment:"备注信息，例如失败原因、核销说明"`
    common.ControlBy
}

func (s *OrdGiftcardWriteoffsInsertReq) Generate(model *models.OrdGiftcardWriteoffs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.OrderId = s.OrderId
    model.GiftCardId = s.GiftCardId
    model.GiftCardCode = s.GiftCardCode
    model.Status = s.Status
    model.Remark = s.Remark
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftcardWriteoffsInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftcardWriteoffsUpdateReq struct {
    Id int `uri:"id" comment:"主键ID，自增"` // 主键ID，自增
    UserId string `json:"userId" comment:"用户ID，表示提交/使用礼品卡的用户"`
    OrderId string `json:"orderId" comment:"订单ID，关联 ord_user_orders.id，用于核销对应的订单"`
    GiftCardId string `json:"giftCardId" comment:"礼品卡ID，关联礼品卡主表（若有）"`
    GiftCardCode string `json:"giftCardCode" comment:"礼品卡卡号/兑换码，保证唯一性"`
    Status string `json:"status" comment:"核销状态：0=待核销，1=已核销，2=失败"`
    Remark string `json:"remark" comment:"备注信息，例如失败原因、核销说明"`
    common.ControlBy
}

func (s *OrdGiftcardWriteoffsUpdateReq) Generate(model *models.OrdGiftcardWriteoffs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.OrderId = s.OrderId
    model.GiftCardId = s.GiftCardId
    model.GiftCardCode = s.GiftCardCode
    model.Status = s.Status
    model.Remark = s.Remark
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftcardWriteoffsUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardWriteoffsGetReq 功能获取请求参数
type OrdGiftcardWriteoffsGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdGiftcardWriteoffsGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardWriteoffsDeleteReq 功能删除请求参数
type OrdGiftcardWriteoffsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftcardWriteoffsDeleteReq) GetId() interface{} {
	return s.Ids
}
