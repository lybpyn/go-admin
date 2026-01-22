package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigWithdrawRulesGetPageReq struct {
	dto.Pagination     `search:"-"`
    CurrencyCode string `form:"currencyCode"  search:"type:exact;column:currency_code;table:hs_config_withdraw_rules" comment:"币种代码，如 USD、CNY、USDT、BTC"`
    CurrencyType string `form:"currencyType"  search:"type:exact;column:currency_type;table:hs_config_withdraw_rules" comment:"币种类型：fiat=法币，crypto=虚拟币"`
    ChainType string `form:"chainType"  search:"type:exact;column:chain_type;table:hs_config_withdraw_rules" comment:"链类型（仅虚拟币适用），如 ERC20/TRC20/BEP20"`
    HsConfigWithdrawRulesOrder
}

type HsConfigWithdrawRulesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_withdraw_rules"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:hs_config_withdraw_rules"`
    CurrencyType string `form:"currencyTypeOrder"  search:"type:order;column:currency_type;table:hs_config_withdraw_rules"`
    ChainType string `form:"chainTypeOrder"  search:"type:order;column:chain_type;table:hs_config_withdraw_rules"`
    SingleMin string `form:"singleMinOrder"  search:"type:order;column:single_min;table:hs_config_withdraw_rules"`
    SingleMax string `form:"singleMaxOrder"  search:"type:order;column:single_max;table:hs_config_withdraw_rules"`
    DailyLimitAmount string `form:"dailyLimitAmountOrder"  search:"type:order;column:daily_limit_amount;table:hs_config_withdraw_rules"`
    DailyLimitCount string `form:"dailyLimitCountOrder"  search:"type:order;column:daily_limit_count;table:hs_config_withdraw_rules"`
    FeeType string `form:"feeTypeOrder"  search:"type:order;column:fee_type;table:hs_config_withdraw_rules"`
    FeeFixed string `form:"feeFixedOrder"  search:"type:order;column:fee_fixed;table:hs_config_withdraw_rules"`
    FeeRate string `form:"feeRateOrder"  search:"type:order;column:fee_rate;table:hs_config_withdraw_rules"`
    MinFee string `form:"minFeeOrder"  search:"type:order;column:min_fee;table:hs_config_withdraw_rules"`
    MaxFee string `form:"maxFeeOrder"  search:"type:order;column:max_fee;table:hs_config_withdraw_rules"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_config_withdraw_rules"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_withdraw_rules"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_withdraw_rules"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_withdraw_rules"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_withdraw_rules"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_withdraw_rules"`
    
}

func (m *HsConfigWithdrawRulesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// FeeTierItem 阶梯手续费配置项
type FeeTierItem struct {
    MinAmount string `json:"minAmount" comment:"区间最小金额（包含）"`
    MaxAmount string `json:"maxAmount" comment:"区间最大金额（包含），0表示无上限"`
    FeeAmount string `json:"feeAmount" comment:"该区间的手续费金额"`
    SortOrder int    `json:"sortOrder" comment:"排序顺序，值越小越优先"`
}

type HsConfigWithdrawRulesInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD、CNY、USDT、BTC"`
    CurrencyType string `json:"currencyType" comment:"币种类型：fiat=法币，crypto=虚拟币"`
    ChainType string `json:"chainType" comment:"链类型（仅虚拟币适用），如 ERC20/TRC20/BEP20"`
    SingleMin string `json:"singleMin" comment:"单笔最小提现数量/金额"`
    SingleMax string `json:"singleMax" comment:"单笔最大提现数量/金额"`
    DailyLimitAmount string `json:"dailyLimitAmount" comment:"每日累计提现上限"`
    DailyLimitCount string `json:"dailyLimitCount" comment:"每日提现次数上限"`
    FeeType string `json:"feeType" comment:"手续费类型：fixed=固定，rate=按比例，mixed=固定+比例，tiered=阶梯收费"`
    FeeFixed string `json:"feeFixed" comment:"固定手续费数量/金额"`
    FeeRate string `json:"feeRate" comment:"手续费率（例如 0.015 表示 1.5%）"`
    MinFee string `json:"minFee" comment:"最小手续费"`
    MaxFee string `json:"maxFee" comment:"最高手续费"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    FeeTiers []FeeTierItem `json:"feeTiers" comment:"阶梯手续费配置（仅当feeType=tiered时有效）"`
    common.ControlBy
}

func (s *HsConfigWithdrawRulesInsertReq) Generate(model *models.HsConfigWithdrawRules)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.CurrencyType = s.CurrencyType
    model.ChainType = s.ChainType
    model.SingleMin = s.SingleMin
    model.SingleMax = s.SingleMax
    model.DailyLimitAmount = s.DailyLimitAmount
    model.DailyLimitCount = s.DailyLimitCount
    model.FeeType = s.FeeType
    model.FeeFixed = s.FeeFixed
    model.FeeRate = s.FeeRate
    model.MinFee = s.MinFee
    model.MaxFee = s.MaxFee
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigWithdrawRulesInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigWithdrawRulesUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD、CNY、USDT、BTC"`
    CurrencyType string `json:"currencyType" comment:"币种类型：fiat=法币，crypto=虚拟币"`
    ChainType string `json:"chainType" comment:"链类型（仅虚拟币适用），如 ERC20/TRC20/BEP20"`
    SingleMin string `json:"singleMin" comment:"单笔最小提现数量/金额"`
    SingleMax string `json:"singleMax" comment:"单笔最大提现数量/金额"`
    DailyLimitAmount string `json:"dailyLimitAmount" comment:"每日累计提现上限"`
    DailyLimitCount string `json:"dailyLimitCount" comment:"每日提现次数上限"`
    FeeType string `json:"feeType" comment:"手续费类型：fixed=固定，rate=按比例，mixed=固定+比例，tiered=阶梯收费"`
    FeeFixed string `json:"feeFixed" comment:"固定手续费数量/金额"`
    FeeRate string `json:"feeRate" comment:"手续费率（例如 0.015 表示 1.5%）"`
    MinFee string `json:"minFee" comment:"最小手续费"`
    MaxFee string `json:"maxFee" comment:"最高手续费"`
    IsActive string `json:"isActive" comment:"是否启用：1=启用，0=禁用"`
    FeeTiers []FeeTierItem `json:"feeTiers" comment:"阶梯手续费配置（仅当feeType=tiered时有效）"`
    common.ControlBy
}

func (s *HsConfigWithdrawRulesUpdateReq) Generate(model *models.HsConfigWithdrawRules)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.CurrencyCode = s.CurrencyCode
    model.CurrencyType = s.CurrencyType
    model.ChainType = s.ChainType
    model.SingleMin = s.SingleMin
    model.SingleMax = s.SingleMax
    model.DailyLimitAmount = s.DailyLimitAmount
    model.DailyLimitCount = s.DailyLimitCount
    model.FeeType = s.FeeType
    model.FeeFixed = s.FeeFixed
    model.FeeRate = s.FeeRate
    model.MinFee = s.MinFee
    model.MaxFee = s.MaxFee
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigWithdrawRulesUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawRulesGetReq 功能获取请求参数
type HsConfigWithdrawRulesGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigWithdrawRulesGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigWithdrawRulesDeleteReq 功能删除请求参数
type HsConfigWithdrawRulesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigWithdrawRulesDeleteReq) GetId() interface{} {
	return s.Ids
}
