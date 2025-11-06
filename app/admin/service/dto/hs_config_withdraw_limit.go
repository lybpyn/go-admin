package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigWithdrawLimitGetPageReq struct {
	dto.Pagination     `search:"-"`
    CurrencyCode string `form:"currencyCode"  search:"type:exact;column:currency_code;table:hs_config_withdraw_limit" comment:"币种代码，如 USD/CNY/USDT"`
    SingleLimit string `form:"singleLimit"  search:"type:exact;column:single_limit;table:hs_config_withdraw_limit" comment:"单笔提现限额"`
    DailyLimitAmount string `form:"dailyLimitAmount"  search:"type:exact;column:daily_limit_amount;table:hs_config_withdraw_limit" comment:"每日累计提现金额上限"`
    DailyLimitCount string `form:"dailyLimitCount"  search:"type:exact;column:daily_limit_count;table:hs_config_withdraw_limit" comment:"每日提现次数上限"`
    IsActive string `form:"isActive"  search:"type:exact;column:is_active;table:hs_config_withdraw_limit" comment:"是否启用：1=启用，0=禁用"`
    HsConfigWithdrawLimitOrder
}

type HsConfigWithdrawLimitOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_withdraw_limit"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:hs_config_withdraw_limit"`
    SingleLimit string `form:"singleLimitOrder"  search:"type:order;column:single_limit;table:hs_config_withdraw_limit"`
    DailyLimitAmount string `form:"dailyLimitAmountOrder"  search:"type:order;column:daily_limit_amount;table:hs_config_withdraw_limit"`
    DailyLimitCount string `form:"dailyLimitCountOrder"  search:"type:order;column:daily_limit_count;table:hs_config_withdraw_limit"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_config_withdraw_limit"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_withdraw_limit"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_withdraw_limit"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_withdraw_limit"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_withdraw_limit"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_withdraw_limit"`
    
}

func (m *HsConfigWithdrawLimitGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigWithdrawLimitInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    SingleLimit string `json:"singleLimit" comment:"单笔提现限额"`
    DailyLimitAmount string `json:"dailyLimitAmount" comment:"每日累计提现金额上限"`
    DailyLimitCount string `json:"dailyLimitCount" comment:"每日提现次数上限"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigWithdrawLimitInsertReq) Generate(model *models.HsConfigWithdrawLimit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.SingleLimit = s.SingleLimit
    model.DailyLimitAmount = s.DailyLimitAmount
    model.DailyLimitCount = s.DailyLimitCount
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigWithdrawLimitInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigWithdrawLimitUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    SingleLimit string `json:"singleLimit" comment:"单笔提现限额"`
    DailyLimitAmount string `json:"dailyLimitAmount" comment:"每日累计提现金额上限"`
    DailyLimitCount string `json:"dailyLimitCount" comment:"每日提现次数上限"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    common.ControlBy
}

func (s *HsConfigWithdrawLimitUpdateReq) Generate(model *models.HsConfigWithdrawLimit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.SingleLimit = s.SingleLimit
    model.DailyLimitAmount = s.DailyLimitAmount
    model.DailyLimitCount = s.DailyLimitCount
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigWithdrawLimitUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawLimitGetReq 功能获取请求参数
type HsConfigWithdrawLimitGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigWithdrawLimitGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawLimitDeleteReq 功能删除请求参数
type HsConfigWithdrawLimitDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigWithdrawLimitDeleteReq) GetId() interface{} {
	return s.Ids
}
