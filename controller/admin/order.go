package admin

import (
	"Goshop/model"
	"Goshop/utils/common"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderList(ctx *gin.Context) {
	queryParams := common.ParseFromQuery(ctx)
	data, dataTotal := model.CreateOrderFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    queryParams["page_no"],
		"page_size":  queryParams["page_size"],
	})
}

func ExportOrderList(ctx *gin.Context) {

}

func ConfirmOrder(ctx *gin.Context) {

}

func CancelOrder(ctx *gin.Context) {

}

func ListOrderLog(ctx *gin.Context) {

}

func OrderDetail(ctx *gin.Context) {
	orderId := ctx.Param("order_id")
	fmt.Println(orderId)
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

func FindOneOrderComplains(ctx *gin.Context) {

}

func OrderComplainsAuth(ctx *gin.Context) {

}

func OrderComplainsComplete(ctx *gin.Context) {

}

func OrderComplainsCommunication(ctx *gin.Context) {

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

func ExportOrderPayLogList(ctx *gin.Context) {

}

func ListOrderBillStatistic(ctx *gin.Context) {
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

func InitOrderBill(ctx *gin.Context) {

}

func ListOrderBill(ctx *gin.Context) {

}

func FindOneOrderBill(ctx *gin.Context) {

}

func ExportOrderBill(ctx *gin.Context) {

}

func NextOrderBill(ctx *gin.Context) {

}

func QueryBillItems(ctx *gin.Context) {

}
