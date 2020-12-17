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
		// 修改管理员密码及头像
		adminGroup.PUT("admin/systems/admin-users", admin.UpdateAdminUserBase)
		// done 用户名（手机号）/密码登录API
		adminGroup.GET("systems/admin-users/login", admin.Login)
		// done 刷新token
		adminGroup.POST("systems/admin-users/token", admin.Refresh)
		// done 注销管理员登录
		adminGroup.POST("systems/admin-users/logout", admin.Logout)
		// 查询平台管理员列表
		adminGroup.GET("admin/systems/manager/admin-users", admin.ListAdminUser)
		// 添加平台管理员
		adminGroup.POST("admin/systems/manager/admin-users", admin.CreateAdminUser)
		// 修改平台管理员
		adminGroup.PUT("admin/systems/manager/admin-users/:id", admin.UpdateAdminUser)
		// 删除平台管理员
		adminGroup.DELETE("admin/systems/manager/admin-users/:id", admin.DelAdminUser)
		// 查询一个平台管理员
		adminGroup.GET("admin/systems/manager/admin-users/:id", admin.FindOneAdminUser)

		adminGroup.GET("systems/roles/:roleId/checked", admin.RoleCheck) // done
		// done 查询站内消息列表
		adminGroup.GET("admin/systems/messages", admin.MessageList)
		// 添加站内消息
		adminGroup.POST("admin/systems/messages", admin.CreateMessage)
		// done 查询投诉主题列表
		adminGroup.GET("admin/systems/complain-topics", admin.ComplainTopicsList)
		// 添加投诉主题
		adminGroup.POST("admin/systems/complain-topics", admin.CreateComplainTopics)
		// 修改投诉主题
		adminGroup.PUT("admin/systems/complain-topics/:id", admin.UpdateComplainTopics)
		// 删除投诉主题
		adminGroup.DELETE("admin/systems/complain-topics/:id", admin.DelComplainTopics)
		// 查询一个投诉主题
		adminGroup.GET("admin/systems/complain-topics/:id", admin.FindOneComplainTopics)
		// 查询快递平台列表
		adminGroup.GET("admin/systems/express-platforms", admin.ListExpressPlatform)
		// 修改快递平台
		adminGroup.PUT("admin/systems/express-platforms/:bean", admin.UpdateExpressPlatform)
		// 获取快递平台的配置
		adminGroup.GET("admin/systems/express-platforms/:bean", admin.FindOneExpressPlatform)
		// 开启某个快递平台方案
		adminGroup.GET("admin/systems/express-platforms/:bean/open", admin.OpenExpressPlatform)
		// 查询物流详细
		adminGroup.GET("admin/systems/express-platforms/express", admin.FindExpressPlatform)
		// 添加菜单
		adminGroup.POST("admin/systems/menus", admin.CreateSystemMenu)
		// 修改菜单
		adminGroup.PUT("admin/systems/menus/:parent_id", admin.UpdateSystemMenu)
		// 删除菜单
		adminGroup.DELETE("admin/systems/menus/:parent_id", admin.DelSystemMenu)
		// 查询一个菜单
		adminGroup.GET("admin/systems/menus/:parent_id", admin.FindOneSystemMenu)
		// 根据父id查询所有菜单
		adminGroup.GET("admin/systems/menus/:parent_id/children", admin.ListSystemMenuById)
		// done 查询消息模版列表
		adminGroup.GET("admin/systems/message-templates", admin.MessageTemplate)
		// 修改消息模版
		adminGroup.PUT("admin/systems/message-templates/:id", admin.UpdateMessageTemplate)

		adminGroup.GET("admin/systems/wechat-msg-tmp/sync", admin.WechatMsgSync) // done
		adminGroup.GET("admin/systems/wechat-msg-tmp", admin.WechatMsg)          // done
		// done 查询物流公司列表
		adminGroup.GET("admin/systems/logi-companies", admin.ListLogiCompany)
		// 添加物流公司
		adminGroup.POST("admin/systems/logi-companies", admin.CreateLogiCompany)
		// 修改物流公司
		adminGroup.PUT("admin/systems/logi-companies/:id", admin.UpdateLogiCompany)
		// 删除物流公司
		adminGroup.DELETE("admin/systems/logi-companies/:id", admin.DelLogiCompany)
		// 查询一个物流公司
		adminGroup.GET("admin/systems/logi-companies/:id", admin.FindOneLogiCompany)
		// 开启或禁用物流公司
		adminGroup.POST("admin/systems/logi-companies/:id", admin.OpenCloseLogiCompany)
		// 查询付款单列表
		adminGroup.GET("admin/trade/orders/pay-log", admin.OrderPayLogList)
		// 收款单导出Excel
		adminGroup.GET("admin/trade/orders/pay-log/list", admin.ExportOrderPayLogList)
		// 查询交易投诉表列表
		adminGroup.GET("admin/trade/order-complains", admin.OrderComplainsList)
		// 查询一个交易投诉
		adminGroup.GET("admin/trade/order-complains/:id", admin.FindOneOrderComplains)
		// 审核并交由商家申诉
		adminGroup.PUT("admin/trade/order-complains/:id/to-appeal", admin.OrderComplainsAuth)
		// 直接仲裁结束流程
		adminGroup.PUT("admin/trade/order-complains/:id/complete", admin.OrderComplainsComplete)
		// 提交对话
		adminGroup.PUT("admin/trade/order-complains/:id/communication", admin.OrderComplainsCommunication)
		// 查询订单列表
		adminGroup.GET("admin/trade/orders", admin.OrderList)
		// 导出订单列表
		adminGroup.GET("admin/trade/orders/export", admin.ExportOrderList)
		// 查询单个订单明细 origin: admin/admin/trade/orders/:order_id -> admin/admin/r/trade/orders/:order_id
		adminGroup.GET("admin/r/trade/orders/:order_id", admin.OrderDetail)
		// 确认收款
		adminGroup.POST("admin/trade/orders/:order_id/pay", admin.ConfirmOrder)
		// 取消订单
		adminGroup.POST("admin/trade/orders/:order_id/cancelled", admin.CancelOrder)
		// 查询订单日志
		adminGroup.GET("admin/trade/orders/:order_id/log", admin.ListOrderLog)
		// 查询团购活动表列表
		adminGroup.GET("admin/promotion/group-buy-actives", admin.ListGroupBuy)
		// 添加团购活动表
		adminGroup.POST("admin/promotion/group-buy-actives", admin.CreateGroupBuy)
		// 编辑团购活动表
		adminGroup.PUT("admin/promotion/group-buy-actives/:id", admin.UpdateGroupBuy)
		// 删除团购活动表
		adminGroup.DELETE("admin/promotion/group-buy-actives/:id", admin.DelGroupBuy)
		// 查找团购活动表
		adminGroup.GET("admin/promotion/group-buy-actives/:id", admin.FindOneGroupBuy)
		// 批量审核商品
		adminGroup.GET("admin/promotion/group-buy-actives/batch/audit", admin.BatchAuditGoods)
		// 查询团购分类列表
		adminGroup.GET("admin/promotion/group-buy-cats", admin.ListGroupBuyCategory)
		// 添加团购分类
		adminGroup.POST("admin/promotion/group-buy-cats", admin.CreateGroupBuyCategory)
		// 修改团购分类
		adminGroup.PUT("admin/promotion/group-buy-cats/:id", admin.UpdateGroupBuyCategory)
		// 删除团购分类
		adminGroup.DELETE("admin/promotion/group-buy-cats/:id", admin.DelGroupBuyCategory)
		// 查找团购分类
		adminGroup.GET("admin/promotion/group-buy-cats/:id", admin.FindOneGroupBuyCategory)
		// 查询团购商品列表
		adminGroup.GET("admin/promotion/group-buy-goods", admin.ListGroupBuyGoods)
		// 查询团购商品信息
		adminGroup.GET("admin/promotion/group-buy-goods/:gb_id", admin.FindOneGroupBuyGoods)
		// 查询积分商品列表
		adminGroup.GET("admin/promotion/exchange-goods", admin.ListExchangeGoods)
		// 查询某分类下的子分类列表
		adminGroup.GET("admin/promotion/exchange-cats/:cat_id/children", admin.PointCategory)
		// 添加积分兑换分类
		adminGroup.POST("admin/promotion/exchange-cats", admin.CreatePointCategory)
		// 修改积分兑换分类
		adminGroup.PUT("admin/promotion/exchange-cats/:cat_id", admin.UpdatePointCategory)
		// 山吹积分兑换分类
		adminGroup.DELETE("admin/promotion/exchange-cats/:cat_id", admin.DelPointCategory)
		// 获取积分兑换分类
		adminGroup.GET("admin/promotion/exchange-cats/:cat_id", admin.FindOnePointCategory)
		// done 查询拼团列表
		adminGroup.GET("admin/promotion/pintuan", admin.ListPinTuan)
		// 获取活动参与的商品
		adminGroup.GET("admin/promotion/pintuan/goods/:id", admin.ListPinTuanGoods)
		// 查询一个拼团入库
		adminGroup.PUT("admin/promotion/pintuan/:id", admin.FindOnePinTuan)
		// 关闭拼团
		adminGroup.PUT("admin/promotion/pintuan/:id/close", admin.ClosePinTuan)
		// 开启拼团
		adminGroup.PUT("admin/promotion/pintuan/:id/open", admin.OpenPinTuan)
		// done 查询优惠券列表
		adminGroup.GET("admin/promotion/coupons", admin.ListCoupon)
		// done 查询限时抢购列表
		adminGroup.GET("admin/promotion/seckills", admin.ListSeckill)
		// 添加限时抢购入库
		adminGroup.POST("admin/promotion/seckills", admin.CreateSeckill)
		// 批量审核商品
		adminGroup.POST("admin/promotion/seckills/batch/audit", admin.BatchAuditSeckill)
		// 修改限时抢购入库
		adminGroup.PUT("admin/promotion/seckills/:id", admin.UpdateSeckill)
		// 发布限时抢购活动
		adminGroup.POST("admin/promotion/seckills/:id/release", admin.ReleaseSeckill)
		// 删除限时抢购入库
		adminGroup.DELETE("admin/promotion/seckills/:id", admin.DelSeckill)
		// 关闭限时抢购
		adminGroup.DELETE("admin/promotion/seckills/:id/close", admin.CloseSeckill)
		// 查询一个限时抢购入库
		adminGroup.GET("admin/promotion/seckills/:id", admin.FindOneSeckill)
		// 查询限时抢购商品列表
		adminGroup.GET("admin/promotion/seckill-applys", admin.ListSeckillApply)
		// 添加优惠券
		adminGroup.POST("admin/promotion/coupons", admin.CreateCoupon)
		// 编辑优惠券
		adminGroup.PUT("admin/promotion/coupons/:coupon_id", admin.UpdateCoupon)
		// 删除优惠券
		adminGroup.DELETE("admin/promotion/coupons/:coupon_id", admin.DelCoupon)
		// 获取优惠券
		adminGroup.GET("admin/promotion/coupons/:coupon_id", admin.FindOneCoupon)
		// 查询会员开票历史记录信息列表
		adminGroup.GET("admin/members/receipts", admin.ListMemberReceipt)
		// 查询会员开票历史记录详细
		adminGroup.GET("admin/members/receipts/:history_id", admin.FindOneMemberReceipt)
		// 查询会员增票资质信息列表
		adminGroup.GET("admin/members/zpzz", admin.ListZpzz)
		// 查询会员增票资质详细
		adminGroup.GET("admin/members/zpzz/:id", admin.FindOneZpzz)
		// 平台审核会员增票资质申请
		adminGroup.POST("admin/members/zpzz/audit/:id/:status", admin.AuditZpzz)
		//查询会员列表
		adminGroup.GET("admin/members", admin.ListMember)
		// 创建会员
		adminGroup.POST("admin/members", admin.CreateMember)
		// 删除会员
		adminGroup.DELETE("admin/members/:id", admin.DelMember)
		// 查询会员
		adminGroup.GET("admin/members/:id", admin.FindOneMember)
		// 查询多个会员的基本信息
		adminGroup.GET("admin/members/:member_ids/list", admin.FindMoreMember)
		// 修改会员
		adminGroup.PUT("admin/members/:id", admin.UpdateMember)
		// 恢复会员
		adminGroup.POST("admin/members/:id", admin.RecoveryMember)
		// 查询评论列表
		adminGroup.GET("admin/members/comments", admin.ListMemberComments)
		// 批量审核商品评论
		adminGroup.POST("admin/members/comments/batch/audit", admin.BatchAuditMemberComments)
		// 删除评论
		adminGroup.DELETE("admin/members/comments/:comment_id", admin.DelMemberComments)
		// 查询会员商品评论详请
		adminGroup.GET("admin/members/comments/:comment_id", admin.FindOneMemberComments)
		// 查询会员积分列表
		adminGroup.GET("admin/members/point/:member_id", admin.ListMemberPoint)
		// 修改会消费积分
		adminGroup.PUT("admin/members/point/:member_id", admin.UpdateMemberPoint)
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
		// 添加店铺模版
		adminGroup.POST("admin/shops/themes", admin.CreateShopThemes)
		// 修改店铺模版
		adminGroup.PUT("admin/shops/themes/:id", admin.UpdateShopThemes)
		// 删除店铺模版
		adminGroup.DELETE("admin/shops/themes/:id", admin.DelShopThemes)
		// 查找店铺模版
		adminGroup.GET("admin/shops/themes/:id", admin.FindOneShopThemes)
		// done 管理员禁用店铺
		adminGroup.PUT("admin/shops/disable/:shop_id", admin.DisableShop)
		// done 管理员恢复店铺使用
		adminGroup.PUT("admin/shops/enable/:shop_id", admin.EnableShop)
		// done 管理员获取店铺详细 origin: admin/admin/shops/:shop_id -> admin/admin/r/shops/:shop_id
		adminGroup.GET("admin/r/shops/:shop_id", admin.ShopDetail)
		// 管理员修改审核店铺信息 origin: admin/admin/shops/:shop_id -> admin/admin/r/shops/:shop_id
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
		// 查询支付方式列表
		adminGroup.GET("admin/payment/payment-methods", admin.ListPaymentMethod)
		// 修改支付方式
		adminGroup.PUT("admin/payment/payment-methods/:payment_plugin_id", admin.UpdatePaymentMethod)
		// 查询支付方式
		adminGroup.GET("admin/payment/payment-methods/:plugin_id", admin.FindOnePaymentMethod)
		// done 首页响应
		adminGroup.GET("admin/index/page", admin.Index)
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
		// 管理员查询某周期结算列表
		adminGroup.GET("admin/order/bills", admin.ListOrderBill)
		// 初始化账单
		adminGroup.GET("admin/order/bills/init", admin.InitOrderBill)
		// 查看某账单详细
		adminGroup.GET("admin/order/bills/:bill_id", admin.FindOneOrderBill)
		// 导出某账单详细
		adminGroup.GET("admin/order/bills/:bill_id/export", admin.ExportOrderBill)
		// 对账单进行下一步操作
		adminGroup.GET("admin/order/bills/:bill_id/next", admin.NextOrderBill)
		// 查看账单中的订单列表或者退款单列表
		adminGroup.GET("admin/order/bills/:bill_id/:bill_type", admin.QueryBillItems)
		// 管理员查询所有周期结算单列表统计
		adminGroup.GET("admin/order/bills/statistics", admin.ListOrderBillStatistic)
		// done 查询导航栏列表
		adminGroup.GET("admin/pages/site-navigations", admin.ListSiteNavigation)
		// 添加导航栏
		adminGroup.POST("admin/pages/site-navigations", admin.CreateSiteNavigation)
		// 编辑导航栏
		adminGroup.PUT("admin/pages/site-navigations/:id", admin.UpdateSiteNavigation)
		// 删除导航栏
		adminGroup.DELETE("admin/pages/site-navigations/:id", admin.DelSiteNavigation)
		// 获取导航栏
		adminGroup.GET("admin/pages/site-navigations/:id", admin.FindOneSiteNavigation)
		// 上下移动导航栏菜单
		adminGroup.PUT("admin/pages/site-navigations/:id/:sort", admin.UpdateSiteNavigationSort)
		// 查询焦点图列表
		adminGroup.GET("admin/focus-pictures", admin.ListFocusPicture)
		// 添加焦点图
		adminGroup.POST("admin/focus-pictures", admin.CreateFocusPicture)
		// 修改焦点图
		adminGroup.PUT("admin/focus-pictures/:id", admin.UpdateFocusPicture)
		// 删除焦点图
		adminGroup.DELETE("admin/focus-pictures/:id", admin.DelFocusPicture)
		// 获取焦点图
		adminGroup.GET("admin/focus-pictures/:id", admin.FindOneFocusPicture)
		// 查询热门关键字列表
		adminGroup.GET("admin/pages/hot-keywords", admin.HotKeyWordsList)
		// 增加热门关键字
		adminGroup.POST("admin/pages/hot-keywords", admin.CreateHotKeyWords)
		// 获取一个热门关键字
		adminGroup.GET("admin/pages/hot-keywords/:id", admin.FindOneHotKeyWords)
		// 更新一个热门关键字
		adminGroup.PUT("admin/pages/hot-keywords/:id", admin.UpdateHotKeyWords)
		// 删除一个热门关键字
		adminGroup.DELETE("admin/pages/hot-keywords/:Id", admin.DelHotKeyWords)
		// 会员下单量统计-》下单量
		adminGroup.GET("admin/statistics/member/order/quantity", admin.StatisticMemberOrderQuantity)
		// 会员下单量统计-》下单金额
		adminGroup.GET("admin/statistics/member/order/money", admin.StatisticMemberOrderMoney)
		// 下单金额page
		adminGroup.GET("admin/statistics/member/order/money/page", admin.StatisticMemberOrderMoneyPage)
		// 会员下单量统计-》下单商品数
		adminGroup.GET("admin/statistics/member/order/goods/num", admin.StatisticMemberOrderGoodsNum)
		// 会员下单量统计-》下单商品数page
		adminGroup.GET("admin/statistics/member/order/goods/num/page", admin.StatisticMemberOrderGoodsNumPage)
		// 会员下单量统计-》下单量 page
		adminGroup.GET("admin/statistics/member/order/quantity/page", admin.StatisticMemberOrderQuantityPage)
		// 新增会员统计
		adminGroup.GET("admin/statistics/member/increase/member", admin.StatisticMemberIncrease)
		// 新增会员统计 page
		adminGroup.GET("admin/statistics/member/increase/member/page", admin.StatisticMemberIncreasePage)
		// 价格销量统计
		adminGroup.GET("admin/statistics/goods/price/sales", admin.StatisticGoodsPrice)
		// 热卖商品按金额统计
		adminGroup.GET("admin/statistics/goods/hot/money", admin.StatisticGoodsHot)
		// 热卖商品按金额统计
		adminGroup.GET("admin/statistics/goods/hot/money/page", admin.StatisticGoodsHotPage)
		// 热卖商品按数量统计
		adminGroup.GET("admin/statistics/goods/hot/num", admin.StatisticGoodsHotNum)
		// 热卖商品按数量统计
		adminGroup.GET("admin/statistics/goods/hot/num/page", admin.StatisticGoodsHotNumPage)
		// 商品销售明细
		adminGroup.GET("admin/statistics/goods/sale/details", admin.StatisticGoodsSaleDetail)
		// 商品收藏排行
		adminGroup.GET("admin/statistics/goods/collect", admin.StatisticGoodsCollect)
		// 商品收藏排行
		adminGroup.GET("admin/statistics/goods/collect/page", admin.StatisticGoodsCollectPage)
		// 按分类统计下单量
		adminGroup.GET("admin/statistics/industry/order/quantity", admin.StatisticIndustryOrderQuantity)
		// 按分类统计下单商品数量
		adminGroup.GET("admin/statistics/industry/goods/num", admin.StatisticIndustryGoodsNum)
		// 按分类统计下单金额
		adminGroup.GET("admin/statistics/industry/order/money", admin.StatisticIndustryOrderMoney)
		// 概括总览
		adminGroup.GET("admin/statistics/industry/overview", admin.StatisticIndustryOverView)
		// 获取店铺访问量数据
		adminGroup.GET("admin/statistics/page_view/shop", admin.StatisticPageViewShop)
		// 获取商品访问量数据，只取前30
		adminGroup.GET("admin/statistics/page_view/goods", admin.StatisticPageViewGoods)
		// 其他统计=》订单统计=》下单金额
		adminGroup.GET("admin/statistics/order/order/money", admin.StatisticOrderMoney)
		// 其他统计=》订单统计=》下单数量
		adminGroup.GET("admin/statistics/order/order/num", admin.StatisticOrderNum)
		// 其他统计=》订单统计=》下单数量
		adminGroup.GET("admin/statistics/order/order/page", admin.StatisticOrderPage)
		// 其他统计=》销售收入统计 page
		adminGroup.GET("admin/statistics/order/sales/money", admin.StatisticOrderSalesMoney)
		// 其他统计=》销售收入 退款统计 page
		adminGroup.GET("admin/statistics/order/aftersales/money", admin.StatisticOrderAfterSalesMoney)
		// 其他统计=》销售收入总览
		adminGroup.GET("admin/statistics/order/sales/total", admin.StatisticOrderSalesTotal)
		// 区域分析=>下单量
		adminGroup.GET("admin/statistics/order/region/num", admin.StatisticOrderRegionNum)
		// 区域分析=>下单金额
		adminGroup.GET("admin/statistics/order/region/money", admin.StatisticOrderRegionMoney)
		// 区域分析表格=>page
		adminGroup.GET("admin/statistics/order/region/form", admin.StatisticOrderRegionForm)
		// 区域分析=>下单会员数
		adminGroup.GET("admin/statistics/order/region/member", admin.StatisticOrderRegionMember)
		// 客单价分布=>客单价分布
		adminGroup.GET("admin/statistics/order/unit/price", admin.StatisticOrderUnitPrice)
		// 客单价分布=>购买频次分析
		adminGroup.GET("admin/statistics/order/unit/num", admin.StatisticOrderUnitNum)
		// 客单价分布=>购买时段分析
		adminGroup.GET("admin/statistics/order/unit/time", admin.StatisticOrderUnitTime)
		// 退款统计
		adminGroup.GET("admin/statistics/order/return/money", admin.StatisticOrderReturnMoney)
		// 检测是否有任务正在进行,有任务返回任务id,无任务返回404
		adminGroup.GET("admin/task/:task_type", admin.AdminTask)
		// 查看任务进度
		adminGroup.GET("admin/task/:task_type/progress", admin.AdminTaskProgress)
		// 清除某任务
		adminGroup.DELETE("admin/task/:task_type", admin.DelAdminTask)
		// 获取当前静态页面设置参数
		adminGroup.GET("admin/page-create/input", admin.PageCreateInput)
		// 页面生成
		adminGroup.POST("admin/page-create/create", admin.CreatePageCreate)
		// 参数保存
		adminGroup.POST("admin/page-create/save", admin.SavePageCreate)
		// 查询文章列表
		adminGroup.GET("admin/pages/articles", admin.ListArticle)
		// 创建文章
		adminGroup.POST("admin/pages/articles", admin.CreateArticle)
		// 获取文章
		adminGroup.GET("admin/pages/articles/:article_id", admin.FindOneArticle)
		// 更新文章
		adminGroup.PUT("admin/pages/articles/:article_id", admin.UpdateArticle)
		// 删除文章
		adminGroup.DELETE("admin/pages/articles/:article_id", admin.DelArticle)
		// 查询文章一级分类列表
		adminGroup.GET("admin/pages/article-categories", admin.ArticleCategoriesList)
		// 添加文章分类
		adminGroup.POST("admin/pages/article-categories", admin.CreateArticleCategories)
		// 修改文章分类
		adminGroup.PUT("admin/pages/article-categories/:id", admin.UpdateArticleCategories)
		// 删除文章分类
		adminGroup.DELETE("admin/pages/article-categories/:id", admin.DelArticleCategories)
		// 获取文章分类
		adminGroup.GET("admin/pages/article-categories/:id", admin.FindOneArticleCategories)
		// 查询文章二级分类列表
		adminGroup.GET("admin/pages/article-categories/:id/children", admin.ArticleCategoriesList)
		// 查询文章分类树
		adminGroup.GET("admin/pages/article-categories/childrens", admin.ArticleCategoryChildrenList) // wait to do
		// 修改楼层
		adminGroup.PUT("admin/pages/:page_id", admin.UpdatePage)
		// 查询楼层
		adminGroup.GET("admin/pages/:page_id", admin.FindOnePage)
		// 使用客户端类型和页面类型修改楼层 替换原来admin/admin/:client_type/:page_type
		adminGroup.GET("admin/pages/client_type/page_type", admin.FindOneClientPage)
		// 使用客户端类型和页面类型查询一个楼层
		adminGroup.PUT("admin/pages/client_type/page_type", admin.UpdateClientPage)
		// 商品推送
		adminGroup.GET("admin/systems/push/:goods_id", admin.PushGoods)
		// 获取推送设置
		adminGroup.GET("admin/systems/push", admin.AppPushSetting)
		// 修改推送设置
		adminGroup.PUT("admin/systems/push", admin.SaveAppPushSetting)

		adminGroup.GET("admin/services", admin.ServiceList)                                                         // wait to do
		adminGroup.GET("admin/services/live-video-api/instances", admin.ServiceLiveVideo)                           // wait to do
		adminGroup.GET("admin/services/live-video-api/instances/:instance_id/logs", admin.ServiceLiveVideoInstance) // wait to do

		// 未发现
		adminGroup.GET("admin/members/deposit/recharge", admin.MemberDepositRechargeList)
		adminGroup.GET("admin/members/deposit/log", admin.MemberDepositLogList)
	}
}
