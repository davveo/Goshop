package seller

import (
	"Goshop/global/consts"
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ask(ctx *gin.Context) {
	queryParams := make(map[string]interface{})
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["status"] = consts.NORMAL
	queryParams["auth_status"] = consts.PASS_AUDIT
	queryParams["seller_id"] = "" // TODO 从当前登录用户中获取用户id
	data, dataTotal := model.CreateMemberAskFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func AskReply(ctx *gin.Context) {
	// TODO
}

func AskDetail(ctx *gin.Context) {
	// TODO

}
