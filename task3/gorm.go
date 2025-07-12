package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

// ====================== 模型结构 ======================

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	PostCount int    // 文章数量统计字段
	Posts     []Post // 一对多关联
}

type Post struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Content  string
	UserID   uint
	Comments []Comment
}

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	PostID  uint
}

// ================ 钩子：Post 创建时更新 User.PostCount ================

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
	return
}

// ================ 钩子：Comment 删除时检查并更新 Post 状态 ================

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("title", gorm.Expr("CONCAT(title, '【无评论】')")).Error
	}
	return
}

// ================ 主入口 ================

func main() {
	initDB()

	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	seedData()
	queryUserWithPostsAndComments(1)
	queryMostCommentedPost()
}

func initDB() {
	var err error
	username := "root"
	password := "root"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local", username, password)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}
}

func seedData() {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}

	user := User{Name: "张三"}
	post1 := Post{Title: "第一篇文章", Content: "内容A"}
	post2 := Post{Title: "第二篇文章", Content: "内容B"}
	comment1 := Comment{Content: "评论1"}
	comment2 := Comment{Content: "评论2"}
	comment3 := Comment{Content: "评论3"}

	post1.Comments = []Comment{comment1, comment2}
	post2.Comments = []Comment{comment3}
	user.Posts = []Post{post1, post2}

	if err := db.Create(&user).Error; err != nil {
		log.Println("插入失败:", err)
	}
}

func queryUserWithPostsAndComments(userID uint) {
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		log.Println("查询失败:", err)
		return
	}

	fmt.Printf("用户: %s，共发表 %d 篇文章：\n", user.Name, len(user.Posts))
	for _, post := range user.Posts {
		fmt.Printf("  - 文章: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("    * 评论: %s\n", comment.Content)
		}
	}
}

func queryMostCommentedPost() {
	type resultWithCount struct {
		Post
		CommentCount int64
	}

	var result resultWithCount
	err := db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Preload(clause.Associations).
		Scan(&result).Error

	if err != nil {
		log.Println("查询失败:", err)
		return
	}

	fmt.Printf("\n评论最多的文章: %s，共 %d 条评论\n", result.Title, result.CommentCount)
	for _, c := range result.Comments {
		fmt.Printf(" - %s\n", c.Content)
	}
}
