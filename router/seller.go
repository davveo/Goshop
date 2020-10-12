package router

import (
	seller "Goshop/controller/seller"

	"github.com/gin-gonic/gin"
)

func SellerApi(router *gin.RouterGroup) {
	sellerGroup := router.Group("seller/seller/")
	{
		sellerGroup.GET("members/reply", seller.MemberReply) // 查询会员商品咨询回复列表
	}
}
