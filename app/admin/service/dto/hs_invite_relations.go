package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsInviteRelationsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_invite_relations" comment:"用户ID"`
    Level1InviterId string `form:"level1InviterId"  search:"type:exact;column:level1_inviter_id;table:hs_invite_relations" comment:"一级邀请人ID"`
    Level2InviterId string `form:"level2InviterId"  search:"type:exact;column:level2_inviter_id;table:hs_invite_relations" comment:"二级邀请人ID"`
    HsInviteRelationsOrder
}

type HsInviteRelationsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_invite_relations"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_invite_relations"`
    Level1InviterId string `form:"level1InviterIdOrder"  search:"type:order;column:level1_inviter_id;table:hs_invite_relations"`
    Level2InviterId string `form:"level2InviterIdOrder"  search:"type:order;column:level2_inviter_id;table:hs_invite_relations"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_invite_relations"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_invite_relations"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_invite_relations"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_invite_relations"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_invite_relations"`
    
}

func (m *HsInviteRelationsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsInviteRelationsInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:"用户ID"`
    Level1InviterId string `json:"level1InviterId" comment:"一级邀请人ID"`
    Level2InviterId string `json:"level2InviterId" comment:"二级邀请人ID"`
    common.ControlBy
}

func (s *HsInviteRelationsInsertReq) Generate(model *models.HsInviteRelations)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.Level1InviterId = s.Level1InviterId
    model.Level2InviterId = s.Level2InviterId
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsInviteRelationsInsertReq) GetId() interface{} {
	return s.Id
}

type HsInviteRelationsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:"用户ID"`
    Level1InviterId string `json:"level1InviterId" comment:"一级邀请人ID"`
    Level2InviterId string `json:"level2InviterId" comment:"二级邀请人ID"`
    common.ControlBy
}

func (s *HsInviteRelationsUpdateReq) Generate(model *models.HsInviteRelations)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.Level1InviterId = s.Level1InviterId
    model.Level2InviterId = s.Level2InviterId
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsInviteRelationsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsInviteRelationsGetReq 功能获取请求参数
type HsInviteRelationsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsInviteRelationsGetReq) GetId() interface{} {
	return s.Id
}

// HsInviteRelationsDeleteReq 功能删除请求参数
type HsInviteRelationsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsInviteRelationsDeleteReq) GetId() interface{} {
	return s.Ids
}
