package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 允许跨域
func Cors() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func NoCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		ctx.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		ctx.Next()
	}
}
