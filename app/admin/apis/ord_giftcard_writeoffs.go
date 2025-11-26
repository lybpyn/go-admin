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

type OrdGiftcardWriteoffs struct {
	api.Api
}

// GetPage 获取礼品卡核销记录表列表
// @Summary 获取礼品卡核销记录表列表
// @Description 获取礼品卡核销记录表列表
// @Tags 礼品卡核销记录表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response{data=models.Page{list=[]models.OrdGiftcardWriteoffs}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-writeoffs [get]
// @Security Bearer
func (e OrdGiftcardWriteoffs) GetPage(c *gin.Context) {
	req := dto.OrdGiftcardWriteoffsGetPageReq{}
	s := service.OrdGiftcardWriteoffs{}
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
	list := make([]models.OrdGiftcardWriteoffs, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡核销记录表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取礼品卡核销记录表
// @Summary 获取礼品卡核销记录表
// @Description 获取礼品卡核销记录表
// @Tags 礼品卡核销记录表
// @Param id path int false "id"
// @Success 200 {object} models.Response{data=models.OrdGiftcardWriteoffs} "{"code": 200, "data": [...]}"
// @Router /api/v1/ord-giftcard-writeoffs/{id} [get]
// @Security Bearer
func (e OrdGiftcardWriteoffs) Get(c *gin.Context) {
	req := dto.OrdGiftcardWriteoffsGetReq{}
	s := service.OrdGiftcardWriteoffs{}
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
	var object models.OrdGiftcardWriteoffs

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取礼品卡核销记录表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建礼品卡核销记录表
// @Summary 创建礼品卡核销记录表
// @Description 创建礼品卡核销记录表
// @Tags 礼品卡核销记录表
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdGiftcardWriteoffsInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ord-giftcard-writeoffs [post]
// @Security Bearer
func (e OrdGiftcardWriteoffs) Insert(c *gin.Context) {
	return
	req := dto.OrdGiftcardWriteoffsInsertReq{}
	s := service.OrdGiftcardWriteoffs{}
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
		e.Error(500, err, fmt.Sprintf("创建礼品卡核销记录表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改礼品卡核销记录表
// @Summary 修改礼品卡核销记录表
// @Description 修改礼品卡核销记录表
// @Tags 礼品卡核销记录表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.OrdGiftcardWriteoffsUpdateReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ord-giftcard-writeoffs/{id} [put]
// @Security Bearer
func (e OrdGiftcardWriteoffs) Update(c *gin.Context) {
	return
	req := dto.OrdGiftcardWriteoffsUpdateReq{}
	s := service.OrdGiftcardWriteoffs{}
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
		e.Error(500, err, fmt.Sprintf("修改礼品卡核销记录表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除礼品卡核销记录表
// @Summary 删除礼品卡核销记录表
// @Description 删除礼品卡核销记录表
// @Tags 礼品卡核销记录表
// @Param data body dto.OrdGiftcardWriteoffsDeleteReq true "body"
// @Success 200 {object} models.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ord-giftcard-writeoffs [delete]
// @Security Bearer
func (e OrdGiftcardWriteoffs) Delete(c *gin.Context) {
	return
	s := service.OrdGiftcardWriteoffs{}
	req := dto.OrdGiftcardWriteoffsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除礼品卡核销记录表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// BatchInsert 批量创建礼品卡核销记录
// @Summary 批量创建礼品卡核销记录
// @Description 管理员批量核销礼品卡订单，可以针对同一个订单创建多条核销记录
// @Tags 礼品卡核销记录表
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdGiftcardWriteoffsBatchInsertReq true "data"
// @Success 200 {object} models.Response	"{"code": 200, "message": "批量核销成功"}"
// @Router /api/v1/ord-giftcard-writeoffs/batch [post]
// @Security Bearer
func (e OrdGiftcardWriteoffs) BatchInsert(c *gin.Context) {
	req := dto.OrdGiftcardWriteoffsBatchInsertReq{}
	
	s := service.OrdGiftcardWriteoffs{}
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

	err = s.BatchInsert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("批量创建礼品卡核销记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(nil, fmt.Sprintf("批量核销成功，共创建 %d 条记录", len(req.WriteoffList)))
}

// CalculateUserLocalCurrency 计算用户入账金额（辅助接口）
// @Summary 计算用户入账金额
// @Description 根据订单ID、卡片面值和折扣率，计算用户将获得的本地货币金额。计算公式：用户入账金额 = 卡片面值 × 折扣率 × 汇率。可选传入折扣率参数，如果不传则使用礼品卡折扣配置中的折扣率
// @Tags 礼品卡核销记录表
// @Accept application/json
// @Product application/json
// @Param data body dto.OrdGiftcardWriteoffsCalculateReq true "data"
// @Success 200 {object} models.Response{data=dto.OrdGiftcardWriteoffsCalculateResp} "{"code": 200, "data": {...}}"
// @Router /api/v1/ord-giftcard-writeoffs/calculate [post]
// @Security Bearer
func (e OrdGiftcardWriteoffs) CalculateUserLocalCurrency(c *gin.Context) {
	req := dto.OrdGiftcardWriteoffsCalculateReq{}
	s := service.OrdGiftcardWriteoffs{}
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

	resp, err := s.CalculateUserLocalCurrency(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("计算用户入账金额失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(resp, "计算成功")
}
