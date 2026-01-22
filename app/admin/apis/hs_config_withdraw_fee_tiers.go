package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type HsConfigWithdrawFeeTiers struct {
	api.Api
}

// GetPage 获取提现阶梯手续费配置列表
// @Summary 获取提现阶梯手续费配置列表
// @Description 获取提现阶梯手续费配置列表
// @Tags 提现阶梯手续费配置
// @Param ruleId query int false "关联的提现规则ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsConfigWithdrawFeeTiers}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-fee-tiers [get]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) GetPage(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeTiersGetPageReq{}
	s := service.HsConfigWithdrawFeeTiers{}
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

	list := make([]models.HsConfigWithdrawFeeTiers, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取提现阶梯手续费配置
// @Summary 获取提现阶梯手续费配置
// @Description 获取提现阶梯手续费配置
// @Tags 提现阶梯手续费配置
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsConfigWithdrawFeeTiers} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-fee-tiers/{id} [get]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) Get(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeTiersGetReq{}
	s := service.HsConfigWithdrawFeeTiers{}
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
	var object models.HsConfigWithdrawFeeTiers

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// GetByRuleId 根据规则ID获取阶梯手续费配置
// @Summary 根据规则ID获取阶梯手续费配置
// @Description 根据规则ID获取阶梯手续费配置列表
// @Tags 提现阶梯手续费配置
// @Param ruleId path int true "提现规则ID"
// @Success 200 {object} response.Response{data=[]models.HsConfigWithdrawFeeTiers} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-config-withdraw-fee-tiers/rule/{ruleId} [get]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) GetByRuleId(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeTiersByRuleIdReq{}
	s := service.HsConfigWithdrawFeeTiers{}
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

	list := make([]models.HsConfigWithdrawFeeTiers, 0)
	err = s.GetByRuleId(req.GetRuleId(), &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(list, "查询成功")
}

// Insert 创建提现阶梯手续费配置
// @Summary 创建提现阶梯手续费配置
// @Description 创建提现阶梯手续费配置
// @Tags 提现阶梯手续费配置
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigWithdrawFeeTiersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-config-withdraw-fee-tiers [post]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) Insert(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeTiersInsertReq{}
	s := service.HsConfigWithdrawFeeTiers{}
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

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改提现阶梯手续费配置
// @Summary 修改提现阶梯手续费配置
// @Description 修改提现阶梯手续费配置
// @Tags 提现阶梯手续费配置
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsConfigWithdrawFeeTiersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-config-withdraw-fee-tiers/{id} [put]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) Update(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeTiersUpdateReq{}
	s := service.HsConfigWithdrawFeeTiers{}
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

	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除提现阶梯手续费配置
// @Summary 删除提现阶梯手续费配置
// @Description 删除提现阶梯手续费配置
// @Tags 提现阶梯手续费配置
// @Param data body dto.HsConfigWithdrawFeeTiersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-config-withdraw-fee-tiers [delete]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) Delete(c *gin.Context) {
	s := service.HsConfigWithdrawFeeTiers{}
	req := dto.HsConfigWithdrawFeeTiersDeleteReq{}
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

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// BatchSave 批量保存提现阶梯手续费配置
// @Summary 批量保存提现阶梯手续费配置
// @Description 批量保存提现阶梯手续费配置（会先删除该规则下的所有旧配置）
// @Tags 提现阶梯手续费配置
// @Accept application/json
// @Product application/json
// @Param data body dto.HsConfigWithdrawFeeTiersBatchReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "保存成功"}"
// @Router /api/v1/hs-config-withdraw-fee-tiers/batch [post]
// @Security Bearer
func (e HsConfigWithdrawFeeTiers) BatchSave(c *gin.Context) {
	req := dto.HsConfigWithdrawFeeTiersBatchReq{}
	s := service.HsConfigWithdrawFeeTiers{}
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

	err = s.BatchSave(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("批量保存提现阶梯手续费配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(nil, "保存成功")
}
