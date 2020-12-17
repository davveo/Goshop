package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Region(ctx *gin.Context) {
	regionId := ctx.Param("region_id")
	fmt.Println(regionId)
}

func CreateRegions(ctx *gin.Context) {

}

func UpdateRegions(ctx *gin.Context) {

}

func DelRegions(ctx *gin.Context) {

}

func FindOneRegions(ctx *gin.Context) {

}

func FindChildrenRegions(ctx *gin.Context) {

}
