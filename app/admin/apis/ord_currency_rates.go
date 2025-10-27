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

type OrdCurrencyRates struct {
	api.Api
}

// GetPage 获取货币对汇率表列表
// @Summary 获取货币对汇率表列表
// @Description 获取货币对汇率表列表
// @Tags 货币对汇率表
// @Param quoteCurrency query string false "计价货币 ISO4217，例如 CNY"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.OrdCurrencyRates}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-currency-rates [get]
// @Security Bearer
func (e OrdCurrencyRates) GetPage(c *gin.Context) {
    req := dto.OrdCurrencyRatesGetPageReq{}
    s := service.OrdCurrencyRates{}
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
	list := make([]models.OrdCurrencyRates, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取货币对汇率表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取货币对汇率表
// @Summary 获取货币对汇率表
// @Description 获取货币对汇率表
// @Tags 货币对汇率表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.OrdCurrencyRates} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-currency-rates/{id} [get]
// @Security Bearer
func (e OrdCurrencyRates) Get(c *gin.Context) {
	req := dto.OrdCurrencyRatesGetReq{}
	s := service.OrdCurrencyRates{}
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
	var object models.OrdCurrencyRates

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取货币对汇率表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建货币对汇率表
// @Summary 创建货币对汇率表
// @Description 创建货币对汇率表
// @Tags 货币对汇率表
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdCurrencyRatesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-currency-rates [post]
// @Security Bearer
func (e OrdCurrencyRates) Insert(c *gin.Context) {
    req := dto.OrdCurrencyRatesInsertReq{}
    s := service.OrdCurrencyRates{}
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
		e.Error(500, err, fmt.Sprintf("创建货币对汇率表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改货币对汇率表
// @Summary 修改货币对汇率表
// @Description 修改货币对汇率表
// @Tags 货币对汇率表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdCurrencyRatesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-currency-rates/{id} [put]
// @Security Bearer
func (e OrdCurrencyRates) Update(c *gin.Context) {
    req := dto.OrdCurrencyRatesUpdateReq{}
    s := service.OrdCurrencyRates{}
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
		e.Error(500, err, fmt.Sprintf("修改货币对汇率表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除货币对汇率表
// @Summary 删除货币对汇率表
// @Description 删除货币对汇率表
// @Tags 货币对汇率表
// @Param data body dto.OrdCurrencyRatesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-currency-rates [delete]
// @Security Bearer
func (e OrdCurrencyRates) Delete(c *gin.Context) {
    s := service.OrdCurrencyRates{}
    req := dto.OrdCurrencyRatesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除货币对汇率表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
