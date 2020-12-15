package admin

import (
	"Goshop/model"
	"Goshop/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DistributionCommissionTplList(ctx *gin.Context) {
	queryParams := common.ParseFromQuery(ctx)

	pageNo, _ := strconv.Atoi(queryParams["page_no"])
	pageSize, _ := strconv.Atoi(queryParams["page_size"])
	data, dataTotal := model.CreateCommissionTplFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func DistributionCommissionTplDetail(ctx *gin.Context) {
	tplId := ctx.Param("tplId")

	commissionTpl, err := model.CreateCommissionTplFactory("").GetModel(tplId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, commissionTpl)
}

func DistributionUpgradeLog(ctx *gin.Context) {

}

func DistributionMember(ctx *gin.Context) {

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

func DistributionSetting(ctx *gin.Context) {

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
