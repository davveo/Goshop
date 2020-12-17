package admin

import (
	"Goshop/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var (
		uuid      = ctx.Query("uuid")
		captcha   = ctx.Query("captcha")
		username  = ctx.Query("username")
		password  = ctx.Query("password")
		uniqueKey = fmt.Sprintf("%s_%s", "LOGIN", uuid)
	)

	if !Store.Verify(uniqueKey, captcha, true) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "图片验证码不正确",
		})
		return
	}
	user := model.CreateUserFactory("").Login(uuid, username, password)
	if user == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "管理员账号密码错误",
		})
		return
	}
	ctx.Header("uuid", uuid)
	ctx.Header("Authorization", user.AccessToken)
	ctx.JSON(http.StatusOK, user)
}

func Logout(ctx *gin.Context) {
	uid := ctx.Query("uid")
	uuid := ctx.GetHeader("uuid")

	if uid == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "会员id不能为空",
		})
		return
	}
	model.CreateUserFactory("").Logout(uuid, uid)
	ctx.JSON(http.StatusOK, gin.H{})
}

func Refresh(ctx *gin.Context) {
	refreshToken := ctx.PostForm("refresh_token")
	err, m := model.CreateUserFactory("").ExchangeToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, m)
}

func UpdateAdminUserBase(ctx *gin.Context) {

}

func ListAdminUser(ctx *gin.Context) {

}

func CreateAdminUser(ctx *gin.Context) {

}

func UpdateAdminUser(ctx *gin.Context) {

}

func DelAdminUser(ctx *gin.Context) {

}

func FindOneAdminUser(ctx *gin.Context) {

}
