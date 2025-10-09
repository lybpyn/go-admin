package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsReceivingAccountsGetPageReq struct {
	dto.Pagination     `search:"-"`
    HsReceivingAccountsOrder
}

type HsReceivingAccountsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_receiving_accounts"`
    MerchantId string `form:"merchantIdOrder"  search:"type:order;column:merchant_id;table:hs_receiving_accounts"`
    BankId string `form:"bankIdOrder"  search:"type:order;column:bank_id;table:hs_receiving_accounts"`
    AccountName string `form:"accountNameOrder"  search:"type:order;column:account_name;table:hs_receiving_accounts"`
    AccountNumber string `form:"accountNumberOrder"  search:"type:order;column:account_number;table:hs_receiving_accounts"`
    Currency string `form:"currencyOrder"  search:"type:order;column:currency;table:hs_receiving_accounts"`
    MinAmount string `form:"minAmountOrder"  search:"type:order;column:min_amount;table:hs_receiving_accounts"`
    MaxAmount string `form:"maxAmountOrder"  search:"type:order;column:max_amount;table:hs_receiving_accounts"`
    DailyLimit string `form:"dailyLimitOrder"  search:"type:order;column:daily_limit;table:hs_receiving_accounts"`
    IsDefault string `form:"isDefaultOrder"  search:"type:order;column:is_default;table:hs_receiving_accounts"`
    AutoTransferEnabled string `form:"autoTransferEnabledOrder"  search:"type:order;column:auto_transfer_enabled;table:hs_receiving_accounts"`
    AutoTransferThreshold string `form:"autoTransferThresholdOrder"  search:"type:order;column:auto_transfer_threshold;table:hs_receiving_accounts"`
    RiskThresholdAmount string `form:"riskThresholdAmountOrder"  search:"type:order;column:risk_threshold_amount;table:hs_receiving_accounts"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_receiving_accounts"`
    InterfaceType string `form:"interfaceTypeOrder"  search:"type:order;column:interface_type;table:hs_receiving_accounts"`
    InterfaceConfig string `form:"interfaceConfigOrder"  search:"type:order;column:interface_config;table:hs_receiving_accounts"`
    AllowedCountries string `form:"allowedCountriesOrder"  search:"type:order;column:allowed_countries;table:hs_receiving_accounts"`
    Note string `form:"noteOrder"  search:"type:order;column:note;table:hs_receiving_accounts"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_receiving_accounts"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_receiving_accounts"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_receiving_accounts"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_receiving_accounts"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_receiving_accounts"`
    
}

