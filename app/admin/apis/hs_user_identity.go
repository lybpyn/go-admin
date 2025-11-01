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

type HsUserIdentity struct {
	api.Api
}

// GetPage 获取HsUserIdentity列表
// @Summary 获取HsUserIdentity列表
// @Description 获取HsUserIdentity列表
// @Tags HsUserIdentity
// @Param userId query string false ""
// @Param identityType query string false "凭证类型: email, phone, google, apple, twitter, tiktok"
// @Param identifier query string false "凭证值: 邮箱, 手机号, 第三方平台的唯一ID"
// @Param credential query string false "存储密码哈希(email/phone)或refresh_token(social)"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserIdentity}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-identity [get]
// @Security Bearer
func (e HsUserIdentity) GetPage(c *gin.Context) {
    req := dto.HsUserIdentityGetPageReq{}
    s := service.HsUserIdentity{}
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
	list := make([]models.HsUserIdentity, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取HsUserIdentity失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取HsUserIdentity
// @Summary 获取HsUserIdentity
// @Description 获取HsUserIdentity
// @Tags HsUserIdentity
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsUserIdentity} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-identity/{id} [get]
// @Security Bearer
func (e HsUserIdentity) Get(c *gin.Context) {
	req := dto.HsUserIdentityGetReq{}
	s := service.HsUserIdentity{}
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
	var object models.HsUserIdentity

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取HsUserIdentity失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建HsUserIdentity
// @Summary 创建HsUserIdentity
// @Description 创建HsUserIdentity
// @Tags HsUserIdentity
// @Accept application/json
// @Product application/json
// @Param data body dto.HsUserIdentityInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-user-identity [post]
// @Security Bearer
func (e HsUserIdentity) Insert(c *gin.Context) {
    req := dto.HsUserIdentityInsertReq{}
    s := service.HsUserIdentity{}
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
		e.Error(500, err, fmt.Sprintf("创建HsUserIdentity失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改HsUserIdentity
// @Summary 修改HsUserIdentity
// @Description 修改HsUserIdentity
// @Tags HsUserIdentity
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsUserIdentityUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-identity/{id} [put]
// @Security Bearer
func (e HsUserIdentity) Update(c *gin.Context) {
    req := dto.HsUserIdentityUpdateReq{}
    s := service.HsUserIdentity{}
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
		e.Error(500, err, fmt.Sprintf("修改HsUserIdentity失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除HsUserIdentity
// @Summary 删除HsUserIdentity
// @Description 删除HsUserIdentity
// @Tags HsUserIdentity
// @Param data body dto.HsUserIdentityDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-identity [delete]
// @Security Bearer
func (e HsUserIdentity) Delete(c *gin.Context) {
    s := service.HsUserIdentity{}
    req := dto.HsUserIdentityDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除HsUserIdentity失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
