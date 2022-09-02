package presentation

import (
	"os"

	"github.com/cawauchi6204/qiita-copy/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func Main() {
	r := middleware.Router()
	r.GET("/users", func(c *gin.Context) {
		// users := service.GetAllUser()
		c.JSON(200, "hoge")
	})
	// r.GET("/user/:id", func(c *gin.Context) {
	// 	userId := c.Param("id")
	// 	i, _ := strconv.Atoi(userId)
	// 	user := service.GetUserById(i)
	// 	c.JSON(200, user)
	// })
	// r.POST("/user", func(c *gin.Context) {
	// 	var user UserRepository.User
	// 	c.BindJSON(&user)
	// 	u := service.SaveUser(user)
	// 	c.JSON(200, u)
	// })
	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
