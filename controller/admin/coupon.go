package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListCoupon(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyword := ctx.DefaultQuery("keyword", "0")
	endTime := ctx.DefaultQuery("end_time", "0")
	sellerId := ctx.DefaultQuery("seller_id", "0")
	startTime := ctx.DefaultQuery("start_time", "0")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyword
	queryParams["end_time"] = endTime
	queryParams["page_size"] = pageSize
	queryParams["seller_id"] = sellerId
	queryParams["start_time"] = startTime
	data, dataTotal := model.CreateCouponFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateCoupon(ctx *gin.Context) {

}

func UpdateCoupon(ctx *gin.Context) {

}

func DelCoupon(ctx *gin.Context) {

}

func FindOneCoupon(ctx *gin.Context) {

}
