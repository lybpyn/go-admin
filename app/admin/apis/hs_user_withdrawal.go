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

type HsUserWithdrawal struct {
	api.Api
}

// GetPage 获取用户余额提现申请表列表
// @Summary 获取用户余额提现申请表列表
// @Description 获取用户余额提现申请表列表
// @Tags 用户余额提现申请表
// @Param withdrawNo query string false "提现单号，唯一"
// @Param userId query string false "用户ID"
// @Param channelTxnId query string false "通道回执流水号"
// @Param requestedAt query time.Time false "发起时间"
// @Param processedAt query time.Time false "处理完成时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserWithdrawal}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdrawal [get]
// @Security Bearer
func (e HsUserWithdrawal) GetPage(c *gin.Context) {
    req := dto.HsUserWithdrawalGetPageReq{}
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.HsUserWithdrawal, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
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
// @Success 200 {object} response.Response{data=models.HsUserWithdrawal} "{"code": 200, "data": [...]}"
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
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
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
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-withdrawal/{id} [put]
// @Security Bearer
func (e HsUserWithdrawal) Update(c *gin.Context) {
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
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-withdrawal [delete]
// @Security Bearer
func (e HsUserWithdrawal) Delete(c *gin.Context) {
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
	e.OK( req.GetId(), "删除成功")
}

// GetAvailable 获取可接单的提现订单列表
// @Summary 获取可接单的提现订单列表
// @Description 获取所有未被接单的提现订单（is_claimed=0, status=pending/review）
// @Tags 用户余额提现申请表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserWithdrawal}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdrawal/available [get]
// @Security Bearer
func (e HsUserWithdrawal) GetAvailable(c *gin.Context) {
	req := dto.HsUserWithdrawalAvailableReq{}
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

	list := make([]models.HsUserWithdrawal, 0)
	var count int64

	err = s.GetAvailable(&req, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取可接单列表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// GetMyOrders 获取我的处理订单列表
// @Summary 获取我的处理订单列表
// @Description 获取当前管理员接单处理的订单列表
// @Tags 用户余额提现申请表
// @Param status query string false "订单状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserWithdrawal}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdrawal/my-orders [get]
// @Security Bearer
func (e HsUserWithdrawal) GetMyOrders(c *gin.Context) {
	req := dto.HsUserWithdrawalMyOrdersReq{}
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

	handlerId := user.GetUserId(c)
	list := make([]models.HsUserWithdrawal, 0)
	var count int64

	err = s.GetMyOrders(&req, handlerId, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取我的订单列表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Claim 接单
// @Summary 接单
// @Description 管理员接单处理提现订单
// @Tags 用户余额提现申请表
// @Param id path int true "订单ID"
// @Success 200 {object} response.Response "{"code": 200, "message": "接单成功"}"
// @Router /api/v1/hs-user-withdrawal/{id}/claim [post]
// @Security Bearer
func (e HsUserWithdrawal) Claim(c *gin.Context) {
	req := dto.HsUserWithdrawalClaimReq{}
	s := service.HsUserWithdrawal{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Uri).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	handlerId := user.GetUserId(c)
	handlerName := user.GetUserName(c)

	err = s.Claim(req.Id, handlerId, handlerName)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "接单成功")
}

// Release 释放订单（取消接单）
// @Summary 释放订单
// @Description 管理员取消接单，释放订单
// @Tags 用户余额提现申请表
// @Param id path int true "订单ID"
// @Success 200 {object} response.Response "{"code": 200, "message": "释放订单成功"}"
// @Router /api/v1/hs-user-withdrawal/{id}/release [post]
// @Security Bearer
func (e HsUserWithdrawal) Release(c *gin.Context) {
	req := dto.HsUserWithdrawalReleaseReq{}
	s := service.HsUserWithdrawal{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Uri).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	handlerId := user.GetUserId(c)

	err = s.Release(req.Id, handlerId)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "释放订单成功")
}

