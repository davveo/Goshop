package router

import (
	seller "Goshop/controller/seller"

	"github.com/gin-gonic/gin"
)

func SellerApi(router *gin.RouterGroup) {
	sellerGroup := router.Group("seller/seller/")
	{
		sellerGroup.GET("members/reply", seller.Reply) // 查询会员商品咨询回复列表

		// 咨询相关API
		sellerGroup.GET("members/asks", seller.Ask)               // 查询咨询列表
		sellerGroup.GET("members/:ask_id/reply", seller.AskReply) // 商家回复会员商品咨询
		sellerGroup.GET("members/:ask_id", seller.AskDetail)      // 查询会员商品咨询详请

		//商家登录API
		sellerGroup.GET("login", seller.Login)                             // 用户名（手机号）/密码登录API
		sellerGroup.POST("login/smscode/:mobile", seller.SendLoginSmsCode) //发送验证码
		sellerGroup.POST("login/{mobile}", seller.MobileLogin)             // 手机号码登录API

		sellerGroup.POST("register/smscode/:mobile", seller.SendRegisterSmsCode) //
	}
}
