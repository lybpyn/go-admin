package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type HsUserWithdrawal struct {
	api.Api
}

// GetPage 获取用户余额提现申请表列表
// @Summary 获取用户余额提现申请表列表
// @Description 获取用户余额提现申请表列表
// @Tags 用户余额提现申请表
// @Param withdrawNo query string false "提现单号，唯一"
// @Param userId query string false "用户ID"
// @Param currencyCode query string false "ISO 4217币种代码，如 USD/CNY"
// @Param amount query string false "提现金额"
// @Param fee query string false "提现手续费"
// @Param netAmount query string false "实际出账金额（amount - fee）"
// @Param method query string false "提现方式：bank/crypto"
// @Param accountInfo query string false "提现账户信息（脱敏）"
// @Param status query string false "状态：pending/review/processing/success/failed/canceled"
// @Param handlerId query string false "接单管理员ID（关联sys_user.user_id）"
// @Param handlerName query string false "接单管理员名称（冗余字段）"
// @Param claimedAt query time.Time false "接单时间"
// @Param isClaimed query string false "是否已接单：0=未接单(可接单), 1=已接单(已锁定)"
// @Param reason query string false "失败/取消原因"
// @Param channelTxnId query string false "通道回执流水号"
// @Param requestedAt query time.Time false "发起时间"
// @Param processedAt query time.Time false "处理完成时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]dto.HsUserWithdrawalWithStats}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdrawal [get]
// @Security Bearer
func (e HsUserWithdrawal) GetPage(c *gin.Context) {
    req := dto.HsUserWithdrawalGetPageReq{}
    s := service.HsUserWithdrawal{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req, binding.Form, binding.Query).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]dto.HsUserWithdrawalWithStats, 0)
	var count int64

	err = s.GetPageWithStats(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户余额提现申请表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户余额提现申请表
// @Summary 获取用户余额提现申请表
// @Description 获取用户余额提现申请表
// @Tags 用户余额提现申请表
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.HsUserWithdrawal} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdrawal/{id} [get]
// @Security Bearer
func (e HsUserWithdrawal) Get(c *gin.Context) {
	req := dto.HsUserWithdrawalGetReq{}
	s := service.HsUserWithdrawal{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.HsUserWithdrawal

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户余额提现申请表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户余额提现申请表
// @Summary 创建用户余额提现申请表
// @Description 创建用户余额提现申请表
// @Tags 用户余额提现申请表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsUserWithdrawalInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-user-withdrawal [post]
// @Security Bearer
func (e HsUserWithdrawal) Insert(c *gin.Context) {
    req := dto.HsUserWithdrawalInsertReq{}
    s := service.HsUserWithdrawal{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建用户余额提现申请表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户余额提现申请表
// @Summary 修改用户余额提现申请表
// @Description 修改用户余额提现申请表
// @Tags 用户余额提现申请表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsUserWithdrawalUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-withdrawal/{id} [put]
// @Security Bearer
func (e HsUserWithdrawal) Update(c *gin.Context) {
	return
    req := dto.HsUserWithdrawalUpdateReq{}
    s := service.HsUserWithdrawal{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改用户余额提现申请表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户余额提现申请表
// @Summary 删除用户余额提现申请表
// @Description 删除用户余额提现申请表
// @Tags 用户余额提现申请表
// @Param data body dto.HsUserWithdrawalDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-withdrawal [delete]
// @Security Bearer
func (e HsUserWithdrawal) Delete(c *gin.Context) {
	return
	s := service.HsUserWithdrawal{}
	req := dto.HsUserWithdrawalDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除用户余额提现申请表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Approve 审核通过提现申请
// @Summary 审核通过提现申请
// @Description 审核通过提现申请，bank方式将自动调用PandaPay代付接口转账，crypto方式不调用
// @Tags 用户余额提现申请表
// @Accept application/json
// @Product application/json
// @Param id path int true "提现记录ID"
// @Param data body dto.HsUserWithdrawalApproveReq true "审核请求"
// @Success 200 {object} models.Response	"{"code": 200, "message": "审核成功"}"
// @Router /api/v1/hs-user-withdrawal/{id}/approve [post]
// @Security Bearer
func (e HsUserWithdrawal) Approve(c *gin.Context) {
	req := dto.HsUserWithdrawalApproveReq{}
	s := service.HsUserWithdrawal{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Approve(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("审核提现申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "审核通过成功")
}

// Reject 拒绝提现申请
// @Summary 拒绝提现申请
// @Description 拒绝提现申请，需要提供拒绝原因
// @Tags 用户余额提现申请表
// @Accept application/json
// @Product application/json
// @Param id path int true "提现记录ID"
// @Param data body dto.HsUserWithdrawalRejectReq true "拒绝请求"
// @Success 200 {object} models.Response	"{"code": 200, "message": "拒绝成功"}"
// @Router /api/v1/hs-user-withdrawal/{id}/reject [post]
// @Security Bearer
func (e HsUserWithdrawal) Reject(c *gin.Context) {
	req := dto.HsUserWithdrawalRejectReq{}
	s := service.HsUserWithdrawal{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Reject(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("拒绝提现申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "拒绝成功")
}

// ManualTransfer 手动处理提现转账
// @Summary 手动处理提现转账
// @Description 手动完成提现转账，记录转账截图和结果
// @Tags 提现管理
// @Accept json
// @Produce json
// @Param id path int true "提现记录ID"
// @Param data body dto.HsUserWithdrawalManualTransferReq true "手动转账信息"
// @Success 200 {object} response.Response{data=int} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdrawal/{id}/manual-transfer [post]
// @Security Bearer
func (e HsUserWithdrawal) ManualTransfer(c *gin.Context) {
	req := dto.HsUserWithdrawalManualTransferReq{}
	s := service.HsUserWithdrawal{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.ManualTransfer(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("手动处理提现转账失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "处理成功")
}
