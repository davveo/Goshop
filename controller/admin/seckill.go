package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SeckillList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	status := ctx.DefaultQuery("status", "")
	endTime := ctx.DefaultQuery("end_time", "")
	startTime := ctx.DefaultQuery("start_time", "")
	seckillName := ctx.DefaultQuery("seckill_name", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["status"] = status
	queryParams["page_no"] = pageNo
	queryParams["end_time"] = endTime
	queryParams["page_size"] = pageSize
	queryParams["start_time"] = startTime
	queryParams["seckill_name"] = seckillName
	queryParams["delete_status"] = consts.NORMAL
	data, dataTotal := model.CreateSeckillFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
