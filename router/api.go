package router

import (
	"orange/controller"
	"orange/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 中间件管理
	router.Use(middleware.Cors())
	//router.Use(middleware.JwtMiddleWare())

	// 路由管理
	AdminGroup := router.Group("")
	{
		AdminGroup.GET("health", controller.Health)                                 // done
		AdminGroup.GET("site-show", controller.SiteShow)                            // done
		AdminGroup.GET("captcha/:uuid/:scene", controller.NewCaptcha)               // done
		AdminGroup.GET("admin/systems/admin-users/login", controller.Login)         // done
		AdminGroup.GET("admin/systems/roles/:roleId/checked", controller.RoleCheck) // done
		AdminGroup.POST("admin/systems/admin-users/logout", controller.Logout)      // done
		AdminGroup.GET("admin/index/page", controller.Index)
		AdminGroup.GET("admin/admin/goods", controller.GoodsList)
		AdminGroup.GET("admin/admin/shops/list", controller.ShopList)        // done
		AdminGroup.GET("admin/admin/goods/brands/all", controller.BrandList) // done
		AdminGroup.GET("admin/admin/goods/categories/0/children", controller.CategoryList)
	}

	return router
}
