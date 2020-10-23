package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GoodsList(context *gin.Context) {
	queryParams := make(map[string]interface{})
	IsAuth, _ := strconv.Atoi(context.Query("is_auth"))
	supplierGoodsType := context.Query("supplier_goods_type")
	goodsType := context.Query("goods_type") // POINT
	pageNo, _ := strconv.Atoi(context.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["is_auth"] = IsAuth
	queryParams["page_size"] = pageSize
	queryParams["goods_type"] = goodsType
	queryParams["supplier_goods_type"] = supplierGoodsType
	data, dataTotal := model.CreateGoodsFactory("").List(queryParams)

	context.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func GoodsUp(ctx *gin.Context) {
	var (
		err       error
		goodId, _ = strconv.Atoi(ctx.Param("good_id"))
	)

	err = model.CreateGoodsFactory("").Up(goodId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func GoodsUnder(ctx *gin.Context) {
	var (
		err        error
		goodIds, _ = strconv.Atoi(ctx.Param("good_ids"))
		reason     = ctx.PostForm("reason")
	)

	err = model.CreateGoodsFactory(ctx, "").Under([]int{goodIds}, reason, consts.PermissionADMIN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

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

func CategoryList(context *gin.Context) {
	// parent_id = 0 说明为顶级

	var (
		err          error
		categoryList []model.CategoryTree
		parentID, _  = strconv.Atoi(context.Param("parent_id"))
	)

	err, categoryList = model.CreateCategoryFactory("").List(parentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, categoryList)
}
