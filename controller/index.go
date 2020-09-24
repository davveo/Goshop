package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"sales_total": map[string]interface{}{"receive_money": 0.03, "refund_money": 0.01, "real_money": 0.02},
		"goods_vos":   []string{},
		"member_vos":  []string{},
	})
}
