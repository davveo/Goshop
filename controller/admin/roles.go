package admin

import (
	"Goshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListRoles(ctx *gin.Context) {

}

func CreateRoles(ctx *gin.Context) {

}

func UpdateRoles(ctx *gin.Context) {

}

func DelRoles(ctx *gin.Context) {

}

func FindOneRoles(ctx *gin.Context) {

}

func RoleCheck(context *gin.Context) {
	var (
		err       error
		roleList  []string
		roleID, _ = strconv.Atoi(context.Param("roleId"))
	)

	roleList, err = model.CreateRoleFactory("").GetRoleMenu(roleID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, roleList)
}
