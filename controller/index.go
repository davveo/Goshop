package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Eshop/global/consts"
	"Eshop/model"
	"Eshop/utils/time_utils"
)

func Index(context *gin.Context) {
	// map[string]interface{}{"receive_money": 0.03, "refund_money": 0.01, "real_money": 0.02}

	startTime, endTime := time_utils.GetStartTimeAndEndTime(consts.MONTH)
	context.JSON(http.StatusOK, gin.H{
		"goods_vos":   model.CreateGoodsFactory("").NewGoods(5),
		"member_vos":  model.CreateMemberFactory("").NewMember(5),
		"sales_total": model.CreateOrderStatisticFactory("").GetSalesMoneyTotal(startTime, endTime),
	})
}
