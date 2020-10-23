package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllShopList(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		model.CreateShopFactory("").All())
}

func ShopList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	shopType := ctx.Query("shop_type")
	shopDisable := ctx.Query("shop_disable")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["shop_type"] = shopType
	queryParams["shop_disable"] = shopDisable
	data, dataTotal := model.CreateShopFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})

}

func ShopThemesList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	_type := ctx.Query("type")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["type"] = _type
	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreatGoshopThemeFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
