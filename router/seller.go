package router

import (
	seller "Goshop/controller/seller"

	"github.com/gin-gonic/gin"
)

func SellerApi(router *gin.RouterGroup) {
	sellerGroup := router.Group("seller/")
	{
		// 查询会员商品咨询回复列表
		sellerGroup.GET("seller/members/reply", seller.Reply)
		// 查询咨询列表
		sellerGroup.GET("seller/members/asks", seller.Ask)
		// 查询会员商品咨询详请
		sellerGroup.GET("seller/members/s/:ask_id", seller.AskDetail)
		// 商家回复会员商品咨询
		sellerGroup.GET("seller/members/d/:ask_id/reply", seller.AskReply)
		// 查询评论列表
		sellerGroup.GET("seller/members/comments", seller.ListComments)
		// 回复评论
		sellerGroup.POST("seller/members/comments/:comment_id/reply", seller.CommentReply)
		// 查询商品评论详请
		sellerGroup.GET("seller/members/comments/:comment_id", seller.FindOneComment)
		// 注销会员登录
		sellerGroup.GET("seller/members/logout", seller.MemberLogout)
		// 用户名（手机号）/密码登录API
		sellerGroup.GET("seller/login", seller.Login)
		//发送验证码
		sellerGroup.POST("seller/login/smscode/:mobile", seller.SendLoginSmsCode)
		// 手机号码登录API
		sellerGroup.POST("seller/login/{mobile}", seller.MobileLogin)
		// 发送验证码
		sellerGroup.POST("seller/register/smscode/:mobile", seller.SendRegisterSmsCode)
		// PC注册
		sellerGroup.POST("seller/register/pc", seller.RegisterForPC)
		// 验证手机验证码
		sellerGroup.GET("seller/check/smscode/{mobile}", seller.CheckSmsCode)
		// 用户名重复校验
		sellerGroup.GET("seller/check/username/{username}", seller.CheckUserName)
		// 手机号重复校验
		sellerGroup.GET("seller/check/mobile/{mobile}", seller.CheckMobile)
		// 刷新token
		sellerGroup.POST("seller/check/token", seller.CheckToken)
		// 获取申请售后服务记录列表
		sellerGroup.GET("seller/after-sales", seller.ListAfterSales)
		// 获取售后服务详细信息
		sellerGroup.GET("seller/after-sales/detail/:service_sn", seller.FindOneAfterSales)
		// 商家审核售后服务申请
		sellerGroup.POST("seller/after-sales/audit/:service_sn", seller.AuditAfterSales)
		// 商家将申请售后服务退还的商品入库
		sellerGroup.POST("seller/after-sales/put-in/warehouse", seller.PutInWareHouseAfterSales)
		// 导出售后服务信息
		sellerGroup.GET("seller/after-sales/export", seller.ExportAfterSales)
		// 商家为售后服务手动创建新订单
		sellerGroup.POST("seller/after-sales/create-order/:service_sn", seller.CreateAfterSales)
		// 商家关闭售后服务单
		sellerGroup.POST("seller/after-sales/close/:service_sn", seller.CloseAfterSales)
		// 获取售后退款单列表
		sellerGroup.GET("seller/after-sales/refund", seller.ListAfterSalesRefund)
		// 在线支付订单商家退款
		sellerGroup.POST("seller/after-sales/refund/:service_sn", seller.UpdateAfterSalesRefund)
		// 货到付款订单商家退款
		sellerGroup.POST("seller/after-sales/refund/cod/:service_sn", seller.CodAfterSalesRefund)
		// 分销商品返利获取
		sellerGroup.GET("seller/distribution", seller.ListDistribution)
		// 分销商品返利设置
		sellerGroup.GET("seller/distribution/goods/:goods_id", seller.FindOneDistributionGoods)
		// 获取分销设置:1开启/0关闭
		sellerGroup.GET("seller/distribution/setting", seller.DistributionSetting)
		// 商品发布，获取当前登录用户选择经营类目的所有父
		sellerGroup.GET("seller/goods/category/:category_id/children", seller.ListTargetGoodsCategory)
		// 商家分类
		sellerGroup.GET("seller/goods/category/seller/children", seller.ListGoodsCategoryChildren)
		// 商品发布，获取所选分类关联的参数信息
		sellerGroup.GET("seller/goods/category/:category_id/:goods_id/params", seller.ListTargetGoodsCategoryParams)
		// 发布商品，获取所选分类关联的参数信息
		sellerGroup.GET("seller/goods/category/:category_id/params", seller.ListGoodsCategoryParams)
		// 修改商品，获取所选分类关联的品牌信息
		sellerGroup.GET("seller/goods/category/:category_id/brands", seller.ListGoodsCategoryBrands)
		// 查询草稿商品列表
		sellerGroup.GET("seller/goods/draft-goods", seller.ListGoodsDraft)
		// 添加商品
		sellerGroup.POST("seller/goods/draft-goods", seller.CreateGoodsDraft)
		// 修改草稿商品
		sellerGroup.PUT("seller/goods/draft-goods/:draft_goods_id", seller.UpdateGoodsDraft)
		// 删除草稿商品
		sellerGroup.DELETE("seller/goods/draft-goods/:draft_goods_id", seller.DelGoodsDraft)
		// 查询一个草稿商品,商家编辑草稿商品使用
		sellerGroup.GET("seller/goods/draft-goods/:draft_goods_id", seller.FindOneGoodsDraft)
		// 查询草稿商品关联的参数，包括没有添加的参数
		sellerGroup.GET("seller/goods/draft-goods/:draft_goods_id/params", seller.FindOneGoodsDraftParams)
		// 查询草稿箱商品sku信息
		sellerGroup.GET("seller/goods/draft-goods/:draft_goods_id/skus", seller.FindOneGoodsDraftSku)
		// 草稿箱商品上架接口
		sellerGroup.PUT("seller/goods/draft-goods/:draft_goods_id/market", seller.UpdateGoodsDraftMarket)
		// 商家单独维护库存接口
		sellerGroup.PUT("seller/goods/:goods_id/quantity", seller.UpdateGoodsQuantity)
		// 查询商品列表
		sellerGroup.GET("seller/goods", seller.ListGoods)
		// 查询预警商品列表
		sellerGroup.GET("seller/goods/warning", seller.ListGoodsWarning)
		// 添加商品
		sellerGroup.POST("seller/goods", seller.CreateGoods)
		// 修改商品
		//sellerGroup.PUT("seller/goods/:id", seller.UpdateGoods)
		// 查询一个商品,商家编辑时使用
		//sellerGroup.GET("seller/goods/:id", seller.FindOneGoods)
		// 商家下架商品
		//sellerGroup.PUT("seller/goods/:goods_ids/under", seller.UpdateGoodsUnder)
		// 商家将商品放入回收站
		//sellerGroup.PUT("seller/goods/:goods_ids/recycle", seller.UpdateGoodsRecycle)
		// 商家还原商品
		//sellerGroup.PUT("seller/goods/:goods_ids/revert", seller.UpdateGoodsRevert)
		// 商家彻底删除商品
		sellerGroup.DELETE("seller/goods/:goods_ids", seller.DelGoods)
		// 查询多个商品的基本信息
		sellerGroup.GET("seller/goods/:goods_ids/details", seller.FindMoreGoods)
		// 商品sku信息信息获取api
		//sellerGroup.GET("seller/goods/:goods_id/skus", seller.ListTargetGoodsSku)
		// 查询多个商品的基本信息
		sellerGroup.GET("seller/goods/skus/:sku_ids/details", seller.FindOneGoodsSku)
		// 查询SKU列表
		sellerGroup.GET("seller/goods/skus", seller.ListGoodsSku)
		// 根据分类id查询规格包括规格值
		sellerGroup.GET("seller/goods/categories/{category_id}/specs", seller.ListGoodsCategorySpecs)
		// 商家自定义某分类的规格项
		sellerGroup.POST("seller/goods/categories/{category_id}/specs", seller.UpdateGoodsCategorySpecs)
		// 商家自定义某规格的规格值
		sellerGroup.POST("seller/goods/specs/{spec_id}/values", seller.UpdateGoodsCategorySpecsValue)
		// 查询商品标签列表
		sellerGroup.GET("seller/goods/tags", seller.ListGoodsTag)
		// 查询某标签下的商品
		sellerGroup.GET("seller/goods/tags/:tag_id/goods", seller.ListTargetGoodsTag)
		// 保存某标签下的商品
		//sellerGroup.PUT("seller/goods/:tag_id/goods/:goods_ids", seller.UpdateGoodsTag)
		// 商家查看我的账单列表
		sellerGroup.GET("seller/order/bills", seller.ListOrderBill)
		// 商家查看某账单详细
		sellerGroup.GET("seller/order/bills/:bill_id", seller.FindOneOrderBill)
		// 卖家对账单进行下一步操作
		sellerGroup.PUT("seller/order/bills/:bill_id/next", seller.OrderBillNext)
		// 查看账单中的订单列表或者退款单列表
		sellerGroup.GET("seller/order/bills/:bill_id/:bill_type", seller.FindOrderBillByType)
		// 导出某账单详细
		sellerGroup.GET("seller/order/bills/:bill_id/export", seller.ExportOrderBill)
		// 查询会员开票历史记录信息列表
		sellerGroup.GET("seller/members/receipts", seller.ListMemberReceipts)
		// 查询会员开票历史记录详细
		sellerGroup.GET("seller/members/receipts/:history_id", seller.FindOneMemberReceipts)
		// 商家开具发票-增值税普通发票和增值税专用发票
		sellerGroup.POST("seller/members/receipts/:history_id/logi", seller.CreateMemberReceipts)
		// 商家开具发票-上传电子普通发票附件
		sellerGroup.POST("seller/members/receipts/upload/files", seller.UploadMemberReceipts)
		// 查询优惠券列表
		sellerGroup.GET("seller/promotion/coupons", seller.ListCoupons)
		// 添加优惠券
		sellerGroup.POST("seller/promotion/coupons", seller.CreateCoupons)
		// 修改优惠券
		sellerGroup.PUT("seller/promotion/coupons/:id", seller.UpdateCoupons)
		// 删除优惠券
		sellerGroup.DELETE("seller/promotion/coupons/:id", seller.DelCoupons)
		// 查询一个优惠券
		sellerGroup.GET("seller/promotion/coupons/:id", seller.FindOneCoupons)
		// 根据状态获取优惠券数据集合
		//sellerGroup.GET("seller/promotion/coupons/:status/list", seller.FindCouponsByStatus)
		// 查询某分类下的子分类列表
		sellerGroup.GET("seller/promotion/exchange-cats/:parent_id/children", seller.ListExchangeCat)
		// 查询满优惠赠品列表
		sellerGroup.GET("seller/promotion/full-discount-gifts", seller.ListFullDiscountGifts)
		// 添加满优惠赠品
		sellerGroup.POST("seller/promotion/full-discount-gifts", seller.CreateFullDiscountGifts)
		// 修改满优惠赠品
		sellerGroup.PUT("seller/promotion/full-discount-gifts/:id", seller.UpdateFullDiscountGifts)
		// 删除满优惠赠品
		sellerGroup.DELETE("seller/promotion/full-discount-gifts/:id", seller.DelFullDiscountGifts)
		// 查询一个满优惠赠品
		sellerGroup.GET("seller/promotion/full-discount-gifts/:id", seller.FindOneFullDiscountGifts)
		// 查询满优惠赠品集合
		sellerGroup.GET("seller/promotion/full-discount-gifts/all", seller.FindAllFullDiscountGifts)
		// 查询满优惠活动列表
		sellerGroup.GET("seller/promotion/full-discounts", seller.ListFullDiscount)
		// 添加满优惠活动
		sellerGroup.POST("seller/promotion/full-discounts", seller.CreateFullDiscount)
		// 修改满优惠活动
		sellerGroup.PUT("seller/promotion/full-discounts/:id", seller.UpdateFullDiscount)
		// 删除满优惠活动
		sellerGroup.DELETE("seller/promotion/full-discounts/:id", seller.DelFullDiscount)
		// 查询一个满优惠活动
		sellerGroup.GET("seller/promotion/full-discounts/:id", seller.FindOneFullDiscount)
		// 查询团购分类列表
		sellerGroup.GET("seller/promotion/group-buy-cats", seller.ListGroupBuyCat)
		// 查询团购商品列表
		sellerGroup.GET("seller/promotion/group-buy-goods", seller.ListGroupBuyGoods)
		// 添加团购商品
		sellerGroup.POST("seller/promotion/group-buy-goods", seller.CreateGroupBuyGoods)
		// 修改团购商品
		sellerGroup.PUT("seller/promotion/group-buy-goods/:id", seller.UpdateGroupBuyGoods)
		// 删除团购商品
		sellerGroup.DELETE("seller/promotion/group-buy-goods/:id", seller.DelGroupBuyGoods)
		// 查询一个团购商品
		sellerGroup.GET("seller/promotion/group-buy-goods/:id", seller.FindOneGroupBuyGoods)
		// 查询可以参与的团购活动列表
		sellerGroup.GET("seller/promotion/group-buy-goods/active", seller.FindActiveGroupBuyGoods)
		// 查询第二件半价列表
		sellerGroup.GET("seller/promotion/half-prices", seller.ListHalfPrices)
		// 添加第二件半价
		sellerGroup.POST("seller/promotion/half-prices", seller.CreateHalfPrices)
		// 修改第二件半价
		sellerGroup.PUT("seller/promotion/half-prices/:id", seller.UpdateHalfPrices)
		// 删除第二件半价
		sellerGroup.DELETE("seller/promotion/half-prices/:id", seller.DelHalfPrices)
		// 查询一个第二件半价
		sellerGroup.GET("seller/promotion/half-prices/:id", seller.FindOneHalfPrices)
		// 查询单品立减列表
		sellerGroup.GET("seller/promotion/minus", seller.ListMinus)
		// 添加单品立减
		sellerGroup.POST("seller/promotion/minus", seller.CreateMinus)
		// 修改单品立减
		sellerGroup.PUT("seller/promotion/minus/:id", seller.UpdateMinus)
		// 删除单品立减
		sellerGroup.DELETE("seller/promotion/minus/:id", seller.DelMinus)
		// 查询一个单品立减
		sellerGroup.GET("seller/promotion/minus/:id", seller.FindOneMinus)
		// 查询活动表列
		sellerGroup.GET("seller/promotion/pintuan", seller.ListPinTuan)
		// 添加活动
		sellerGroup.POST("seller/promotion/pintuan", seller.CreatePinTuan)
		// 修改活动
		sellerGroup.PUT("seller/promotion/pintuan/:id", seller.UpdatePinTuan)
		// 删除活动
		sellerGroup.DELETE("seller/promotion/pintuan/:id", seller.DelPinTuan)
		// 查询一个活动
		sellerGroup.GET("seller/promotion/pintuan/:id", seller.FindOnePinTuan)
		// 修改活动参与的商品
		sellerGroup.POST("seller/promotion/pintuan/goods/:id", seller.UpdatePinTuanGoods)
		// 获取活动参与的商品
		sellerGroup.GET("seller/promotion/pintuan/goods/:id", seller.ListPinTuanGoods)
		// 查询限时抢购申请商品列表
		sellerGroup.GET("seller/promotion/seckill-applys", seller.ListSeckillApply)
		// 添加限时抢购申请
		sellerGroup.POST("seller/promotion/seckill-applys", seller.CreateSeckillApply)
		// 删除限时抢购申请
		sellerGroup.DELETE("seller/promotion/seckill-applys/:id", seller.DelSeckillApply)
		// 查询一个限时抢购申请
		sellerGroup.GET("seller/promotion/seckill-applys/:id", seller.FindOneSeckillApply)
		// 查询所有的限时抢购活动
		sellerGroup.GET("seller/promotion/seckill-applys/seckill", seller.ListAllSeckill)
		// 查询一个限时抢购活动
		sellerGroup.POST("seller/promotion/seckill-applys/:seckill_id/seckill", seller.FindOneTargetSeckillApply)
	}
}
