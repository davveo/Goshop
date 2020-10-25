package router

import (
	"Goshop/controller/admin"
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
		ApiGroup.POST("base/uploaders", admin.Upload)                  // wait to do
	}
	AdminApi(ApiGroup)
	BuyerApi(ApiGroup)
	SellerApi(ApiGroup)
	return router
}
