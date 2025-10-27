package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftcardRegionGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrdGiftcardRegionOrder
}

type OrdGiftcardRegionOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_giftcard_region"`
    CategoryId string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:ord_giftcard_region"`
    RegionCode string `form:"regionCodeOrder"  search:"type:order;column:region_code;table:ord_giftcard_region"`
    QuoteCurrency string `form:"quoteCurrencyOrder"  search:"type:order;column:quote_currency;table:ord_giftcard_region"`
    QuoteCurrencySymbol string `form:"quoteCurrencySymbolOrder"  search:"type:order;column:quote_currency_symbol;table:ord_giftcard_region"`
    Rate string `form:"rateOrder"  search:"type:order;column:rate;table:ord_giftcard_region"`
    IsMain string `form:"isMainOrder"  search:"type:order;column:is_main;table:ord_giftcard_region"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_giftcard_region"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_giftcard_region"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_giftcard_region"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_giftcard_region"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_giftcard_region"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_giftcard_region"`
    
}

func (m *OrdGiftcardRegionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftcardRegionInsertReq struct {
    Id int `json:"-" comment:"地区ID"` // 地区ID
    CategoryId string `json:"categoryId" comment:"所属分类ID，外键 -> hs_giftcard_category.id"`
    RegionCode string `json:"regionCode" comment:"地区代码，如 US、UK、CN"`
    QuoteCurrency string `json:"quoteCurrency" comment:"货币代码，如 USD, GBP, CNY"`
    QuoteCurrencySymbol string `json:"quoteCurrencySymbol" comment:""`
    Rate string `json:"rate" comment:""`
    IsMain string `json:"isMain" comment:"是否默认区域"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    common.ControlBy
}

func (s *OrdGiftcardRegionInsertReq) Generate(model *models.OrdGiftcardRegion)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CategoryId = s.CategoryId
    model.RegionCode = s.RegionCode
    model.QuoteCurrency = s.QuoteCurrency
    model.QuoteCurrencySymbol = s.QuoteCurrencySymbol
    model.Rate = s.Rate
    model.IsMain = s.IsMain
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftcardRegionInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftcardRegionUpdateReq struct {
    Id int `uri:"id" comment:"地区ID"` // 地区ID
    CategoryId string `json:"categoryId" comment:"所属分类ID，外键 -> hs_giftcard_category.id"`
    RegionCode string `json:"regionCode" comment:"地区代码，如 US、UK、CN"`
    QuoteCurrency string `json:"quoteCurrency" comment:"货币代码，如 USD, GBP, CNY"`
    QuoteCurrencySymbol string `json:"quoteCurrencySymbol" comment:""`
    Rate string `json:"rate" comment:""`
    IsMain string `json:"isMain" comment:"是否默认区域"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    common.ControlBy
}

func (s *OrdGiftcardRegionUpdateReq) Generate(model *models.OrdGiftcardRegion)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CategoryId = s.CategoryId
    model.RegionCode = s.RegionCode
    model.QuoteCurrency = s.QuoteCurrency
    model.QuoteCurrencySymbol = s.QuoteCurrencySymbol
    model.Rate = s.Rate
    model.IsMain = s.IsMain
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftcardRegionUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardRegionGetReq 功能获取请求参数
type OrdGiftcardRegionGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdGiftcardRegionGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftcardRegionDeleteReq 功能删除请求参数
type OrdGiftcardRegionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftcardRegionDeleteReq) GetId() interface{} {
	return s.Ids
}
