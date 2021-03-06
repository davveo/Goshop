package admin

import (
	"Goshop/model"
	"Goshop/model/request"
	"Goshop/utils/transfer"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CategoryList(ctx *gin.Context) {
	// parent_id = 0 说明为顶级

	var (
		err          error
		categoryList []model.CategoryTree
		parentID, _  = strconv.Atoi(ctx.Param("parent_id"))
	)

	err, categoryList = model.CreateCategoryFactory("").List(parentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, categoryList)
}

func CategoryAllList(ctx *gin.Context) {

}

func CreateCategory(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "")
	image := ctx.DefaultPostForm("image", "")
	isShow := ctx.DefaultPostForm("isShow", "")
	advImage := ctx.DefaultPostForm("advImage", "")
	parentId := ctx.DefaultPostForm("parent_id", "")
	advImageLink := ctx.DefaultPostForm("advImageLink", "")
	categoryOrder := ctx.DefaultPostForm("category_order", "")

	categoryRequest := request.CategoryRequest{
		Name: name, ParentId: parentId,
		CategoryOrder: categoryOrder,
		Image:         image, IsShow: isShow,
		AdvImage:     advImage,
		AdvImageLink: advImageLink,
	}
	mapData := transfer.StructToMap(categoryRequest)
	category, err := model.CreateCategoryFactory("").Add(mapData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func EditCategory(ctx *gin.Context) {

}

func DelCategory(ctx *gin.Context) {

}

func Category(ctx *gin.Context) {

}

func CategoryBrand(ctx *gin.Context) {

}

func SaveCategoryBrand(ctx *gin.Context) {

}

func CategorySpecs(ctx *gin.Context) {

}

func SaveCategorySpecs(ctx *gin.Context) {

}

func CategoryParam(ctx *gin.Context) {

}
