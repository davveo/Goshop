package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StatisticMemberOrderQuantity(ctx *gin.Context) {

}

func StatisticMemberOrderQuantityPage(ctx *gin.Context) {
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")
	sellerId := ctx.Query("seller_id")

	fmt.Println(year, month, cycleType, sellerId)
}

func StatisticMemberIncrease(ctx *gin.Context) {
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType)
}

func StatisticMemberIncreasePage(ctx *gin.Context) {
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType)
}

func StatisticGoodsPrice(ctx *gin.Context) {
	prices := ctx.Query("prices")
	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId, prices)
}

func StatisticGoodsHot(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId)
}

func StatisticGoodsHotPage(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId)
}

func StatisticGoodsSaleDetail(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId, queryParams, pageSize, pageNo)
}

func StatisticGoodsCollect(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId)
}

func StatisticGoodsCollectPage(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId)
}

func StatisticIndustry(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId)
}

func StatisticIndustryOverView(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	sellerId := ctx.Query("seller_id")
	categoryId := ctx.Query("category_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId, categoryId, queryParams, pageSize, pageNo)
}

func StatisticPageViewShop(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId)
}

func StatisticPageViewGoods(ctx *gin.Context) {
	sellerId := ctx.Query("seller_id")
	year := ctx.Query("year")
	month := ctx.Query("month")
	cycleType := ctx.Query("cycle_type")

	fmt.Println(year, month, cycleType, sellerId)
}

func StatisticPageOrder(ctx *gin.Context) {

}

func StatisticPageMoney(ctx *gin.Context) {

}

func StatisticOrderSalesMoney(ctx *gin.Context) {

}

func StatisticOrderSalesTotal(ctx *gin.Context) {

}

func StatisticOrderRegionForm(ctx *gin.Context) {

}

func StatisticOrderRegionMember(ctx *gin.Context) {

}
func StatisticOrderUnitPrice(ctx *gin.Context) {

}

func StatisticOrderReturnMoney(ctx *gin.Context) {

}
