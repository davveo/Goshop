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

func BrandList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	name := ctx.DefaultQuery("name", "")

	queryParams["name"] = name
	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateBrandFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func BrandAllList(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		model.CreateBrandFactory("").GetALllBrands())
}

func Brand(ctx *gin.Context) {
	brandId, _ := strconv.Atoi(ctx.Param("brand_id"))
	brand, err := model.CreateBrandFactory("").GetModel(brandId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

func CreateBrand(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "")
	logo := ctx.DefaultPostForm("logo", "")

	brandRequest := request.BrandRequest{
		Name: name, Logo: logo, Disabled: "1",
	}
	mapData := transfer.StructToMap(brandRequest)
	brand, err := model.CreateBrandFactory("").Add(mapData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

func UpdateBrand(ctx *gin.Context) {
	var (
		err      error
		logo     = ctx.DefaultPostForm("logo", "")
		name     = ctx.DefaultPostForm("name", "")
		disabled = ctx.DefaultPostForm("disabled", "")
		brandId  = ctx.DefaultPostForm("brand_id", "")
	)

	brandRequest := request.BrandRequest{
		Name: name, Logo: logo, BrandId: brandId, Disabled: disabled,
	}
	mapData := transfer.StructToMap(brandRequest)
	brand, err := model.CreateBrandFactory("").Edit(mapData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

func DeleteBrand(ctx *gin.Context) {
	brandIds := ctx.Param("brand_id")
	brandIdList := strings.Split(brandIds, ",")
	err := model.CreateBrandFactory("").Delete(transfer.StringToInt(brandIdList))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
