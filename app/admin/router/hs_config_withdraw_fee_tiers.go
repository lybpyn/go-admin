package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerHsConfigWithdrawFeeTiersRouter)
}

// registerHsConfigWithdrawFeeTiersRouter 注册提现阶梯手续费配置路由
func registerHsConfigWithdrawFeeTiersRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.HsConfigWithdrawFeeTiers{}
	r := v1.Group("/hs-config-withdraw-fee-tiers").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.GET("/rule/:ruleId", api.GetByRuleId)
		r.POST("", api.Insert)
		r.POST("/batch", api.BatchSave)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
