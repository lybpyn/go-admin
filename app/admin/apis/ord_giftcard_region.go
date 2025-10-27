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

type OrdGiftcardRegion struct {
	api.Api
}

// GetPage 获取礼品卡地区列表
// @Summary 获取礼品卡地区列表
// @Description 获取礼品卡地区列表
// @Tags 礼品卡地区
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrdGiftcardRegion}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-region [get]
// @Security Bearer
func (e OrdGiftcardRegion) GetPage(c *gin.Context) {
    req := dto.OrdGiftcardRegionGetPageReq{}
    s := service.OrdGiftcardRegion{}
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
	list := make([]models.OrdGiftcardRegion, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡地区失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取礼品卡地区
// @Summary 获取礼品卡地区
// @Description 获取礼品卡地区
// @Tags 礼品卡地区
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrdGiftcardRegion} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-region/{id} [get]
// @Security Bearer
func (e OrdGiftcardRegion) Get(c *gin.Context) {
	req := dto.OrdGiftcardRegionGetReq{}
	s := service.OrdGiftcardRegion{}
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
	var object models.OrdGiftcardRegion

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡地区失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建礼品卡地区
// @Summary 创建礼品卡地区
// @Description 创建礼品卡地区
// @Tags 礼品卡地区
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdGiftcardRegionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-giftcard-region [post]
// @Security Bearer
func (e OrdGiftcardRegion) Insert(c *gin.Context) {
    req := dto.OrdGiftcardRegionInsertReq{}
    s := service.OrdGiftcardRegion{}
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
		e.Error(500, err, fmt.Sprintf("创建礼品卡地区失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改礼品卡地区
// @Summary 修改礼品卡地区
// @Description 修改礼品卡地区
// @Tags 礼品卡地区
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdGiftcardRegionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-giftcard-region/{id} [put]
// @Security Bearer
func (e OrdGiftcardRegion) Update(c *gin.Context) {
    req := dto.OrdGiftcardRegionUpdateReq{}
    s := service.OrdGiftcardRegion{}
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
		e.Error(500, err, fmt.Sprintf("修改礼品卡地区失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除礼品卡地区
// @Summary 删除礼品卡地区
// @Description 删除礼品卡地区
// @Tags 礼品卡地区
// @Param data body dto.OrdGiftcardRegionDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-giftcard-region [delete]
// @Security Bearer
func (e OrdGiftcardRegion) Delete(c *gin.Context) {
    s := service.OrdGiftcardRegion{}
    req := dto.OrdGiftcardRegionDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除礼品卡地区失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
