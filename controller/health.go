package controller

import (
	"Goshop/global/errno"
	"Goshop/model"
	"Goshop/utils/response"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	if !model.CreateHealthFactory("").Check() {
		response.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}
	response.SendResponse(ctx, nil, nil)
}
