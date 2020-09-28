package controller

import (
	"net/http"
	"Goshop/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CouponList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	sellerId := ctx.Query("seller_id")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["seller_id"] = sellerId
	data, dataTotal := model.CreateCouponFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
