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

type HsConfigWithdrawLimit struct {
	api.Api
}

// GetPage 获取用户提现限额配置表列表
// @Summary 获取用户提现限额配置表列表
// @Description 获取用户提现限额配置表列表
// @Tags 用户提现限额配置表
// @Param currencyCode query string false "币种代码，如 USD/CNY/USDT"
// @Param singleLimit query string false "单笔提现限额"
// @Param dailyLimitAmount query string false "每日累计提现金额上限"
// @Param dailyLimitCount query string false "每日提现次数上限"
// @Param isActive query string false "是否启用：1=启用，0=禁用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsConfigWithdrawLimit}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-limit [get]
// @Security Bearer
func (e HsConfigWithdrawLimit) GetPage(c *gin.Context) {
    req := dto.HsConfigWithdrawLimitGetPageReq{}
    s := service.HsConfigWithdrawLimit{}
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
	list := make([]models.HsConfigWithdrawLimit, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户提现限额配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户提现限额配置表
// @Summary 获取用户提现限额配置表
// @Description 获取用户提现限额配置表
// @Tags 用户提现限额配置表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsConfigWithdrawLimit} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-limit/{id} [get]
// @Security Bearer
func (e HsConfigWithdrawLimit) Get(c *gin.Context) {
	req := dto.HsConfigWithdrawLimitGetReq{}
	s := service.HsConfigWithdrawLimit{}
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
	var object models.HsConfigWithdrawLimit

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户提现限额配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户提现限额配置表
// @Summary 创建用户提现限额配置表
// @Description 创建用户提现限额配置表
// @Tags 用户提现限额配置表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigWithdrawLimitInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-withdraw-limit [post]
// @Security Bearer
func (e HsConfigWithdrawLimit) Insert(c *gin.Context) {
    req := dto.HsConfigWithdrawLimitInsertReq{}
    s := service.HsConfigWithdrawLimit{}
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
		e.Error(500, err, fmt.Sprintf("创建用户提现限额配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户提现限额配置表
// @Summary 修改用户提现限额配置表
// @Description 修改用户提现限额配置表
// @Tags 用户提现限额配置表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigWithdrawLimitUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-withdraw-limit/{id} [put]
// @Security Bearer
func (e HsConfigWithdrawLimit) Update(c *gin.Context) {
    req := dto.HsConfigWithdrawLimitUpdateReq{}
    s := service.HsConfigWithdrawLimit{}
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
		e.Error(500, err, fmt.Sprintf("修改用户提现限额配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户提现限额配置表
// @Summary 删除用户提现限额配置表
// @Description 删除用户提现限额配置表
// @Tags 用户提现限额配置表
// @Param data body dto.HsConfigWithdrawLimitDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-withdraw-limit [delete]
// @Security Bearer
func (e HsConfigWithdrawLimit) Delete(c *gin.Context) {
    s := service.HsConfigWithdrawLimit{}
    req := dto.HsConfigWithdrawLimitDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户提现限额配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
