package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardCategoryGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrdGiftcardCategoryOrder
}

type OrdGiftcardCategoryOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard_category"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:ord_giftcard_category"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_giftcard_category"`
    DiscountRate string `form:"discountRateOrder"  search:"type:order;column:discount_rate;table:ord_giftcard_category"`
    SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:ord_giftcard_category"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard_category"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard_category"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard_category"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard_category"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard_category"`
    
}

func (m *OrdGiftcardCategoryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardCategoryInsertReq struct {
    Id int `json:"-" comment:"分类ID"` // 分类ID
    Name string `json:"name" comment:"分类名称，如 Steam、eBay"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    DiscountRate string `json:"discountRate" comment:"汇率折扣展示用"`
    SortOrder string `json:"sortOrder" comment:""`
    common.ControlBy
}

func (s *OrdGiftcardCategoryInsertReq) Generate(model *models.OrdGiftcardCategory)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Status = s.Status
    model.DiscountRate = s.DiscountRate
    model.SortOrder = s.SortOrder
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftcardCategoryInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftcardCategoryUpdateReq struct {
    Id int `uri:"id" comment:"分类ID"` // 分类ID
    Name string `json:"name" comment:"分类名称，如 Steam、eBay"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    DiscountRate string `json:"discountRate" comment:"汇率折扣展示用"`
    SortOrder string `json:"sortOrder" comment:""`
    common.ControlBy
}

func (s *OrdGiftcardCategoryUpdateReq) Generate(model *models.OrdGiftcardCategory)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Status = s.Status
    model.DiscountRate = s.DiscountRate
    model.SortOrder = s.SortOrder
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftcardCategoryUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardCategoryGetReq 功能获取请求参数
type OrdGiftcardCategoryGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdGiftcardCategoryGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardCategoryDeleteReq 功能删除请求参数
type OrdGiftcardCategoryDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftcardCategoryDeleteReq) GetId() interface{} {
	return s.Ids
}
