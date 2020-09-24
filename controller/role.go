package controller

import (
	"net/http"
	"orange/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
