package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SpecsValues(ctx *gin.Context) {
	var (
		specId, _ = strconv.Atoi(ctx.Param("spec_id"))
	)

	ctx.JSON(http.StatusOK, model.CreateSpecValuesFactory("").
		ListBySpecId(specId, consts.PermissionADMIN))
}

func UpdateSpecsValues(ctx *gin.Context) {
	var (
		specId, _ = strconv.Atoi(ctx.Param("spec_id"))
	)
	valueLst := ctx.PostFormArray("value_list")

	//fmt.Println(ctx.Request.PostForm)
	valueList, err := model.CreateSpecValuesFactory("").SaveSpecValue(specId, valueLst)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, valueList)
}
