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
// @Param orderNo query string false "订单号"
// @Param userId query string false "用户ID"
// @Param status query int false "订单状态:0=待处理,1=已经接单,2=已完成,3=已取消,4=已经驳回"
// @Param processingStatus query int false "管理员处理状态:1=正在处理,2=取消,3=完成"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.OrdUserOrders}} "{"code": 200, "data": [...]}"
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
// @Success 200 {object} models.Response{data=models.OrdUserOrders} "{"code": 200, "data": [...]}"
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
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
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
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
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
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
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

// GetPageByAssign 根据接单人和状态查询订单列表
// @Summary 根据接单人和状态查询订单列表
// @Description 根据当前登录用户作为接单人，查询不同状态的订单列表（包含兑换码信息，仅接单人可访问）
// @Tags 礼品卡订单表
// @Param status query string false "订单状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.OrdUserOrders}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-user-orders/my-assigned [get]
// @Security Bearer
func (e OrdUserOrders) GetPageByAssign(c *gin.Context) {
    req := dto.OrdUserOrdersGetByAssignReq{}
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

	// 自动设置接单人ID为当前登录用户ID，确保只能查询自己接的单
	req.AssignBy = user.GetUserId(c)

	p := actions.GetPermissionFromContext(c)
	list := make([]models.OrdUserOrders, 0)
	var count int64

	err = s.GetPageByAssign(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查询接单人订单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// AcceptOrder 管理员接单
// @Summary 管理员接单
// @Description 管理员接单，将订单分配给当前登录的管理员
// @Tags 礼品卡订单表
// @Param id path int true "订单ID"
// @Success 200 {object} models.Response "{"code": 200, "message": "接单成功"}"
// @Router /api/v1/ord-user-orders/{id}/accept [post]
// @Security Bearer
func (e OrdUserOrders) AcceptOrder(c *gin.Context) {
    req := dto.OrdUserOrdersAcceptReq{}
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

	// 获取当前管理员ID
	adminId := user.GetUserId(c)

	// 查询当前管理员信息以获取名称
	var admin models.SysUser
	err = e.Orm.Model(&models.SysUser{}).Where("user_id = ?", adminId).First(&admin).Error
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "获取管理员信息失败")
		return
	}

	p := actions.GetPermissionFromContext(c)

	// 调用Service层接单方法
	err = s.AcceptOrder(&req, adminId, admin.NickName, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("接单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "接单成功")
}

// CancelAcceptOrder 取消接单
// @Summary 取消接单
// @Description 管理员取消接单，只能取消自己接的单
// @Tags 礼品卡订单表
// @Param id path int true "订单ID"
// @Param data body dto.OrdUserOrdersCancelAcceptReq true "body"
// @Success 200 {object} models.Response "{"code": 200, "message": "取消接单成功"}"
// @Router /api/v1/ord-user-orders/{id}/cancel-accept [post]
// @Security Bearer
func (e OrdUserOrders) CancelAcceptOrder(c *gin.Context) {
    req := dto.OrdUserOrdersCancelAcceptReq{}
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

	// 获取当前管理员ID
	adminId := user.GetUserId(c)

	p := actions.GetPermissionFromContext(c)

	// 调用Service层取消接单方法
	err = s.CancelAcceptOrder(&req, adminId, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("取消接单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "取消接单成功")
}
