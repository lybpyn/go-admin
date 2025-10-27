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

type HsConfigCheckinRewards struct {
	api.Api
}

// GetPage 获取签到奖励配置表列表
// @Summary 获取签到奖励配置表列表
// @Description 获取签到奖励配置表列表
// @Tags 签到奖励配置表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsConfigCheckinRewards}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-checkin-rewards [get]
// @Security Bearer
func (e HsConfigCheckinRewards) GetPage(c *gin.Context) {
    req := dto.HsConfigCheckinRewardsGetPageReq{}
    s := service.HsConfigCheckinRewards{}
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
	list := make([]models.HsConfigCheckinRewards, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取签到奖励配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取签到奖励配置表
// @Summary 获取签到奖励配置表
// @Description 获取签到奖励配置表
// @Tags 签到奖励配置表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsConfigCheckinRewards} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-checkin-rewards/{id} [get]
// @Security Bearer
func (e HsConfigCheckinRewards) Get(c *gin.Context) {
	req := dto.HsConfigCheckinRewardsGetReq{}
	s := service.HsConfigCheckinRewards{}
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
	var object models.HsConfigCheckinRewards

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取签到奖励配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建签到奖励配置表
// @Summary 创建签到奖励配置表
// @Description 创建签到奖励配置表
// @Tags 签到奖励配置表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigCheckinRewardsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-checkin-rewards [post]
// @Security Bearer
func (e HsConfigCheckinRewards) Insert(c *gin.Context) {
    req := dto.HsConfigCheckinRewardsInsertReq{}
    s := service.HsConfigCheckinRewards{}
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
		e.Error(500, err, fmt.Sprintf("创建签到奖励配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改签到奖励配置表
// @Summary 修改签到奖励配置表
// @Description 修改签到奖励配置表
// @Tags 签到奖励配置表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigCheckinRewardsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-checkin-rewards/{id} [put]
// @Security Bearer
func (e HsConfigCheckinRewards) Update(c *gin.Context) {
    req := dto.HsConfigCheckinRewardsUpdateReq{}
    s := service.HsConfigCheckinRewards{}
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
		e.Error(500, err, fmt.Sprintf("修改签到奖励配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除签到奖励配置表
// @Summary 删除签到奖励配置表
// @Description 删除签到奖励配置表
// @Tags 签到奖励配置表
// @Param data body dto.HsConfigCheckinRewardsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-checkin-rewards [delete]
// @Security Bearer
func (e HsConfigCheckinRewards) Delete(c *gin.Context) {
    s := service.HsConfigCheckinRewards{}
    req := dto.HsConfigCheckinRewardsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除签到奖励配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
