package controllers

import (
	"exchange_app/global"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToggleLike(ctx *gin.Context) {
	aid := ctx.Param("id")
	_, err := strconv.Atoi(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章ID无效"})
        return
	}

	uid, ok := ctx.Get("userID")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
        return
	}
	userID := uid.(uint)

	key := "article:likes:" + aid
	liked, _ := global.Rdb.SIsMember(ctx, key, userID).Result()

	if liked {
		global.Rdb.SRem(ctx, key, userID)
		count, _ := global.Rdb.SCard(ctx, key).Result()
		ctx.JSON(http.StatusOK, gin.H{
			"liked": false,
            "like_count": count,
            "message": "取消点赞",
		})
		return
	}

	global.Rdb.SAdd(ctx, key, userID)
    count, _ := global.Rdb.SCard(ctx, key).Result()
    ctx.JSON(http.StatusOK, gin.H{
        "liked": true,
        "like_count": count,
        "message": "点赞成功",
    })
}

func GetLikes(ctx *gin.Context) {

	aid := ctx.Param("id")
    key := "article:likes:" + aid

    count, _ := global.Rdb.SCard(ctx, key).Result()

    ctx.JSON(http.StatusOK, gin.H{
        "article_id": aid,
        "like_count": count,
    })

}