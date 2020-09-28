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

	ApiGroup := router.Group("")
	{
		ApiGroup.GET("health", controller.Health)                   // done
		ApiGroup.GET("site-show", controller.SiteShow)              // done
		ApiGroup.GET("captcha/:uuid/:scene", controller.NewCaptcha) // done
	}

	// 路由管理
	AdminGroup := router.Group("admin/")
	{
		AdminGroup.GET("systems/admin-users/login", controller.Login)         // done
		AdminGroup.POST("systems/admin-users/logout", controller.Logout)      // done
		AdminGroup.GET("systems/roles/:roleId/checked", controller.RoleCheck) // done
		AdminGroup.GET("admin/systems/messages", controller.MessageList)
		AdminGroup.GET("admin/systems/admin-users/token", controller.Refresh) // done
		AdminGroup.GET("admin/systems/complain-topics", controller.ComplainTopicsList)
		AdminGroup.GET("index/page", controller.Index)
		AdminGroup.GET("admin/goods", controller.GoodsList)                                   // done
		AdminGroup.GET("admin/goods/specs", controller.SpecsList)                             // done
		AdminGroup.GET("admin/goods/brands", controller.BrandList)                            // done
		AdminGroup.GET("admin/goods/brands/all", controller.BrandAllList)                     // done
		AdminGroup.GET("admin/goods/categories/:parent_id/children", controller.CategoryList) // done
		AdminGroup.GET("admin/trade/orders", controller.OrderList)
		AdminGroup.GET("admin/after-sales", controller.AfterSalesList)
		AdminGroup.GET("admin/after-sales/refund", controller.AfterSalesRefundList)
		AdminGroup.GET("admin/trade/orders/pay-log", controller.OrderPayLogList)
		AdminGroup.GET("admin/trade/order-complains", controller.OrderComplainsList)
		AdminGroup.GET("admin/members/receipts", controller.MemberReceiptList)
		AdminGroup.GET("admin/members/zpzz", controller.ZpzzList)
		AdminGroup.GET("admin/members", controller.MemberList)
		AdminGroup.GET("admin/members/comments", controller.MemberCommentsList)
		AdminGroup.GET("admin/members/asks", controller.MemberAskList)

		AdminGroup.GET("admin/shops/list", controller.AllShopList) // done
		AdminGroup.GET("admin/shops", controller.ShopList)
		AdminGroup.GET("admin/shops/themes", controller.ShopThemesList)

		AdminGroup.GET("admin/order/bills/statistics", controller.OrderBillStatisticList)
		AdminGroup.GET("admin/promotion/group-buy-actives", controller.GroupBuyList)
		AdminGroup.GET("admin/promotion/group-buy-cats", controller.GroupBuyCategoryList)
		AdminGroup.GET("admin/promotion/exchange-cats/:cat_id/children", controller.PointCategory)
		AdminGroup.GET("admin/promotion/seckills", controller.SeckillList)
		AdminGroup.GET("admin/promotion/coupons", controller.CouponList)
		AdminGroup.GET("admin/promotion/pintuan", controller.PinTuanList)
		AdminGroup.GET("admin/pages/:client_type/:page_type", controller.Page)
		AdminGroup.GET("admin/focus-pictures", controller.FocusPicture)

		// 未发现
		AdminGroup.GET("admin/members/deposit/recharge", controller.MemberDepositRechargeList)
		AdminGroup.GET("admin/members/deposit/log", controller.MemberDepositLogList)
	}

	return router
}
