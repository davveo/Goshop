package controller

import (
	"Goshop/global/consts"
	"Goshop/model"
	"Goshop/utils/time_utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	// map[string]interface{}{"receive_money": 0.03, "refund_money": 0.01, "real_money": 0.02}
	nowYear := strconv.Itoa(time.Now().Year())
	startTime, endTime := time_utils.GetStartTimeAndEndTime(consts.MONTH)
	context.JSON(http.StatusOK, gin.H{
		"goods_vos":   model.CreateGoodsFactory("").NewGoods(5),
		"member_vos":  model.CreateMemberFactory("").NewMember(5),
		"sales_total": model.CreateOrderStatisticFactory("").GetSalesMoneyTotal(nowYear, startTime, endTime),
	})
}
