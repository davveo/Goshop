package controller

import (
	"net/http"
	"orange/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GoodsList(context *gin.Context) {
	queryParams := make(map[string]interface{})
	pageNo, _ := strconv.Atoi(context.Query("page_no"))
	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	IsAuth, _ := strconv.Atoi(context.Query("is_auth"))
	supplierGoodsType := context.Query("supplier_goods_type")

	queryParams["page_no"] = pageNo
	queryParams["is_auth"] = IsAuth
	queryParams["page_size"] = pageSize
	queryParams["supplier_goods_type"] = supplierGoodsType
	data, dataTotal := model.CreateGoodsFactory("").List(queryParams)

	context.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func BrandList(context *gin.Context) {
	queryParams := make(map[string]interface{})
	pageNo, _ := strconv.Atoi(context.Query("page_no"))
	pageSize, _ := strconv.Atoi(context.Query("page_size"))

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
