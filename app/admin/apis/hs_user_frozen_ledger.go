package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type HsUserFrozenLedger struct {
	api.Api
}

// GetPage 获取用户冻结余额流水列表
// @Summary 获取用户冻结余额流水列表
// @Description 获取用户冻结余额流水列表
// @Tags 用户冻结余额流水
// @Param userId query string false "用户ID"
// @Param currencyCode query string false "币种代码，如 USD/CNY/USDT"
// @Param direction query string false "1=冻结增加，-1=冻结减少(解冻)"
// @Param amount query string false "冻结或解冻金额"
// @Param frozenBefore query string false "变动前冻结余额"
// @Param frozenAfter query string false "变动后冻结余额"
// @Param bizType query string false "业务类型：invite_commissions/order_rebate等"
// @Param bizId query string false "业务单号"
// @Param idempotencyKey query string false "幂等键"
// @Param remark query string false "备注"
// @Param status query string false "1=已冻结或解冻，0=待处理，-1=冲正"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.HsUserFrozenLedger}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-frozen-ledger [get]
// @Security Bearer
func (e HsUserFrozenLedger) GetPage(c *gin.Context) {
    req := dto.HsUserFrozenLedgerGetPageReq{}
    s := service.HsUserFrozenLedger{}
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.HsUserFrozenLedger, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户冻结余额流水失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户冻结余额流水
// @Summary 获取用户冻结余额流水
// @Description 获取用户冻结余额流水
// @Tags 用户冻结余额流水
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.HsUserFrozenLedger} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-frozen-ledger/{id} [get]
// @Security Bearer
func (e HsUserFrozenLedger) Get(c *gin.Context) {
	req := dto.HsUserFrozenLedgerGetReq{}
	s := service.HsUserFrozenLedger{}
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
	var object models.HsUserFrozenLedger

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户冻结余额流水失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户冻结余额流水
// @Summary 创建用户冻结余额流水
// @Description 创建用户冻结余额流水
// @Tags 用户冻结余额流水
// @Accept application/json
// @Product application/json
// @Param data body dto.HsUserFrozenLedgerInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-user-frozen-ledger [post]
// @Security Bearer
func (e HsUserFrozenLedger) Insert(c *gin.Context) {
    req := dto.HsUserFrozenLedgerInsertReq{}
    s := service.HsUserFrozenLedger{}
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
		e.Error(500, err, fmt.Sprintf("创建用户冻结余额流水失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户冻结余额流水
// @Summary 修改用户冻结余额流水
// @Description 修改用户冻结余额流水
// @Tags 用户冻结余额流水
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsUserFrozenLedgerUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-frozen-ledger/{id} [put]
// @Security Bearer
func (e HsUserFrozenLedger) Update(c *gin.Context) {
    req := dto.HsUserFrozenLedgerUpdateReq{}
    s := service.HsUserFrozenLedger{}
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
		e.Error(500, err, fmt.Sprintf("修改用户冻结余额流水失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户冻结余额流水
// @Summary 删除用户冻结余额流水
// @Description 删除用户冻结余额流水
// @Tags 用户冻结余额流水
// @Param data body dto.HsUserFrozenLedgerDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-frozen-ledger [delete]
// @Security Bearer
func (e HsUserFrozenLedger) Delete(c *gin.Context) {
    s := service.HsUserFrozenLedger{}
    req := dto.HsUserFrozenLedgerDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户冻结余额流水失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
