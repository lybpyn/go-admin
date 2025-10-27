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

type OrdGiftcardCategory struct {
	api.Api
}

// GetPage 获取礼品卡分类列表
// @Summary 获取礼品卡分类列表
// @Description 获取礼品卡分类列表
// @Tags 礼品卡分类
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrdGiftcardCategory}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-category [get]
// @Security Bearer
func (e OrdGiftcardCategory) GetPage(c *gin.Context) {
    req := dto.OrdGiftcardCategoryGetPageReq{}
    s := service.OrdGiftcardCategory{}
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
	list := make([]models.OrdGiftcardCategory, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡分类失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取礼品卡分类
// @Summary 获取礼品卡分类
// @Description 获取礼品卡分类
// @Tags 礼品卡分类
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrdGiftcardCategory} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-category/{id} [get]
// @Security Bearer
func (e OrdGiftcardCategory) Get(c *gin.Context) {
	req := dto.OrdGiftcardCategoryGetReq{}
	s := service.OrdGiftcardCategory{}
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
	var object models.OrdGiftcardCategory

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡分类失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建礼品卡分类
// @Summary 创建礼品卡分类
// @Description 创建礼品卡分类
// @Tags 礼品卡分类
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdGiftcardCategoryInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-giftcard-category [post]
// @Security Bearer
func (e OrdGiftcardCategory) Insert(c *gin.Context) {
    req := dto.OrdGiftcardCategoryInsertReq{}
    s := service.OrdGiftcardCategory{}
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
		e.Error(500, err, fmt.Sprintf("创建礼品卡分类失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改礼品卡分类
// @Summary 修改礼品卡分类
// @Description 修改礼品卡分类
// @Tags 礼品卡分类
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdGiftcardCategoryUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-giftcard-category/{id} [put]
// @Security Bearer
func (e OrdGiftcardCategory) Update(c *gin.Context) {
    req := dto.OrdGiftcardCategoryUpdateReq{}
    s := service.OrdGiftcardCategory{}
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
		e.Error(500, err, fmt.Sprintf("修改礼品卡分类失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除礼品卡分类
// @Summary 删除礼品卡分类
// @Description 删除礼品卡分类
// @Tags 礼品卡分类
// @Param data body dto.OrdGiftcardCategoryDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-giftcard-category [delete]
// @Security Bearer
func (e OrdGiftcardCategory) Delete(c *gin.Context) {
    s := service.OrdGiftcardCategory{}
    req := dto.OrdGiftcardCategoryDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除礼品卡分类失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
