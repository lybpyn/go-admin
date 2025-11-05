package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
	"go-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerHsUserWithdrawalRouter)
}

// registerHsUserWithdrawalRouter
func registerHsUserWithdrawalRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.HsUserWithdrawal{}
	r := v1.Group("/hs-user-withdrawal").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)

		// 接单相关路由
		r.GET("/available", api.GetAvailable)         // 获取可接单列表
		r.GET("/my-orders", api.GetMyOrders)          // 获取我的处理订单
		r.POST("/:id/claim", api.Claim)               // 接单
		r.POST("/:id/release", api.Release)           // 释放订单
	}
}