package presentation

import (
	"net/http"
	"os"

	"github.com/cawauchi6204/qiita-copy/cmd/application/coordinator"
	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/middleware"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := middleware.Middleware()

	r.POST("/signup", func(c *gin.Context) {
		coordinator.SignUp(c)
		c.JSON(200, "成功したかも")
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
	r.GET("user/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		user := service.GetUserById(userId)
		c.JSON(200, user)
	})
	r.GET("user/:userId/posts", func(c *gin.Context) {
		userId := c.Param("userId")
		posts := service.GetAllPostsByUserId(userId)
		c.JSON(200, posts)
	})
	r.GET("user/:userId/items/:postId", func(c *gin.Context) {
		userId := c.Param("userId")
		postId := c.Param("postId")
		post := service.GetPostByUserId(userId, postId)
		c.JSON(200, post)
	})
	r.GET("user/:userId/items/:postId/likes", func(c *gin.Context) {
		postId := c.Param("postId")
		likes := service.GetLikesByPostId(postId)

		userIds := []string{}
		for _, v := range likes {
			userIds = append(userIds, v.LikeUserId)
		}
		users := service.GetUsersById(userIds)
		c.JSON(200, users)
	})
	r.GET("/post/:postId/likes", func(c *gin.Context) {
		postId := c.Param("postId")
		likes := service.GetLikesByPostId(postId)
		c.JSON(200, likes)
	})
	r.GET("/post/:postId/tags", func(c *gin.Context) {
		postId := c.Param("postId")
		tags := service.GetTagsByPostId(postId)
		c.JSON(200, tags)
	})
	r.GET("/tag/:tagId", func(c *gin.Context) {
		tagId := c.Param("tagId")
		tag := service.GetTagById(tagId)
		c.JSON(200, tag)
	})

	// ↓↓↓↓ログイン済みのときに叩けるAPI↓↓↓↓
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
			post := service.CreatePost(c)
			c.JSON(200, post)
		})
		authUserGroup.PUT("/user", func(c *gin.Context) {
			service.UpdateUser(c)
		})
		authUserGroup.PUT("/like", func(c *gin.Context) {
			service.UpdateLike(c)
		})
		authUserGroup.POST("/tag", func(c *gin.Context) {
			var request request.CreateTagRequest
			err := c.BindJSON(&request)
			if err != nil {
				c.Status(http.StatusBadRequest)
			} else {
				service.CreateTag(request.ID, request.ImgUrl)
				service.CreatePostTag(request.PostId, request.ID)
			}
		})
	}
	// ↑↑↑↑ここまで↑↑↑↑

	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
