package presentation

import (
	"os"
	"strconv"

	"github.com/cawauchi6204/qiita-copy/cmd/application/coordinator"
	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := middleware.Middleware()

	r.POST("/login", func(c *gin.Context) {
		coordinator.Login(c)
		c.JSON(200, "成功したかも")
	})
	r.POST("/logout", func(c *gin.Context) {
		coordinator.Logout(c)
		c.JSON(200, "成功したかも")
	})
	r.GET("/posts", func(c *gin.Context) {
		posts := service.GetAllPosts()
		c.JSON(200, posts)
	})
	r.GET("/users", func(c *gin.Context) {
		users := service.GetAllUsers()
		c.JSON(200, users)
	})
	r.GET("/:userId/posts", func(c *gin.Context) {
		userId := c.Param("userId")
		posts := service.GetAllPostsByUserId(userId)
		c.JSON(200, posts)
	})
	r.GET("/:userId/items/:postId", func(c *gin.Context) {
		userId := c.Param("userId")
		postId := c.Param("postId")
		post := service.GetPostByUserId(userId, postId)
		c.JSON(200, post)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		i, _ := strconv.Atoi(userId)
		user := service.GetUserById(i)
		c.JSON(200, user)
	})
	r.POST("/signup", func(c *gin.Context) {
		coordinator.SignUp(c)
		c.JSON(200, "成功したかも")
	})
	authUserGroup := r.Group("/")
	authUserGroup.Use(middleware.LoginCheckMiddleware())
	{
		authUserGroup.GET("/mypage", func(c *gin.Context) {
			c.JSON(200, "認証されています")
		})
		authUserGroup.GET("/myinfo", func(c *gin.Context) {
			user := coordinator.GetMyInfo(c)
			c.JSON(200, user)
		})
		authUserGroup.POST("/post", func(c *gin.Context) {
			service.CreatePost(c)
		})
	}

	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
