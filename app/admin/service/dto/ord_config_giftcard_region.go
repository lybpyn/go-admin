package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdConfigGiftcardRegionGetPageReq struct {
	dto.Pagination     `search:"-"`
    RegionName string `form:"regionName"  search:"type:exact;column:region_name;table:ord_config_giftcard_region" comment:"区域名称"`
    OrdConfigGiftcardRegionOrder
}

type OrdConfigGiftcardRegionOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_config_giftcard_region"`
    RegionName string `form:"regionNameOrder"  search:"type:order;column:region_name;table:ord_config_giftcard_region"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_config_giftcard_region"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_config_giftcard_region"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_config_giftcard_region"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_config_giftcard_region"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_config_giftcard_region"`
    
}

func (m *OrdConfigGiftcardRegionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdConfigGiftcardRegionInsertReq struct {
    Id int `json:"-" comment:""` //
    RegionName     string `json:"regionName" comment:"区域名称"`
    RateType       string `json:"rateType" comment:"汇率类型" default:"standard"`
    CurrencySymbol string `json:"currencySymbol" comment:"货币符号"`
    common.ControlBy
}

func (s *OrdConfigGiftcardRegionInsertReq) Generate(model *models.OrdConfigGiftcardRegion)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.RegionName = s.RegionName
    // 设置 RateType 默认值
    if s.RateType == "" {
        model.RateType = "standard"
    } else {
        model.RateType = s.RateType
    }
    model.CurrencySymbol = s.CurrencySymbol
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdConfigGiftcardRegionInsertReq) GetId() interface{} {
	return s.Id
}

type OrdConfigGiftcardRegionUpdateReq struct {
    Id int `uri:"id" comment:""` //
    RegionName     string `json:"regionName" comment:"区域名称"`
    RateType       string `json:"rateType" comment:"汇率类型"`
    CurrencySymbol string `json:"currencySymbol" comment:"货币符号"`
    common.ControlBy
}

func (s *OrdConfigGiftcardRegionUpdateReq) Generate(model *models.OrdConfigGiftcardRegion)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.RegionName = s.RegionName
    if s.RateType != "" {
        model.RateType = s.RateType
    }
    if s.CurrencySymbol != "" {
        model.CurrencySymbol = s.CurrencySymbol
    }
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdConfigGiftcardRegionUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdConfigGiftcardRegionGetReq 功能获取请求参数
type OrdConfigGiftcardRegionGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdConfigGiftcardRegionGetReq) GetId() interface{} {
	return s.Id
}

// OrdConfigGiftcardRegionDeleteReq 功能删除请求参数
type OrdConfigGiftcardRegionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdConfigGiftcardRegionDeleteReq) GetId() interface{} {
	return s.Ids
}
