package admin

import (
	"Goshop/model"
	"Goshop/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListArticle(ctx *gin.Context) {
	queryParams := common.ParseFromQuery(ctx)
	data, dataTotal := model.CreateArticleFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    queryParams["page_no"],
		"page_size":  queryParams["page_size"],
	})
}

func CreateArticle(ctx *gin.Context) {
	body := common.ParseFromBody(ctx)
	article, err := model.CreateArticleFactory("").Add(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, article)
}

func UpdateArticle(ctx *gin.Context) {
	body := common.ParseFromBody(ctx)
	article, err := model.CreateArticleFactory("").Edit(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, article)
}

func DelArticle(ctx *gin.Context) {
	articleId, _ := strconv.Atoi(ctx.Param("article_id"))
	err := model.CreateArticleFactory("").Delete(articleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func FindOneArticle(ctx *gin.Context) {
	articleId, _ := strconv.Atoi(ctx.Param("article_id"))
	article, err := model.CreateArticleFactory("").GetModel(articleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, article)
}
