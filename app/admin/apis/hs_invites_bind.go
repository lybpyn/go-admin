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

type HsInvitesBind struct {
	api.Api
}

// GetPage 获取用户邀请关系表列表
// @Summary 获取用户邀请关系表列表
// @Description 获取用户邀请关系表列表
// @Tags 用户邀请关系表
// @Param userId query string false "被邀请用户ID"
// @Param inviterId query string false "直接邀请人ID (一级)"
// @Param level query string false "层级: 1=一级, 2=二级"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsInvitesBind}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-invites-bind [get]
// @Security Bearer
func (e HsInvitesBind) GetPage(c *gin.Context) {
    req := dto.HsInvitesBindGetPageReq{}
    s := service.HsInvitesBind{}
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
	list := make([]models.HsInvitesBind, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户邀请关系表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户邀请关系表
// @Summary 获取用户邀请关系表
// @Description 获取用户邀请关系表
// @Tags 用户邀请关系表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsInvitesBind} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-invites-bind/{id} [get]
// @Security Bearer
func (e HsInvitesBind) Get(c *gin.Context) {
	req := dto.HsInvitesBindGetReq{}
	s := service.HsInvitesBind{}
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
	var object models.HsInvitesBind

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户邀请关系表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户邀请关系表
// @Summary 创建用户邀请关系表
// @Description 创建用户邀请关系表
// @Tags 用户邀请关系表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsInvitesBindInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-invites-bind [post]
// @Security Bearer
func (e HsInvitesBind) Insert(c *gin.Context) {
    req := dto.HsInvitesBindInsertReq{}
    s := service.HsInvitesBind{}
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
		e.Error(500, err, fmt.Sprintf("创建用户邀请关系表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户邀请关系表
// @Summary 修改用户邀请关系表
// @Description 修改用户邀请关系表
// @Tags 用户邀请关系表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsInvitesBindUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-invites-bind/{id} [put]
// @Security Bearer
func (e HsInvitesBind) Update(c *gin.Context) {
    req := dto.HsInvitesBindUpdateReq{}
    s := service.HsInvitesBind{}
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
		e.Error(500, err, fmt.Sprintf("修改用户邀请关系表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户邀请关系表
// @Summary 删除用户邀请关系表
// @Description 删除用户邀请关系表
// @Tags 用户邀请关系表
// @Param data body dto.HsInvitesBindDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-invites-bind [delete]
// @Security Bearer
func (e HsInvitesBind) Delete(c *gin.Context) {
    s := service.HsInvitesBind{}
    req := dto.HsInvitesBindDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户邀请关系表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
