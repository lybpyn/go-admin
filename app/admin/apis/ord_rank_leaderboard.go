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

type OrdRankLeaderboard struct {
	api.Api
}

// GetPage 获取OrdRankLeaderboard列表
// @Summary 获取OrdRankLeaderboard列表
// @Description 获取OrdRankLeaderboard列表
// @Tags OrdRankLeaderboard
// @Param rank query int false ""
// @Param userId query string false ""
// @Param username query string false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrdRankLeaderboard}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-rank-leaderboard [get]
// @Security Bearer
func (e OrdRankLeaderboard) GetPage(c *gin.Context) {
    req := dto.OrdRankLeaderboardGetPageReq{}
    s := service.OrdRankLeaderboard{}
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
	list := make([]models.OrdRankLeaderboard, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取OrdRankLeaderboard失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取OrdRankLeaderboard
// @Summary 获取OrdRankLeaderboard
// @Description 获取OrdRankLeaderboard
// @Tags OrdRankLeaderboard
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrdRankLeaderboard} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-rank-leaderboard/{id} [get]
// @Security Bearer
func (e OrdRankLeaderboard) Get(c *gin.Context) {
	req := dto.OrdRankLeaderboardGetReq{}
	s := service.OrdRankLeaderboard{}
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
	var object models.OrdRankLeaderboard

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取OrdRankLeaderboard失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建OrdRankLeaderboard
// @Summary 创建OrdRankLeaderboard
// @Description 创建OrdRankLeaderboard
// @Tags OrdRankLeaderboard
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdRankLeaderboardInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-rank-leaderboard [post]
// @Security Bearer
func (e OrdRankLeaderboard) Insert(c *gin.Context) {
    req := dto.OrdRankLeaderboardInsertReq{}
    s := service.OrdRankLeaderboard{}
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
		e.Error(500, err, fmt.Sprintf("创建OrdRankLeaderboard失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改OrdRankLeaderboard
// @Summary 修改OrdRankLeaderboard
// @Description 修改OrdRankLeaderboard
// @Tags OrdRankLeaderboard
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdRankLeaderboardUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-rank-leaderboard/{id} [put]
// @Security Bearer
func (e OrdRankLeaderboard) Update(c *gin.Context) {
    req := dto.OrdRankLeaderboardUpdateReq{}
    s := service.OrdRankLeaderboard{}
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
		e.Error(500, err, fmt.Sprintf("修改OrdRankLeaderboard失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除OrdRankLeaderboard
// @Summary 删除OrdRankLeaderboard
// @Description 删除OrdRankLeaderboard
// @Tags OrdRankLeaderboard
// @Param data body dto.OrdRankLeaderboardDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-rank-leaderboard [delete]
// @Security Bearer
func (e OrdRankLeaderboard) Delete(c *gin.Context) {
    s := service.OrdRankLeaderboard{}
    req := dto.OrdRankLeaderboardDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除OrdRankLeaderboard失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
