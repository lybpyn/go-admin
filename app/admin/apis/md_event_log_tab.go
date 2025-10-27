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

type MdEventLogTab struct {
	api.Api
}

// GetPage 获取埋点事件日志表列表
// @Summary 获取埋点事件日志表列表
// @Description 获取埋点事件日志表列表
// @Tags 埋点事件日志表
// @Param eventId query string false "事件ID（唯一标识）"
// @Param eventCode query string false "事件代码，如：xx_Installation, xx_register_success"
// @Param eventName query string false "事件名称，如：安装、注册成功"
// @Param moduleName query string false "模块名称"
// @Param userId query string false "用户ID（未登录为0）"
// @Param deviceId query string false "设备唯一标识（IDFA/GAID等）"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.MdEventLogTab}} "{"code": 200, "data": [...]}"
// @Router /api/v1/md-event-log-tab [get]
// @Security Bearer
func (e MdEventLogTab) GetPage(c *gin.Context) {
    req := dto.MdEventLogTabGetPageReq{}
    s := service.MdEventLogTab{}
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
	list := make([]models.MdEventLogTab, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取埋点事件日志表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取埋点事件日志表
// @Summary 获取埋点事件日志表
// @Description 获取埋点事件日志表
// @Tags 埋点事件日志表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.MdEventLogTab} "{"code": 200, "data": [...]}"
// @Router /api/v1/md-event-log-tab/{id} [get]
// @Security Bearer
func (e MdEventLogTab) Get(c *gin.Context) {
	req := dto.MdEventLogTabGetReq{}
	s := service.MdEventLogTab{}
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
	var object models.MdEventLogTab

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取埋点事件日志表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建埋点事件日志表
// @Summary 创建埋点事件日志表
// @Description 创建埋点事件日志表
// @Tags 埋点事件日志表
// @Accept application/json
// @Product application/json
// @Param data body dto.MdEventLogTabInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/md-event-log-tab [post]
// @Security Bearer
func (e MdEventLogTab) Insert(c *gin.Context) {
    req := dto.MdEventLogTabInsertReq{}
    s := service.MdEventLogTab{}
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
		e.Error(500, err, fmt.Sprintf("创建埋点事件日志表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改埋点事件日志表
// @Summary 修改埋点事件日志表
// @Description 修改埋点事件日志表
// @Tags 埋点事件日志表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.MdEventLogTabUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/md-event-log-tab/{id} [put]
// @Security Bearer
func (e MdEventLogTab) Update(c *gin.Context) {
    req := dto.MdEventLogTabUpdateReq{}
    s := service.MdEventLogTab{}
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
		e.Error(500, err, fmt.Sprintf("修改埋点事件日志表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除埋点事件日志表
// @Summary 删除埋点事件日志表
// @Description 删除埋点事件日志表
// @Tags 埋点事件日志表
// @Param data body dto.MdEventLogTabDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/md-event-log-tab [delete]
// @Security Bearer
func (e MdEventLogTab) Delete(c *gin.Context) {
    s := service.MdEventLogTab{}
    req := dto.MdEventLogTabDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除埋点事件日志表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
