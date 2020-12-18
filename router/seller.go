package router

import (
	seller "Goshop/controller/seller"

	"github.com/gin-gonic/gin"
)

func SellerApi(router *gin.RouterGroup) {
	sellerGroup := router.Group("seller/")
	{
		sellerGroup.GET("seller/members/reply", seller.Reply) // 查询会员商品咨询回复列表

		// 咨询相关API
		sellerGroup.GET("seller/members/asks", seller.Ask)                 // 查询咨询列表
		sellerGroup.GET("seller/members/s/:ask_id", seller.AskDetail)      // 查询会员商品咨询详请  members/{ask_id}
		sellerGroup.GET("seller/members/d/:ask_id/reply", seller.AskReply) // 商家回复会员商品咨询 members/{ask_id}/reply

		//商家登录API
		sellerGroup.GET("seller/login", seller.Login)                             // 用户名（手机号）/密码登录API
		sellerGroup.POST("seller/login/smscode/:mobile", seller.SendLoginSmsCode) //发送验证码
		sellerGroup.POST("seller/login/{mobile}", seller.MobileLogin)             // 手机号码登录API

		sellerGroup.POST("seller/register/smscode/:mobile", seller.SendRegisterSmsCode) //

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
	}
}
