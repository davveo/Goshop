package admin

import (
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MessageList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateMessageFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateMessage(ctx *gin.Context) {

}

func MessageTemplate(ctx *gin.Context) {
	queryParams := make(map[string]interface{})
	messageType := ctx.DefaultQuery("type", "SHOP") //MEMBER/SHOP
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["type"] = messageType

	data, dataTotal := model.CreateMessageTemplateFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func UpdateMessageTemplate(ctx *gin.Context) {

}

func Sync(ctx *gin.Context) {
	// 查询微信服务消息模板是否已经同步

	isSync := model.CreateWechatMessageTemplateFactory("").IsSync()

	ctx.JSON(http.StatusOK, isSync)

}

func SyncMsgTmp(ctx *gin.Context) {

}

func ListWechatMsg(ctx *gin.Context) {
	queryParams := make(map[string]interface{})
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize

	data, dataTotal := model.CreateWechatMessageTemplateFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func FindOneWechatMsg(ctx *gin.Context) {

}

func UpdateWechatMsg(ctx *gin.Context) {

}

func DelWechatMsg(ctx *gin.Context) {

}
