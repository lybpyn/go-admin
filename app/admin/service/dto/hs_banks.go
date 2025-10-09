package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsBanksGetPageReq struct {
	dto.Pagination     `search:"-"`
    HsBanksOrder
}

type HsBanksOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_banks"`
    BankCode string `form:"bankCodeOrder"  search:"type:order;column:bank_code;table:hs_banks"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:hs_banks"`
    Country string `form:"countryOrder"  search:"type:order;column:country;table:hs_banks"`
    SwiftCode string `form:"swiftCodeOrder"  search:"type:order;column:swift_code;table:hs_banks"`
    RoutingNumber string `form:"routingNumberOrder"  search:"type:order;column:routing_number;table:hs_banks"`
    SupportsInternational string `form:"supportsInternationalOrder"  search:"type:order;column:supports_international;table:hs_banks"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_banks"`
    Note string `form:"noteOrder"  search:"type:order;column:note;table:hs_banks"`
    Extra string `form:"extraOrder"  search:"type:order;column:extra;table:hs_banks"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_banks"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_banks"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_banks"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_banks"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_banks"`
    
}

func (m *HsBanksGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsBanksInsertReq struct {
    Id int `json:"-" comment:""` // 
    BankCode string `json:"bankCode" comment:"银行编码/银行行号，如 SWIFT/BIC 或自定义编码"`
    Name string `json:"name" comment:"银行名称"`
    Country string `json:"country" comment:"国家 ISO2"`
    SwiftCode string `json:"swiftCode" comment:"SWIFT/BIC (如适用)"`
    RoutingNumber string `json:"routingNumber" comment:"路由/清算号(如 ABA)"`
    SupportsInternational string `json:"supportsInternational" comment:"是否支持国际汇款: 0否,1是"`
    Status string `json:"status" comment:"状态: 0=禁用,1=启用"`
    Note string `json:"note" comment:""`
    Extra string `json:"extra" comment:"扩展字段，例如银行支付说明、证件要求等"`
    common.ControlBy
}

func (s *HsBanksInsertReq) Generate(model *models.HsBanks)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.Name = s.Name
    model.Country = s.Country
    model.SwiftCode = s.SwiftCode
    model.RoutingNumber = s.RoutingNumber
    model.SupportsInternational = s.SupportsInternational
    model.Status = s.Status
    model.Note = s.Note
    model.Extra = s.Extra
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsBanksInsertReq) GetId() interface{} {
	return s.Id
}

type HsBanksUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    BankCode string `json:"bankCode" comment:"银行编码/银行行号，如 SWIFT/BIC 或自定义编码"`
    Name string `json:"name" comment:"银行名称"`
    Country string `json:"country" comment:"国家 ISO2"`
    SwiftCode string `json:"swiftCode" comment:"SWIFT/BIC (如适用)"`
    RoutingNumber string `json:"routingNumber" comment:"路由/清算号(如 ABA)"`
    SupportsInternational string `json:"supportsInternational" comment:"是否支持国际汇款: 0否,1是"`
    Status string `json:"status" comment:"状态: 0=禁用,1=启用"`
    Note string `json:"note" comment:""`
    Extra string `json:"extra" comment:"扩展字段，例如银行支付说明、证件要求等"`
    common.ControlBy
}

func (s *HsBanksUpdateReq) Generate(model *models.HsBanks)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.Name = s.Name
    model.Country = s.Country
    model.SwiftCode = s.SwiftCode
    model.RoutingNumber = s.RoutingNumber
    model.SupportsInternational = s.SupportsInternational
    model.Status = s.Status
    model.Note = s.Note
    model.Extra = s.Extra
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsBanksUpdateReq) GetId() interface{} {
	return s.Id
}

// HsBanksGetReq 功能获取请求参数
type HsBanksGetReq struct {
     Id int `uri:"id"`
}
func (s *HsBanksGetReq) GetId() interface{} {
	return s.Id
}

// HsBanksDeleteReq 功能删除请求参数
type HsBanksDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsBanksDeleteReq) GetId() interface{} {
	return s.Ids
}
