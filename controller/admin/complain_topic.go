package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ComplainTopicsList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateComplainTopicFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateComplainTopics(ctx *gin.Context) {

}

func UpdateComplainTopics(ctx *gin.Context) {

}

func DelComplainTopics(ctx *gin.Context) {

}

func FindOneComplainTopics(ctx *gin.Context) {

}
