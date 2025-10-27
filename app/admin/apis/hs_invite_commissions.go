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

type HsInviteCommissions struct {
	api.Api
}

// GetPage 获取订单分成流水表列表
// @Summary 获取订单分成流水表列表
// @Description 获取订单分成流水表列表
// @Tags 订单分成流水表
// @Param orderId query string false "订单ID"
// @Param userId query string false "分成获得者ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsInviteCommissions}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-invite-commissions [get]
// @Security Bearer
func (e HsInviteCommissions) GetPage(c *gin.Context) {
    req := dto.HsInviteCommissionsGetPageReq{}
    s := service.HsInviteCommissions{}
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
	list := make([]models.HsInviteCommissions, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取订单分成流水表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取订单分成流水表
// @Summary 获取订单分成流水表
// @Description 获取订单分成流水表
// @Tags 订单分成流水表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsInviteCommissions} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-invite-commissions/{id} [get]
// @Security Bearer
func (e HsInviteCommissions) Get(c *gin.Context) {
	req := dto.HsInviteCommissionsGetReq{}
	s := service.HsInviteCommissions{}
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
	var object models.HsInviteCommissions

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取订单分成流水表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建订单分成流水表
// @Summary 创建订单分成流水表
// @Description 创建订单分成流水表
// @Tags 订单分成流水表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsInviteCommissionsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-invite-commissions [post]
// @Security Bearer
func (e HsInviteCommissions) Insert(c *gin.Context) {
    req := dto.HsInviteCommissionsInsertReq{}
    s := service.HsInviteCommissions{}
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
		e.Error(500, err, fmt.Sprintf("创建订单分成流水表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改订单分成流水表
// @Summary 修改订单分成流水表
// @Description 修改订单分成流水表
// @Tags 订单分成流水表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsInviteCommissionsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-invite-commissions/{id} [put]
// @Security Bearer
func (e HsInviteCommissions) Update(c *gin.Context) {
    req := dto.HsInviteCommissionsUpdateReq{}
    s := service.HsInviteCommissions{}
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
		e.Error(500, err, fmt.Sprintf("修改订单分成流水表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除订单分成流水表
// @Summary 删除订单分成流水表
// @Description 删除订单分成流水表
// @Tags 订单分成流水表
// @Param data body dto.HsInviteCommissionsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-invite-commissions [delete]
// @Security Bearer
func (e HsInviteCommissions) Delete(c *gin.Context) {
    s := service.HsInviteCommissions{}
    req := dto.HsInviteCommissionsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除订单分成流水表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
