package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"Goshop/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SiteSetting(ctx *gin.Context) {
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

func SaveSiteSetting(ctx *gin.Context) {
	body := common.ParseFromBody(ctx)
	config, err := model.CreateSettingFactory("").Save(consts.SITE, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
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

func SaveGoodsSetting(ctx *gin.Context) {
	// 获取商品审核设置信息
	body := common.ParseFromBody(ctx)
	config, err := model.CreateSettingFactory("").Save(consts.GOODS, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func TradeOrderSetting(ctx *gin.Context) {
	config := model.CreateSettingFactory("").Get(consts.TRADE)
	if config == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func SaveTradeOrderSetting(ctx *gin.Context) {
	body := common.ParseFromBody(ctx)
	config, err := model.CreateSettingFactory("").Save(consts.TRADE, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func PointSetting(ctx *gin.Context) {
	config := model.CreateSettingFactory("").Get(consts.POINT)
	if config == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func SavePointSetting(ctx *gin.Context) {
	body := common.ParseFromBody(ctx)
	config, err := model.CreateSettingFactory("").Save(consts.POINT, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func DistributionSetting(ctx *gin.Context) {
	config := model.CreateSettingFactory("").Get(consts.DISTRIBUTION)
	if config == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func SaveDistributionSetting(ctx *gin.Context) {
	body := common.ParseFromBody(ctx)
	config, err := model.CreateSettingFactory("").Save(consts.DISTRIBUTION, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}