func (m *HsReceivingAccountsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsReceivingAccountsInsertReq struct {
    Id int `json:"-" comment:""` // 
    MerchantId string `json:"merchantId" comment:"关联 merchants.id"`
    BankId string `json:"bankId" comment:"关联 banks.id (如果是银行账户)"`
    AccountName string `json:"accountName" comment:"账户名/收款人名称"`
    AccountNumber string `json:"accountNumber" comment:"账号（可能为卡号/IBAN/虚拟账号）"`
    Currency string `json:"currency" comment:"货币代码 (USD/CNY/...)"`
    MinAmount string `json:"minAmount" comment:"单次最小收款金额"`
    MaxAmount string `json:"maxAmount" comment:"单次最大收款金额 (NULL 表示无限制)"`
    DailyLimit string `json:"dailyLimit" comment:"每日上限 (平台/风控)"`
    IsDefault string `json:"isDefault" comment:"是否默认收款账号: 0否,1是"`
    AutoTransferEnabled string `json:"autoTransferEnabled" comment:"是否启用自动转出/自动结算"`
    AutoTransferThreshold string `json:"autoTransferThreshold" comment:"自动转出阈值 (达到该金额触发)"`
    RiskThresholdAmount string `json:"riskThresholdAmount" comment:"风控阈值(超过则人工审核) —— 可与汇率表联动计算"`
    Status string `json:"status" comment:"状态:0=禁用,1=启用,2=待审核"`
    InterfaceType string `json:"interfaceType" comment:"接口/通道类型，如 bank_transfer, third_party_api 等"`
    InterfaceConfig string `json:"interfaceConfig" comment:"接口配置信息（JSON）: 如 api_key, secret, callback_url, 接口码等"`
    AllowedCountries string `json:"allowedCountries" comment:"允许接收国家数组，例如 ["CN","US"]"`
    Note string `json:"note" comment:""`
    common.ControlBy
}

func (s *HsReceivingAccountsInsertReq) Generate(model *models.HsReceivingAccounts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.MerchantId = s.MerchantId
    model.BankId = s.BankId
    model.AccountName = s.AccountName
    model.AccountNumber = s.AccountNumber
    model.Currency = s.Currency
    model.MinAmount = s.MinAmount
    model.MaxAmount = s.MaxAmount
    model.DailyLimit = s.DailyLimit
    model.IsDefault = s.IsDefault
    model.AutoTransferEnabled = s.AutoTransferEnabled
    model.AutoTransferThreshold = s.AutoTransferThreshold
    model.RiskThresholdAmount = s.RiskThresholdAmount
    model.Status = s.Status
    model.InterfaceType = s.InterfaceType
    model.InterfaceConfig = s.InterfaceConfig
    model.AllowedCountries = s.AllowedCountries
    model.Note = s.Note
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsReceivingAccountsInsertReq) GetId() interface{} {
	return s.Id
}

type HsReceivingAccountsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    MerchantId string `json:"merchantId" comment:"关联 merchants.id"`
    BankId string `json:"bankId" comment:"关联 banks.id (如果是银行账户)"`
    AccountName string `json:"accountName" comment:"账户名/收款人名称"`
    AccountNumber string `json:"accountNumber" comment:"账号（可能为卡号/IBAN/虚拟账号）"`
    Currency string `json:"currency" comment:"货币代码 (USD/CNY/...)"`
    MinAmount string `json:"minAmount" comment:"单次最小收款金额"`
    MaxAmount string `json:"maxAmount" comment:"单次最大收款金额 (NULL 表示无限制)"`
    DailyLimit string `json:"dailyLimit" comment:"每日上限 (平台/风控)"`
    IsDefault string `json:"isDefault" comment:"是否默认收款账号: 0否,1是"`
    AutoTransferEnabled string `json:"autoTransferEnabled" comment:"是否启用自动转出/自动结算"`
    AutoTransferThreshold string `json:"autoTransferThreshold" comment:"自动转出阈值 (达到该金额触发)"`
    RiskThresholdAmount string `json:"riskThresholdAmount" comment:"风控阈值(超过则人工审核) —— 可与汇率表联动计算"`
    Status string `json:"status" comment:"状态:0=禁用,1=启用,2=待审核"`
    InterfaceType string `json:"interfaceType" comment:"接口/通道类型，如 bank_transfer, third_party_api 等"`
    InterfaceConfig string `json:"interfaceConfig" comment:"接口配置信息（JSON）: 如 api_key, secret, callback_url, 接口码等"`
    AllowedCountries string `json:"allowedCountries" comment:"允许接收国家数组，例如 ["CN","US"]"`
    Note string `json:"note" comment:""`
    common.ControlBy
}

func (s *HsReceivingAccountsUpdateReq) Generate(model *models.HsReceivingAccounts)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.MerchantId = s.MerchantId
    model.BankId = s.BankId
    model.AccountName = s.AccountName
    model.AccountNumber = s.AccountNumber
    model.Currency = s.Currency
    model.MinAmount = s.MinAmount
    model.MaxAmount = s.MaxAmount
    model.DailyLimit = s.DailyLimit
    model.IsDefault = s.IsDefault
    model.AutoTransferEnabled = s.AutoTransferEnabled
    model.AutoTransferThreshold = s.AutoTransferThreshold
    model.RiskThresholdAmount = s.RiskThresholdAmount
    model.Status = s.Status
    model.InterfaceType = s.InterfaceType
    model.InterfaceConfig = s.InterfaceConfig
    model.AllowedCountries = s.AllowedCountries
    model.Note = s.Note
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsReceivingAccountsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsReceivingAccountsGetReq 功能获取请求参数
type HsReceivingAccountsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsReceivingAccountsGetReq) GetId() interface{} {
	return s.Id
}

// HsReceivingAccountsDeleteReq 功能删除请求参数
type HsReceivingAccountsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsReceivingAccountsDeleteReq) GetId() interface{} {
	return s.Ids
}
