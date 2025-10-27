package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsInviteCommissionsGetPageReq struct {
	dto.Pagination     `search:"-"`
    OrderId string `form:"orderId"  search:"type:exact;column:order_id;table:hs_invite_commissions" comment:"订单ID"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_invite_commissions" comment:"分成获得者ID"`
    HsInviteCommissionsOrder
}

type HsInviteCommissionsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_invite_commissions"`
    OrderId string `form:"orderIdOrder"  search:"type:order;column:order_id;table:hs_invite_commissions"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_invite_commissions"`
    CommissionLevel string `form:"commissionLevelOrder"  search:"type:order;column:commission_level;table:hs_invite_commissions"`
    CommissionRate string `form:"commissionRateOrder"  search:"type:order;column:commission_rate;table:hs_invite_commissions"`
    CommissionAmount string `form:"commissionAmountOrder"  search:"type:order;column:commission_amount;table:hs_invite_commissions"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_invite_commissions"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_invite_commissions"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_invite_commissions"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_invite_commissions"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_invite_commissions"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_invite_commissions"`
    
}

func (m *HsInviteCommissionsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsInviteCommissionsInsertReq struct {
    Id int `json:"-" comment:""` // 
    OrderId string `json:"orderId" comment:"订单ID"`
    UserId string `json:"userId" comment:"分成获得者ID"`
    CommissionLevel string `json:"commissionLevel" comment:"分成层级: 1=一级 5%, 2=二级 3%"`
    CommissionRate string `json:"commissionRate" comment:"分成比例，例如 0.05 表示5%"`
    CommissionAmount string `json:"commissionAmount" comment:"分成金额"`
    Status string `json:"status" comment:"状态: 0=待结算,1=已结算,2=取消"`
    common.ControlBy
}

func (s *HsInviteCommissionsInsertReq) Generate(model *models.HsInviteCommissions)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.OrderId = s.OrderId
    model.UserId = s.UserId
    model.CommissionLevel = s.CommissionLevel
    model.CommissionRate = s.CommissionRate
    model.CommissionAmount = s.CommissionAmount
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsInviteCommissionsInsertReq) GetId() interface{} {
	return s.Id
}

type HsInviteCommissionsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    OrderId string `json:"orderId" comment:"订单ID"`
    UserId string `json:"userId" comment:"分成获得者ID"`
    CommissionLevel string `json:"commissionLevel" comment:"分成层级: 1=一级 5%, 2=二级 3%"`
    CommissionRate string `json:"commissionRate" comment:"分成比例，例如 0.05 表示5%"`
    CommissionAmount string `json:"commissionAmount" comment:"分成金额"`
    Status string `json:"status" comment:"状态: 0=待结算,1=已结算,2=取消"`
    common.ControlBy
}

func (s *HsInviteCommissionsUpdateReq) Generate(model *models.HsInviteCommissions)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.OrderId = s.OrderId
    model.UserId = s.UserId
    model.CommissionLevel = s.CommissionLevel
    model.CommissionRate = s.CommissionRate
    model.CommissionAmount = s.CommissionAmount
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsInviteCommissionsUpdateReq) GetId() interface{} {
	return s.Id
}

// HsInviteCommissionsGetReq 功能获取请求参数
type HsInviteCommissionsGetReq struct {
     Id int `uri:"id"`
}
func (s *HsInviteCommissionsGetReq) GetId() interface{} {
	return s.Id
}

// HsInviteCommissionsDeleteReq 功能删除请求参数
type HsInviteCommissionsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsInviteCommissionsDeleteReq) GetId() interface{} {
	return s.Ids
}
