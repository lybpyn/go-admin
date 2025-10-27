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

type MdEventDefinitionTab struct {
	api.Api
}

// GetPage 获取事件定义表列表
// @Summary 获取事件定义表列表
// @Description 获取事件定义表列表
// @Tags 事件定义表
// @Param eventCode query string false "事件代码（唯一）"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.MdEventDefinitionTab}} "{"code": 200, "data": [...]}"
// @Router /api/v1/md-event-definition-tab [get]
// @Security Bearer
func (e MdEventDefinitionTab) GetPage(c *gin.Context) {
    req := dto.MdEventDefinitionTabGetPageReq{}
    s := service.MdEventDefinitionTab{}
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
	list := make([]models.MdEventDefinitionTab, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取事件定义表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取事件定义表
// @Summary 获取事件定义表
// @Description 获取事件定义表
// @Tags 事件定义表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.MdEventDefinitionTab} "{"code": 200, "data": [...]}"
// @Router /api/v1/md-event-definition-tab/{id} [get]
// @Security Bearer
func (e MdEventDefinitionTab) Get(c *gin.Context) {
	req := dto.MdEventDefinitionTabGetReq{}
	s := service.MdEventDefinitionTab{}
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
	var object models.MdEventDefinitionTab

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取事件定义表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建事件定义表
// @Summary 创建事件定义表
// @Description 创建事件定义表
// @Tags 事件定义表
// @Accept application/json
// @Product application/json
// @Param data body dto.MdEventDefinitionTabInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/md-event-definition-tab [post]
// @Security Bearer
func (e MdEventDefinitionTab) Insert(c *gin.Context) {
    req := dto.MdEventDefinitionTabInsertReq{}
    s := service.MdEventDefinitionTab{}
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
		e.Error(500, err, fmt.Sprintf("创建事件定义表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改事件定义表
// @Summary 修改事件定义表
// @Description 修改事件定义表
// @Tags 事件定义表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.MdEventDefinitionTabUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/md-event-definition-tab/{id} [put]
// @Security Bearer
func (e MdEventDefinitionTab) Update(c *gin.Context) {
    req := dto.MdEventDefinitionTabUpdateReq{}
    s := service.MdEventDefinitionTab{}
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
		e.Error(500, err, fmt.Sprintf("修改事件定义表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除事件定义表
// @Summary 删除事件定义表
// @Description 删除事件定义表
// @Tags 事件定义表
// @Param data body dto.MdEventDefinitionTabDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/md-event-definition-tab [delete]
// @Security Bearer
func (e MdEventDefinitionTab) Delete(c *gin.Context) {
    s := service.MdEventDefinitionTab{}
    req := dto.MdEventDefinitionTabDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除事件定义表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
