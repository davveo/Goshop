package admin

import (
	"Goshop/model"
	"Goshop/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DistributionUpgradeLog(ctx *gin.Context) {

}

func DistributionMember(ctx *gin.Context) {

}

func DistributionMemberChangeTpl(ctx *gin.Context) {

}

func DistributionBillOrder(ctx *gin.Context) {

}

func DistributionBillSellbackOrder(ctx *gin.Context) {

}

func DistributionBillTotalList(ctx *gin.Context) {
	queryParams := common.ParseFromQuery(ctx)

	pageNo, _ := strconv.Atoi(queryParams["page_no"])
	pageSize, _ := strconv.Atoi(queryParams["page_size"])
	data, dataTotal := model.CreateBillTotalFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func DistributionWithdraw(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	uname := ctx.DefaultQuery("uname", "")
	status := ctx.DefaultQuery("status", "") // APPLY:申请中/VIA_AUDITING:审核通过/FAIL_AUDITING:审核未通过/RANSFER_ACCOUNTS:已转账
	startTime := ctx.DefaultQuery("start_time", "")
	endTime := ctx.DefaultQuery("end_time", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["uname"] = uname
	queryParams["status"] = status
	queryParams["page_no"] = pageNo
	queryParams["end_time"] = endTime
	queryParams["page_size"] = pageSize
	queryParams["start_time"] = startTime
	data, dataTotal := model.CreateWithDrawFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})

}

func DistributionWithdrawExport(ctx *gin.Context) {

}

func DistributionWithdrawBatchAuditing(ctx *gin.Context) {

}

func DistributionWithdrawBatchAccountPaid(ctx *gin.Context) {

}
