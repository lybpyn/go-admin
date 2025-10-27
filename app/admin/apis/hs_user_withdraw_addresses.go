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

type HsUserWithdrawAddresses struct {
	api.Api
}

// GetPage 获取用户虚拟币提现地址表列表
// @Summary 获取用户虚拟币提现地址表列表
// @Description 获取用户虚拟币提现地址表列表
// @Tags 用户虚拟币提现地址表
// @Param userId query string false "所属用户ID"
// @Param address query string false "钱包提现地址"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserWithdrawAddresses}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdraw-addresses [get]
// @Security Bearer
func (e HsUserWithdrawAddresses) GetPage(c *gin.Context) {
    req := dto.HsUserWithdrawAddressesGetPageReq{}
    s := service.HsUserWithdrawAddresses{}
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
	list := make([]models.HsUserWithdrawAddresses, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户虚拟币提现地址表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户虚拟币提现地址表
// @Summary 获取用户虚拟币提现地址表
// @Description 获取用户虚拟币提现地址表
// @Tags 用户虚拟币提现地址表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsUserWithdrawAddresses} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-withdraw-addresses/{id} [get]
// @Security Bearer
func (e HsUserWithdrawAddresses) Get(c *gin.Context) {
	req := dto.HsUserWithdrawAddressesGetReq{}
	s := service.HsUserWithdrawAddresses{}
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
	var object models.HsUserWithdrawAddresses

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户虚拟币提现地址表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户虚拟币提现地址表
// @Summary 创建用户虚拟币提现地址表
// @Description 创建用户虚拟币提现地址表
// @Tags 用户虚拟币提现地址表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsUserWithdrawAddressesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-user-withdraw-addresses [post]
// @Security Bearer
func (e HsUserWithdrawAddresses) Insert(c *gin.Context) {
    req := dto.HsUserWithdrawAddressesInsertReq{}
    s := service.HsUserWithdrawAddresses{}
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
		e.Error(500, err, fmt.Sprintf("创建用户虚拟币提现地址表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户虚拟币提现地址表
// @Summary 修改用户虚拟币提现地址表
// @Description 修改用户虚拟币提现地址表
// @Tags 用户虚拟币提现地址表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsUserWithdrawAddressesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-withdraw-addresses/{id} [put]
// @Security Bearer
func (e HsUserWithdrawAddresses) Update(c *gin.Context) {
    req := dto.HsUserWithdrawAddressesUpdateReq{}
    s := service.HsUserWithdrawAddresses{}
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
		e.Error(500, err, fmt.Sprintf("修改用户虚拟币提现地址表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户虚拟币提现地址表
// @Summary 删除用户虚拟币提现地址表
// @Description 删除用户虚拟币提现地址表
// @Tags 用户虚拟币提现地址表
// @Param data body dto.HsUserWithdrawAddressesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-withdraw-addresses [delete]
// @Security Bearer
func (e HsUserWithdrawAddresses) Delete(c *gin.Context) {
    s := service.HsUserWithdrawAddresses{}
    req := dto.HsUserWithdrawAddressesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户虚拟币提现地址表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
