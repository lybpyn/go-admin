package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardDiscountsGetPageReq struct {
	dto.Pagination     `search:"-"`
    GiftcardId int `form:"giftcardId" search:"type:exact;column:giftcard_id;table:ord_giftcard_discounts" comment:"礼品卡ID"`
  
  //  OrdGiftcardDiscountsOrder
}

type OrdGiftcardDiscountsOrder struct {
    Id int `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard_discounts"`
    GiftcardId int `form:"giftcardIdOrder"  search:"type:order;column:giftcard_id;table:ord_giftcard_discounts"`
    CardType string `form:"cardTypeOrder"  search:"type:order;column:card_type;table:ord_giftcard_discounts"`
    DiscountRate string `form:"discountRateOrder"  search:"type:order;column:discount_rate;table:ord_giftcard_discounts"`
    CreateBy int `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard_discounts"`
    UpdateBy int `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard_discounts"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard_discounts"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard_discounts"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard_discounts"`
    
}

func (m *OrdGiftcardDiscountsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardDiscountsInsertReq struct {
    Id int `json:"-" comment:""` // 
    GiftcardId int `json:"giftcardId" comment:"礼品卡ID -> ord_giftcard.id"`
    CardType string `json:"cardType" comment:"卡类型"`
    DiscountRate string `json:"discountRate" comment:"折扣汇率，例如0.95 表示95折"`
    common.ControlBy
}

func (s *OrdGiftcardDiscountsInsertReq) Generate(model *models.OrdGiftcardDiscounts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.GiftcardId = s.GiftcardId
    model.CardType = s.CardType
    model.DiscountRate = s.DiscountRate
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftcardDiscountsInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftcardDiscountsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    GiftcardId int `json:"giftcardId" comment:"礼品卡ID -> ord_giftcard.id"`
    CardType string `json:"cardType" comment:"卡类型"`
    DiscountRate string `json:"discountRate" comment:"折扣汇率，例如0.95 表示95折"`
    common.ControlBy
}

func (s *OrdGiftcardDiscountsUpdateReq) Generate(model *models.OrdGiftcardDiscounts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.GiftcardId = s.GiftcardId
    model.CardType = s.CardType
    model.DiscountRate = s.DiscountRate
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftcardDiscountsUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardDiscountsGetReq 功能获取请求参数
type OrdGiftcardDiscountsGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdGiftcardDiscountsGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardDiscountsDeleteReq 功能删除请求参数
type OrdGiftcardDiscountsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftcardDiscountsDeleteReq) GetId() interface{} {
	return s.Ids
}

// OrdGiftcardDiscountsBatchUpdateReq 批量修改折扣率请求参数
type OrdGiftcardDiscountsBatchUpdateReq struct {
	Items []OrdGiftcardDiscountsBatchUpdateItem `json:"items" binding:"required,min=1"`
	common.ControlBy
}

type OrdGiftcardDiscountsBatchUpdateItem struct {
	Id           int    `json:"id" binding:"required"`
	DiscountRate string `json:"discountRate" binding:"required"`
}

// OrdGiftcardDiscountsBatchInsertReq 批量新增折扣率请求参数
type OrdGiftcardDiscountsBatchInsertReq struct {
	Items []OrdGiftcardDiscountsBatchInsertItem `json:"items" binding:"required,min=1"`
	common.ControlBy
}

type OrdGiftcardDiscountsBatchInsertItem struct {
	GiftcardId   int    `json:"giftcardId" binding:"required"`
	CardType     string `json:"cardType" binding:"required"`
	DiscountRate string `json:"discountRate" binding:"required"`
}
