package admin

import (
	"Goshop/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var (
		uuid      = context.Query("uuid")
		captcha   = context.Query("captcha")
		username  = context.Query("username")
		password  = context.Query("password")
		uniqueKey = fmt.Sprintf("%s_%s", "LOGIN", uuid)
	)

	if !Store.Verify(uniqueKey, captcha, true) {
		context.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "图片验证码不正确",
		})
		return
	}
	user := model.CreateUserFactory("").Login(uuid, username, password)
	if user == nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "管理员账号密码错误",
		})
		return
	}
	context.Header("uuid", uuid)
	context.Header("Authorization", user.AccessToken)
	context.JSON(http.StatusOK, user)
}

func Logout(context *gin.Context) {
	uid := context.Query("uid")
	uuid := context.GetHeader("uuid")

	if uid == "" {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "会员id不能为空",
		})
		return
	}
	model.CreateUserFactory("").Logout(uuid, uid)
	context.JSON(http.StatusOK, gin.H{})
}

func Refresh(context *gin.Context) {
	refreshToken := context.PostForm("refresh_token")
	err, m := model.CreateUserFactory("").ExchangeToken(refreshToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, m)
}
