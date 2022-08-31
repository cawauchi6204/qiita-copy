package presentation

import (
	"os"

	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/gin-gonic/gin"
)

func Main() {
	r := gin.Default()
	users := service.GetAllUser()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
