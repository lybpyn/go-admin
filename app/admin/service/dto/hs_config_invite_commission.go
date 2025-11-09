package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfigInviteCommissionGetPageReq struct {
	dto.Pagination     `search:"-"`
    ConfigName string `form:"configName"  search:"type:exact;column:config_name;table:hs_config_invite_commission" comment:"配置名称"`
    Status string `form:"status"  search:"type:exact;column:status;table:hs_config_invite_commission" comment:"状态：1=启用，0=禁用"`
    HsConfigInviteCommissionOrder
}

type HsConfigInviteCommissionOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_config_invite_commission"`
    ConfigName string `form:"configNameOrder"  search:"type:order;column:config_name;table:hs_config_invite_commission"`
    FirstLevelRate string `form:"firstLevelRateOrder"  search:"type:order;column:first_level_rate;table:hs_config_invite_commission"`
    SecondLevelRate string `form:"secondLevelRateOrder"  search:"type:order;column:second_level_rate;table:hs_config_invite_commission"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_config_invite_commission"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:hs_config_invite_commission"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_config_invite_commission"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_config_invite_commission"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_config_invite_commission"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_config_invite_commission"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_config_invite_commission"`
    
}

func (m *HsConfigInviteCommissionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfigInviteCommissionInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    ConfigName string `json:"configName" comment:"配置名称"`
    FirstLevelRate string `json:"firstLevelRate" comment:"一级邀请分成比例（百分比，如5.0000表示5%）"`
    SecondLevelRate string `json:"secondLevelRate" comment:"二级邀请分成比例（百分比，如3.0000表示3%）"`
    Status string `json:"status" comment:"状态：1=启用，0=禁用"`
    Remark string `json:"remark" comment:"备注说明"`
    common.ControlBy
}

func (s *HsConfigInviteCommissionInsertReq) Generate(model *models.HsConfigInviteCommission)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ConfigName = s.ConfigName
    model.FirstLevelRate = s.FirstLevelRate
    model.SecondLevelRate = s.SecondLevelRate
    model.Status = s.Status
    model.Remark = s.Remark
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfigInviteCommissionInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfigInviteCommissionUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    ConfigName string `json:"configName" comment:"配置名称"`
    FirstLevelRate string `json:"firstLevelRate" comment:"一级邀请分成比例（百分比，如5.0000表示5%）"`
    SecondLevelRate string `json:"secondLevelRate" comment:"二级邀请分成比例（百分比，如3.0000表示3%）"`
    Status string `json:"status" comment:"状态：1=启用，0=禁用"`
    Remark string `json:"remark" comment:"备注说明"`
    common.ControlBy
}

func (s *HsConfigInviteCommissionUpdateReq) Generate(model *models.HsConfigInviteCommission)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ConfigName = s.ConfigName
    model.FirstLevelRate = s.FirstLevelRate
    model.SecondLevelRate = s.SecondLevelRate
    model.Status = s.Status
    model.Remark = s.Remark
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfigInviteCommissionUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfigInviteCommissionGetReq 功能获取请求参数
type HsConfigInviteCommissionGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfigInviteCommissionGetReq) GetId() interface{} {
	return s.Id
}

// HsConfigInviteCommissionDeleteReq 功能删除请求参数
type HsConfigInviteCommissionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfigInviteCommissionDeleteReq) GetId() interface{} {
	return s.Ids
}
