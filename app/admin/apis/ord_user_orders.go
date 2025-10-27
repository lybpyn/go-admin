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

type OrdUserOrders struct {
	api.Api
}

// GetPage 获取礼品卡订单表列表
// @Summary 获取礼品卡订单表列表
// @Description 获取礼品卡订单表列表
// @Tags 礼品卡订单表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrdUserOrders}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-user-orders [get]
// @Security Bearer
func (e OrdUserOrders) GetPage(c *gin.Context) {
    req := dto.OrdUserOrdersGetPageReq{}
    s := service.OrdUserOrders{}
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
	list := make([]models.OrdUserOrders, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡订单表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取礼品卡订单表
// @Summary 获取礼品卡订单表
// @Description 获取礼品卡订单表
// @Tags 礼品卡订单表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrdUserOrders} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-user-orders/{id} [get]
// @Security Bearer
func (e OrdUserOrders) Get(c *gin.Context) {
	req := dto.OrdUserOrdersGetReq{}
	s := service.OrdUserOrders{}
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
	var object models.OrdUserOrders

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡订单表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建礼品卡订单表
// @Summary 创建礼品卡订单表
// @Description 创建礼品卡订单表
// @Tags 礼品卡订单表
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdUserOrdersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-user-orders [post]
// @Security Bearer
func (e OrdUserOrders) Insert(c *gin.Context) {
    req := dto.OrdUserOrdersInsertReq{}
    s := service.OrdUserOrders{}
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
		e.Error(500, err, fmt.Sprintf("创建礼品卡订单表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改礼品卡订单表
// @Summary 修改礼品卡订单表
// @Description 修改礼品卡订单表
// @Tags 礼品卡订单表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdUserOrdersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-user-orders/{id} [put]
// @Security Bearer
func (e OrdUserOrders) Update(c *gin.Context) {
    req := dto.OrdUserOrdersUpdateReq{}
    s := service.OrdUserOrders{}
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
		e.Error(500, err, fmt.Sprintf("修改礼品卡订单表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除礼品卡订单表
// @Summary 删除礼品卡订单表
// @Description 删除礼品卡订单表
// @Tags 礼品卡订单表
// @Param data body dto.OrdUserOrdersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-user-orders [delete]
// @Security Bearer
func (e OrdUserOrders) Delete(c *gin.Context) {
    s := service.OrdUserOrders{}
    req := dto.OrdUserOrdersDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除礼品卡订单表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
