package middleware

import (
	"log"
	"net/http"
	"Goshop/utils/enum"
	ojwt "Goshop/utils/jwt"

	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//  模拟验证token
		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			log.Println(enum.ErrorsValidatorBindParamsFail, err.Error())
			context.Abort()
		}

		j := ojwt.NewJwt()
		claims, err := j.ParseToken(headerParams.Authorization)
		if err != nil {
			if err == ojwt.TokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": enum.ErrorsNoAuthorization,
				})
				context.Abort()
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": enum.ErrorsNoAuthorization,
			})
			//暂停执行
			context.Abort()
		}

		context.Set("claims", claims)
		context.Next()
	}
}
