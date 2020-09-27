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
		AdminGroup.GET("admin/admin/systems/admin-users/token", controller.Refresh) // done
		AdminGroup.GET("admin/index/page", controller.Index)
		AdminGroup.GET("admin/admin/goods", controller.GoodsList)                                   // done
		AdminGroup.GET("admin/admin/goods/specs", controller.SpecsList)                             // done
		AdminGroup.GET("admin/admin/shops/list", controller.ShopList)                               // done
		AdminGroup.GET("admin/admin/goods/brands", controller.BrandList)                            // done
		AdminGroup.GET("admin/admin/goods/brands/all", controller.BrandAllList)                     // done
		AdminGroup.GET("admin/admin/goods/categories/:parent_id/children", controller.CategoryList) // done

		AdminGroup.GET("admin/admin/trade/orders", controller.OrderList)
		AdminGroup.GET("admin/admin/after-sales", controller.AfterSalesList)
		AdminGroup.GET("admin/admin/after-sales/refund", controller.AfterSalesRefundList)
		AdminGroup.GET("admin/admin/trade/orders/pay-log", controller.TradeOrderPayLogList)
		AdminGroup.GET("admin/admin/members/receipts", controller.MemberReceiptList)
		AdminGroup.GET("admin/admin/members/zpzz", controller.ZpzzList)
		AdminGroup.GET("admin/admin/trade/order-complains", controller.OrderComplainsList)
		AdminGroup.GET("admin/admin/systems/complain-topics", controller.ComplainTopicsList)
	}

	return router
}
