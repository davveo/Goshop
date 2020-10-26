package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SiteSetting(ctx *gin.Context) {

}

func GoodsSetting(ctx *gin.Context) {
	// 获取商品审核设置信息
	config := model.CreateSettingFactory("").Get(consts.GOODS)
	if config == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func TradeOrderSetting(ctx *gin.Context) {

}

func PoingSetting(ctx *gin.Context) {

}
