package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardDiscountsGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrdGiftcardDiscountsOrder
}

type OrdGiftcardDiscountsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard_discounts"`
    GiftcardId string `form:"giftcardIdOrder"  search:"type:order;column:giftcard_id;table:ord_giftcard_discounts"`
    CardType string `form:"cardTypeOrder"  search:"type:order;column:card_type;table:ord_giftcard_discounts"`
    DiscountRate string `form:"discountRateOrder"  search:"type:order;column:discount_rate;table:ord_giftcard_discounts"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard_discounts"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard_discounts"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard_discounts"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard_discounts"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard_discounts"`
    
}

func (m *OrdGiftcardDiscountsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardDiscountsInsertReq struct {
    Id int `json:"-" comment:""` // 
    GiftcardId string `json:"giftcardId" comment:"礼品卡ID -> ord_giftcard.id"`
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
    GiftcardId string `json:"giftcardId" comment:"礼品卡ID -> ord_giftcard.id"`
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
