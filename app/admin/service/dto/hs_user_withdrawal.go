package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type HsUserWithdrawalGetPageReq struct {
	dto.Pagination     `search:"-"`
    WithdrawNo string `form:"withdrawNo"  search:"type:exact;column:withdraw_no;table:hs_user_withdrawal" comment:"提现单号，唯一"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:hs_user_withdrawal" comment:"用户ID"`
    ChannelTxnId string `form:"channelTxnId"  search:"type:exact;column:channel_txn_id;table:hs_user_withdrawal" comment:"通道回执流水号"`
    RequestedAt time.Time `form:"requestedAt"  search:"type:gte;column:requested_at;table:hs_user_withdrawal" comment:"发起时间"`
    ProcessedAt time.Time `form:"processedAt"  search:"type:lte;column:processed_at;table:hs_user_withdrawal" comment:"处理完成时间"`
    HsUserWithdrawalOrder
}

type HsUserWithdrawalOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_withdrawal"`
    WithdrawNo string `form:"withdrawNoOrder"  search:"type:order;column:withdraw_no;table:hs_user_withdrawal"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_withdrawal"`
    Currency string `form:"currencyOrder"  search:"type:order;column:currency;table:hs_user_withdrawal"`
    Amount string `form:"amountOrder"  search:"type:order;column:amount;table:hs_user_withdrawal"`
    Fee string `form:"feeOrder"  search:"type:order;column:fee;table:hs_user_withdrawal"`
    NetAmount string `form:"netAmountOrder"  search:"type:order;column:net_amount;table:hs_user_withdrawal"`
    Method string `form:"methodOrder"  search:"type:order;column:method;table:hs_user_withdrawal"`
    AccountInfo string `form:"accountInfoOrder"  search:"type:order;column:account_info;table:hs_user_withdrawal"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_user_withdrawal"`
    Reason string `form:"reasonOrder"  search:"type:order;column:reason;table:hs_user_withdrawal"`
    ChannelTxnId string `form:"channelTxnIdOrder"  search:"type:order;column:channel_txn_id;table:hs_user_withdrawal"`
    RequestedAt string `form:"requestedAtOrder"  search:"type:order;column:requested_at;table:hs_user_withdrawal"`
    ProcessedAt string `form:"processedAtOrder"  search:"type:order;column:processed_at;table:hs_user_withdrawal"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:hs_user_withdrawal"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:hs_user_withdrawal"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:hs_user_withdrawal"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:hs_user_withdrawal"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:hs_user_withdrawal"`
    
}

func (m *HsUserWithdrawalGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type HsUserWithdrawalInsertReq struct {
    Id int `json:"-" comment:""` // 
    WithdrawNo string `json:"withdrawNo" comment:"提现单号，唯一"`
    UserId string `json:"userId" comment:"用户ID"`
    Currency string `json:"currency" comment:"ISO 4217币种代码，如 USD/CNY"`
    Amount string `json:"amount" comment:"提现金额"`
    Fee string `json:"fee" comment:"提现手续费"`
    NetAmount string `json:"netAmount" comment:"实际出账金额（amount - fee）"`
    Method string `json:"method" comment:"提现方式：bank/paypal/crypto/…"`
    AccountInfo string `json:"accountInfo" comment:"提现账户信息（脱敏）"`
    Status string `json:"status" comment:"状态：pending/review/processing/success/failed/canceled"`
    Reason string `json:"reason" comment:"失败/取消原因"`
    ChannelTxnId string `json:"channelTxnId" comment:"通道回执流水号"`
    RequestedAt time.Time `json:"requestedAt" comment:"发起时间"`
    ProcessedAt time.Time `json:"processedAt" comment:"处理完成时间"`
    common.ControlBy
}

func (s *HsUserWithdrawalInsertReq) Generate(model *models.HsUserWithdrawal)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.WithdrawNo = s.WithdrawNo
    model.UserId = s.UserId
    model.Currency = s.Currency
    model.Amount = s.Amount
    model.Fee = s.Fee
    model.NetAmount = s.NetAmount
    model.Method = s.Method
    model.AccountInfo = s.AccountInfo
    model.Status = s.Status
    model.Reason = s.Reason
    model.ChannelTxnId = s.ChannelTxnId
    model.RequestedAt = s.RequestedAt
    model.ProcessedAt = s.ProcessedAt
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *HsUserWithdrawalInsertReq) GetId() interface{} {
	return s.Id
}

type HsUserWithdrawalUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    WithdrawNo string `json:"withdrawNo" comment:"提现单号，唯一"`
    UserId string `json:"userId" comment:"用户ID"`
    Currency string `json:"currency" comment:"ISO 4217币种代码，如 USD/CNY"`
    Amount string `json:"amount" comment:"提现金额"`
    Fee string `json:"fee" comment:"提现手续费"`
    NetAmount string `json:"netAmount" comment:"实际出账金额（amount - fee）"`
    Method string `json:"method" comment:"提现方式：bank/paypal/crypto/…"`
    AccountInfo string `json:"accountInfo" comment:"提现账户信息（脱敏）"`
    Status string `json:"status" comment:"状态：pending/review/processing/success/failed/canceled"`
    Reason string `json:"reason" comment:"失败/取消原因"`
    ChannelTxnId string `json:"channelTxnId" comment:"通道回执流水号"`
    RequestedAt time.Time `json:"requestedAt" comment:"发起时间"`
    ProcessedAt time.Time `json:"processedAt" comment:"处理完成时间"`
    common.ControlBy
}

func (s *HsUserWithdrawalUpdateReq) Generate(model *models.HsUserWithdrawal)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.WithdrawNo = s.WithdrawNo
    model.UserId = s.UserId
    model.Currency = s.Currency
    model.Amount = s.Amount
    model.Fee = s.Fee
    model.NetAmount = s.NetAmount
    model.Method = s.Method
    model.AccountInfo = s.AccountInfo
    model.Status = s.Status
    model.Reason = s.Reason
    model.ChannelTxnId = s.ChannelTxnId
    model.RequestedAt = s.RequestedAt
    model.ProcessedAt = s.ProcessedAt
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *HsUserWithdrawalUpdateReq) GetId() interface{} {
	return s.Id
}

// HsUserWithdrawalGetReq 功能获取请求参数
type HsUserWithdrawalGetReq struct {
     Id int `uri:"id"`
}
func (s *HsUserWithdrawalGetReq) GetId() interface{} {
	return s.Id
}

// HsUserWithdrawalDeleteReq 功能删除请求参数
type HsUserWithdrawalDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *HsUserWithdrawalDeleteReq) GetId() interface{} {
	return s.Ids
}
