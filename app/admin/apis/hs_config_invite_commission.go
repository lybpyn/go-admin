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

type HsConfigInviteCommission struct {
	api.Api
}

// GetPage 获取邀请分成比例配置表列表
// @Summary 获取邀请分成比例配置表列表
// @Description 获取邀请分成比例配置表列表
// @Tags 邀请分成比例配置表
// @Param configName query string false "配置名称"
// @Param status query string false "状态：1=启用，0=禁用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.HsConfigInviteCommission}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-invite-commission [get]
// @Security Bearer
func (e HsConfigInviteCommission) GetPage(c *gin.Context) {
    req := dto.HsConfigInviteCommissionGetPageReq{}
    s := service.HsConfigInviteCommission{}
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
	list := make([]models.HsConfigInviteCommission, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取邀请分成比例配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取邀请分成比例配置表
// @Summary 获取邀请分成比例配置表
// @Description 获取邀请分成比例配置表
// @Tags 邀请分成比例配置表
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.HsConfigInviteCommission} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-invite-commission/{id} [get]
// @Security Bearer
func (e HsConfigInviteCommission) Get(c *gin.Context) {
	req := dto.HsConfigInviteCommissionGetReq{}
	s := service.HsConfigInviteCommission{}
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
	var object models.HsConfigInviteCommission

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取邀请分成比例配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建邀请分成比例配置表
// @Summary 创建邀请分成比例配置表
// @Description 创建邀请分成比例配置表
// @Tags 邀请分成比例配置表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigInviteCommissionInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-invite-commission [post]
// @Security Bearer
func (e HsConfigInviteCommission) Insert(c *gin.Context) {
    req := dto.HsConfigInviteCommissionInsertReq{}
    s := service.HsConfigInviteCommission{}
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
		e.Error(500, err, fmt.Sprintf("创建邀请分成比例配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改邀请分成比例配置表
// @Summary 修改邀请分成比例配置表
// @Description 修改邀请分成比例配置表
// @Tags 邀请分成比例配置表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigInviteCommissionUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-invite-commission/{id} [put]
// @Security Bearer
func (e HsConfigInviteCommission) Update(c *gin.Context) {
    req := dto.HsConfigInviteCommissionUpdateReq{}
    s := service.HsConfigInviteCommission{}
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
		e.Error(500, err, fmt.Sprintf("修改邀请分成比例配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除邀请分成比例配置表
// @Summary 删除邀请分成比例配置表
// @Description 删除邀请分成比例配置表
// @Tags 邀请分成比例配置表
// @Param data body dto.HsConfigInviteCommissionDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-invite-commission [delete]
// @Security Bearer
func (e HsConfigInviteCommission) Delete(c *gin.Context) {
    s := service.HsConfigInviteCommission{}
    req := dto.HsConfigInviteCommissionDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除邀请分成比例配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
