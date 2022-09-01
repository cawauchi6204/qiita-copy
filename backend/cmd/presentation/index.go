package presentation

import (
	"os"
	"strconv"

	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		users := service.GetAllUser()
		c.JSON(200, users)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		i, _ := strconv.Atoi(userId)
		user := service.GetUserById(i)
		c.JSON(200, user)
	})
	r.POST("/user/create", func(c *gin.Context) {

	})
	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
