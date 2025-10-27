package controllers

import (
	"exchange_app/global"
	"exchange_app/models"
	"exchange_app/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var user models.User
	
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate((&user)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

func Login(ctx *gin.Context) {
	var input models.User
	
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	
	if err := global.Db.Where("user_name = ?", input.UserName).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "用户不存在",
		})
		return		
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "密码错误",
		})
		return		
	}

	token, _ := utils.GenerateToken(user.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token": token,
	})
}

