package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// path解析

	router.GET("/:name/:id", func(c *gin.Context) {
		var name = c.Param("name")
		var id = c.Param("id")

		c.JSON(200, gin.H{"name": name, "uuid": id})
	})

	// 参数绑定
	router.POST("/login", func(c *gin.Context) {
		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": u.Username,
				"password": u.Password,
			})
		}
	})

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload/index.tmpl", nil)
	})

	// 文件上传
	router.POST("/upload", func(c *gin.Context) {
		r, err := c.FormFile("example")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
		} else {
			// 保存文件
			dst := fmt.Sprintf("./%s", r.Filename)
			c.SaveUploadedFile(r, dst)
		}
	})

	router.Run(":8080")
}
