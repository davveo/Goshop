package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListGroupBuy(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateGroupBuyFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateGroupBuy(ctx *gin.Context) {

}

func UpdateGroupBuy(ctx *gin.Context) {

}

func DelGroupBuy(ctx *gin.Context) {

}

func FindOneGroupBuy(ctx *gin.Context) {

}

func BatchAuditGoods(ctx *gin.Context) {

}

func ListGroupBuyCategory(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateGroupBuyCategoryFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateGroupBuyCategory(ctx *gin.Context) {

}

func UpdateGroupBuyCategory(ctx *gin.Context) {

}

func DelGroupBuyCategory(ctx *gin.Context) {

}

func FindOneGroupBuyCategory(ctx *gin.Context) {

}
