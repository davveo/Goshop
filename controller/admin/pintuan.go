package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListPinTuan(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	name := ctx.DefaultQuery("name", "")
	status := ctx.DefaultQuery("status", "") // WAIT:待开始，UNDERWAY：进行中，END：已结束
	endTime := ctx.DefaultQuery("end_time", "")
	sellerId := ctx.DefaultQuery("seller_id", "")
	startTime := ctx.DefaultQuery("start_time", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["name"] = name
	queryParams["status"] = status
	queryParams["page_no"] = pageNo
	queryParams["end_time"] = endTime
	queryParams["seller_id"] = sellerId
	queryParams["page_size"] = pageSize
	queryParams["start_time"] = startTime
	data, dataTotal := model.CreatePinTuanFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func ListPinTuanGoods(ctx *gin.Context) {

}

func FindOnePinTuan(ctx *gin.Context) {

}

func ClosePinTuan(ctx *gin.Context) {

}

func OpenPinTuan(ctx *gin.Context) {

}
