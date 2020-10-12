package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Region(ctx *gin.Context) {
	regionId := ctx.Param("region_id")
	fmt.Println(regionId)
}
