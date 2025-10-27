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

type HsUserCheckins struct {
	api.Api
}

// GetPage 获取用户签到表列表
// @Summary 获取用户签到表列表
// @Description 获取用户签到表列表
// @Tags 用户签到表
// @Param userId query string false "用户ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsUserCheckins}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-checkins [get]
// @Security Bearer
func (e HsUserCheckins) GetPage(c *gin.Context) {
    req := dto.HsUserCheckinsGetPageReq{}
    s := service.HsUserCheckins{}
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
	list := make([]models.HsUserCheckins, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户签到表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户签到表
// @Summary 获取用户签到表
// @Description 获取用户签到表
// @Tags 用户签到表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsUserCheckins} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-user-checkins/{id} [get]
// @Security Bearer
func (e HsUserCheckins) Get(c *gin.Context) {
	req := dto.HsUserCheckinsGetReq{}
	s := service.HsUserCheckins{}
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
	var object models.HsUserCheckins

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户签到表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户签到表
// @Summary 创建用户签到表
// @Description 创建用户签到表
// @Tags 用户签到表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsUserCheckinsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-user-checkins [post]
// @Security Bearer
func (e HsUserCheckins) Insert(c *gin.Context) {
    req := dto.HsUserCheckinsInsertReq{}
    s := service.HsUserCheckins{}
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
		e.Error(500, err, fmt.Sprintf("创建用户签到表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户签到表
// @Summary 修改用户签到表
// @Description 修改用户签到表
// @Tags 用户签到表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsUserCheckinsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-user-checkins/{id} [put]
// @Security Bearer
func (e HsUserCheckins) Update(c *gin.Context) {
    req := dto.HsUserCheckinsUpdateReq{}
    s := service.HsUserCheckins{}
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
		e.Error(500, err, fmt.Sprintf("修改用户签到表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户签到表
// @Summary 删除用户签到表
// @Description 删除用户签到表
// @Tags 用户签到表
// @Param data body dto.HsUserCheckinsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-user-checkins [delete]
// @Security Bearer
func (e HsUserCheckins) Delete(c *gin.Context) {
    s := service.HsUserCheckins{}
    req := dto.HsUserCheckinsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户签到表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
