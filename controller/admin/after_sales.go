package admin

import (
	"Goshop/model"
	"Goshop/utils/common"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func buildAfterSaleQueryParam(ctx *gin.Context) map[string]interface{} {
	queryParams := make(map[string]interface{})

	keyword := ctx.DefaultQuery("keyword", "")
	pageNo := ctx.DefaultQuery("page_no", "1")
	endTime := ctx.DefaultQuery("end_time", "")
	orderSn := ctx.DefaultQuery("order_sn", "")
	memberId := ctx.DefaultQuery("member_id", "")
	sellerId := ctx.DefaultQuery("seller_id", "")
	pageSize := ctx.DefaultQuery("page_size", "20")
	startTime := ctx.DefaultQuery("start_time", "")
	goodsName := ctx.DefaultQuery("goods_name", "")
	serviceSn := ctx.DefaultQuery("service_sn", "")
	serviceType := ctx.DefaultQuery("service_type", "")
	serviceStatus := ctx.DefaultQuery("service_status", "")
	createChannel := ctx.DefaultQuery("create_channel", "")

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyword
	queryParams["end_time"] = endTime
	queryParams["order_sn"] = orderSn
	queryParams["page_size"] = pageSize
	queryParams["member_id"] = memberId
	queryParams["seller_id"] = sellerId
	queryParams["start_time"] = startTime
	queryParams["service_sn"] = serviceSn
	queryParams["goods_name"] = goodsName
	queryParams["service_type"] = serviceType
	queryParams["service_status"] = serviceStatus
	queryParams["create_channel"] = createChannel

	return queryParams
}

func buildRefundQueryParam(ctx *gin.Context) map[string]interface{} {
	queryParams := make(map[string]interface{})
	common.ParseFromQuery(ctx)
	keyword := ctx.DefaultQuery("keyword", "")
	pageNo := ctx.DefaultQuery("page_no", "1")
	endTime := ctx.DefaultQuery("end_time", "")
	orderSn := ctx.DefaultQuery("order_sn", "")
	memberId := ctx.DefaultQuery("member_id", "")
	sellerId := ctx.DefaultQuery("seller_id", "")
	refundWay := ctx.DefaultQuery("refund_way", "")
	pageSize := ctx.DefaultQuery("page_size", "20")
	goodsName := ctx.DefaultQuery("goods_name", "")
	serviceSn := ctx.DefaultQuery("service_sn", "")
	startTime := ctx.DefaultQuery("start_time", "")
	refundStatus := ctx.DefaultQuery("refund_status", "")
	createChannel := ctx.DefaultQuery("create_channel", "")

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyword
	queryParams["end_time"] = endTime
	queryParams["order_sn"] = orderSn
	queryParams["page_size"] = pageSize
	queryParams["member_id"] = memberId
	queryParams["seller_id"] = sellerId
	queryParams["start_time"] = startTime
	queryParams["service_sn"] = serviceSn
	queryParams["goods_name"] = goodsName
	queryParams["refund_way"] = refundWay
	queryParams["refund_status"] = refundStatus
	queryParams["create_channel"] = createChannel

	return queryParams
}

func AfterSalesList(ctx *gin.Context) {
	queryParams := buildAfterSaleQueryParam(ctx)

	pageNo := queryParams["page_no"].(int)
	pageSize := queryParams["page_size"].(int)
	data, dataTotal := model.CreateAfterSalesFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func AfterSalesDetail(ctx *gin.Context) {
	serviceSn := ctx.Param("service_sn")
	afterSales, err := model.CreateAfterSalesFactory("").Detail(serviceSn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, afterSales)
}

func AfterSalesExport(ctx *gin.Context) {
	queryParams := buildAfterSaleQueryParam(ctx)
	dataList := model.CreateAfterSalesFactory("").ExportAfterSale(queryParams)
	ctx.JSON(http.StatusOK, dataList)
}

func AfterSalesRefundList(ctx *gin.Context) {
	queryParams := common.ParseFromQuery(ctx)
	pageNo, _ := strconv.Atoi(queryParams["page_no"])
	pageSize, _ := strconv.Atoi(queryParams["page_size"])

	data, dataTotal := model.CreateAfterSalesRefundFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func AfterSalesRefund(ctx *gin.Context) {
	serviceSn := ctx.Param("service_sn")
	fmt.Println(serviceSn)
}
