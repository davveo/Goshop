package admin

import (
	"Goshop/model"
	"Goshop/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func DistributionCommissionTplDel(ctx *gin.Context) {
	tplId := ctx.Param("tplId")
	err := model.CreateCommissionTplFactory("").Delete(tplId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type DRequest struct {
	Sex    int    `json:"sex"`
	Region string `json:"region"`
}

func DistributionCommissionTpl(ctx *gin.Context) {
	// body数据获取
	//data, err := model.CreateCommissionTplFactory("").Add(body)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"code":    http.StatusInternalServerError,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, data)
}

func DistributionCommissionTplEdit(ctx *gin.Context) {
	// TODO body数据获取
	//data, err := model.CreateCommissionTplFactory("").Edit(body)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"code":    http.StatusInternalServerError,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, data)
}
