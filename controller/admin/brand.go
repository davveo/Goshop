package admin

import (
	"Goshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BrandList(context *gin.Context) {
	queryParams := make(map[string]interface{})
	pageNo, _ := strconv.Atoi(context.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("page_size", "20"))

	name := context.Query("name")

	queryParams["name"] = name
	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateBrandFactory("").List(queryParams, name)

	context.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func BrandAllList(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		model.CreateBrandFactory("").GetALllBrands())
}
