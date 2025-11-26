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
    BankName string `form:"bankNameOrder"  search:"type:order;column:bank_name;table:hs_banks"`
    BankNameShort string `form:"bankNameShortOrder"  search:"type:order;column:bank_name_short;table:hs_banks"`
    ChannelType string `form:"channelTypeOrder"  search:"type:order;column:channel_type;table:hs_banks"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_banks"`
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
    BankName string `json:"bankName" comment:"银行名称"`
    BankNameShort string `json:"bankNameShort" comment:"银行名称缩写"`
    ChannelType int64 `json:"channelType" comment:"通道类型"`
    Status int `json:"status" comment:"状态: 0=禁用,1=启用"`
    common.ControlBy
}

func (s *HsBanksInsertReq) Generate(model *models.HsBanks)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.BankName = s.BankName
    model.BankNameShort = s.BankNameShort
    model.ChannelType = s.ChannelType
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsBanksInsertReq) GetId() interface{} {
	return s.Id
}

type HsBanksUpdateReq struct {
    Id int `uri:"id" comment:""` //
    BankCode string `json:"bankCode" comment:"银行编码/银行行号，如 SWIFT/BIC 或自定义编码"`
    BankName string `json:"bankName" comment:"银行名称"`
    BankNameShort string `json:"bankNameShort" comment:"银行名称缩写"`
    ChannelType int64 `json:"channelType" comment:"通道类型"`
    Status int `json:"status" comment:"状态: 0=禁用,1=启用"`
    common.ControlBy
}

func (s *HsBanksUpdateReq) Generate(model *models.HsBanks)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.BankCode = s.BankCode
    model.BankName = s.BankName
    model.BankNameShort = s.BankNameShort
    model.ChannelType = s.ChannelType
    model.Status = s.Status
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
