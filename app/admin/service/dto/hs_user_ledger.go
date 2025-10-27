package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserLedgerGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_ledger" comment:""`
    BizId string `form:"bizId"  search:"type:exact;column:biz_id;table:hs_user_ledger" comment:"业务单号：例如订单号/提现单号"`
    HsUserLedgerOrder
}

type HsUserLedgerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_ledger"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_ledger"`
    Currency string `form:"currencyOrder"  search:"type:order;column:currency;table:hs_user_ledger"`
    Direction string `form:"directionOrder"  search:"type:order;column:direction;table:hs_user_ledger"`
    Amount string `form:"amountOrder"  search:"type:order;column:amount;table:hs_user_ledger"`
    BalanceBefore string `form:"balanceBeforeOrder"  search:"type:order;column:balance_before;table:hs_user_ledger"`
    BalanceAfter string `form:"balanceAfterOrder"  search:"type:order;column:balance_after;table:hs_user_ledger"`
    BizType string `form:"bizTypeOrder"  search:"type:order;column:biz_type;table:hs_user_ledger"`
    BizId string `form:"bizIdOrder"  search:"type:order;column:biz_id;table:hs_user_ledger"`
    IdempotencyKey string `form:"idempotencyKeyOrder"  search:"type:order;column:idempotency_key;table:hs_user_ledger"`
    RefTable string `form:"refTableOrder"  search:"type:order;column:ref_table;table:hs_user_ledger"`
    RefId string `form:"refIdOrder"  search:"type:order;column:ref_id;table:hs_user_ledger"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:hs_user_ledger"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_user_ledger"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_ledger"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_ledger"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_ledger"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_ledger"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_ledger"`
    
}

func (m *HsUserLedgerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserLedgerInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:""`
    Currency string `json:"currency" comment:"ISO 4217，例如 USD/CNY"`
    Direction string `json:"direction" comment:"1=入账(credit), -1=出账(debit)"`
    Amount string `json:"amount" comment:"本次发生额，>0"`
    BalanceBefore string `json:"balanceBefore" comment:"记账前余额"`
    BalanceAfter string `json:"balanceAfter" comment:"记账后余额"`
    BizType string `json:"bizType" comment:"业务类型：order_settlement/withdraw/withdraw_fee/withdraw_reversal/manual_adjust_* 等"`
    BizId string `json:"bizId" comment:"业务单号：例如订单号/提现单号"`
    IdempotencyKey string `json:"idempotencyKey" comment:"用于幂等控制：如 ORDER_SETTLED:{order_no}"`
    RefTable string `json:"refTable" comment:"可选：引用表名"`
    RefId string `json:"refId" comment:"可选：引用ID"`
    Remark string `json:"remark" comment:""`
    Status string `json:"status" comment:"1=已入账，0=待入账，-1=冲正"`
    common.ControlBy
}

func (s *HsUserLedgerInsertReq) Generate(model *models.HsUserLedger)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.Currency = s.Currency
    model.Direction = s.Direction
    model.Amount = s.Amount
    model.BalanceBefore = s.BalanceBefore
    model.BalanceAfter = s.BalanceAfter
    model.BizType = s.BizType
    model.BizId = s.BizId
    model.IdempotencyKey = s.IdempotencyKey
    model.RefTable = s.RefTable
    model.RefId = s.RefId
    model.Remark = s.Remark
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserLedgerInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserLedgerUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:""`
    Currency string `json:"currency" comment:"ISO 4217，例如 USD/CNY"`
    Direction string `json:"direction" comment:"1=入账(credit), -1=出账(debit)"`
    Amount string `json:"amount" comment:"本次发生额，>0"`
    BalanceBefore string `json:"balanceBefore" comment:"记账前余额"`
    BalanceAfter string `json:"balanceAfter" comment:"记账后余额"`
    BizType string `json:"bizType" comment:"业务类型：order_settlement/withdraw/withdraw_fee/withdraw_reversal/manual_adjust_* 等"`
    BizId string `json:"bizId" comment:"业务单号：例如订单号/提现单号"`
    IdempotencyKey string `json:"idempotencyKey" comment:"用于幂等控制：如 ORDER_SETTLED:{order_no}"`
    RefTable string `json:"refTable" comment:"可选：引用表名"`
    RefId string `json:"refId" comment:"可选：引用ID"`
    Remark string `json:"remark" comment:""`
    Status string `json:"status" comment:"1=已入账，0=待入账，-1=冲正"`
    common.ControlBy
}

func (s *HsUserLedgerUpdateReq) Generate(model *models.HsUserLedger)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.Currency = s.Currency
    model.Direction = s.Direction
    model.Amount = s.Amount
    model.BalanceBefore = s.BalanceBefore
    model.BalanceAfter = s.BalanceAfter
    model.BizType = s.BizType
    model.BizId = s.BizId
    model.IdempotencyKey = s.IdempotencyKey
    model.RefTable = s.RefTable
    model.RefId = s.RefId
    model.Remark = s.Remark
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserLedgerUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserLedgerGetReq 功能获取请求参数
type HsUserLedgerGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserLedgerGetReq) GetId() interface{} {
	return s.Id
}

// HsUserLedgerDeleteReq 功能删除请求参数
type HsUserLedgerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserLedgerDeleteReq) GetId() interface{} {
	return s.Ids
}
