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

type HsConfigWithdrawRules struct {
	api.Api
}

// GetPage 获取提现规则配置表（兼容法币与虚拟币）列表
// @Summary 获取提现规则配置表（兼容法币与虚拟币）列表
// @Description 获取提现规则配置表（兼容法币与虚拟币）列表
// @Tags 提现规则配置表（兼容法币与虚拟币）
// @Param currencyCode query string false "币种代码，如 USD、CNY、USDT、BTC"
// @Param currencyType query string false "币种类型：fiat=法币，crypto=虚拟币"
// @Param chainType query string false "链类型（仅虚拟币适用），如 ERC20/TRC20/BEP20"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsConfigWithdrawRules}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-rules [get]
// @Security Bearer
func (e HsConfigWithdrawRules) GetPage(c *gin.Context) {
    req := dto.HsConfigWithdrawRulesGetPageReq{}
    s := service.HsConfigWithdrawRules{}
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
	list := make([]models.HsConfigWithdrawRules, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现规则配置表（兼容法币与虚拟币）失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取提现规则配置表（兼容法币与虚拟币）
// @Summary 获取提现规则配置表（兼容法币与虚拟币）
// @Description 获取提现规则配置表（兼容法币与虚拟币）
// @Tags 提现规则配置表（兼容法币与虚拟币）
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsConfigWithdrawRules} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-rules/{id} [get]
// @Security Bearer
func (e HsConfigWithdrawRules) Get(c *gin.Context) {
	req := dto.HsConfigWithdrawRulesGetReq{}
	s := service.HsConfigWithdrawRules{}
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
	var object models.HsConfigWithdrawRules

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现规则配置表（兼容法币与虚拟币）失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建提现规则配置表（兼容法币与虚拟币）
// @Summary 创建提现规则配置表（兼容法币与虚拟币）
// @Description 创建提现规则配置表（兼容法币与虚拟币）
// @Tags 提现规则配置表（兼容法币与虚拟币）
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigWithdrawRulesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-withdraw-rules [post]
// @Security Bearer
func (e HsConfigWithdrawRules) Insert(c *gin.Context) {
    req := dto.HsConfigWithdrawRulesInsertReq{}
    s := service.HsConfigWithdrawRules{}
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
		e.Error(500, err, fmt.Sprintf("创建提现规则配置表（兼容法币与虚拟币）失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改提现规则配置表（兼容法币与虚拟币）
// @Summary 修改提现规则配置表（兼容法币与虚拟币）
// @Description 修改提现规则配置表（兼容法币与虚拟币）
// @Tags 提现规则配置表（兼容法币与虚拟币）
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigWithdrawRulesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-withdraw-rules/{id} [put]
// @Security Bearer
func (e HsConfigWithdrawRules) Update(c *gin.Context) {
    req := dto.HsConfigWithdrawRulesUpdateReq{}
    s := service.HsConfigWithdrawRules{}
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
		e.Error(500, err, fmt.Sprintf("修改提现规则配置表（兼容法币与虚拟币）失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除提现规则配置表（兼容法币与虚拟币）
// @Summary 删除提现规则配置表（兼容法币与虚拟币）
// @Description 删除提现规则配置表（兼容法币与虚拟币）
// @Tags 提现规则配置表（兼容法币与虚拟币）
// @Param data body dto.HsConfigWithdrawRulesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-withdraw-rules [delete]
// @Security Bearer
func (e HsConfigWithdrawRules) Delete(c *gin.Context) {
    s := service.HsConfigWithdrawRules{}
    req := dto.HsConfigWithdrawRulesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除提现规则配置表（兼容法币与虚拟币）失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
