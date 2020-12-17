package admin

import (
	"Goshop/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PointCategory(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	queryParams["cat_id"] = ctx.DefaultQuery("cat_id", "0")
	data := model.CreatePointCateGoryFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, data)
}

func CreatePointCategory(ctx *gin.Context) {

}

func UpdatePointCategory(ctx *gin.Context) {

}

func DelPointCategory(ctx *gin.Context) {

}

func FindOnePointCategory(ctx *gin.Context) {

}
