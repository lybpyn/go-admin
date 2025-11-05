package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdOrderGiftcardImagesGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrderId string `form:"orderId"  search:"type:exact;column:order_id;table:ord_order_giftcard_images" comment:"关联的订单ID"`
    ImageUrl string `form:"imageUrl"  search:"type:exact;column:image_url;table:ord_order_giftcard_images" comment:"礼品卡图片URL"`
    SortOrder string `form:"sortOrder"  search:"type:exact;column:sort_order;table:ord_order_giftcard_images" comment:"排序顺序"`
    OrdOrderGiftcardImagesOrder
}

type OrdOrderGiftcardImagesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_order_giftcard_images"`
    OrderId string `form:"orderIdOrder"  search:"type:order;column:order_id;table:ord_order_giftcard_images"`
    ImageUrl string `form:"imageUrlOrder"  search:"type:order;column:image_url;table:ord_order_giftcard_images"`
    SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:ord_order_giftcard_images"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_order_giftcard_images"`
    
}

func (m *OrdOrderGiftcardImagesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdOrderGiftcardImagesInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    OrderId string `json:"orderId" comment:"关联的订单ID"`
    ImageUrl string `json:"imageUrl" comment:"礼品卡图片URL"`
    SortOrder string `json:"sortOrder" comment:"排序顺序"`
    common.ControlBy
}

func (s *OrdOrderGiftcardImagesInsertReq) Generate(model *models.OrdOrderGiftcardImages)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.OrderId = s.OrderId
    model.ImageUrl = s.ImageUrl
    model.SortOrder = s.SortOrder
}

func (s *OrdOrderGiftcardImagesInsertReq) GetId() interface{} {
	return s.Id
}

type OrdOrderGiftcardImagesUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    OrderId string `json:"orderId" comment:"关联的订单ID"`
    ImageUrl string `json:"imageUrl" comment:"礼品卡图片URL"`
    SortOrder string `json:"sortOrder" comment:"排序顺序"`
    common.ControlBy
}

func (s *OrdOrderGiftcardImagesUpdateReq) Generate(model *models.OrdOrderGiftcardImages)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.OrderId = s.OrderId
    model.ImageUrl = s.ImageUrl
    model.SortOrder = s.SortOrder
}

func (s *OrdOrderGiftcardImagesUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdOrderGiftcardImagesGetReq 功能获取请求参数
type OrdOrderGiftcardImagesGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdOrderGiftcardImagesGetReq) GetId() interface{} {
	return s.Id
}

// OrdOrderGiftcardImagesDeleteReq 功能删除请求参数
type OrdOrderGiftcardImagesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdOrderGiftcardImagesDeleteReq) GetId() interface{} {
	return s.Ids
}
