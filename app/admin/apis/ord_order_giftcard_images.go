package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type OrdOrderGiftcardImages struct {
	api.Api
}

// GetPage 获取礼品卡订单图片表列表
// @Summary 获取礼品卡订单图片表列表
// @Description 获取礼品卡订单图片表列表
// @Tags 礼品卡订单图片表
// @Param orderId query string false "关联的订单ID"
// @Param imageUrl query string false "礼品卡图片URL"
// @Param sortOrder query string false "排序顺序"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.OrdOrderGiftcardImages}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-order-giftcard-images [get]
// @Security Bearer
func (e OrdOrderGiftcardImages) GetPage(c *gin.Context) {
    req := dto.OrdOrderGiftcardImagesGetPageReq{}
    s := service.OrdOrderGiftcardImages{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req, binding.Form, binding.Query).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	// 获取当前管理员ID
	adminId := user.GetUserId(c)

	p := actions.GetPermissionFromContext(c)
	list := make([]models.OrdOrderGiftcardImages, 0)
	var count int64

	err = s.GetPage(&req, p, adminId, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡订单图片表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取礼品卡订单图片表
// @Summary 获取礼品卡订单图片表
// @Description 获取礼品卡订单图片表
// @Tags 礼品卡订单图片表
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.OrdOrderGiftcardImages} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-order-giftcard-images/{id} [get]
// @Security Bearer
func (e OrdOrderGiftcardImages) Get(c *gin.Context) {
	req := dto.OrdOrderGiftcardImagesGetReq{}
	s := service.OrdOrderGiftcardImages{}
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
	var object models.OrdOrderGiftcardImages

	// 获取当前管理员ID
	adminId := user.GetUserId(c)

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, adminId, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡订单图片表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建礼品卡订单图片表
// @Summary 创建礼品卡订单图片表
// @Description 创建礼品卡订单图片表
// @Tags 礼品卡订单图片表
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdOrderGiftcardImagesInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-order-giftcard-images [post]
// @Security Bearer
func (e OrdOrderGiftcardImages) Insert(c *gin.Context) {
    req := dto.OrdOrderGiftcardImagesInsertReq{}
    s := service.OrdOrderGiftcardImages{}
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
		e.Error(500, err, fmt.Sprintf("创建礼品卡订单图片表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改礼品卡订单图片表
// @Summary 修改礼品卡订单图片表
// @Description 修改礼品卡订单图片表
// @Tags 礼品卡订单图片表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdOrderGiftcardImagesUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-order-giftcard-images/{id} [put]
// @Security Bearer
func (e OrdOrderGiftcardImages) Update(c *gin.Context) {
    req := dto.OrdOrderGiftcardImagesUpdateReq{}
    s := service.OrdOrderGiftcardImages{}
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
		e.Error(500, err, fmt.Sprintf("修改礼品卡订单图片表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除礼品卡订单图片表
// @Summary 删除礼品卡订单图片表
// @Description 删除礼品卡订单图片表
// @Tags 礼品卡订单图片表
// @Param data body dto.OrdOrderGiftcardImagesDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-order-giftcard-images [delete]
// @Security Bearer
func (e OrdOrderGiftcardImages) Delete(c *gin.Context) {
    s := service.OrdOrderGiftcardImages{}
    req := dto.OrdOrderGiftcardImagesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除礼品卡订单图片表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
