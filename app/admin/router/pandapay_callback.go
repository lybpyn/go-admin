package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	// 回调接口不需要认证，注册到 routerNoCheckRole
	routerNoCheckRole = append(routerNoCheckRole, registerPandaPayCallbackNoAuthRouter)
	// 测试接口需要认证，注册到 routerCheckRole
	routerCheckRole = append(routerCheckRole, registerPandaPayCallbackAuthRouter)
}

// registerPandaPayCallbackNoAuthRouter 注册无需认证的PandaPay回调路由
func registerPandaPayCallbackNoAuthRouter(v1 *gin.RouterGroup) {
	api := apis.PandaPayCallback{}

	// 回调接口完全不需要任何认证
	r := v1.Group("/callback")
	{
		r.POST("/pandapay", api.HandleCallback)
	}
}

// registerPandaPayCallbackAuthRouter 注册需要认证的PandaPay测试路由
func registerPandaPayCallbackAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.PandaPayCallback{}

	// 测试接口需要认证
	rAuth := v1.Group("/callback").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		rAuth.GET("/pandapay/test", api.GetCallbackTest)
	}
}

