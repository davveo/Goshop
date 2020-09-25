package controller

import (
	"net/http"
	"orange/model"

	"github.com/gin-gonic/gin"
)

func ShopList(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		model.CreateShopFactory("").All())
}
