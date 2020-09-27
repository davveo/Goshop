package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orange/model"
	"strconv"
)

func PointCategory(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	queryParams["cat_id"] = ctx.DefaultQuery("cat_id", "0")
	data := model.CreatePointCateGoryFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, data)
}
