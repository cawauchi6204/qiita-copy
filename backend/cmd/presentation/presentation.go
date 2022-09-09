package presentation

import (
	"os"
	"strconv"

	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := middleware.Middleware()
	r.GET("/posts", func(c *gin.Context) {
		posts := service.GetAllPosts()
		c.JSON(200, posts)
	})
	r.GET("/users", func(c *gin.Context) {
		users := service.GetAllUsers()
		c.JSON(200, users)
	})
	r.GET("/users/:userId/posts", func(c *gin.Context) {
		userId := c.Param("userId")
		posts := service.GetAllPostsByUserId(userId)
		c.JSON(200, posts)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		i, _ := strconv.Atoi(userId)
		user := service.GetUserById(i)
		c.JSON(200, user)
	})
	r.POST("/user", func(c *gin.Context) {
		var user repository.User
		c.BindJSON(&user)
		service.Signup(user.Name, user.Email, user.Password)
		c.JSON(200, "成功したかも")
	})

	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
