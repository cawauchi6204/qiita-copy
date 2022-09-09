package presentation

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cawauchi6204/qiita-copy/cmd/application/coordinator"
	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := middleware.Middleware()

	// http://localhost/set_cookieでクッキーをセットし、
	// http://localhost/get_cookieでクッキーを取得し、
	// http://localhost/get_cookieでクッキーを削除することができる
	r.GET("/set_cookie", func(c *gin.Context) {
		// セット
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

		c.String(http.StatusOK, "Set cookie")
	})

	r.GET("/get_cookie", func(c *gin.Context) {
		// 取得
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			c.String(http.StatusOK, "Failed to get cookie")
			return
		}
		r.GET("/delete_cookie", func(c *gin.Context) {
			c.SetCookie("gin_cookie", "", -1, "/", "localhost", false, false)
		})

		c.String(http.StatusOK, fmt.Sprintf("Cookie value: %s \n", cookie))
	})

	// http://localhost：3000/set_cookie-by-frontでクッキーをセットし、
	// http://localhost：3000/get_cookie-by-frontでクッキーを取得し、
	// http://localhost：3000/get_cookie-by-frontでクッキーを削除することができ「ない」
	r.GET("/set_cookie-by-front", func(c *gin.Context) {
		// セット
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

		c.String(http.StatusOK, "Set cookie")
	})

	r.GET("/get_cookie-by-front", func(c *gin.Context) {
		// 取得
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			c.String(http.StatusOK, "Failed to get cookie")
			return
		}
		r.GET("/delete_cookie-by-front", func(c *gin.Context) {
			c.SetCookie("gin_cookie", "", -1, "/", "localhost", false, false)
		})

		c.String(http.StatusOK, fmt.Sprintf("Cookie value: %s \n", cookie))
	})
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
