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

type HsConfigFrozenLimit struct {
	api.Api
}

// GetPage 获取用户冻结金额限制配置表列表
// @Summary 获取用户冻结金额限制配置表列表
// @Description 获取用户冻结金额限制配置表列表
// @Tags 用户冻结金额限制配置表
// @Param currencyCode query string false "币种代码，如 USD/CNY/USDT"
// @Param frozenLimitAmount query string false "可冻结金额上限 / 提现限制金额"
// @Param isActive query string false "是否启用：1=启用，0=禁用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.HsConfigFrozenLimit}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-frozen-limit [get]
// @Security Bearer
func (e HsConfigFrozenLimit) GetPage(c *gin.Context) {
    req := dto.HsConfigFrozenLimitGetPageReq{}
    s := service.HsConfigFrozenLimit{}
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
	list := make([]models.HsConfigFrozenLimit, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户冻结金额限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户冻结金额限制配置表
// @Summary 获取用户冻结金额限制配置表
// @Description 获取用户冻结金额限制配置表
// @Tags 用户冻结金额限制配置表
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.HsConfigFrozenLimit} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-frozen-limit/{id} [get]
// @Security Bearer
func (e HsConfigFrozenLimit) Get(c *gin.Context) {
	req := dto.HsConfigFrozenLimitGetReq{}
	s := service.HsConfigFrozenLimit{}
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
	var object models.HsConfigFrozenLimit

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户冻结金额限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户冻结金额限制配置表
// @Summary 创建用户冻结金额限制配置表
// @Description 创建用户冻结金额限制配置表
// @Tags 用户冻结金额限制配置表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigFrozenLimitInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-frozen-limit [post]
// @Security Bearer
func (e HsConfigFrozenLimit) Insert(c *gin.Context) {
    req := dto.HsConfigFrozenLimitInsertReq{}
    s := service.HsConfigFrozenLimit{}
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
		e.Error(500, err, fmt.Sprintf("创建用户冻结金额限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户冻结金额限制配置表
// @Summary 修改用户冻结金额限制配置表
// @Description 修改用户冻结金额限制配置表
// @Tags 用户冻结金额限制配置表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigFrozenLimitUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-frozen-limit/{id} [put]
// @Security Bearer
func (e HsConfigFrozenLimit) Update(c *gin.Context) {
    req := dto.HsConfigFrozenLimitUpdateReq{}
    s := service.HsConfigFrozenLimit{}
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
		e.Error(500, err, fmt.Sprintf("修改用户冻结金额限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户冻结金额限制配置表
// @Summary 删除用户冻结金额限制配置表
// @Description 删除用户冻结金额限制配置表
// @Tags 用户冻结金额限制配置表
// @Param data body dto.HsConfigFrozenLimitDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-frozen-limit [delete]
// @Security Bearer
func (e HsConfigFrozenLimit) Delete(c *gin.Context) {
    s := service.HsConfigFrozenLimit{}
    req := dto.HsConfigFrozenLimitDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户冻结金额限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
