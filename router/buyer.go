package router

import "github.com/gin-gonic/gin"

func BuyerApi(router *gin.RouterGroup) {
	buyerGroup := router.Group("buyer/buyer/")
	{
		buyerGroup.GET("")
	}
}
