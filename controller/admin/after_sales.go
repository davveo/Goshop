package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AfterSalesList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyword := ctx.DefaultQuery("keyword", "")
	endTime := ctx.DefaultQuery("end_time", "")
	orderSn := ctx.DefaultQuery("order_sn", "")
	memberId := ctx.DefaultQuery("member_id", "")
	sellerId := ctx.DefaultQuery("seller_id", "")
	startTime := ctx.DefaultQuery("start_time", "")
	goodsName := ctx.DefaultQuery("goods_name", "")
	serviceSn := ctx.DefaultQuery("service_sn", "")
	serviceType := ctx.DefaultQuery("service_type", "")
	serviceStatus := ctx.DefaultQuery("service_status", "")
	createChannel := ctx.DefaultQuery("create_channel", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

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

	data, dataTotal := model.CreateAfterSalesFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func AfterSalesRefundList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateAfterSalesRefundFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func AfterSalesDetail(ctx *gin.Context) {

}
