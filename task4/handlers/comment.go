package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task4/models"
)

func CreateComment(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	userID, _ := c.Get("user_id")
	comment.UserID = uint(userID.(float64))
	comment.PostID = uint(postID)

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论创建失败"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func GetComments(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	var comments []models.Comment
	if err := db.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论列表失败"})
		return
	}
	c.JSON(http.StatusOK, comments)
}
