package admin

import (
	"Goshop/model"
	"Goshop/model/request"
	"Goshop/utils/transfer"
	"net/http"
	"strconv"
	"strings"

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
	data, dataTotal := model.CreateSpecFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateSpecs(ctx *gin.Context) {
	specName := ctx.DefaultPostForm("spec_name", "")
	specMemo := ctx.DefaultPostForm("spec_memo", "")

	specRequest := request.SpecsRequest{
		SpecMemo: specMemo, SpecName: specName,
	}
	mapData := transfer.StructToMap(specRequest)

	// extra params
	mapData["seller_id"] = "0"
	mapData["disabled"] = 1

	spec, err := model.CreateSpecFactory("").Add(mapData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, spec)
}

func UpdateSpecs(ctx *gin.Context) {
	var (
		err      error
		specMemo = ctx.DefaultPostForm("spec_memo", "")
		specName = ctx.DefaultPostForm("spec_name", "")
		disabled = ctx.DefaultPostForm("disabled", "")
		specId   = ctx.DefaultPostForm("spec_id", "")
		sellerId = ctx.DefaultPostForm("seller_id", "")
	)

	specRequest := request.SpecsRequest{
		SpecMemo: specMemo, SpecName: specName,
	}
	mapData := transfer.StructToMap(specRequest)

	// extra params
	mapData["spec_id"] = specId
	mapData["disabled"] = disabled
	mapData["seller_id"] = sellerId

	spec, err := model.CreateSpecFactory("").Edit(mapData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, spec)
}

func DeleteSpecs(ctx *gin.Context) {

	specIds := ctx.Param("spec_id")
	split := strings.Split(specIds, ",")
	err := model.CreateSpecFactory("").Delete(transfer.StringToInt(split))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func SpecsValues(ctx *gin.Context) {
	//var (
	//	err       error
	//	specId, _ = strconv.Atoi(ctx.Param("spec_id"))
	//)

}

func UpdateSpecsValues(ctx *gin.Context) {
	//var (
	//	err       error
	//	specId, _ = strconv.Atoi(ctx.Param("spec_id"))
	//)

}
