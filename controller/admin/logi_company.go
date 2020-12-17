package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListLogiCompany(ctx *gin.Context) {
	queryParams := make(map[string]interface{})
	name := ctx.DefaultQuery("name", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["name"] = name
	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["status"] = consts.NORMAL

	data, dataTotal := model.CreateLogisticsCompanyFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateLogiCompany(ctx *gin.Context) {
}

func UpdateLogiCompany(ctx *gin.Context) {
}

func DelLogiCompany(ctx *gin.Context) {
}

func FindOneLogiCompany(ctx *gin.Context) {
}

func OpenCloseLogiCompany(ctx *gin.Context) {
}
