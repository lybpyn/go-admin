package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUsersGetPageReq struct {
	dto.Pagination     `search:"-"`
    Username string `form:"username"  search:"type:exact;column:username;table:hs_users" comment:"用户名（可选展示用）"`
    Firstname string `form:"firstname"  search:"type:exact;column:firstname;table:hs_users" comment:""`
    Lastname string `form:"lastname"  search:"type:exact;column:lastname;table:hs_users" comment:""`
    InviteCode string `form:"inviteCode"  search:"type:exact;column:invite_code;table:hs_users" comment:""`
    HsUsersOrder
}

type HsUsersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_users"`
    Username string `form:"usernameOrder"  search:"type:order;column:username;table:hs_users"`
    PasswordHash string `form:"passwordHashOrder"  search:"type:order;column:password_hash;table:hs_users"`
    Firstname string `form:"firstnameOrder"  search:"type:order;column:firstname;table:hs_users"`
    Lastname string `form:"lastnameOrder"  search:"type:order;column:lastname;table:hs_users"`
    Avatar string `form:"avatarOrder"  search:"type:order;column:avatar;table:hs_users"`
    Balance string `form:"balanceOrder"  search:"type:order;column:balance;table:hs_users"`
    LevelId string `form:"levelIdOrder"  search:"type:order;column:level_id;table:hs_users"`
    Experience string `form:"experienceOrder"  search:"type:order;column:experience;table:hs_users"`
    TotalExperience string `form:"totalExperienceOrder"  search:"type:order;column:total_experience;table:hs_users"`
    InviteCode string `form:"inviteCodeOrder"  search:"type:order;column:invite_code;table:hs_users"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_users"`
    Version string `form:"versionOrder"  search:"type:order;column:version;table:hs_users"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_users"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_users"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_users"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_users"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_users"`
    
}

func (m *HsUsersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUsersInsertReq struct {
    Id int `json:"-" comment:""` // 
    Username string `json:"username" comment:"用户名（可选展示用）"`
    PasswordHash string `json:"passwordHash" comment:""`
    Firstname string `json:"firstname" comment:""`
    Lastname string `json:"lastname" comment:""`
    Avatar string `json:"avatar" comment:"头像URL"`
    Balance string `json:"balance" comment:""`
    LevelId string `json:"levelId" comment:"用户等级ID"`
    Experience string `json:"experience" comment:"当前经验"`
    TotalExperience string `json:"totalExperience" comment:"累计经验"`
    InviteCode string `json:"inviteCode" comment:""`
    Status string `json:"status" comment:"状态：1正常，0封禁"`
    Version string `json:"version" comment:""`
    common.ControlBy
}

func (s *HsUsersInsertReq) Generate(model *models.HsUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Username = s.Username
    model.PasswordHash = s.PasswordHash
    model.Firstname = s.Firstname
    model.Lastname = s.Lastname
    model.Avatar = s.Avatar
    model.Balance = s.Balance
    model.LevelId = s.LevelId
    model.Experience = s.Experience
    model.TotalExperience = s.TotalExperience
    model.InviteCode = s.InviteCode
    model.Status = s.Status
    model.Version = s.Version
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUsersInsertReq) GetId() interface{} {
	return s.Id
}

type HsUsersUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Username string `json:"username" comment:"用户名（可选展示用）"`
    PasswordHash string `json:"passwordHash" comment:""`
    Firstname string `json:"firstname" comment:""`
    Lastname string `json:"lastname" comment:""`
    Avatar string `json:"avatar" comment:"头像URL"`
    Balance string `json:"balance" comment:""`
    LevelId string `json:"levelId" comment:"用户等级ID"`
    Experience string `json:"experience" comment:"当前经验"`
    TotalExperience string `json:"totalExperience" comment:"累计经验"`
    InviteCode string `json:"inviteCode" comment:""`
    Status string `json:"status" comment:"状态：1正常，0封禁"`
    Version string `json:"version" comment:""`
    common.ControlBy
}

func (s *HsUsersUpdateReq) Generate(model *models.HsUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Username = s.Username
    model.PasswordHash = s.PasswordHash
    model.Firstname = s.Firstname
    model.Lastname = s.Lastname
    model.Avatar = s.Avatar
    model.Balance = s.Balance
    model.LevelId = s.LevelId
    model.Experience = s.Experience
    model.TotalExperience = s.TotalExperience
    model.InviteCode = s.InviteCode
    model.Status = s.Status
    model.Version = s.Version
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUsersUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUsersGetReq 功能获取请求参数
type HsUsersGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUsersGetReq) GetId() interface{} {
	return s.Id
}

// HsUsersDeleteReq 功能删除请求参数
type HsUsersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUsersDeleteReq) GetId() interface{} {
	return s.Ids
}
