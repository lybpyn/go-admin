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

type HsCarouselBanner struct {
	api.Api
}

// GetPage 获取首页轮播广告列表
// @Summary 获取首页轮播广告列表
// @Description 获取首页轮播广告列表
// @Tags 首页轮播广告
// @Param startTime query time.Time false "开始展示时间"
// @Param endTime query time.Time false "结束展示时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsCarouselBanner}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-carousel-banner [get]
// @Security Bearer
func (e HsCarouselBanner) GetPage(c *gin.Context) {
    req := dto.HsCarouselBannerGetPageReq{}
    s := service.HsCarouselBanner{}
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
	list := make([]models.HsCarouselBanner, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取首页轮播广告失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取首页轮播广告
// @Summary 获取首页轮播广告
// @Description 获取首页轮播广告
// @Tags 首页轮播广告
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsCarouselBanner} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-carousel-banner/{id} [get]
// @Security Bearer
func (e HsCarouselBanner) Get(c *gin.Context) {
	req := dto.HsCarouselBannerGetReq{}
	s := service.HsCarouselBanner{}
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
	var object models.HsCarouselBanner

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取首页轮播广告失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建首页轮播广告
// @Summary 创建首页轮播广告
// @Description 创建首页轮播广告
// @Tags 首页轮播广告
// @Accept application/json
// @Product application/json
// @Param data body dto.HsCarouselBannerInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-carousel-banner [post]
// @Security Bearer
func (e HsCarouselBanner) Insert(c *gin.Context) {
    req := dto.HsCarouselBannerInsertReq{}
    s := service.HsCarouselBanner{}
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
		e.Error(500, err, fmt.Sprintf("创建首页轮播广告失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改首页轮播广告
// @Summary 修改首页轮播广告
// @Description 修改首页轮播广告
// @Tags 首页轮播广告
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsCarouselBannerUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-carousel-banner/{id} [put]
// @Security Bearer
func (e HsCarouselBanner) Update(c *gin.Context) {
    req := dto.HsCarouselBannerUpdateReq{}
    s := service.HsCarouselBanner{}
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
		e.Error(500, err, fmt.Sprintf("修改首页轮播广告失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除首页轮播广告
// @Summary 删除首页轮播广告
// @Description 删除首页轮播广告
// @Tags 首页轮播广告
// @Param data body dto.HsCarouselBannerDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-carousel-banner [delete]
// @Security Bearer
func (e HsCarouselBanner) Delete(c *gin.Context) {
    s := service.HsCarouselBanner{}
    req := dto.HsCarouselBannerDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除首页轮播广告失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
