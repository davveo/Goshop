package admin

import (
	"Goshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
