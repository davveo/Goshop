package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"Goshop/model/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GoodsList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})
	IsAuth, _ := strconv.Atoi(ctx.Query("is_auth"))
	supplierGoodsType := ctx.Query("supplier_goods_type")
	goodsType := ctx.Query("goods_type") // POINT
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["is_auth"] = IsAuth
	queryParams["page_size"] = pageSize
	queryParams["goods_type"] = goodsType
	queryParams["supplier_goods_type"] = supplierGoodsType
	data, dataTotal := model.CreateGoodsFactory(ctx, "").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func GoodsUp(ctx *gin.Context) {
	var (
		err        error
		goodsId, _ = strconv.Atoi(ctx.Param("goods_id"))
	)

	err = model.CreateGoodsFactory(ctx, "").Up(goodsId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func GoodsUnder(ctx *gin.Context) {
	var (
		err        error
		goodsId, _ = strconv.Atoi(ctx.Param("goodsId"))
		reason     = ctx.PostForm("reason")
	)

	err = model.CreateGoodsFactory(ctx, "").Under([]int{goodsId}, reason, consts.PermissionADMIN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func GoodsBatchAudit(ctx *gin.Context) {
	var (
		batchAuditRequest request.BatchAuditRequest
	)
	if err := ctx.BindJSON(&batchAuditRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if err := model.CreateGoodsFactory(ctx, "").
		BatchAuditGoods(map[string]interface{}{
			"goods_ids": batchAuditRequest.GoodsIds,
			"message":   batchAuditRequest.Message,
			"pass":      batchAuditRequest.Pass,
		}); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func GoodsListDetail(ctx *gin.Context) {

}
