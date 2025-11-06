package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigFrozenLimitGetPageReq struct {
	dto.Pagination     `search:"-"`
    CurrencyCode string `form:"currencyCode"  search:"type:exact;column:currency_code;table:hs_config_frozen_limit" comment:"币种代码，如 USD/CNY/USDT"`
    FrozenLimitAmount string `form:"frozenLimitAmount"  search:"type:exact;column:frozen_limit_amount;table:hs_config_frozen_limit" comment:"可冻结金额上限 / 提现限制金额"`
    IsActive string `form:"isActive"  search:"type:exact;column:is_active;table:hs_config_frozen_limit" comment:"是否启用：1=启用，0=禁用"`
    HsConfigFrozenLimitOrder
}

type HsConfigFrozenLimitOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_frozen_limit"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:hs_config_frozen_limit"`
    FrozenLimitAmount string `form:"frozenLimitAmountOrder"  search:"type:order;column:frozen_limit_amount;table:hs_config_frozen_limit"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_config_frozen_limit"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_frozen_limit"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_frozen_limit"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_frozen_limit"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_frozen_limit"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_frozen_limit"`
    
}

func (m *HsConfigFrozenLimitGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigFrozenLimitInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    FrozenLimitAmount string `json:"frozenLimitAmount" comment:"可冻结金额上限 / 提现限制金额"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigFrozenLimitInsertReq) Generate(model *models.HsConfigFrozenLimit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.FrozenLimitAmount = s.FrozenLimitAmount
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigFrozenLimitInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigFrozenLimitUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    FrozenLimitAmount string `json:"frozenLimitAmount" comment:"可冻结金额上限 / 提现限制金额"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigFrozenLimitUpdateReq) Generate(model *models.HsConfigFrozenLimit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.FrozenLimitAmount = s.FrozenLimitAmount
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigFrozenLimitUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigFrozenLimitGetReq 功能获取请求参数
type HsConfigFrozenLimitGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigFrozenLimitGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigFrozenLimitDeleteReq 功能删除请求参数
type HsConfigFrozenLimitDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigFrozenLimitDeleteReq) GetId() interface{} {
	return s.Ids
}
