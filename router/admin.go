package router

import (
	"Goshop/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminApi(router *gin.RouterGroup) {
	adminGroup := router.Group("admin/")
	{
		// done 查询规格项列表
		adminGroup.GET("admin/goods/specs", admin.SpecsList)
		// done 添加规格项
		adminGroup.POST("admin/goods/specs", admin.CreateSpecs)
		// done 修改规格项
		adminGroup.PUT("admin/goods/specs/:spec_id", admin.UpdateSpecs)
		// done 删除规格项
		adminGroup.DELETE("admin/goods/specs/:spec_id", admin.DeleteSpecs)
		// done 查询一个规格项
		adminGroup.GET("admin/goods/specs/:spec_id", admin.Specs)
		// done 查询规格值列表
		adminGroup.GET("admin/goods/specs/:spec_id/values", admin.SpecsValues)
		// done 添加某规格的规格值
		adminGroup.POST("admin/goods/specs/:spec_id/values", admin.UpdateSpecsValues)
		// done 商品列表
		adminGroup.GET("admin/goods", admin.GoodsList)
		// done origin: admin/admin/goods/:goods_id/up -> admin/admin/r/goods/:goods_id/up
		adminGroup.PUT("admin/r/goods/:goods_id/up", admin.GoodsUp)
		// done origin: admin/admin/goods/:goods_id/under -> admin/admin/r/goods/:goods_id/under
		adminGroup.PUT("admin/r/goods/:goods_id/under", admin.GoodsUnder)
		// done 查询多个商品的基本信息
		adminGroup.GET("admin/goods/:goods_id/detail", admin.GoodsListDetail)
		// done 管理员批量审核商品
		adminGroup.POST("admin/goods/batch/audit", admin.GoodsBatchAudit)

		// done 查询品牌列表
		adminGroup.GET("admin/goods/brands", admin.BrandList)
		// done 添加品牌
		adminGroup.POST("admin/goods/brands", admin.CreateBrand)
		// done 查询一个品牌
		adminGroup.GET("admin/goods/brands/:brand_id", admin.Brand)
		// done 修改品牌
		adminGroup.PUT("admin/goods/brands/:brand_id", admin.UpdateBrand)
		// done 删除品牌
		adminGroup.DELETE("admin/goods/brands/:brand_id", admin.DeleteBrand)
		// done 查询所有品牌 // origin: admin/admin/goods/brands/all -> admin/admin/r/goods/brands/all
		adminGroup.GET("admin/r/goods/brands/all", admin.BrandAllList)

		// done 查询某分类下的子分类列表
		adminGroup.GET("admin/goods/categories/:parent_id/children", admin.CategoryList)
		// 查询某分类下的全部子分类列表
		adminGroup.GET("admin/goods/categories/:parent_id/all-children", admin.CategoryAllList)
		// done 添加商品分类
		adminGroup.POST("admin/goods/categories", admin.CreateCategory)
		// 修改商品分类
		adminGroup.PUT("admin/goods/categories/:id", admin.EditCategory)
		// 删除商品分类
		adminGroup.DELETE("admin/goods/categories/:id", admin.DelCategory)
		// 查询商品分类
		adminGroup.GET("admin/goods/categories/:id", admin.Category)
		// 查询分类品牌
		adminGroup.GET("admin/goods/categories/:category_id/brands", admin.CategoryBrand)
		// 管理员操作分类绑定品牌
		adminGroup.PUT("admin/goods/categories/:category_id/brands", admin.SaveCategoryBrand)
		// 查询分类规格
		adminGroup.GET("admin/goods/categories/:category_id/specs", admin.CategorySpecs)
		// 管理员操作分类绑定规格
		adminGroup.PUT("admin/goods/categories/:category_id/specs", admin.SaveCategorySpecs)
		// 查询分类参数
		adminGroup.GET("admin/goods/categories/:category_id/param", admin.CategoryParam)
		// 商品索引初始化
		adminGroup.GET("admin/goods/search", admin.GoodsSearchCreate)

		// done 查询自定义分词列表
		adminGroup.GET("admin/goodssearch/custom-words", admin.GoodsSearchCustomWord)
		// 添加自定义分词
		adminGroup.POST("admin/goodssearch/custom-words", admin.CreateGoodsSearchCustomWord)
		// 修改自定义分词
		adminGroup.PUT("admin/goodssearch/custom-words/:id", admin.EditGoodsSearchCustomWord)
		// 删除自定义分词
		adminGroup.DELETE("admin/goodssearch/custom-words/:id", admin.DelGoodsSearchCustomWord)
		// 查询一个自定义分词
		adminGroup.GET("admin/goodssearch/custom-words/:id", admin.FindOneGoodsSearchCustomWord)
		// 设置ES分词库秘钥
		adminGroup.PUT("admin/goodssearch/custom-words/secret-key", admin.CreateEsCustomWordSecretKey)
		// 获取ES分词库秘钥
		adminGroup.GET("admin/goodssearch/custom-words/secret-key", admin.FindEsCustomWordSecretKey)
		// 查询商品优先级列表
		adminGroup.GET("admin/goodssearch/priority", admin.ListGoodsSearchPriority)
		// 修改商品优先级
		adminGroup.PUT("admin/goodssearch/priority", admin.UpdateGoodsSearchPriority)

		// done 查询关键字历史列表
		adminGroup.GET("admin/goodssearch/keywords", admin.ListGoodsSearchKeyWord)
		// done 查询提示词列表
		adminGroup.GET("admin/goodssearch/goods-words", admin.GoodsSearchGoodsWord)
		// 添加自定义提示词
		adminGroup.POST("admin/goodssearch/goods-words", admin.CreateGoodsSearchGoodsWord)
		// 删除提示词
		adminGroup.POST("admin/goodssearch/goods-words/:id", admin.DelGoodsSearchGoodsWord)
		// 修改自定义提示词
		adminGroup.PUT("admin/goodssearch/goods-words/:id/words", admin.EditGoodsSearchGoodsWord)
		// 修改提示词排序
		adminGroup.PUT("admin/goodssearch/goods-words/:id/sort", admin.SortGoodsSearchGoodsWord)

		// 查询会员商品咨询回复列表
		adminGroup.GET("admin/members/reply", admin.ListMemberReply)
		// 批量审核会员商品咨询回复
		adminGroup.POST("admin/members/reply/batch/audit", admin.BatchAuditMemberReply)
		// 删除会员商品咨询回复
		adminGroup.DELETE("admin/members/reply/:id", admin.DelMemberReply)

		// 查询指定会员的地址列表
		adminGroup.GET("admin/members/addresses/:member_id", admin.ListMemberAddress)

		// 系统相关
		adminGroup.GET("systems/admin-users/login", admin.Login)                  // done
		adminGroup.POST("systems/admin-users/logout", admin.Logout)               // done
		adminGroup.GET("systems/roles/:roleId/checked", admin.RoleCheck)          // done
		adminGroup.GET("admin/systems/messages", admin.MessageList)               // done
		adminGroup.GET("admin/systems/admin-users/token", admin.Refresh)          // done
		adminGroup.GET("admin/systems/complain-topics", admin.ComplainTopicsList) // done
		adminGroup.GET("admin/systems/message-templates", admin.MessageTemplate)  // done
		adminGroup.GET("admin/systems/wechat-msg-tmp/sync", admin.WechatMsgSync)  // done
		adminGroup.GET("admin/systems/wechat-msg-tmp", admin.WechatMsg)           // done
		adminGroup.GET("admin/systems/logi-companies", admin.LogiCompany)         // done 物流公司相关API

		// 交易相关
		adminGroup.GET("admin/trade/orders", admin.OrderList)
		adminGroup.GET("admin/trade/orders/pay-log", admin.OrderPayLogList)
		adminGroup.GET("admin/trade/order-complains", admin.OrderComplainsList)
		// origin: admin/admin/trade/orders/:order_id -> admin/admin/r/trade/orders/:order_id
		adminGroup.GET("admin/r/trade/orders/:order_id", admin.OrderDetail)

		// promotion相关
		adminGroup.GET("admin/promotion/group-buy-actives", admin.GroupBuyList)
		adminGroup.GET("admin/promotion/group-buy-cats", admin.GroupBuyCategoryList)
		adminGroup.GET("admin/promotion/exchange-cats/:cat_id/children", admin.PointCategory)
		adminGroup.GET("admin/promotion/seckills", admin.SeckillList) // done
		adminGroup.GET("admin/promotion/coupons", admin.CouponList)   // done
		adminGroup.GET("admin/promotion/pintuan", admin.PinTuanList)  // done

		// 会员相关
		adminGroup.GET("admin/members/receipts", admin.MemberReceiptList)
		adminGroup.GET("admin/members/zpzz", admin.ZpzzList)
		adminGroup.GET("admin/members", admin.MemberList)

		// 查询评论列表
		adminGroup.GET("admin/members/comments", admin.ListMemberComments)
		// 批量审核商品评论
		adminGroup.POST("admin/members/comments/batch/audit", admin.BatchAuditMemberComments)
		// 删除评论
		adminGroup.DELETE("admin/members/comments/:comment_id", admin.DelMemberComments)
		// 查询会员商品评论详请
		adminGroup.GET("admin/members/comments/:comment_id", admin.FindOneMemberComments)

		// 查询咨询列表
		adminGroup.GET("admin/members/asks", admin.ListMemberAsk)
		// 批量审核会员商品咨询
		adminGroup.POST("admin/members/asks/batch/audit", admin.BatchAuditMemberAsk)
		// 删除咨询
		adminGroup.DELETE("admin/members/asks/:ask_id", admin.DelMemberAsk)
		// 查询会员商品咨询详请
		adminGroup.GET("admin/members/asks/:ask_id", admin.FindOneMemberAsk)

		// done 分页查询列表
		adminGroup.GET("admin/shops/list", admin.AllShopList)
		// done 分页查询店铺列表
		adminGroup.GET("admin/shops", admin.ShopList)
		// done 获取商店主题
		adminGroup.GET("admin/shops/themes", admin.ShopThemesList)
		// done 管理员禁用店铺
		adminGroup.PUT("admin/shops/disable/:shop_id", admin.DisableShop)
		// done 管理员恢复店铺使用
		adminGroup.PUT("admin/shops/enable/:shop_id", admin.EnableShop)
		// done 管理员获取店铺详细
		// origin: admin/admin/shops/:shop_id -> admin/admin/r/shops/:shop_id
		adminGroup.GET("admin/r/shops/:shop_id", admin.ShopDetail)
		//管理员修改审核店铺信息
		// origin: admin/admin/shops/:shop_id -> admin/admin/r/shops/:shop_id
		adminGroup.PUT("admin/r/shops/:shop_id", admin.EditShop)
		//后台添加店铺
		adminGroup.POST("admin/shops", admin.CreateShop)

		// done 获取站点设置
		adminGroup.GET("admin/settings/site", admin.SiteSetting)
		// done 修改站点设置
		adminGroup.POST("admin/settings/site", admin.SaveSiteSetting)
		// done 交易订单设置
		adminGroup.GET("admin/trade/orders/setting", admin.TradeOrderSetting)
		// done 修改交易订单设置
		adminGroup.POST("admin/trade/orders/setting", admin.SaveTradeOrderSetting)
		// done 获取积分设置
		adminGroup.GET("admin/settings/point", admin.PointSetting)
		// done 修改积分设置
		adminGroup.POST("admin/settings/point", admin.SavePointSetting)
		// done 获取商品设置
		adminGroup.GET("admin/goods/settings", admin.GoodsSetting)
		// done 修改商品设置
		adminGroup.POST("admin/goods/settings", admin.SaveGoodsSetting)
		// 添加参数组
		adminGroup.POST("admin/goods/parameter-groups", admin.CreateParameterGroups)
		// 修改参数组
		adminGroup.PUT("admin/goods/parameter-groups/:id", admin.EditParameterGroups)
		// 删除参数组
		adminGroup.DELETE("admin/goods/parameter-groups/:id", admin.DelParameterGroups)
		// 查询参数组
		adminGroup.GET("admin/goods/parameter-groups/:id", admin.FindParameterGroups)
		// 参数组上移或者下移
		adminGroup.PUT("admin/goods/parameter-groups/:group_id/sort", admin.SortParameterGroups)
		// 添加参数
		adminGroup.POST("admin/goods/parameters", admin.CreateParameters)
		// 修改参数
		adminGroup.PUT("admin/goods/parameters/:id", admin.EditParameters)
		// 删除参数
		adminGroup.DELETE("admin/goods/parameters/:id", admin.DelParameters)
		// 查询参数
		adminGroup.GET("admin/goods/parameters/:id", admin.FindParameters)
		// 参数上移或者下移
		adminGroup.PUT("admin/goods/parameters/:param_id/sort", admin.SortParameters)

		// done 分销商分页
		adminGroup.GET("admin/distribution/bill/member", admin.BillMemberList)
		// done 获取某个业绩详情  admin/distribution/bill/member/:id->admin/r/distribution/bill/member/:id
		adminGroup.GET("admin/r/distribution/bill/o/member/:id", admin.BillMemberDetail)
		// 获取某个分销商下级业绩
		adminGroup.GET("admin/distribution/bill/member/down", admin.DownBillMember)
		// done 导出会员结算单
		adminGroup.GET("admin/distribution/bill/member/export", admin.ExportBillMember)

		// done 模板列表
		adminGroup.GET("admin/distribution/commission-tpl", admin.DistributionCommissionTplList)
		// done 获取模版
		adminGroup.GET("admin/distribution/commission-tpl/:tplId", admin.DistributionCommissionTplDetail)
		// 编辑模版
		adminGroup.PUT("admin/distribution/commission-tpl/:tplId", admin.DistributionCommissionTplEdit)
		// done 删除模版
		adminGroup.DELETE("admin/distribution/commission-tpl/:tplId", admin.DistributionCommissionTplDel)
		// 添加模版
		adminGroup.POST("admin/distribution/commission-tpl", admin.DistributionCommissionTpl)
		// 获取信任登录配置参数
		adminGroup.GET("admin/members/connect", admin.ListConnectSetting)
		// 修改信任登录参数
		adminGroup.PUT("admin/members/connect/:type", admin.EditConnectSetting)

		// 查询某优惠券领取列表
		adminGroup.GET("admin/members/coupon", admin.ListMemberCoupon)
		// 废弃某优惠券
		adminGroup.PUT("admin/members/coupon/:member_coupon_id/cancel", admin.CancelMemberCoupon)

		// 分销商列表
		adminGroup.GET("admin/distribution/member", admin.DistributionMember)
		// 修改分销商模版
		adminGroup.PUT("admin/distribution/tpl", admin.DistributionMemberChangeTpl)

		// 结算单 分销订单查询
		adminGroup.GET("admin/distribution/order", admin.DistributionBillOrder)
		// 结算单 分销退款订单查询
		adminGroup.GET("admin/distribution/order/sellback", admin.DistributionBillSellbackOrder)

		// done 结算单分页
		adminGroup.GET("admin/distribution/bill/total", admin.DistributionBillTotalList)

		// done 分销设置
		adminGroup.GET("admin/distribution/settings", admin.DistributionSetting)

		// done 修改分销设置
		adminGroup.PUT("admin/distribution/settings", admin.SaveDistributionSetting)

		// 订单金额统计
		adminGroup.GET("admin/distribution/statistic/order", admin.DistributionStatisticOrder)
		// 订单数量统计
		adminGroup.GET("admin/distribution/statistic/count", admin.DistributionStatisticCount)
		// 订单返现统计
		adminGroup.GET("admin/distribution/statistic/push", admin.DistributionStatisticPush)
		// 店铺返现统计
		adminGroup.GET("admin/distribution/statistic/push/seller", admin.DistributionStatisticPushSeller)

		// 获取升级日志
		adminGroup.GET("admin/distribution/upgradelog", admin.DistributionUpgradeLog)

		// done 提现申请审核列表
		adminGroup.GET("admin/distribution/withdraw/apply", admin.DistributionWithdraw)
		// 导出提现申请
		adminGroup.GET("admin/distribution/withdraw/export", admin.DistributionWithdrawExport)
		// 批量审核提现申请
		adminGroup.POST("admin/distribution/withdraw/batch/auditing", admin.DistributionWithdrawBatchAuditing)
		// 批量设为已转账
		adminGroup.POST("admin/distribution/withdraw/batch/account/paid", admin.DistributionWithdrawBatchAccountPaid)

		adminGroup.GET("admin/payment/payment-methods", admin.PaymentMethod) // wait to do
		adminGroup.GET("admin/index/page", admin.Index)                      // done
		// done 获取申请售后服务记录列表
		adminGroup.GET("admin/after-sales", admin.AfterSalesList)
		// done 获取售后服务详细信息
		adminGroup.GET("admin/after-sales/detail/:service_sn", admin.AfterSalesDetail)
		// done 导出售后服务信息
		adminGroup.GET("admin/after-sales/export", admin.AfterSalesExport)
		// done 获取售后退款单列表
		adminGroup.GET("admin/after-sales/refund", admin.AfterSalesRefundList)
		// done 平台退款操作
		adminGroup.GET("admin/after-sales/refund/:service_sn", admin.AfterSalesRefund)

		adminGroup.GET("admin/order/bills/statistics", admin.OrderBillStatisticList)

		adminGroup.GET("admin/pages/site-navigations", admin.PageSiteNavigationList)
		adminGroup.GET("admin/pages/client_type/page_type", admin.Page) // 替换原来admin/admin/pages/PC/INDEX
		adminGroup.GET("admin/focus-pictures", admin.FocusPicture)
		adminGroup.GET("admin/pages/hot-keywords", admin.HotKeyWordsList)

		// 数据统计相关
		adminGroup.GET("admin/statistics/member/order/quantity", admin.StatisticMemberOrderQuantity)          // wait to do
		adminGroup.GET("admin/statistics/member/order/quantity/page", admin.StatisticMemberOrderQuantityPage) // wait to do
		adminGroup.GET("admin/statistics/member/increase/member", admin.StatisticMemberIncrease)              // wait to do
		adminGroup.GET("admin/statistics/member/increase/member/page", admin.StatisticMemberIncreasePage)     // wait to do
		adminGroup.GET("admin/statistics/goods/price/sales", admin.StatisticGoodsPrice)                       // wait to do
		adminGroup.GET("admin/statistics/goods/hot/money", admin.StatisticGoodsHot)                           // wait to do
		adminGroup.GET("admin/statistics/goods/hot/money/page", admin.StatisticGoodsHotPage)                  // wait to do
		adminGroup.GET("admin/statistics/goods/sale/details", admin.StatisticGoodsSaleDetail)                 // wait to do
		adminGroup.GET("admin/statistics/goods/collect", admin.StatisticGoodsCollect)                         // wait to do
		adminGroup.GET("admin/statistics/goods/collect/page", admin.StatisticGoodsCollectPage)                // wait to do
		adminGroup.GET("admin/statistics/industry/order/quantity", admin.StatisticIndustry)                   // wait to do
		adminGroup.GET("admin/statistics/industry/overview", admin.StatisticIndustryOverView)                 // wait to do
		adminGroup.GET("admin/statistics/page_view/shop", admin.StatisticPageViewShop)                        // wait to do
		adminGroup.GET("admin/statistics/page_view/goods", admin.StatisticPageViewGoods)                      // wait to do
		adminGroup.GET("admin/statistics/order/order/page", admin.StatisticPageOrder)                         // wait to do
		adminGroup.GET("admin/statistics/order/order/money", admin.StatisticPageMoney)                        // wait to do
		adminGroup.GET("admin/statistics/order/sales/money", admin.StatisticOrderSalesMoney)                  // wait to do
		adminGroup.GET("admin/statistics/order/sales/total", admin.StatisticOrderSalesTotal)                  // wait to do
		adminGroup.GET("admin/statistics/order/region/form", admin.StatisticOrderRegionForm)                  // wait to do
		adminGroup.GET("admin/statistics/order/region/member", admin.StatisticOrderRegionMember)              // wait to do
		adminGroup.GET("admin/statistics/order/unit/price", admin.StatisticOrderUnitPrice)                    // wait to do
		adminGroup.GET("admin/statistics/order/return/money", admin.StatisticOrderReturnMoney)                // wait to do

		adminGroup.GET("admin/task/:task_type", admin.AdminTask)         // wait to do
		adminGroup.GET("admin/page-create/input", admin.PageCreateInput) // wait to do

		adminGroup.GET("admin/pages/articles", admin.ArticleList)                  // todo
		adminGroup.POST("admin/pages/articles", admin.CreateArticle)               // todo
		adminGroup.GET("admin/pages/articles/:article_id", admin.GetArticle)       // todo
		adminGroup.PUT("admin/pages/articles/:article_id", admin.UpdateArticle)    // todo
		adminGroup.DELETE("admin/pages/articles/:article_id", admin.DeleteArticle) // todo

		adminGroup.GET("admin/pages/article-categories", admin.ArticleCategoriesList)                 // wait to do
		adminGroup.GET("admin/pages/article-categories/childrens", admin.ArticleCategoryChildrenList) // wait to do

		adminGroup.GET("admin/services", admin.ServiceList)                                                         // wait to do
		adminGroup.GET("admin/services/live-video-api/instances", admin.ServiceLiveVideo)                           // wait to do
		adminGroup.GET("admin/services/live-video-api/instances/:instance_id/logs", admin.ServiceLiveVideoInstance) // wait to do

		// 未发现
		adminGroup.GET("admin/members/deposit/recharge", admin.MemberDepositRechargeList)
		adminGroup.GET("admin/members/deposit/log", admin.MemberDepositLogList)
	}
}
