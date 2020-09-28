package controller

import (
	"net/http"
	"Eshop/model"

	"github.com/gin-gonic/gin"
)

func Health(context *gin.Context) {
	if !model.CreateHealthFactory("").Check() {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "数据库链接异常",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}
