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
    CurrencyCode string `form:"currencyCode"  search:"type:exact;column:currency_code;table:hs_user_withdrawal" comment:"ISO 4217币种代码，如 USD/CNY"`
    Amount string `form:"amount"  search:"type:exact;column:amount;table:hs_user_withdrawal" comment:"提现金额"`
    Fee string `form:"fee"  search:"type:exact;column:fee;table:hs_user_withdrawal" comment:"提现手续费"`
    NetAmount string `form:"netAmount"  search:"type:exact;column:net_amount;table:hs_user_withdrawal" comment:"实际出账金额（amount - fee）"`
    Method string `form:"method"  search:"type:exact;column:method;table:hs_user_withdrawal" comment:"提现方式：bank/crypto"`
    AccountInfo string `form:"accountInfo"  search:"type:exact;column:account_info;table:hs_user_withdrawal" comment:"提现账户信息（脱敏）"`
    Status string `form:"status"  search:"type:exact;column:status;table:hs_user_withdrawal" comment:"状态：pending/review/processing/success/failed/canceled"`
    HandlerId string `form:"handlerId"  search:"type:exact;column:handler_id;table:hs_user_withdrawal" comment:"接单管理员ID（关联sys_user.user_id）"`
    HandlerName string `form:"handlerName"  search:"type:exact;column:handler_name;table:hs_user_withdrawal" comment:"接单管理员名称（冗余字段）"`
    ClaimedAt time.Time `form:"claimedAt"  search:"type:exact;column:claimed_at;table:hs_user_withdrawal" comment:"接单时间"`
    IsClaimed string `form:"isClaimed"  search:"type:exact;column:is_claimed;table:hs_user_withdrawal" comment:"是否已接单：0=未接单(可接单), 1=已接单(已锁定)"`
    Reason string `form:"reason"  search:"type:exact;column:reason;table:hs_user_withdrawal" comment:"失败/取消原因"`
    ChannelTxnId string `form:"channelTxnId"  search:"type:exact;column:channel_txn_id;table:hs_user_withdrawal" comment:"通道回执流水号"`
    RequestedAt time.Time `form:"requestedAt"  search:"type:exact;column:requested_at;table:hs_user_withdrawal" comment:"发起时间"`
    ProcessedAt time.Time `form:"processedAt"  search:"type:exact;column:processed_at;table:hs_user_withdrawal" comment:"处理完成时间"`
    HsUserWithdrawalOrder
}

type HsUserWithdrawalOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:hs_user_withdrawal"`
    WithdrawNo string `form:"withdrawNoOrder"  search:"type:order;column:withdraw_no;table:hs_user_withdrawal"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:hs_user_withdrawal"`
    CurrencyCode string `form:"currencyCodeOrder"  search:"type:order;column:currency_code;table:hs_user_withdrawal"`
    Amount string `form:"amountOrder"  search:"type:order;column:amount;table:hs_user_withdrawal"`
    Fee string `form:"feeOrder"  search:"type:order;column:fee;table:hs_user_withdrawal"`
    NetAmount string `form:"netAmountOrder"  search:"type:order;column:net_amount;table:hs_user_withdrawal"`
    Method string `form:"methodOrder"  search:"type:order;column:method;table:hs_user_withdrawal"`
    AccountInfo string `form:"accountInfoOrder"  search:"type:order;column:account_info;table:hs_user_withdrawal"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:hs_user_withdrawal"`
    HandlerId string `form:"handlerIdOrder"  search:"type:order;column:handler_id;table:hs_user_withdrawal"`
    HandlerName string `form:"handlerNameOrder"  search:"type:order;column:handler_name;table:hs_user_withdrawal"`
    ClaimedAt string `form:"claimedAtOrder"  search:"type:order;column:claimed_at;table:hs_user_withdrawal"`
    IsClaimed string `form:"isClaimedOrder"  search:"type:order;column:is_claimed;table:hs_user_withdrawal"`
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
    CurrencyCode string `json:"currencyCode" comment:"ISO 4217币种代码，如 USD/CNY"`
    Amount string `json:"amount" comment:"提现金额"`
    Fee string `json:"fee" comment:"提现手续费"`
    NetAmount string `json:"netAmount" comment:"实际出账金额（amount - fee）"`
    Method string `json:"method" comment:"提现方式：bank/crypto"`
    AccountInfo string `json:"accountInfo" comment:"提现账户信息（脱敏）"`
    Status string `json:"status" comment:"状态：pending/review/processing/success/failed/canceled"`
    HandlerId string `json:"handlerId" comment:"接单管理员ID（关联sys_user.user_id）"`
    HandlerName string `json:"handlerName" comment:"接单管理员名称（冗余字段）"`
    ClaimedAt time.Time `json:"claimedAt" comment:"接单时间"`
    IsClaimed string `json:"isClaimed" comment:"是否已接单：0=未接单(可接单), 1=已接单(已锁定)"`
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
    model.CurrencyCode = s.CurrencyCode
    model.Amount = s.Amount
    model.Fee = s.Fee
    model.NetAmount = s.NetAmount
    model.Method = s.Method
    model.AccountInfo = s.AccountInfo
    model.Status = s.Status
    model.HandlerId = s.HandlerId
    model.HandlerName = s.HandlerName
    model.ClaimedAt = s.ClaimedAt
    model.IsClaimed = s.IsClaimed
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
    CurrencyCode string `json:"currencyCode" comment:"ISO 4217币种代码，如 USD/CNY"`
    Amount string `json:"amount" comment:"提现金额"`
    Fee string `json:"fee" comment:"提现手续费"`
    NetAmount string `json:"netAmount" comment:"实际出账金额（amount - fee）"`
    Method string `json:"method" comment:"提现方式：bank/crypto"`
    AccountInfo string `json:"accountInfo" comment:"提现账户信息（脱敏）"`
    Status string `json:"status" comment:"状态：pending/review/processing/success/failed/canceled"`
    HandlerId string `json:"handlerId" comment:"接单管理员ID（关联sys_user.user_id）"`
    HandlerName string `json:"handlerName" comment:"接单管理员名称（冗余字段）"`
    ClaimedAt time.Time `json:"claimedAt" comment:"接单时间"`
    IsClaimed string `json:"isClaimed" comment:"是否已接单：0=未接单(可接单), 1=已接单(已锁定)"`
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
    model.CurrencyCode = s.CurrencyCode
    model.Amount = s.Amount
    model.Fee = s.Fee
    model.NetAmount = s.NetAmount
    model.Method = s.Method
    model.AccountInfo = s.AccountInfo
    model.Status = s.Status
    model.HandlerId = s.HandlerId
    model.HandlerName = s.HandlerName
    model.ClaimedAt = s.ClaimedAt
    model.IsClaimed = s.IsClaimed
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

// HsUserWithdrawalWithStats 带统计信息的提现记录响应
type HsUserWithdrawalWithStats struct {
	models.HsUserWithdrawal
	TotalWithdrawn string `json:"totalWithdrawn"` // 累计提现金额
	TotalSales     string `json:"totalSales"`     // 累计售卖金额
	WithdrawCount  int64  `json:"withdrawCount"`  // 提现笔数
	OrderCount     int64  `json:"orderCount"`     // 订单笔数
}
