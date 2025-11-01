package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsConfgiUserLevelsGetPageReq struct {
	dto.Pagination     `search:"-"`
    HsConfgiUserLevelsOrder
}

type HsConfgiUserLevelsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_confgi_user_levels"`
    LevelName string `form:"levelNameOrder"  search:"type:order;column:level_name;table:hs_confgi_user_levels"`
    UpExperience string `form:"upExperienceOrder"  search:"type:order;column:up_experience;table:hs_confgi_user_levels"`
    LevelIcon string `form:"levelIconOrder"  search:"type:order;column:level_icon;table:hs_confgi_user_levels"`
    LevelPrivileges string `form:"levelPrivilegesOrder"  search:"type:order;column:level_privileges;table:hs_confgi_user_levels"`
    FirstInviteCommossions string `form:"firstInviteCommossionsOrder"  search:"type:order;column:first_invite_commossions;table:hs_confgi_user_levels"`
    SecondInviteCommossions string `form:"secondInviteCommossionsOrder"  search:"type:order;column:second_invite_commossions;table:hs_confgi_user_levels"`
    SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:hs_confgi_user_levels"`
    IsActive string `form:"isActiveOrder"  search:"type:order;column:is_active;table:hs_confgi_user_levels"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_confgi_user_levels"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_confgi_user_levels"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_confgi_user_levels"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_confgi_user_levels"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_confgi_user_levels"`
    
}

func (m *HsConfgiUserLevelsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsConfgiUserLevelsInsertReq struct {
    Id int `json:"-" comment:"等级ID"` // 等级ID
    LevelName string `json:"levelName" comment:"等级名称"`
    UpExperience string `json:"upExperience" comment:"升级所需经验值"`
    LevelIcon string `json:"levelIcon" comment:"等级图标URL"`
    LevelPrivileges string `json:"levelPrivileges" comment:"等级特权配置(JSON格式)"`
    FirstInviteCommossions string `json:"firstInviteCommossions" comment:"一级分成奖励"`
    SecondInviteCommossions string `json:"secondInviteCommossions" comment:"二级分成奖励"`
    SortOrder string `json:"sortOrder" comment:"排序顺序"`
    IsActive string `json:"isActive" comment:"是否启用"`
    common.ControlBy
}

func (s *HsConfgiUserLevelsInsertReq) Generate(model *models.HsConfgiUserLevels)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.LevelName = s.LevelName
    model.UpExperience = s.UpExperience
    model.LevelIcon = s.LevelIcon
    model.LevelPrivileges = s.LevelPrivileges
    model.FirstInviteCommossions = s.FirstInviteCommossions
    model.SecondInviteCommossions = s.SecondInviteCommossions
    model.SortOrder = s.SortOrder
    model.IsActive = s.IsActive
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsConfgiUserLevelsInsertReq) GetId() interface{} {
	return s.Id
}

type HsConfgiUserLevelsUpdateReq struct {
    Id int `uri:"id" comment:"等级ID"` // 等级ID
    LevelName string `json:"levelName" comment:"等级名称"`
    UpExperience string `json:"upExperience" comment:"升级所需经验值"`
    LevelIcon string `json:"levelIcon" comment:"等级图标URL"`
    LevelPrivileges string `json:"levelPrivileges" comment:"等级特权配置(JSON格式)"`
    FirstInviteCommossions string `json:"firstInviteCommossions" comment:"一级分成奖励"`
    SecondInviteCommossions string `json:"secondInviteCommossions" comment:"二级分成奖励"`
    SortOrder string `json:"sortOrder" comment:"排序顺序"`
    IsActive string `json:"isActive" comment:"是否启用"`
    common.ControlBy
}

func (s *HsConfgiUserLevelsUpdateReq) Generate(model *models.HsConfgiUserLevels)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.LevelName = s.LevelName
    model.UpExperience = s.UpExperience
    model.LevelIcon = s.LevelIcon
    model.LevelPrivileges = s.LevelPrivileges
    model.FirstInviteCommossions = s.FirstInviteCommossions
    model.SecondInviteCommossions = s.SecondInviteCommossions
    model.SortOrder = s.SortOrder
    model.IsActive = s.IsActive
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsConfgiUserLevelsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsConfgiUserLevelsGetReq 功能获取请求参数
type HsConfgiUserLevelsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsConfgiUserLevelsGetReq) GetId() interface{} {
	return s.Id
}

// HsConfgiUserLevelsDeleteReq 功能删除请求参数
type HsConfgiUserLevelsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsConfgiUserLevelsDeleteReq) GetId() interface{} {
	return s.Ids
}
