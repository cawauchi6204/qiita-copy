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
	r.GET("/users", func(c *gin.Context) {
		users := service.GetAllUsers()
		c.JSON(200, users)
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
	}

	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
