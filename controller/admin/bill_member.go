package admin

import (
	"Goshop/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BillMemberList(ctx *gin.Context) {
	queryParams := buildAfterSaleQueryParam(ctx)

	pageNo := queryParams["page_no"].(int)
	pageSize := queryParams["page_size"].(int)
	data, dataTotal := model.CreateBillMemberFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func BillMemberDetail(ctx *gin.Context) {

}

func DownBillMember(ctx *gin.Context) {

}

func ExportBillMember(ctx *gin.Context) {

}
