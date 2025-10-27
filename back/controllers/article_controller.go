package controllers

import (
	"exchange_app/global"
	"exchange_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uid, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未携带 Token"})
		return
	}

	userID := uid.(uint)  
	article.UserID = userID

	if err := global.Db.AutoMigrate((&article)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "发布成功"})
}

func GetArtiles(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "文章",
	})
}

