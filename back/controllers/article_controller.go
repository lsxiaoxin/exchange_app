package controllers

import (
	"exchange_app/global"
	"exchange_app/models"
	"net/http"
	"strconv"

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
	var articles []models.Article

	if err := global.Db.Order("created_at desc").Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return		
	}

	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})

}

func DeleteArtile(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id , err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "文章id无效",
		})
		return
	}

	c, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未携带 Token"})
		return
	}
	userID := c.(uint)

	var artilce models.Article
	if err := global.Db.First(&artilce, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	if artilce.UserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "你没有权限删除这篇文章"})
		return
	}

	if err := global.Db.Delete(&artilce).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})

}

