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

type HsUserSocialTokens struct {
	api.Api
}

// GetPage 获取社交登录Token表列表
// @Summary 获取社交登录Token表列表
// @Description 获取社交登录Token表列表
// @Tags 社交登录Token表
// @Param userId query string false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserSocialTokens}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-social-tokens [get]
// @Security Bearer
func (e HsUserSocialTokens) GetPage(c *gin.Context) {
    req := dto.HsUserSocialTokensGetPageReq{}
    s := service.HsUserSocialTokens{}
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
	list := make([]models.HsUserSocialTokens, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取社交登录Token表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取社交登录Token表
// @Summary 获取社交登录Token表
// @Description 获取社交登录Token表
// @Tags 社交登录Token表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsUserSocialTokens} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-social-tokens/{id} [get]
// @Security Bearer
func (e HsUserSocialTokens) Get(c *gin.Context) {
	req := dto.HsUserSocialTokensGetReq{}
	s := service.HsUserSocialTokens{}
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
	var object models.HsUserSocialTokens

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取社交登录Token表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建社交登录Token表
// @Summary 创建社交登录Token表
// @Description 创建社交登录Token表
// @Tags 社交登录Token表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsUserSocialTokensInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-user-social-tokens [post]
// @Security Bearer
func (e HsUserSocialTokens) Insert(c *gin.Context) {
    req := dto.HsUserSocialTokensInsertReq{}
    s := service.HsUserSocialTokens{}
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
		e.Error(500, err, fmt.Sprintf("创建社交登录Token表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改社交登录Token表
// @Summary 修改社交登录Token表
// @Description 修改社交登录Token表
// @Tags 社交登录Token表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsUserSocialTokensUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-social-tokens/{id} [put]
// @Security Bearer
func (e HsUserSocialTokens) Update(c *gin.Context) {
    req := dto.HsUserSocialTokensUpdateReq{}
    s := service.HsUserSocialTokens{}
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
		e.Error(500, err, fmt.Sprintf("修改社交登录Token表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除社交登录Token表
// @Summary 删除社交登录Token表
// @Description 删除社交登录Token表
// @Tags 社交登录Token表
// @Param data body dto.HsUserSocialTokensDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-social-tokens [delete]
// @Security Bearer
func (e HsUserSocialTokens) Delete(c *gin.Context) {
    s := service.HsUserSocialTokens{}
    req := dto.HsUserSocialTokensDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除社交登录Token表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
