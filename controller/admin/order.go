package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWord := ctx.Query("keyword")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyWord
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateOrderFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func OrderComplainsList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWord := ctx.Query("keyword")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyWord
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateOrderComplainsFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func OrderPayLogList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	payStatus := ctx.DefaultQuery("pay_status", "")
	paymentType := ctx.DefaultQuery("payment_type", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["pay_status"] = payStatus
	queryParams["payment_type"] = paymentType
	data, dataTotal := model.CreateTradePayLogFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func OrderBillStatisticList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateOrderBillStatisticFactory("").GetAllBill(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
