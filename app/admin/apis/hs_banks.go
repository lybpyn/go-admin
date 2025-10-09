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

type HsBanks struct {
	api.Api
}

// GetPage 获取银行列表列表
// @Summary 获取银行列表列表
// @Description 获取银行列表列表
// @Tags 银行列表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.HsBanks}} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-banks [get]
// @Security Bearer
func (e HsBanks) GetPage(c *gin.Context) {
    req := dto.HsBanksGetPageReq{}
    s := service.HsBanks{}
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
	list := make([]models.HsBanks, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取银行列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取银行列表
// @Summary 获取银行列表
// @Description 获取银行列表
// @Tags 银行列表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.HsBanks} "{"code": 200, "data": [...]}"
// @Router /api/v1/hs-banks/{id} [get]
// @Security Bearer
func (e HsBanks) Get(c *gin.Context) {
	req := dto.HsBanksGetReq{}
	s := service.HsBanks{}
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
	var object models.HsBanks

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取银行列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建银行列表
// @Summary 创建银行列表
// @Description 创建银行列表
// @Tags 银行列表
// @Accept application/json
// @Product application/json
// @Param data body dto.HsBanksInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/hs-banks [post]
// @Security Bearer
func (e HsBanks) Insert(c *gin.Context) {
    req := dto.HsBanksInsertReq{}
    s := service.HsBanks{}
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
		e.Error(500, err, fmt.Sprintf("创建银行列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改银行列表
// @Summary 修改银行列表
// @Description 修改银行列表
// @Tags 银行列表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.HsBanksUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/hs-banks/{id} [put]
// @Security Bearer
func (e HsBanks) Update(c *gin.Context) {
    req := dto.HsBanksUpdateReq{}
    s := service.HsBanks{}
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
		e.Error(500, err, fmt.Sprintf("修改银行列表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除银行列表
// @Summary 删除银行列表
// @Description 删除银行列表
// @Tags 银行列表
// @Param data body dto.HsBanksDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/hs-banks [delete]
// @Security Bearer
func (e HsBanks) Delete(c *gin.Context) {
    s := service.HsBanks{}
    req := dto.HsBanksDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除银行列表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
