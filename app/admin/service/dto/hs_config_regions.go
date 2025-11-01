package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigRegionsGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:hs_config_regions" comment:"地区名称"`
    HsConfigRegionsOrder
}

type HsConfigRegionsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_regions"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:hs_config_regions"`
    Code string `form:"codeOrder"  search:"type:order;column:code;table:hs_config_regions"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_config_regions"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_regions"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_regions"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_regions"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_regions"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_regions"`
    
}

func (m *HsConfigRegionsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigRegionsInsertReq struct {
    Id int `json:"-" comment:"地区ID"` // 地区ID
    Name string `json:"name" comment:"地区名称"`
    Code string `json:"code" comment:"地区代码，如 CN、US、JP 等"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigRegionsInsertReq) Generate(model *models.HsConfigRegions)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Code = s.Code
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigRegionsInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigRegionsUpdateReq struct {
    Id int `uri:"id" comment:"地区ID"` // 地区ID
    Name string `json:"name" comment:"地区名称"`
    Code string `json:"code" comment:"地区代码，如 CN、US、JP 等"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigRegionsUpdateReq) Generate(model *models.HsConfigRegions)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Code = s.Code
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigRegionsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigRegionsGetReq 功能获取请求参数
type HsConfigRegionsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigRegionsGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigRegionsDeleteReq 功能删除请求参数
type HsConfigRegionsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigRegionsDeleteReq) GetId() interface{} {
	return s.Ids
}
