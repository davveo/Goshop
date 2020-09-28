package controller

import (
	"net/http"
	"orange/global/consts"
	"orange/model"

	"github.com/gin-gonic/gin"
)

func SiteShow(ctx *gin.Context) {
	config := model.CreateSettingFactory("").Get(consts.SITE)
	if config == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func FocusPicture(ctx *gin.Context) {
	clientType := ctx.Query("client_type") //APP/WAP/PC

	data, err := model.CreateFocusPictureFactory("").List(clientType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func Page(ctx *gin.Context) {
	pageType := ctx.Param("page_type")     // APP/WAP/PC
	clientType := ctx.Param("client_type") // INDEX/SPECIAL
	data, err := model.CreatePageFactory("").GetByType(clientType, pageType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
