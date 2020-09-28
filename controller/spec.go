package controller

import (
	"net/http"
	"Goshop/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SpecsList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWord := ctx.Query("keyword")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keyword"] = keyWord
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateSpecFactory("").List(queryParams, keyWord)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
