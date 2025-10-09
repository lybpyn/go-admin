package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsMerchantsGetPageReq struct {
	dto.Pagination     `search:"-"`
    HsMerchantsOrder
}

type HsMerchantsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_merchants"`
    MerchantCode string `form:"merchantCodeOrder"  search:"type:order;column:merchant_code;table:hs_merchants"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:hs_merchants"`
    ContactName string `form:"contactNameOrder"  search:"type:order;column:contact_name;table:hs_merchants"`
    ContactPhone string `form:"contactPhoneOrder"  search:"type:order;column:contact_phone;table:hs_merchants"`
    ContactEmail string `form:"contactEmailOrder"  search:"type:order;column:contact_email;table:hs_merchants"`
    Country string `form:"countryOrder"  search:"type:order;column:country;table:hs_merchants"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_merchants"`
    DailyLimit string `form:"dailyLimitOrder"  search:"type:order;column:daily_limit;table:hs_merchants"`
    Note string `form:"noteOrder"  search:"type:order;column:note;table:hs_merchants"`
    Extra string `form:"extraOrder"  search:"type:order;column:extra;table:hs_merchants"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_merchants"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_merchants"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_merchants"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_merchants"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_merchants"`
    
}

func (m *HsMerchantsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsMerchantsInsertReq struct {
    Id int `json:"-" comment:""` // 
    MerchantCode string `json:"merchantCode" comment:"外部/内部唯一编码 (可用于对接第三方)"`
    Name string `json:"name" comment:"卡商名称/公司名"`
    ContactName string `json:"contactName" comment:"联系人姓名"`
    ContactPhone string `json:"contactPhone" comment:"联系人电话"`
    ContactEmail string `json:"contactEmail" comment:"联系人邮箱"`
    Country string `json:"country" comment:"国家/地区 ISO2 (如 CN, US)"`
    Status string `json:"status" comment:"状态: 0=禁用,1=启用,2=冻结"`
    DailyLimit string `json:"dailyLimit" comment:"日限额 (可选)"`
    Note string `json:"note" comment:"备注/其他说明"`
    Extra string `json:"extra" comment:"扩展信息: 如资质文件url、合同信息等"`
    common.ControlBy
}

func (s *HsMerchantsInsertReq) Generate(model *models.HsMerchants)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.MerchantCode = s.MerchantCode
    model.Name = s.Name
    model.ContactName = s.ContactName
    model.ContactPhone = s.ContactPhone
    model.ContactEmail = s.ContactEmail
    model.Country = s.Country
    model.Status = s.Status
    model.DailyLimit = s.DailyLimit
    model.Note = s.Note
    model.Extra = s.Extra
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsMerchantsInsertReq) GetId() interface{} {
	return s.Id
}

type HsMerchantsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    MerchantCode string `json:"merchantCode" comment:"外部/内部唯一编码 (可用于对接第三方)"`
    Name string `json:"name" comment:"卡商名称/公司名"`
    ContactName string `json:"contactName" comment:"联系人姓名"`
    ContactPhone string `json:"contactPhone" comment:"联系人电话"`
    ContactEmail string `json:"contactEmail" comment:"联系人邮箱"`
    Country string `json:"country" comment:"国家/地区 ISO2 (如 CN, US)"`
    Status string `json:"status" comment:"状态: 0=禁用,1=启用,2=冻结"`
    DailyLimit string `json:"dailyLimit" comment:"日限额 (可选)"`
    Note string `json:"note" comment:"备注/其他说明"`
    Extra string `json:"extra" comment:"扩展信息: 如资质文件url、合同信息等"`
    common.ControlBy
}

func (s *HsMerchantsUpdateReq) Generate(model *models.HsMerchants)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.MerchantCode = s.MerchantCode
    model.Name = s.Name
    model.ContactName = s.ContactName
    model.ContactPhone = s.ContactPhone
    model.ContactEmail = s.ContactEmail
    model.Country = s.Country
    model.Status = s.Status
    model.DailyLimit = s.DailyLimit
    model.Note = s.Note
    model.Extra = s.Extra
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsMerchantsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsMerchantsGetReq 功能获取请求参数
type HsMerchantsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsMerchantsGetReq) GetId() interface{} {
	return s.Id
}

// HsMerchantsDeleteReq 功能删除请求参数
type HsMerchantsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsMerchantsDeleteReq) GetId() interface{} {
	return s.Ids
}
