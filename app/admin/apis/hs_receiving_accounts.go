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

type HsReceivingAccounts struct {
	api.Api
}

// GetPage 获取收款账号/设置表 (关联卡商与银行或通道)列表
// @Summary 获取收款账号/设置表 (关联卡商与银行或通道)列表
// @Description 获取收款账号/设置表 (关联卡商与银行或通道)列表
// @Tags 收款账号/设置表 (关联卡商与银行或通道)
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsReceivingAccounts}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-receiving-accounts [get]
// @Security Bearer
func (e HsReceivingAccounts) GetPage(c *gin.Context) {
    req := dto.HsReceivingAccountsGetPageReq{}
    s := service.HsReceivingAccounts{}
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
	list := make([]models.HsReceivingAccounts, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取收款账号/设置表 (关联卡商与银行或通道)失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取收款账号/设置表 (关联卡商与银行或通道)
// @Summary 获取收款账号/设置表 (关联卡商与银行或通道)
// @Description 获取收款账号/设置表 (关联卡商与银行或通道)
// @Tags 收款账号/设置表 (关联卡商与银行或通道)
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsReceivingAccounts} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-receiving-accounts/{id} [get]
// @Security Bearer
func (e HsReceivingAccounts) Get(c *gin.Context) {
	req := dto.HsReceivingAccountsGetReq{}
	s := service.HsReceivingAccounts{}
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
	var object models.HsReceivingAccounts

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取收款账号/设置表 (关联卡商与银行或通道)失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建收款账号/设置表 (关联卡商与银行或通道)
// @Summary 创建收款账号/设置表 (关联卡商与银行或通道)
// @Description 创建收款账号/设置表 (关联卡商与银行或通道)
// @Tags 收款账号/设置表 (关联卡商与银行或通道)
// @Accept application/json
// @Product application/json
// @Param data body dto.HsReceivingAccountsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-receiving-accounts [post]
// @Security Bearer
func (e HsReceivingAccounts) Insert(c *gin.Context) {
    req := dto.HsReceivingAccountsInsertReq{}
    s := service.HsReceivingAccounts{}
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
		e.Error(500, err, fmt.Sprintf("创建收款账号/设置表 (关联卡商与银行或通道)失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改收款账号/设置表 (关联卡商与银行或通道)
// @Summary 修改收款账号/设置表 (关联卡商与银行或通道)
// @Description 修改收款账号/设置表 (关联卡商与银行或通道)
// @Tags 收款账号/设置表 (关联卡商与银行或通道)
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsReceivingAccountsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-receiving-accounts/{id} [put]
// @Security Bearer
func (e HsReceivingAccounts) Update(c *gin.Context) {
    req := dto.HsReceivingAccountsUpdateReq{}
    s := service.HsReceivingAccounts{}
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
		e.Error(500, err, fmt.Sprintf("修改收款账号/设置表 (关联卡商与银行或通道)失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除收款账号/设置表 (关联卡商与银行或通道)
// @Summary 删除收款账号/设置表 (关联卡商与银行或通道)
// @Description 删除收款账号/设置表 (关联卡商与银行或通道)
// @Tags 收款账号/设置表 (关联卡商与银行或通道)
// @Param data body dto.HsReceivingAccountsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-receiving-accounts [delete]
// @Security Bearer
func (e HsReceivingAccounts) Delete(c *gin.Context) {
    s := service.HsReceivingAccounts{}
    req := dto.HsReceivingAccountsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除收款账号/设置表 (关联卡商与银行或通道)失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
