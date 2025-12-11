package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPandaPayCallbackRouter)
}

// registerPandaPayCallbackRouter 注册PandaPay回调路由
func registerPandaPayCallbackRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.PandaPayCallback{}

	// 回调接口不需要鉴权
	r := v1.Group("/callback").Use(middleware.AuthCheckRole())
	{
		r.POST("/pandapay", api.HandleCallback)
	}

	// 测试接口需要鉴权（可选）
	rAuth := v1.Group("/callback").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		rAuth.GET("/pandapay/test", api.GetCallbackTest)
	}
}
