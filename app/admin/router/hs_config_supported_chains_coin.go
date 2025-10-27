package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
	"go-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerHsConfigSupportedChainsCoinRouter)
}

// registerHsConfigSupportedChainsCoinRouter
func registerHsConfigSupportedChainsCoinRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.HsConfigSupportedChainsCoin{}
	r := v1.Group("/hs-config-supported-chains-coin").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}