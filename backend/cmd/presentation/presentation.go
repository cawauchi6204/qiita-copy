package presentation

import (
	"net/http"
	"os"
	"strconv"

	"github.com/cawauchi6204/qiita-copy/cmd/application/coordinator"
	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/cawauchi6204/qiita-copy/cmd/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := middleware.Middleware()
	r.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("loginUser", c.PostForm("userId"))
		session.Save()
		c.String(http.StatusOK, "ログイン完了")
	})
	r.POST("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.String(http.StatusOK, "ログアウトしました")
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
		var user repository.User
		c.BindJSON(&user)
		coordinator.SignUp(c)
		c.JSON(200, "成功したかも")
	})

	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
