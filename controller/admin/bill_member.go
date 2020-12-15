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
	id := ctx.Param("id")
	billMember, err := model.CreateBillMemberFactory("").GetBillMember(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, billMember)

}

func DownBillMember(ctx *gin.Context) {
	id := ctx.Query("id")
	memberId := ctx.Query("member_id")
	// TODO
	model.CreateBillMemberFactory("").AllDown(id, memberId)
}

func ExportBillMember(ctx *gin.Context) {
	queryParams := buildAfterSaleQueryParam(ctx)

	queryParams["page_no"] = 1
	queryParams["page_size"] = 99999
	data, _ := model.CreateBillMemberFactory("").List(queryParams)
	ctx.JSON(http.StatusOK, data)
}
