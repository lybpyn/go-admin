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

type NoSystemNotifications struct {
	api.Api
}

// GetPage 获取系统消息提醒表列表
// @Summary 获取系统消息提醒表列表
// @Description 获取系统消息提醒表列表
// @Tags 系统消息提醒表
// @Param targetUserId query string false "目标用户ID（仅 target_type=user 时有效）"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.NoSystemNotifications}} "{"code": 200, "data": [...]}"
// @Router /api/v1/no-system-notifications [get]
// @Security Bearer
func (e NoSystemNotifications) GetPage(c *gin.Context) {
    req := dto.NoSystemNotificationsGetPageReq{}
    s := service.NoSystemNotifications{}
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
	list := make([]models.NoSystemNotifications, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取系统消息提醒表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取系统消息提醒表
// @Summary 获取系统消息提醒表
// @Description 获取系统消息提醒表
// @Tags 系统消息提醒表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.NoSystemNotifications} "{"code": 200, "data": [...]}"
// @Router /api/v1/no-system-notifications/{id} [get]
// @Security Bearer
func (e NoSystemNotifications) Get(c *gin.Context) {
	req := dto.NoSystemNotificationsGetReq{}
	s := service.NoSystemNotifications{}
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
	var object models.NoSystemNotifications

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取系统消息提醒表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建系统消息提醒表
// @Summary 创建系统消息提醒表
// @Description 创建系统消息提醒表
// @Tags 系统消息提醒表
// @Accept application/json
// @Product application/json
// @Param data body dto.NoSystemNotificationsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/no-system-notifications [post]
// @Security Bearer
func (e NoSystemNotifications) Insert(c *gin.Context) {
    req := dto.NoSystemNotificationsInsertReq{}
    s := service.NoSystemNotifications{}
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
		e.Error(500, err, fmt.Sprintf("创建系统消息提醒表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改系统消息提醒表
// @Summary 修改系统消息提醒表
// @Description 修改系统消息提醒表
// @Tags 系统消息提醒表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.NoSystemNotificationsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/no-system-notifications/{id} [put]
// @Security Bearer
func (e NoSystemNotifications) Update(c *gin.Context) {
    req := dto.NoSystemNotificationsUpdateReq{}
    s := service.NoSystemNotifications{}
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
		e.Error(500, err, fmt.Sprintf("修改系统消息提醒表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除系统消息提醒表
// @Summary 删除系统消息提醒表
// @Description 删除系统消息提醒表
// @Tags 系统消息提醒表
// @Param data body dto.NoSystemNotificationsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/no-system-notifications [delete]
// @Security Bearer
func (e NoSystemNotifications) Delete(c *gin.Context) {
    s := service.NoSystemNotifications{}
    req := dto.NoSystemNotificationsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除系统消息提醒表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
