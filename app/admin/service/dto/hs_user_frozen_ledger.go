package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserFrozenLedgerGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_frozen_ledger" comment:""`
    CurrencyCode string `form:"currencyCode"  search:"type:exact;column:currency_code;table:hs_user_frozen_ledger" comment:"币种代码，如 USD/CNY/USDT"`
    Direction string `form:"direction"  search:"type:exact;column:direction;table:hs_user_frozen_ledger" comment:"1=冻结增加，-1=冻结减少(解冻)"`
    Amount string `form:"amount"  search:"type:exact;column:amount;table:hs_user_frozen_ledger" comment:"冻结或解冻金额"`
    FrozenBefore string `form:"frozenBefore"  search:"type:exact;column:frozen_before;table:hs_user_frozen_ledger" comment:"变动前冻结余额"`
    FrozenAfter string `form:"frozenAfter"  search:"type:exact;column:frozen_after;table:hs_user_frozen_ledger" comment:"变动后冻结余额"`
    BizType string `form:"bizType"  search:"type:exact;column:biz_type;table:hs_user_frozen_ledger" comment:"业务类型：invite_commissions/order_rebate等"`
    BizId string `form:"bizId"  search:"type:exact;column:biz_id;table:hs_user_frozen_ledger" comment:"业务单号"`
    IdempotencyKey string `form:"idempotencyKey"  search:"type:exact;column:idempotency_key;table:hs_user_frozen_ledger" comment:"幂等键"`
    Remark string `form:"remark"  search:"type:exact;column:remark;table:hs_user_frozen_ledger" comment:""`
    Status string `form:"status"  search:"type:exact;column:status;table:hs_user_frozen_ledger" comment:"1=已冻结或解冻，0=待处理，-1=冲正"`
    HsUserFrozenLedgerOrder
}

type HsUserFrozenLedgerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_frozen_ledger"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_frozen_ledger"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:hs_user_frozen_ledger"`
    Direction string `form:"directionOrder"  search:"type:order;column:direction;table:hs_user_frozen_ledger"`
    Amount string `form:"amountOrder"  search:"type:order;column:amount;table:hs_user_frozen_ledger"`
    FrozenBefore string `form:"frozenBeforeOrder"  search:"type:order;column:frozen_before;table:hs_user_frozen_ledger"`
    FrozenAfter string `form:"frozenAfterOrder"  search:"type:order;column:frozen_after;table:hs_user_frozen_ledger"`
    BizType string `form:"bizTypeOrder"  search:"type:order;column:biz_type;table:hs_user_frozen_ledger"`
    BizId string `form:"bizIdOrder"  search:"type:order;column:biz_id;table:hs_user_frozen_ledger"`
    IdempotencyKey string `form:"idempotencyKeyOrder"  search:"type:order;column:idempotency_key;table:hs_user_frozen_ledger"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:hs_user_frozen_ledger"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_user_frozen_ledger"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_frozen_ledger"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_frozen_ledger"`
    
}

func (m *HsUserFrozenLedgerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserFrozenLedgerInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:""`
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    Direction string `json:"direction" comment:"1=冻结增加，-1=冻结减少(解冻)"`
    Amount string `json:"amount" comment:"冻结或解冻金额"`
    FrozenBefore string `json:"frozenBefore" comment:"变动前冻结余额"`
    FrozenAfter string `json:"frozenAfter" comment:"变动后冻结余额"`
    BizType string `json:"bizType" comment:"业务类型：invite_commissions/order_rebate等"`
    BizId string `json:"bizId" comment:"业务单号"`
    IdempotencyKey string `json:"idempotencyKey" comment:"幂等键"`
    Remark string `json:"remark" comment:""`
    Status string `json:"status" comment:"1=已冻结或解冻，0=待处理，-1=冲正"`
    common.ControlBy
}

func (s *HsUserFrozenLedgerInsertReq) Generate(model *models.HsUserFrozenLedger)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.CurrencyCode = s.CurrencyCode
    model.Direction = s.Direction
    model.Amount = s.Amount
    model.FrozenBefore = s.FrozenBefore
    model.FrozenAfter = s.FrozenAfter
    model.BizType = s.BizType
    model.BizId = s.BizId
    model.IdempotencyKey = s.IdempotencyKey
    model.Remark = s.Remark
    model.Status = s.Status
}

func (s *HsUserFrozenLedgerInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserFrozenLedgerUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:""`
    CurrencyCode string `json:"currencyCode" comment:"币种代码，如 USD/CNY/USDT"`
    Direction string `json:"direction" comment:"1=冻结增加，-1=冻结减少(解冻)"`
    Amount string `json:"amount" comment:"冻结或解冻金额"`
    FrozenBefore string `json:"frozenBefore" comment:"变动前冻结余额"`
    FrozenAfter string `json:"frozenAfter" comment:"变动后冻结余额"`
    BizType string `json:"bizType" comment:"业务类型：invite_commissions/order_rebate等"`
    BizId string `json:"bizId" comment:"业务单号"`
    IdempotencyKey string `json:"idempotencyKey" comment:"幂等键"`
    Remark string `json:"remark" comment:""`
    Status string `json:"status" comment:"1=已冻结或解冻，0=待处理，-1=冲正"`
    common.ControlBy
}

func (s *HsUserFrozenLedgerUpdateReq) Generate(model *models.HsUserFrozenLedger)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.CurrencyCode = s.CurrencyCode
    model.Direction = s.Direction
    model.Amount = s.Amount
    model.FrozenBefore = s.FrozenBefore
    model.FrozenAfter = s.FrozenAfter
    model.BizType = s.BizType
    model.BizId = s.BizId
    model.IdempotencyKey = s.IdempotencyKey
    model.Remark = s.Remark
    model.Status = s.Status
}

func (s *HsUserFrozenLedgerUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserFrozenLedgerGetReq 功能获取请求参数
type HsUserFrozenLedgerGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserFrozenLedgerGetReq) GetId() interface{} {
	return s.Id
}

// HsUserFrozenLedgerDeleteReq 功能删除请求参数
type HsUserFrozenLedgerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserFrozenLedgerDeleteReq) GetId() interface{} {
	return s.Ids
}
