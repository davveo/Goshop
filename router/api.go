package router

import (
	"Goshop/controller"
	"Goshop/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 中间件管理
	router.Use(middleware.Cors())
	router.Use(middleware.NoCache())
	router.Use(middleware.RequestId())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "The incorrect API route.",
			"data":    nil,
		})
	})

	ApiGroup := router.Group("")
	{
		ApiGroup.GET("health", controller.Health)                           // done
		ApiGroup.GET("site-show", controller.SiteShow)                      // done
		ApiGroup.GET("captcha/:uuid/:scene", controller.NewCaptcha)         // done
		ApiGroup.GET("base/regions/:region_id/children", controller.Region) // wait to do
	}

	// 商品相关
	AdminGoodsGroup := router.Group("admin/admin/goods")
	{
		AdminGoodsGroup.GET("", controller.GoodsList)                                  // done
		AdminGoodsGroup.GET("specs", controller.SpecsList)                             // done
		AdminGoodsGroup.GET("brands", controller.BrandList)                            // done
		AdminGoodsGroup.GET("brands/all", controller.BrandAllList)                     // done
		AdminGoodsGroup.GET("categories/:parent_id/children", controller.CategoryList) // done
		AdminGoodsGroup.GET("settings", controller.GoodsSetting)
	}

	// 会员相关
	AdminMemberGroup := router.Group("admin/admin/members")
	{
		AdminMemberGroup.GET("", controller.MemberList)
		AdminMemberGroup.GET("zpzz", controller.ZpzzList)
		AdminMemberGroup.GET("asks", controller.MemberAskList)
		AdminMemberGroup.GET("receipts", controller.MemberReceiptList)
		AdminMemberGroup.GET("comments", controller.MemberCommentsList)
		// 未发现
		AdminMemberGroup.GET("deposit/log", controller.MemberDepositLogList)
		AdminMemberGroup.GET("deposit/recharge", controller.MemberDepositRechargeList)
	}

	// 促销相关
	AdminPromotionGroup := router.Group("admin/admin/promotion")
	{
		AdminPromotionGroup.GET("coupons", controller.CouponList)
		AdminPromotionGroup.GET("pintuan", controller.PinTuanList)
		AdminPromotionGroup.GET("seckills", controller.SeckillList)
		AdminPromotionGroup.GET("group-buy-actives", controller.GroupBuyList)
		AdminPromotionGroup.GET("group-buy-cats", controller.GroupBuyCategoryList)
		AdminPromotionGroup.GET("exchange-cats/:cat_id/children", controller.PointCategory)
	}

	// 数据统计相关
	AdminStatisticGroup := router.Group("admin/admin/statistics")
	{
		AdminStatisticGroup.GET("member/order/quantity", controller.StatisticMemberOrderQuantity)          // wait to do
		AdminStatisticGroup.GET("member/order/quantity/page", controller.StatisticMemberOrderQuantityPage) // wait to do
		AdminStatisticGroup.GET("member/increase/member", controller.StatisticMemberIncrease)              // wait to do
		AdminStatisticGroup.GET("member/increase/member/page", controller.StatisticMemberIncreasePage)     // wait to do
		AdminStatisticGroup.GET("goods/price/sales", controller.StatisticGoodsPrice)                       // wait to do
		AdminStatisticGroup.GET("goods/hot/money", controller.StatisticGoodsHot)                           // wait to do
		AdminStatisticGroup.GET("goods/hot/money/page", controller.StatisticGoodsHotPage)                  // wait to do
		AdminStatisticGroup.GET("goods/sale/details", controller.StatisticGoodsSaleDetail)                 // wait to do
		AdminStatisticGroup.GET("goods/collect", controller.StatisticGoodsCollect)                         // wait to do
		AdminStatisticGroup.GET("goods/collect/page", controller.StatisticGoodsCollectPage)                // wait to do
		AdminStatisticGroup.GET("industry/order/quantity", controller.StatisticIndustry)                   // wait to do
		AdminStatisticGroup.GET("industry/overview", controller.StatisticIndustryOverView)                 // wait to do
		AdminStatisticGroup.GET("page_view/shop", controller.StatisticPageViewShop)                        // wait to do
		AdminStatisticGroup.GET("page_view/goods", controller.StatisticPageViewGoods)                      // wait to do
		AdminStatisticGroup.GET("order/order/page", controller.StatisticPageOrder)                         // wait to do
		AdminStatisticGroup.GET("order/order/money", controller.StatisticPageMoney)                        // wait to do
		AdminStatisticGroup.GET("order/sales/money", controller.StatisticOrderSalesMoney)                  // wait to do
		AdminStatisticGroup.GET("order/sales/total", controller.StatisticOrderSalesTotal)                  // wait to do
		AdminStatisticGroup.GET("order/region/form", controller.StatisticOrderRegionForm)                  // wait to do
		AdminStatisticGroup.GET("order/region/member", controller.StatisticOrderRegionMember)              // wait to do
		AdminStatisticGroup.GET("order/unit/price", controller.StatisticOrderUnitPrice)                    // wait to do
		AdminStatisticGroup.GET("order/return/money", controller.StatisticOrderReturnMoney)                // wait to do
	}

	// 分销相关
	AdminDistributionGroup := router.Group("admin/admin/distribution")
	{
		AdminDistributionGroup.GET("commission-tpl", controller.DistributionCommissionTpl) // wait to do
		AdminDistributionGroup.GET("upgradelog", controller.DistributionUpgradeLog)        // wait to do
		AdminDistributionGroup.GET("member", controller.DistributionMember)                // wait to do
		AdminDistributionGroup.GET("bill/total", controller.DistributionBillTotal)         // wait to do
		AdminDistributionGroup.GET("settings", controller.DistributionSetting)             // wait to do
		AdminDistributionGroup.GET("withdraw/apply", controller.DistributionWithdraw)      // wait to do
	}

	// 店铺相关
	AdminShopGroup := router.Group("admin/admin/shops")
	{
		AdminShopGroup.GET("", controller.ShopList)             // admin/admin/shops
		AdminShopGroup.GET("list", controller.AllShopList)      // admin/admin/shops/list done
		AdminShopGroup.GET("themes", controller.ShopThemesList) // admin/admin/shops/themes
	}

	AdminGroup := router.Group("admin/")
	{
		AdminGroup.GET("systems/admin-users/login", controller.Login)         // done
		AdminGroup.POST("systems/admin-users/logout", controller.Logout)      // done
		AdminGroup.GET("systems/roles/:roleId/checked", controller.RoleCheck) // done
		AdminGroup.GET("admin/systems/messages", controller.MessageList)
		AdminGroup.GET("admin/systems/admin-users/token", controller.Refresh) // done
		AdminGroup.GET("admin/systems/complain-topics", controller.ComplainTopicsList)
		AdminGroup.GET("admin/index/page", controller.Index) //done
		AdminGroup.GET("admin/trade/orders", controller.OrderList)
		AdminGroup.GET("admin/after-sales", controller.AfterSalesList)
		AdminGroup.GET("admin/after-sales/refund", controller.AfterSalesRefundList)
		AdminGroup.GET("admin/trade/orders/pay-log", controller.OrderPayLogList)
		AdminGroup.GET("admin/trade/order-complains", controller.OrderComplainsList)
		AdminGroup.GET("admin/order/bills/statistics", controller.OrderBillStatisticList)
		AdminGroup.GET("admin/pages/site-navigations", controller.PageSiteNavigationList)
		AdminGroup.GET("admin/pages/client_type/page_type", controller.Page) // 替换原来admin/admin/pages/PC/INDEX
		AdminGroup.GET("admin/focus-pictures", controller.FocusPicture)
		AdminGroup.GET("admin/pages/hot-keywords", controller.HotKeyWordsList)

		AdminGroup.GET("admin/settings/site", controller.SiteSetting) // wait to do
		// wait to do
		AdminGroup.GET("admin/trade/orders/setting", controller.TradeOrderSetting)                                       // wait to do
		AdminGroup.GET("admin/settings/point", controller.PoingSetting)                                                  // wait to do
		AdminGroup.GET("admin/task/:task_type", controller.AdminTask)                                                    // wait to do
		AdminGroup.GET("admin/page-create/input", controller.PageCreateInput)                                            // wait to do
		AdminGroup.GET("admin/systems/message-templates", controller.MessageTemplate)                                    // wait to do
		AdminGroup.GET("admin/systems/wechat-msg-tmp/sync", controller.WechatMsgSync)                                    // wait to do
		AdminGroup.GET("admin/systems/wechat-msg-tmp", controller.WechatMsg)                                             // wait to do
		AdminGroup.GET("admin/systems/logi-companies", controller.LogiCompany)                                           // wait to do
		AdminGroup.GET("admin/payment/payment-methods", controller.PaymentMethod)                                        // wait to do
		AdminGroup.GET("admin/goodssearch/custom-words", controller.GoodsSearchCustomWord)                               // wait to do
		AdminGroup.GET("admin/goodssearch/keywords", controller.GoodsSearchKeyWord)                                      // wait to do
		AdminGroup.GET("admin/goodssearch/goods-words", controller.GoodsSearchGoodsWord)                                 // wait to do
		AdminGroup.GET("admin/pages/articles", controller.ArticleList)                                                   // wait to do
		AdminGroup.GET("admin/pages/article-categories", controller.ArticleCategoriesList)                               // wait to do
		AdminGroup.GET("admin/pages/article-categories/childrens", controller.ArticleCategoryChildrenList)               // wait to do
		AdminGroup.GET("admin/services", controller.ServiceList)                                                         // wait to do
		AdminGroup.GET("admin/services/live-video-api/instances", controller.ServiceLiveVideo)                           // wait to do
		AdminGroup.GET("admin/services/live-video-api/instances/:instance_id/logs", controller.ServiceLiveVideoInstance) // wait to do

	}

	return router
}
