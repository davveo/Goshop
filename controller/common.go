package controller

import (
	"net/http"
	"orange/global/consts"
	"orange/model"

	"github.com/gin-gonic/gin"
)

func SiteShow(context *gin.Context) {
	config := model.CreateSettingFactory("").Get(consts.SITE)
	if config == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	context.JSON(http.StatusOK, config)
}
