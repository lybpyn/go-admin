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

type HsConfigWithdrawFee struct {
	api.Api
}

// GetPage 获取提现费率与限制配置表列表
// @Summary 获取提现费率与限制配置表列表
// @Description 获取提现费率与限制配置表列表
// @Tags 提现费率与限制配置表
// @Param currencyCode query string false "币种代码，如 USD/CNY/USDT"
// @Param minAmount query string false "最小提现金额"
// @Param maxAmount query string false "最大提现金额"
// @Param feeRate query string false "手续费率（例如 0.015 表示 1.5%）"
// @Param isActive query string false "是否启用：1=启用，0=禁用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsConfigWithdrawFee}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-fee [get]
// @Security Bearer
func (e HsConfigWithdrawFee) GetPage(c *gin.Context) {
    req := dto.HsConfigWithdrawFeeGetPageReq{}
    s := service.HsConfigWithdrawFee{}
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
	list := make([]models.HsConfigWithdrawFee, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现费率与限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取提现费率与限制配置表
// @Summary 获取提现费率与限制配置表
// @Description 获取提现费率与限制配置表
// @Tags 提现费率与限制配置表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsConfigWithdrawFee} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-fee/{id} [get]
// @Security Bearer
func (e HsConfigWithdrawFee) Get(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeGetReq{}
	s := service.HsConfigWithdrawFee{}
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
	var object models.HsConfigWithdrawFee

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现费率与限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建提现费率与限制配置表
// @Summary 创建提现费率与限制配置表
// @Description 创建提现费率与限制配置表
// @Tags 提现费率与限制配置表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigWithdrawFeeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-withdraw-fee [post]
// @Security Bearer
func (e HsConfigWithdrawFee) Insert(c *gin.Context) {
    req := dto.HsConfigWithdrawFeeInsertReq{}
    s := service.HsConfigWithdrawFee{}
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
		e.Error(500, err, fmt.Sprintf("创建提现费率与限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改提现费率与限制配置表
// @Summary 修改提现费率与限制配置表
// @Description 修改提现费率与限制配置表
// @Tags 提现费率与限制配置表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigWithdrawFeeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-withdraw-fee/{id} [put]
// @Security Bearer
func (e HsConfigWithdrawFee) Update(c *gin.Context) {
    req := dto.HsConfigWithdrawFeeUpdateReq{}
    s := service.HsConfigWithdrawFee{}
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
		e.Error(500, err, fmt.Sprintf("修改提现费率与限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除提现费率与限制配置表
// @Summary 删除提现费率与限制配置表
// @Description 删除提现费率与限制配置表
// @Tags 提现费率与限制配置表
// @Param data body dto.HsConfigWithdrawFeeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-withdraw-fee [delete]
// @Security Bearer
func (e HsConfigWithdrawFee) Delete(c *gin.Context) {
    s := service.HsConfigWithdrawFee{}
    req := dto.HsConfigWithdrawFeeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除提现费率与限制配置表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
