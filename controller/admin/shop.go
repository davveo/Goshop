package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllShopList(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		model.CreateShopFactory(ctx, "").All())
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
	data, dataTotal := model.CreateShopFactory(ctx, "").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})

}

func ShopThemesList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	_type := ctx.DefaultQuery("type", "")
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

func DisableShop(ctx *gin.Context) {
	shopId, _ := strconv.Atoi(ctx.Param("shop_id"))
	if err := model.CreateShopFactory(ctx, "").DisableShop(shopId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"seller_id": shopId,
		"operator":  "管理员禁用店铺",
	})
}

func EnableShop(ctx *gin.Context) {
	shopId, _ := strconv.Atoi(ctx.Param("shop_id"))
	if err := model.CreateShopFactory(ctx, "").EnableShop(shopId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"seller_id": shopId,
		"operator":  "管理员恢复店铺",
	})
}

func ShopDetail(ctx *gin.Context) {
	shopId, _ := strconv.Atoi(ctx.Param("shop_id"))
	shop, err := model.CreateShopFactory(ctx, "").GetShop(shopId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, shop)
}

func EditShop(ctx *gin.Context) {

}

func CreateShop(ctx *gin.Context) {

}
