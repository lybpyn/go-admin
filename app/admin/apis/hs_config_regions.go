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

type HsConfigRegions struct {
	api.Api
}

// GetPage 获取系统地区表（注册地区选择）列表
// @Summary 获取系统地区表（注册地区选择）列表
// @Description 获取系统地区表（注册地区选择）列表
// @Tags 系统地区表（注册地区选择）
// @Param name query string false "地区名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsConfigRegions}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-regions [get]
// @Security Bearer
func (e HsConfigRegions) GetPage(c *gin.Context) {
    req := dto.HsConfigRegionsGetPageReq{}
    s := service.HsConfigRegions{}
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
	list := make([]models.HsConfigRegions, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取系统地区表（注册地区选择）失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取系统地区表（注册地区选择）
// @Summary 获取系统地区表（注册地区选择）
// @Description 获取系统地区表（注册地区选择）
// @Tags 系统地区表（注册地区选择）
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsConfigRegions} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-regions/{id} [get]
// @Security Bearer
func (e HsConfigRegions) Get(c *gin.Context) {
	req := dto.HsConfigRegionsGetReq{}
	s := service.HsConfigRegions{}
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
	var object models.HsConfigRegions

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取系统地区表（注册地区选择）失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建系统地区表（注册地区选择）
// @Summary 创建系统地区表（注册地区选择）
// @Description 创建系统地区表（注册地区选择）
// @Tags 系统地区表（注册地区选择）
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigRegionsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-regions [post]
// @Security Bearer
func (e HsConfigRegions) Insert(c *gin.Context) {
    req := dto.HsConfigRegionsInsertReq{}
    s := service.HsConfigRegions{}
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
		e.Error(500, err, fmt.Sprintf("创建系统地区表（注册地区选择）失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改系统地区表（注册地区选择）
// @Summary 修改系统地区表（注册地区选择）
// @Description 修改系统地区表（注册地区选择）
// @Tags 系统地区表（注册地区选择）
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigRegionsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-regions/{id} [put]
// @Security Bearer
func (e HsConfigRegions) Update(c *gin.Context) {
    req := dto.HsConfigRegionsUpdateReq{}
    s := service.HsConfigRegions{}
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
		e.Error(500, err, fmt.Sprintf("修改系统地区表（注册地区选择）失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除系统地区表（注册地区选择）
// @Summary 删除系统地区表（注册地区选择）
// @Description 删除系统地区表（注册地区选择）
// @Tags 系统地区表（注册地区选择）
// @Param data body dto.HsConfigRegionsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-regions [delete]
// @Security Bearer
func (e HsConfigRegions) Delete(c *gin.Context) {
    s := service.HsConfigRegions{}
    req := dto.HsConfigRegionsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除系统地区表（注册地区选择）失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
