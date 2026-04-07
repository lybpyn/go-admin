package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdConfigCurrencyRatesGetPageReq struct {
	dto.Pagination     `search:"-"`
    BaseCurrencyCode string `form:"baseCurrencyCode"  search:"type:exact;column:base_currency_code;table:ord_config_currency_rates" comment:"基准货币代码 (ISO 4217)"`
    QuoteCurrencyCode string `form:"quoteCurrencyCode"  search:"type:exact;column:quote_currency_code;table:ord_config_currency_rates" comment:"报价货币代码 (ISO 4217)"`
    Rate string `form:"rate"  search:"type:exact;column:rate;table:ord_config_currency_rates" comment:"汇率: 1 base_currency = rate quote_currency"`
    Source string `form:"source"  search:"type:exact;column:source;table:ord_config_currency_rates" comment:"汇率来源,如 manual, api, coingecko"`
    Status string `form:"status"  search:"type:exact;column:status;table:ord_config_currency_rates" comment:"状态: 1=启用, 0=禁用"`
    OrdConfigCurrencyRatesOrder
}

type OrdConfigCurrencyRatesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_config_currency_rates"`
    BaseCurrencyCode string `form:"baseCurrencyCodeOrder"  search:"type:order;column:base_currency_code;table:ord_config_currency_rates"`
    QuoteCurrencyCode string `form:"quoteCurrencyCodeOrder"  search:"type:order;column:quote_currency_code;table:ord_config_currency_rates"`
    Rate string `form:"rateOrder"  search:"type:order;column:rate;table:ord_config_currency_rates"`
    Source string `form:"sourceOrder"  search:"type:order;column:source;table:ord_config_currency_rates"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_config_currency_rates"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_config_currency_rates"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_config_currency_rates"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_config_currency_rates"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_config_currency_rates"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_config_currency_rates"`

}

func (m *OrdConfigCurrencyRatesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdConfigCurrencyRatesInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    BaseCurrencyCode string `json:"baseCurrencyCode" comment:"基准货币代码 (ISO 4217)"`
    QuoteCurrencyCode string `json:"quoteCurrencyCode" comment:"报价货币代码 (ISO 4217)"`
    Rate string `json:"rate" comment:"汇率: 1 base_currency = rate quote_currency"`
    Source string `json:"source" comment:"汇率来源,如 manual, api, coingecko"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    common.ControlBy
}

func (s *OrdConfigCurrencyRatesInsertReq) Generate(model *models.OrdConfigCurrencyRates)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BaseCurrencyCode = s.BaseCurrencyCode
    model.QuoteCurrencyCode = s.QuoteCurrencyCode
    model.Rate = s.Rate
    model.Source = s.Source
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdConfigCurrencyRatesInsertReq) GetId() interface{} {
	return s.Id
}

type OrdConfigCurrencyRatesUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    BaseCurrencyCode string `json:"baseCurrencyCode" comment:"基准货币代码 (ISO 4217)"`
    QuoteCurrencyCode string `json:"quoteCurrencyCode" comment:"报价货币代码 (ISO 4217)"`
    Rate string `json:"rate" comment:"汇率: 1 base_currency = rate quote_currency"`
    Source string `json:"source" comment:"汇率来源,如 manual, api, coingecko"`
    Status string `json:"status" comment:"状态: 1=启用, 0=禁用"`
    common.ControlBy
}

func (s *OrdConfigCurrencyRatesUpdateReq) Generate(model *models.OrdConfigCurrencyRates)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BaseCurrencyCode = s.BaseCurrencyCode
    model.QuoteCurrencyCode = s.QuoteCurrencyCode
    model.Rate = s.Rate
    model.Source = s.Source
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdConfigCurrencyRatesUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdConfigCurrencyRatesGetReq 功能获取请求参数
type OrdConfigCurrencyRatesGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdConfigCurrencyRatesGetReq) GetId() interface{} {
	return s.Id
}

// OrdConfigCurrencyRatesDeleteReq 功能删除请求参数
type OrdConfigCurrencyRatesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdConfigCurrencyRatesDeleteReq) GetId() interface{} {
	return s.Ids
}

// OrdConfigCurrencyRatesBatchQueryReq 批量查询汇率请求参数
type OrdConfigCurrencyRatesBatchQueryReq struct {
	CurrencyPairs []CurrencyPair `json:"currencyPairs" binding:"required,min=1,max=500" comment:"货币对列表，最多300个"`
}

// CurrencyPair 货币对
type CurrencyPair struct {
	BaseCurrencyCode  string `json:"baseCurrencyCode" binding:"required" comment:"基准货币代码"`
	QuoteCurrencyCode string `json:"quoteCurrencyCode" binding:"required" comment:"报价货币代码"`
}

// OrdConfigCurrencyRatesBatchQueryResp 批量查询汇率响应
type OrdConfigCurrencyRatesBatchQueryResp struct {
	Rates []RateInfo `json:"rates" comment:"汇率信息列表"`
}

// RateInfo 汇率信息
type RateInfo struct {
	BaseCurrencyCode  string `json:"baseCurrencyCode" comment:"基准货币代码"`
	QuoteCurrencyCode string `json:"quoteCurrencyCode" comment:"报价货币代码"`
	Rate              string `json:"rate" comment:"汇率"`
	Source            string `json:"source" comment:"汇率来源"`
	Status            string `json:"status" comment:"状态"`
}
