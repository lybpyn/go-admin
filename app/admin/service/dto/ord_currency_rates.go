package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdCurrencyRatesGetPageReq struct {
	dto.Pagination     `search:"-"`
    QuoteCurrency string `form:"quoteCurrency"  search:"type:exact;column:quote_currency;table:ord_currency_rates" comment:"计价货币 ISO4217，例如 CNY"`
    OrdCurrencyRatesOrder
}

type OrdCurrencyRatesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_currency_rates"`
    QuoteCurrency string `form:"quoteCurrencyOrder"  search:"type:order;column:quote_currency;table:ord_currency_rates"`
    Rate string `form:"rateOrder"  search:"type:order;column:rate;table:ord_currency_rates"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_currency_rates"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_currency_rates"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_currency_rates"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_currency_rates"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_currency_rates"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_currency_rates"`
    
}

func (m *OrdCurrencyRatesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdCurrencyRatesInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    QuoteCurrency string `json:"quoteCurrency" comment:"计价货币 ISO4217，例如 CNY"`
    Rate string `json:"rate" comment:"汇率，表示 1 base_currency = rate quote_currency"`
    Status string `json:"status" comment:"状态: 1=启用,0=禁用"`
    common.ControlBy
}

func (s *OrdCurrencyRatesInsertReq) Generate(model *models.OrdCurrencyRates)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.QuoteCurrency = s.QuoteCurrency
    model.Rate = s.Rate
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdCurrencyRatesInsertReq) GetId() interface{} {
	return s.Id
}

type OrdCurrencyRatesUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    QuoteCurrency string `json:"quoteCurrency" comment:"计价货币 ISO4217，例如 CNY"`
    Rate string `json:"rate" comment:"汇率，表示 1 base_currency = rate quote_currency"`
    Status string `json:"status" comment:"状态: 1=启用,0=禁用"`
    common.ControlBy
}

func (s *OrdCurrencyRatesUpdateReq) Generate(model *models.OrdCurrencyRates)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.QuoteCurrency = s.QuoteCurrency
    model.Rate = s.Rate
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdCurrencyRatesUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdCurrencyRatesGetReq 功能获取请求参数
type OrdCurrencyRatesGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdCurrencyRatesGetReq) GetId() interface{} {
	return s.Id
}

// OrdCurrencyRatesDeleteReq 功能删除请求参数
type OrdCurrencyRatesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdCurrencyRatesDeleteReq) GetId() interface{} {
	return s.Ids
}
