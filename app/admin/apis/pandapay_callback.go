package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type PandaPayCallback struct {
	api.Api
}

// HandleCallback 处理PandaPay回调通知
// @Summary 处理PandaPay回调通知
// @Description 接收PandaPay的支付回调通知，更新提现订单状态
// @Tags 回调接口
// @Accept json
// @Produce json
// @Param data body dto.PandaPayCallbackReq true "回调数据"
// @Success 200 {object} dto.PandaPayCallbackResp "成功"
// @Failure 400 {object} dto.PandaPayCallbackResp "请求参数错误"
// @Failure 500 {object} dto.PandaPayCallbackResp "服务器内部错误"
// @Router /api/v1/callback/pandapay [post]
func (e *PandaPayCallback) HandleCallback(c *gin.Context) {
	req := dto.PandaPayCallbackReq{}
	s := service.PandaPayCallback{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(http.StatusBadRequest, err, "参数绑定失败")
		c.JSON(http.StatusOK, dto.PandaPayCallbackResp{
			Code:    1,
			Message: "参数绑定失败",
		})
		return
	}

	// 记录回调日志
	e.Logger.Infof("Received PandaPay callback: merchantOrderId=%s, status=%d, orderId=%d",
		req.MerchantOrderId, req.Status, req.OrderId)

	// 处理回调
	err = s.HandleCallback(&req)
	if err != nil {
		e.Logger.Errorf("Handle PandaPay callback failed: %s", err)
		c.JSON(http.StatusOK, dto.PandaPayCallbackResp{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, dto.PandaPayCallbackResp{
		Code:    0,
		Message: "success",
	})
}

// GetCallbackTest 测试回调接口（仅用于开发调试）
// @Summary 测试回调接口
// @Description 测试PandaPay回调接口是否正常工作
// @Tags 回调接口
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {string} string "ok"
// @Router /api/v1/callback/pandapay/test [get]
func (e *PandaPayCallback) GetCallbackTest(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(map[string]interface{}{
		"status":  "ok",
		"message": "PandaPay回调接口已就绪",
	}, "测试成功")
}

