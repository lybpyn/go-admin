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

type OrdGiftcardDiscounts struct {
	api.Api
}

// GetPage 获取礼品卡不同类型折扣列表
// @Summary 获取礼品卡不同类型折扣列表
// @Description 获取礼品卡不同类型折扣列表
// @Tags 礼品卡不同类型折扣
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrdGiftcardDiscounts}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-discounts [get]
// @Security Bearer
func (e OrdGiftcardDiscounts) GetPage(c *gin.Context) {
    req := dto.OrdGiftcardDiscountsGetPageReq{}
    s := service.OrdGiftcardDiscounts{}
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
	list := make([]models.OrdGiftcardDiscounts, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡不同类型折扣失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取礼品卡不同类型折扣
// @Summary 获取礼品卡不同类型折扣
// @Description 获取礼品卡不同类型折扣
// @Tags 礼品卡不同类型折扣
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrdGiftcardDiscounts} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-discounts/{id} [get]
// @Security Bearer
func (e OrdGiftcardDiscounts) Get(c *gin.Context) {
	req := dto.OrdGiftcardDiscountsGetReq{}
	s := service.OrdGiftcardDiscounts{}
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
	var object models.OrdGiftcardDiscounts

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡不同类型折扣失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建礼品卡不同类型折扣
// @Summary 创建礼品卡不同类型折扣
// @Description 创建礼品卡不同类型折扣
// @Tags 礼品卡不同类型折扣
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdGiftcardDiscountsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-giftcard-discounts [post]
// @Security Bearer
func (e OrdGiftcardDiscounts) Insert(c *gin.Context) {
    req := dto.OrdGiftcardDiscountsInsertReq{}
    s := service.OrdGiftcardDiscounts{}
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
		e.Error(500, err, fmt.Sprintf("创建礼品卡不同类型折扣失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改礼品卡不同类型折扣
// @Summary 修改礼品卡不同类型折扣
// @Description 修改礼品卡不同类型折扣
// @Tags 礼品卡不同类型折扣
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdGiftcardDiscountsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-giftcard-discounts/{id} [put]
// @Security Bearer
func (e OrdGiftcardDiscounts) Update(c *gin.Context) {
    req := dto.OrdGiftcardDiscountsUpdateReq{}
    s := service.OrdGiftcardDiscounts{}
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
		e.Error(500, err, fmt.Sprintf("修改礼品卡不同类型折扣失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除礼品卡不同类型折扣
// @Summary 删除礼品卡不同类型折扣
// @Description 删除礼品卡不同类型折扣
// @Tags 礼品卡不同类型折扣
// @Param data body dto.OrdGiftcardDiscountsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-giftcard-discounts [delete]
// @Security Bearer
func (e OrdGiftcardDiscounts) Delete(c *gin.Context) {
    s := service.OrdGiftcardDiscounts{}
    req := dto.OrdGiftcardDiscountsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除礼品卡不同类型折扣失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
