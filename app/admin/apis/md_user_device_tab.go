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

type MdUserDeviceTab struct {
	api.Api
}

// GetPage 获取用户设备信息表列表
// @Summary 获取用户设备信息表列表
// @Description 获取用户设备信息表列表
// @Tags 用户设备信息表
// @Param deviceId query string false "设备唯一标识"
// @Param userId query string false "用户ID（未绑定为0）"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.MdUserDeviceTab}} "{"code": 200, "data": [...]}"
// @Router /api/v1/md-user-device-tab [get]
// @Security Bearer
func (e MdUserDeviceTab) GetPage(c *gin.Context) {
    req := dto.MdUserDeviceTabGetPageReq{}
    s := service.MdUserDeviceTab{}
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
	list := make([]models.MdUserDeviceTab, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户设备信息表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户设备信息表
// @Summary 获取用户设备信息表
// @Description 获取用户设备信息表
// @Tags 用户设备信息表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.MdUserDeviceTab} "{"code": 200, "data": [...]}"
// @Router /api/v1/md-user-device-tab/{id} [get]
// @Security Bearer
func (e MdUserDeviceTab) Get(c *gin.Context) {
	req := dto.MdUserDeviceTabGetReq{}
	s := service.MdUserDeviceTab{}
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
	var object models.MdUserDeviceTab

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户设备信息表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户设备信息表
// @Summary 创建用户设备信息表
// @Description 创建用户设备信息表
// @Tags 用户设备信息表
// @Accept application/json
// @Product application/json
// @Param data body dto.MdUserDeviceTabInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/md-user-device-tab [post]
// @Security Bearer
func (e MdUserDeviceTab) Insert(c *gin.Context) {
    req := dto.MdUserDeviceTabInsertReq{}
    s := service.MdUserDeviceTab{}
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
		e.Error(500, err, fmt.Sprintf("创建用户设备信息表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户设备信息表
// @Summary 修改用户设备信息表
// @Description 修改用户设备信息表
// @Tags 用户设备信息表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.MdUserDeviceTabUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/md-user-device-tab/{id} [put]
// @Security Bearer
func (e MdUserDeviceTab) Update(c *gin.Context) {
    req := dto.MdUserDeviceTabUpdateReq{}
    s := service.MdUserDeviceTab{}
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
		e.Error(500, err, fmt.Sprintf("修改用户设备信息表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户设备信息表
// @Summary 删除用户设备信息表
// @Description 删除用户设备信息表
// @Tags 用户设备信息表
// @Param data body dto.MdUserDeviceTabDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/md-user-device-tab [delete]
// @Security Bearer
func (e MdUserDeviceTab) Delete(c *gin.Context) {
    s := service.MdUserDeviceTab{}
    req := dto.MdUserDeviceTabDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户设备信息表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
