package router

import (
	"Goshop/controller/admin"
	"Goshop/controller/seller"
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
		ApiGroup.GET("health", admin.Health)                           // done
		ApiGroup.GET("site-show", admin.SiteShow)                      // done
		ApiGroup.GET("captcha/:uuid/:scene", admin.NewCaptcha)         // done
		ApiGroup.GET("base/regions/:region_id/children", admin.Region) // wait to do

		// 买家相关
		BuyApi := ApiGroup.Group("buyer/buyer/")
		{
			BuyApi.GET("")
		}

		// 卖家相关
		SellerApi := ApiGroup.Group("seller/seller/")
		{
			SellerApi.GET("members/reply", seller.MemberReply) // 查询会员商品咨询回复列表
		}
	}

	// 路由管理
	//AdminGroup := router.Group("admin/").Use(middleware.JWTAuth())
	AdminGroup := router.Group("admin/")
	{
		AdminGroup.GET("systems/admin-users/login", admin.Login)         // done
		AdminGroup.POST("systems/admin-users/logout", admin.Logout)      // done
		AdminGroup.GET("systems/roles/:roleId/checked", admin.RoleCheck) // done
		AdminGroup.GET("admin/systems/messages", admin.MessageList)
		AdminGroup.GET("admin/systems/admin-users/token", admin.Refresh) // done
		AdminGroup.GET("admin/systems/complain-topics", admin.ComplainTopicsList)
		AdminGroup.GET("admin/index/page", admin.Index)                                  // done
		AdminGroup.GET("admin/goods", admin.GoodsList)                                   // done
		AdminGroup.GET("admin/goods/specs", admin.SpecsList)                             // done
		AdminGroup.GET("admin/goods/brands", admin.BrandList)                            // done
		AdminGroup.GET("admin/goods/brands/all", admin.BrandAllList)                     // done
		AdminGroup.GET("admin/goods/categories/:parent_id/children", admin.CategoryList) // done
		AdminGroup.GET("admin/trade/orders", admin.OrderList)
		AdminGroup.GET("admin/after-sales", admin.AfterSalesList)
		AdminGroup.GET("admin/after-sales/refund", admin.AfterSalesRefundList)
		AdminGroup.GET("admin/trade/orders/pay-log", admin.OrderPayLogList)
		AdminGroup.GET("admin/trade/order-complains", admin.OrderComplainsList)
		AdminGroup.GET("admin/members/receipts", admin.MemberReceiptList)
		AdminGroup.GET("admin/members/zpzz", admin.ZpzzList)
		AdminGroup.GET("admin/members", admin.MemberList)
		AdminGroup.GET("admin/members/comments", admin.MemberCommentsList)
		AdminGroup.GET("admin/members/asks", admin.MemberAskList)

		AdminGroup.GET("admin/shops/list", admin.AllShopList) // done
		AdminGroup.GET("admin/shops", admin.ShopList)
		AdminGroup.GET("admin/shops/themes", admin.ShopThemesList)

		AdminGroup.GET("admin/order/bills/statistics", admin.OrderBillStatisticList)
		AdminGroup.GET("admin/promotion/group-buy-actives", admin.GroupBuyList)
		AdminGroup.GET("admin/promotion/group-buy-cats", admin.GroupBuyCategoryList)
		AdminGroup.GET("admin/promotion/exchange-cats/:cat_id/children", admin.PointCategory)
		AdminGroup.GET("admin/promotion/seckills", admin.SeckillList)
		AdminGroup.GET("admin/promotion/coupons", admin.CouponList)
		AdminGroup.GET("admin/promotion/pintuan", admin.PinTuanList)
		AdminGroup.GET("admin/pages/site-navigations", admin.PageSiteNavigationList)
		AdminGroup.GET("admin/pages/client_type/page_type", admin.Page) // 替换原来admin/admin/pages/PC/INDEX
		AdminGroup.GET("admin/focus-pictures", admin.FocusPicture)
		AdminGroup.GET("admin/pages/hot-keywords", admin.HotKeyWordsList)
		AdminGroup.GET("admin/statistics/member/order/quantity", admin.StatisticMemberOrderQuantity)                // wait to do
		AdminGroup.GET("admin/statistics/member/order/quantity/page", admin.StatisticMemberOrderQuantityPage)       // wait to do
		AdminGroup.GET("admin/statistics/member/increase/member", admin.StatisticMemberIncrease)                    // wait to do
		AdminGroup.GET("admin/statistics/member/increase/member/page", admin.StatisticMemberIncreasePage)           // wait to do
		AdminGroup.GET("admin/statistics/goods/price/sales", admin.StatisticGoodsPrice)                             // wait to do
		AdminGroup.GET("admin/statistics/goods/hot/money", admin.StatisticGoodsHot)                                 // wait to do
		AdminGroup.GET("admin/statistics/goods/hot/money/page", admin.StatisticGoodsHotPage)                        // wait to do
		AdminGroup.GET("admin/statistics/goods/sale/details", admin.StatisticGoodsSaleDetail)                       // wait to do
		AdminGroup.GET("admin/statistics/goods/collect", admin.StatisticGoodsCollect)                               // wait to do
		AdminGroup.GET("admin/statistics/goods/collect/page", admin.StatisticGoodsCollectPage)                      // wait to do
		AdminGroup.GET("admin/statistics/industry/order/quantity", admin.StatisticIndustry)                         // wait to do
		AdminGroup.GET("admin/statistics/industry/overview", admin.StatisticIndustryOverView)                       // wait to do
		AdminGroup.GET("admin/statistics/page_view/shop", admin.StatisticPageViewShop)                              // wait to do
		AdminGroup.GET("admin/statistics/page_view/goods", admin.StatisticPageViewGoods)                            // wait to do
		AdminGroup.GET("admin/statistics/order/order/page", admin.StatisticPageOrder)                               // wait to do
		AdminGroup.GET("admin/statistics/order/order/money", admin.StatisticPageMoney)                              // wait to do
		AdminGroup.GET("admin/statistics/order/sales/money", admin.StatisticOrderSalesMoney)                        // wait to do
		AdminGroup.GET("admin/statistics/order/sales/total", admin.StatisticOrderSalesTotal)                        // wait to do
		AdminGroup.GET("admin/statistics/order/region/form", admin.StatisticOrderRegionForm)                        // wait to do
		AdminGroup.GET("admin/statistics/order/region/member", admin.StatisticOrderRegionMember)                    // wait to do
		AdminGroup.GET("admin/statistics/order/unit/price", admin.StatisticOrderUnitPrice)                          // wait to do
		AdminGroup.GET("admin/statistics/order/return/money", admin.StatisticOrderReturnMoney)                      // wait to do
		AdminGroup.GET("admin/settings/site", admin.SiteSetting)                                                    // wait to do
		AdminGroup.GET("admin/goods/settings", admin.GoodsSetting)                                                  // wait to do
		AdminGroup.GET("admin/trade/orders/setting", admin.TradeOrderSetting)                                       // wait to do
		AdminGroup.GET("admin/settings/point", admin.PoingSetting)                                                  // wait to do
		AdminGroup.GET("admin/task/:task_type", admin.AdminTask)                                                    // wait to do
		AdminGroup.GET("admin/page-create/input", admin.PageCreateInput)                                            // wait to do
		AdminGroup.GET("admin/systems/message-templates", admin.MessageTemplate)                                    // wait to do
		AdminGroup.GET("admin/systems/wechat-msg-tmp/sync", admin.WechatMsgSync)                                    // wait to do
		AdminGroup.GET("admin/systems/wechat-msg-tmp", admin.WechatMsg)                                             // wait to do
		AdminGroup.GET("admin/systems/logi-companies", admin.LogiCompany)                                           // wait to do
		AdminGroup.GET("admin/payment/payment-methods", admin.PaymentMethod)                                        // wait to do
		AdminGroup.GET("admin/goodssearch/custom-words", admin.GoodsSearchCustomWord)                               // wait to do
		AdminGroup.GET("admin/goodssearch/keywords", admin.GoodsSearchKeyWord)                                      // wait to do
		AdminGroup.GET("admin/goodssearch/goods-words", admin.GoodsSearchGoodsWord)                                 // wait to do
		AdminGroup.GET("admin/pages/articles", admin.ArticleList)                                                   // wait to do
		AdminGroup.GET("admin/pages/article-categories", admin.ArticleCategoriesList)                               // wait to do
		AdminGroup.GET("admin/pages/article-categories/childrens", admin.ArticleCategoryChildrenList)               // wait to do
		AdminGroup.GET("admin/services", admin.ServiceList)                                                         // wait to do
		AdminGroup.GET("admin/services/live-video-api/instances", admin.ServiceLiveVideo)                           // wait to do
		AdminGroup.GET("admin/services/live-video-api/instances/:instance_id/logs", admin.ServiceLiveVideoInstance) // wait to do
		AdminGroup.GET("admin/distribution/commission-tpl", admin.DistributionCommissionTpl)                        // wait to do
		AdminGroup.GET("admin/distribution/upgradelog", admin.DistributionUpgradeLog)                               // wait to do
		AdminGroup.GET("admin/distribution/member", admin.DistributionMember)                                       // wait to do
		AdminGroup.GET("admin/distribution/bill/total", admin.DistributionBillTotal)                                // wait to do
		AdminGroup.GET("admin/distribution/settings", admin.DistributionSetting)                                    // wait to do
		AdminGroup.GET("admin/distribution/withdraw/apply", admin.DistributionWithdraw)                             // wait to do

		// 未发现
		AdminGroup.GET("admin/members/deposit/recharge", admin.MemberDepositRechargeList)
		AdminGroup.GET("admin/members/deposit/log", admin.MemberDepositLogList)
	}

	return router
}
