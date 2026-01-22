package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// HsConfigWithdrawFeeTiersGetPageReq 分页查询请求
type HsConfigWithdrawFeeTiersGetPageReq struct {
	dto.Pagination `search:"-"`
	RuleId         int `form:"ruleId" search:"type:exact;column:rule_id;table:hs_config_withdraw_fee_tiers" comment:"关联的提现规则ID"`
	HsConfigWithdrawFeeTiersOrder
}

type HsConfigWithdrawFeeTiersOrder struct {
	Id        string `form:"idOrder" search:"type:order;column:id;table:hs_config_withdraw_fee_tiers"`
	RuleId    string `form:"ruleIdOrder" search:"type:order;column:rule_id;table:hs_config_withdraw_fee_tiers"`
	MinAmount string `form:"minAmountOrder" search:"type:order;column:min_amount;table:hs_config_withdraw_fee_tiers"`
	MaxAmount string `form:"maxAmountOrder" search:"type:order;column:max_amount;table:hs_config_withdraw_fee_tiers"`
	FeeAmount string `form:"feeAmountOrder" search:"type:order;column:fee_amount;table:hs_config_withdraw_fee_tiers"`
	SortOrder string `form:"sortOrderOrder" search:"type:order;column:sort_order;table:hs_config_withdraw_fee_tiers"`
	CreatedAt string `form:"createdAtOrder" search:"type:order;column:created_at;table:hs_config_withdraw_fee_tiers"`
	UpdatedAt string `form:"updatedAtOrder" search:"type:order;column:updated_at;table:hs_config_withdraw_fee_tiers"`
}

func (m *HsConfigWithdrawFeeTiersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// HsConfigWithdrawFeeTiersInsertReq 创建请求
type HsConfigWithdrawFeeTiersInsertReq struct {
	Id        int    `json:"-" comment:"主键ID"`
	RuleId    int    `json:"ruleId" comment:"关联的提现规则ID" binding:"required"`
	MinAmount string `json:"minAmount" comment:"区间最小金额（包含）" binding:"required"`
	MaxAmount string `json:"maxAmount" comment:"区间最大金额（包含），0表示无上限" binding:"required"`
	FeeAmount string `json:"feeAmount" comment:"该区间的手续费金额" binding:"required"`
	SortOrder int    `json:"sortOrder" comment:"排序顺序，值越小越优先"`
}

func (s *HsConfigWithdrawFeeTiersInsertReq) Generate(model *models.HsConfigWithdrawFeeTiers) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RuleId = s.RuleId
	model.MinAmount = s.MinAmount
	model.MaxAmount = s.MaxAmount
	model.FeeAmount = s.FeeAmount
	model.SortOrder = s.SortOrder
}

func (s *HsConfigWithdrawFeeTiersInsertReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawFeeTiersUpdateReq 修改请求
type HsConfigWithdrawFeeTiersUpdateReq struct {
	Id        int    `uri:"id" comment:"主键ID"`
	RuleId    int    `json:"ruleId" comment:"关联的提现规则ID"`
	MinAmount string `json:"minAmount" comment:"区间最小金额（包含）"`
	MaxAmount string `json:"maxAmount" comment:"区间最大金额（包含），0表示无上限"`
	FeeAmount string `json:"feeAmount" comment:"该区间的手续费金额"`
	SortOrder int    `json:"sortOrder" comment:"排序顺序，值越小越优先"`
}

func (s *HsConfigWithdrawFeeTiersUpdateReq) Generate(model *models.HsConfigWithdrawFeeTiers) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RuleId = s.RuleId
	model.MinAmount = s.MinAmount
	model.MaxAmount = s.MaxAmount
	model.FeeAmount = s.FeeAmount
	model.SortOrder = s.SortOrder
}

func (s *HsConfigWithdrawFeeTiersUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawFeeTiersGetReq 获取单条记录请求
type HsConfigWithdrawFeeTiersGetReq struct {
	Id int `uri:"id"`
}

func (s *HsConfigWithdrawFeeTiersGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawFeeTiersDeleteReq 删除请求
type HsConfigWithdrawFeeTiersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigWithdrawFeeTiersDeleteReq) GetId() interface{} {
	return s.Ids
}

// HsConfigWithdrawFeeTiersByRuleIdReq 根据规则ID查询请求
type HsConfigWithdrawFeeTiersByRuleIdReq struct {
	RuleId int `uri:"ruleId" binding:"required"`
}

func (s *HsConfigWithdrawFeeTiersByRuleIdReq) GetRuleId() int {
	return s.RuleId
}

// HsConfigWithdrawFeeTiersBatchReq 批量创建/更新请求
type HsConfigWithdrawFeeTiersBatchReq struct {
	RuleId int                                 `json:"ruleId" comment:"关联的提现规则ID" binding:"required"`
	Tiers  []HsConfigWithdrawFeeTiersBatchItem `json:"tiers" comment:"阶梯配置列表" binding:"required"`
}

// HsConfigWithdrawFeeTiersBatchItem 批量操作的单条记录
type HsConfigWithdrawFeeTiersBatchItem struct {
	Id        int    `json:"id" comment:"主键ID，0表示新增"`
	MinAmount string `json:"minAmount" comment:"区间最小金额（包含）" binding:"required"`
	MaxAmount string `json:"maxAmount" comment:"区间最大金额（包含），0表示无上限" binding:"required"`
	FeeAmount string `json:"feeAmount" comment:"该区间的手续费金额" binding:"required"`
	SortOrder int    `json:"sortOrder" comment:"排序顺序，值越小越优先"`
}
