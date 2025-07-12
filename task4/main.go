package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"task4/handlers"
	"task4/middleware"
	"task4/models"
)

var db *gorm.DB

func main() {
	// 初始化数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("无法连接 MySQL 数据库:", err)
	}

	// 自动迁移模型
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// 初始化 Gin 路由
	r := gin.Default()

	// 将 db 传递给 handlers
	handlers.InitDB(db)

	// 路由分组
	api := r.Group("/api")
	{
		// 认证相关
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		// 需要认证的路由
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// 文章管理
			protected.POST("/posts", handlers.CreatePost)
			protected.GET("/posts", handlers.GetPosts)
			protected.GET("/posts/:id", handlers.GetPost)
			protected.PUT("/posts/:id", handlers.UpdatePost)
			protected.DELETE("/posts/:id", handlers.DeletePost)

			// 评论管理
			protected.POST("/posts/:id/comments", handlers.CreateComment)
			protected.GET("/posts/:id/comments", handlers.GetComments)
		}
	}

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
