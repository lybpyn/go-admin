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

type OrdConfigCurrencyRates struct {
	api.Api
}

// GetPage 获取货币汇率表-支持多币种对和地区化配置列表
// @Summary 获取货币汇率表-支持多币种对和地区化配置列表
// @Description 获取货币汇率表-支持多币种对和地区化配置列表
// @Tags 货币汇率表-支持多币种对和地区化配置
// @Param baseCurrencyCode query string false "基准货币代码 (ISO 4217)"
// @Param quoteCurrencyCode query string false "报价货币代码 (ISO 4217)"
// @Param rate query string false "汇率: 1 base_currency = rate quote_currency"
// @Param regionCode query string false "地区代码,为空表示全局汇率"
// @Param rateType query string false "汇率类型: standard=标准, buying=买入, selling=卖出"
// @Param source query string false "汇率来源,如 manual, api, coingecko"
// @Param status query string false "状态: 1=启用, 0=禁用"
// @Param validFrom query time.Time false "生效开始时间"
// @Param validTo query time.Time false "生效结束时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.OrdConfigCurrencyRates}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-config-currency-rates [get]
// @Security Bearer
func (e OrdConfigCurrencyRates) GetPage(c *gin.Context) {
    req := dto.OrdConfigCurrencyRatesGetPageReq{}
    s := service.OrdConfigCurrencyRates{}
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.OrdConfigCurrencyRates, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取货币汇率表-支持多币种对和地区化配置失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取货币汇率表-支持多币种对和地区化配置
// @Summary 获取货币汇率表-支持多币种对和地区化配置
// @Description 获取货币汇率表-支持多币种对和地区化配置
// @Tags 货币汇率表-支持多币种对和地区化配置
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.OrdConfigCurrencyRates} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-config-currency-rates/{id} [get]
// @Security Bearer
func (e OrdConfigCurrencyRates) Get(c *gin.Context) {
	req := dto.OrdConfigCurrencyRatesGetReq{}
	s := service.OrdConfigCurrencyRates{}
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
	var object models.OrdConfigCurrencyRates

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取货币汇率表-支持多币种对和地区化配置失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建货币汇率表-支持多币种对和地区化配置
// @Summary 创建货币汇率表-支持多币种对和地区化配置
// @Description 创建货币汇率表-支持多币种对和地区化配置
// @Tags 货币汇率表-支持多币种对和地区化配置
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdConfigCurrencyRatesInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-config-currency-rates [post]
// @Security Bearer
func (e OrdConfigCurrencyRates) Insert(c *gin.Context) {
    req := dto.OrdConfigCurrencyRatesInsertReq{}
    s := service.OrdConfigCurrencyRates{}
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
		e.Error(500, err, fmt.Sprintf("创建货币汇率表-支持多币种对和地区化配置失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改货币汇率表-支持多币种对和地区化配置
// @Summary 修改货币汇率表-支持多币种对和地区化配置
// @Description 修改货币汇率表-支持多币种对和地区化配置
// @Tags 货币汇率表-支持多币种对和地区化配置
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdConfigCurrencyRatesUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-config-currency-rates/{id} [put]
// @Security Bearer
func (e OrdConfigCurrencyRates) Update(c *gin.Context) {
    req := dto.OrdConfigCurrencyRatesUpdateReq{}
    s := service.OrdConfigCurrencyRates{}
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
		e.Error(500, err, fmt.Sprintf("修改货币汇率表-支持多币种对和地区化配置失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除货币汇率表-支持多币种对和地区化配置
// @Summary 删除货币汇率表-支持多币种对和地区化配置
// @Description 删除货币汇率表-支持多币种对和地区化配置
// @Tags 货币汇率表-支持多币种对和地区化配置
// @Param data body dto.OrdConfigCurrencyRatesDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-config-currency-rates [delete]
// @Security Bearer
func (e OrdConfigCurrencyRates) Delete(c *gin.Context) {
    s := service.OrdConfigCurrencyRates{}
    req := dto.OrdConfigCurrencyRatesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除货币汇率表-支持多币种对和地区化配置失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}

// BatchQuery 批量查询汇率
// @Summary 批量查询汇率
// @Description 根据多个货币对批量查询汇率信息，最多支持50个货币对
// @Tags 货币汇率表-支持多币种对和地区化配置
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdConfigCurrencyRatesBatchQueryReq true "批量查询请求参数"
// @Success 200 {object} models.Response{data=dto.OrdConfigCurrencyRatesBatchQueryResp} "{"code": 200, "data": {...}}"
// @Router /api/v1/ord-config-currency-rates/batch-query [post]
// @Security Bearer
func (e OrdConfigCurrencyRates) BatchQuery(c *gin.Context) {
    req := dto.OrdConfigCurrencyRatesBatchQueryReq{}
    s := service.OrdConfigCurrencyRates{}
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

	resp, err := s.BatchQuery(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("批量查询汇率失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(resp, "查询成功")
}
