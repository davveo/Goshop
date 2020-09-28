package router

import (
	"Eshop/controller"
	"Eshop/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 中间件管理
	router.Use(middleware.Cors())
	//router.Use(middleware.JwtMiddleWare())

	ApiGroup := router.Group("")
	{
		ApiGroup.GET("health", controller.Health)                           // done
		ApiGroup.GET("site-show", controller.SiteShow)                      // done
		ApiGroup.GET("captcha/:uuid/:scene", controller.NewCaptcha)         // done
		ApiGroup.GET("base/regions/:region_id/children", controller.Region) // wait to do
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
		AdminGroup.GET("admin/pages/site-navigations", controller.PageSiteNavigationList)
		AdminGroup.GET("admin/pages/client_type/page_type", controller.Page) // 替换原来admin/admin/pages/PC/INDEX
		AdminGroup.GET("admin/focus-pictures", controller.FocusPicture)
		AdminGroup.GET("admin/pages/hot-keywords", controller.HotKeyWordsList)
		AdminGroup.GET("admin/statistics/member/order/quantity", controller.StatisticMemberOrderQuantity)                // wait to do
		AdminGroup.GET("admin/statistics/member/order/quantity/page", controller.StatisticMemberOrderQuantityPage)       // wait to do
		AdminGroup.GET("admin/statistics/member/increase/member", controller.StatisticMemberIncrease)                    // wait to do
		AdminGroup.GET("admin/statistics/member/increase/member/page", controller.StatisticMemberIncreasePage)           // wait to do
		AdminGroup.GET("admin/statistics/goods/price/sales", controller.StatisticGoodsPrice)                             // wait to do
		AdminGroup.GET("admin/statistics/goods/hot/money", controller.StatisticGoodsHot)                                 // wait to do
		AdminGroup.GET("admin/statistics/goods/hot/money/page", controller.StatisticGoodsHotPage)                        // wait to do
		AdminGroup.GET("admin/statistics/goods/sale/details", controller.StatisticGoodsSaleDetail)                       // wait to do
		AdminGroup.GET("admin/statistics/goods/collect", controller.StatisticGoodsCollect)                               // wait to do
		AdminGroup.GET("admin/statistics/goods/collect/page", controller.StatisticGoodsCollectPage)                      // wait to do
		AdminGroup.GET("admin/statistics/industry/order/quantity", controller.StatisticIndustry)                         // wait to do
		AdminGroup.GET("admin/statistics/industry/overview", controller.StatisticIndustryOverView)                       // wait to do
		AdminGroup.GET("admin/statistics/page_view/shop", controller.StatisticPageViewShop)                              // wait to do
		AdminGroup.GET("admin/statistics/page_view/goods", controller.StatisticPageViewGoods)                            // wait to do
		AdminGroup.GET("admin/statistics/order/order/page", controller.StatisticPageOrder)                               // wait to do
		AdminGroup.GET("admin/statistics/order/order/money", controller.StatisticPageMoney)                              // wait to do
		AdminGroup.GET("admin/statistics/order/sales/money", controller.StatisticOrderSalesMoney)                        // wait to do
		AdminGroup.GET("admin/statistics/order/sales/total", controller.StatisticOrderSalesTotal)                        // wait to do
		AdminGroup.GET("admin/statistics/order/region/form", controller.StatisticOrderRegionForm)                        // wait to do
		AdminGroup.GET("admin/statistics/order/region/member", controller.StatisticOrderRegionMember)                    // wait to do
		AdminGroup.GET("admin/statistics/order/unit/price", controller.StatisticOrderUnitPrice)                          // wait to do
		AdminGroup.GET("admin/statistics/order/return/money", controller.StatisticOrderReturnMoney)                      // wait to do
		AdminGroup.GET("admin/settings/site", controller.SiteSetting)                                                    // wait to do
		AdminGroup.GET("admin/goods/settings", controller.GoodsSetting)                                                  // wait to do
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
		AdminGroup.GET("admin/distribution/commission-tpl", controller.DistributionCommissionTpl)                        // wait to do
		AdminGroup.GET("admin/distribution/upgradelog", controller.DistributionUpgradeLog)                               // wait to do
		AdminGroup.GET("admin/distribution/member", controller.DistributionMember)                                       // wait to do
		AdminGroup.GET("admin/distribution/bill/total", controller.DistributionBillTotal)                                // wait to do
		AdminGroup.GET("admin/distribution/settings", controller.DistributionSetting)                                    // wait to do
		AdminGroup.GET("admin/distribution/withdraw/apply", controller.DistributionWithdraw)                             // wait to do

		// 未发现
		AdminGroup.GET("admin/members/deposit/recharge", controller.MemberDepositRechargeList)
		AdminGroup.GET("admin/members/deposit/log", controller.MemberDepositLogList)
	}

	return router
}
