package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigWithdrawFeeGetPageReq struct {
	dto.Pagination     `search:"-"`
    CurrencyCode string `form:"currencyCode"  search:"type:exact;column:currency_code;table:hs_config_withdraw_fee" comment:"币种代码，如 USD/CNY/USDT"`
    MinAmount string `form:"minAmount"  search:"type:exact;column:min_amount;table:hs_config_withdraw_fee" comment:"最小提现金额"`
    MaxAmount string `form:"maxAmount"  search:"type:exact;column:max_amount;table:hs_config_withdraw_fee" comment:"最大提现金额"`
    FeeRate string `form:"feeRate"  search:"type:exact;column:fee_rate;table:hs_config_withdraw_fee" comment:"手续费率（例如 0.015 表示 1.5%）"`
    IsActive string `form:"isActive"  search:"type:exact;column:is_active;table:hs_config_withdraw_fee" comment:"是否启用：1=启用，0=禁用"`
    HsConfigWithdrawFeeOrder
}

type HsConfigWithdrawFeeOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_withdraw_fee"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:hs_config_withdraw_fee"`
    MinAmount string `form:"minAmountOrder"  search:"type:order;column:min_amount;table:hs_config_withdraw_fee"`
    MaxAmount string `form:"maxAmountOrder"  search:"type:order;column:max_amount;table:hs_config_withdraw_fee"`
    FeeRate string `form:"feeRateOrder"  search:"type:order;column:fee_rate;table:hs_config_withdraw_fee"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_config_withdraw_fee"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_withdraw_fee"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_withdraw_fee"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_withdraw_fee"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_withdraw_fee"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_withdraw_fee"`
    
}

func (m *HsConfigWithdrawFeeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigWithdrawFeeInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    MinAmount string `json:"minAmount" comment:"最小提现金额"`
    MaxAmount string `json:"maxAmount" comment:"最大提现金额"`
    FeeRate string `json:"feeRate" comment:"手续费率（例如 0.015 表示 1.5%）"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigWithdrawFeeInsertReq) Generate(model *models.HsConfigWithdrawFee)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.MinAmount = s.MinAmount
    model.MaxAmount = s.MaxAmount
    model.FeeRate = s.FeeRate
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigWithdrawFeeInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigWithdrawFeeUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    MinAmount string `json:"minAmount" comment:"最小提现金额"`
    MaxAmount string `json:"maxAmount" comment:"最大提现金额"`
    FeeRate string `json:"feeRate" comment:"手续费率（例如 0.015 表示 1.5%）"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigWithdrawFeeUpdateReq) Generate(model *models.HsConfigWithdrawFee)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.MinAmount = s.MinAmount
    model.MaxAmount = s.MaxAmount
    model.FeeRate = s.FeeRate
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigWithdrawFeeUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawFeeGetReq 功能获取请求参数
type HsConfigWithdrawFeeGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigWithdrawFeeGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawFeeDeleteReq 功能删除请求参数
type HsConfigWithdrawFeeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigWithdrawFeeDeleteReq) GetId() interface{} {
	return s.Ids
}
