package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OrdGiftCardSuppliersGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:ord_gift_card_suppliers" comment:"供应商名称"`
    Code string `form:"code"  search:"type:exact;column:code;table:ord_gift_card_suppliers" comment:"供应商代码"`
    ContactPerson string `form:"contactPerson"  search:"type:exact;column:contact_person;table:ord_gift_card_suppliers" comment:"联系人"`
    ContactEmail string `form:"contactEmail"  search:"type:exact;column:contact_email;table:ord_gift_card_suppliers" comment:"联系邮箱"`
    ContactPhone string `form:"contactPhone"  search:"type:exact;column:contact_phone;table:ord_gift_card_suppliers" comment:"联系电话"`
    OrdGiftCardSuppliersOrder
}

type OrdGiftCardSuppliersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ord_gift_card_suppliers"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:ord_gift_card_suppliers"`
    Code string `form:"codeOrder"  search:"type:order;column:code;table:ord_gift_card_suppliers"`
    ContactPerson string `form:"contactPersonOrder"  search:"type:order;column:contact_person;table:ord_gift_card_suppliers"`
    ContactEmail string `form:"contactEmailOrder"  search:"type:order;column:contact_email;table:ord_gift_card_suppliers"`
    ContactPhone string `form:"contactPhoneOrder"  search:"type:order;column:contact_phone;table:ord_gift_card_suppliers"`
    ApiEndpoint string `form:"apiEndpointOrder"  search:"type:order;column:api_endpoint;table:ord_gift_card_suppliers"`
    ApiKey string `form:"apiKeyOrder"  search:"type:order;column:api_key;table:ord_gift_card_suppliers"`
    ApiConfig string `form:"apiConfigOrder"  search:"type:order;column:api_config;table:ord_gift_card_suppliers"`
    SettlementRate string `form:"settlementRateOrder"  search:"type:order;column:settlement_rate;table:ord_gift_card_suppliers"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ord_gift_card_suppliers"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ord_gift_card_suppliers"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ord_gift_card_suppliers"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ord_gift_card_suppliers"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ord_gift_card_suppliers"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ord_gift_card_suppliers"`
    
}

func (m *OrdGiftCardSuppliersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type OrdGiftCardSuppliersInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    Name string `json:"name" comment:"供应商名称"`
    Code string `json:"code" comment:"供应商代码"`
    ContactPerson string `json:"contactPerson" comment:"联系人"`
    ContactEmail string `json:"contactEmail" comment:"联系邮箱"`
    ContactPhone string `json:"contactPhone" comment:"联系电话"`
    ApiEndpoint string `json:"apiEndpoint" comment:"API接口地址"`
    ApiKey string `json:"apiKey" comment:"API密钥"`
    ApiConfig string `json:"apiConfig" comment:"API配置信息"`
    SettlementRate string `json:"settlementRate" comment:"结算费率"`
    Status string `json:"status" comment:"状态：0=禁用，1=启用"`
    common.ControlBy
}

func (s *OrdGiftCardSuppliersInsertReq) Generate(model *models.OrdGiftCardSuppliers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Code = s.Code
    model.ContactPerson = s.ContactPerson
    model.ContactEmail = s.ContactEmail
    model.ContactPhone = s.ContactPhone
    model.ApiEndpoint = s.ApiEndpoint
    model.ApiKey = s.ApiKey
    model.ApiConfig = s.ApiConfig
    model.SettlementRate = s.SettlementRate
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *OrdGiftCardSuppliersInsertReq) GetId() interface{} {
	return s.Id
}

type OrdGiftCardSuppliersUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    Name string `json:"name" comment:"供应商名称"`
    Code string `json:"code" comment:"供应商代码"`
    ContactPerson string `json:"contactPerson" comment:"联系人"`
    ContactEmail string `json:"contactEmail" comment:"联系邮箱"`
    ContactPhone string `json:"contactPhone" comment:"联系电话"`
    ApiEndpoint string `json:"apiEndpoint" comment:"API接口地址"`
    ApiKey string `json:"apiKey" comment:"API密钥"`
    ApiConfig string `json:"apiConfig" comment:"API配置信息"`
    SettlementRate string `json:"settlementRate" comment:"结算费率"`
    Status string `json:"status" comment:"状态：0=禁用，1=启用"`
    common.ControlBy
}

func (s *OrdGiftCardSuppliersUpdateReq) Generate(model *models.OrdGiftCardSuppliers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.Code = s.Code
    model.ContactPerson = s.ContactPerson
    model.ContactEmail = s.ContactEmail
    model.ContactPhone = s.ContactPhone
    model.ApiEndpoint = s.ApiEndpoint
    model.ApiKey = s.ApiKey
    model.ApiConfig = s.ApiConfig
    model.SettlementRate = s.SettlementRate
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *OrdGiftCardSuppliersUpdateReq) GetId() interface{} {
	return s.Id
}

// OrdGiftCardSuppliersGetReq 功能获取请求参数
type OrdGiftCardSuppliersGetReq struct {
     Id int `uri:"id"`
}
func (s *OrdGiftCardSuppliersGetReq) GetId() interface{} {
	return s.Id
}

// OrdGiftCardSuppliersDeleteReq 功能删除请求参数
type OrdGiftCardSuppliersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *OrdGiftCardSuppliersDeleteReq) GetId() interface{} {
	return s.Ids
}
