package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsInvitesBindGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_invites_bind" comment:"被邀请用户ID"`
    InviterId string `form:"inviterId"  search:"type:exact;column:inviter_id;table:hs_invites_bind" comment:"直接邀请人ID (一级)"`
    Level string `form:"level"  search:"type:exact;column:level;table:hs_invites_bind" comment:"层级: 1=一级, 2=二级"`
    HsInvitesBindOrder
}

type HsInvitesBindOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_invites_bind"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_invites_bind"`
    InviterId string `form:"inviterIdOrder"  search:"type:order;column:inviter_id;table:hs_invites_bind"`
    Level string `form:"levelOrder"  search:"type:order;column:level;table:hs_invites_bind"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_invites_bind"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_invites_bind"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_invites_bind"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_invites_bind"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_invites_bind"`
    
}

func (m *HsInvitesBindGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsInvitesBindInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:"被邀请用户ID"`
    InviterId string `json:"inviterId" comment:"直接邀请人ID (一级)"`
    Level string `json:"level" comment:"层级: 1=一级, 2=二级"`
    common.ControlBy
}

func (s *HsInvitesBindInsertReq) Generate(model *models.HsInvitesBind)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.InviterId = s.InviterId
    model.Level = s.Level
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsInvitesBindInsertReq) GetId() interface{} {
	return s.Id
}

type HsInvitesBindUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:"被邀请用户ID"`
    InviterId string `json:"inviterId" comment:"直接邀请人ID (一级)"`
    Level string `json:"level" comment:"层级: 1=一级, 2=二级"`
    common.ControlBy
}

func (s *HsInvitesBindUpdateReq) Generate(model *models.HsInvitesBind)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.InviterId = s.InviterId
    model.Level = s.Level
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsInvitesBindUpdateReq) GetId() interface{} {
	return s.Id
}

// HsInvitesBindGetReq 功能获取请求参数
type HsInvitesBindGetReq struct {
     Id int `uri:"id"`
}
func (s *HsInvitesBindGetReq) GetId() interface{} {
	return s.Id
}

// HsInvitesBindDeleteReq 功能删除请求参数
type HsInvitesBindDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsInvitesBindDeleteReq) GetId() interface{} {
	return s.Ids
}
