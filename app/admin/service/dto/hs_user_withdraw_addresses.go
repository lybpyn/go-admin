package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserWithdrawAddressesGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_withdraw_addresses" comment:"所属用户ID"`
    Address string `form:"address"  search:"type:exact;column:address;table:hs_user_withdraw_addresses" comment:"钱包提现地址"`
    HsUserWithdrawAddressesOrder
}

type HsUserWithdrawAddressesOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_withdraw_addresses"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_withdraw_addresses"`
    ConfigSupportedChainsCoinId string `form:"configSupportedChainsCoinIdOrder"  search:"type:order;column:config_supported_chains_coin_id;table:hs_user_withdraw_addresses"`
    Address string `form:"addressOrder"  search:"type:order;column:address;table:hs_user_withdraw_addresses"`
    Label string `form:"labelOrder"  search:"type:order;column:label;table:hs_user_withdraw_addresses"`
    IsDefault string `form:"isDefaultOrder"  search:"type:order;column:is_default;table:hs_user_withdraw_addresses"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_user_withdraw_addresses"`
    CreatedIp string `form:"createdIpOrder"  search:"type:order;column:created_ip;table:hs_user_withdraw_addresses"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_withdraw_addresses"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_withdraw_addresses"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_withdraw_addresses"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_withdraw_addresses"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_withdraw_addresses"`
    
}

func (m *HsUserWithdrawAddressesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserWithdrawAddressesInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"所属用户ID"`
    ConfigSupportedChainsCoinId string `json:"configSupportedChainsCoinId" comment:"链类型，例如 TRC20、ERC20、BEP20"`
    Address string `json:"address" comment:"钱包提现地址"`
    Label string `json:"label" comment:"用户自定义备注标签，例如“我的TRC钱包”"`
    IsDefault string `json:"isDefault" comment:"是否默认提现地址：0=否，1=是"`
    Status string `json:"status" comment:"状态：1=启用，0=禁用，2=待审核，-1=拒绝"`
    CreatedIp string `json:"createdIp" comment:"添加时的IP地址"`
    common.ControlBy
}

func (s *HsUserWithdrawAddressesInsertReq) Generate(model *models.HsUserWithdrawAddresses)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.ConfigSupportedChainsCoinId = s.ConfigSupportedChainsCoinId
    model.Address = s.Address
    model.Label = s.Label
    model.IsDefault = s.IsDefault
    model.Status = s.Status
    model.CreatedIp = s.CreatedIp
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserWithdrawAddressesInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserWithdrawAddressesUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    UserId string `json:"userId" comment:"所属用户ID"`
    ConfigSupportedChainsCoinId string `json:"configSupportedChainsCoinId" comment:"链类型，例如 TRC20、ERC20、BEP20"`
    Address string `json:"address" comment:"钱包提现地址"`
    Label string `json:"label" comment:"用户自定义备注标签，例如“我的TRC钱包”"`
    IsDefault string `json:"isDefault" comment:"是否默认提现地址：0=否，1=是"`
    Status string `json:"status" comment:"状态：1=启用，0=禁用，2=待审核，-1=拒绝"`
    CreatedIp string `json:"createdIp" comment:"添加时的IP地址"`
    common.ControlBy
}

func (s *HsUserWithdrawAddressesUpdateReq) Generate(model *models.HsUserWithdrawAddresses)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.ConfigSupportedChainsCoinId = s.ConfigSupportedChainsCoinId
    model.Address = s.Address
    model.Label = s.Label
    model.IsDefault = s.IsDefault
    model.Status = s.Status
    model.CreatedIp = s.CreatedIp
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserWithdrawAddressesUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserWithdrawAddressesGetReq 功能获取请求参数
type HsUserWithdrawAddressesGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserWithdrawAddressesGetReq) GetId() interface{} {
	return s.Id
}

// HsUserWithdrawAddressesDeleteReq 功能删除请求参数
type HsUserWithdrawAddressesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserWithdrawAddressesDeleteReq) GetId() interface{} {
	return s.Ids
}
