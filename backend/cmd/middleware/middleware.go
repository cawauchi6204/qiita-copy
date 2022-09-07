package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Middleware() (r *gin.Engine) {
	r = gin.Default()

	// cors設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			`http://localhost:80`,
			`http://localhost:3000`,
			`http://localhost`,
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	return
}

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "loginUserIdKey"
		id := session.GetSession(c, cookieKey)
		if id == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Next()
			log.Println("通ったわ！！！")
		}
	}
}

func LogoutCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "loginUserIdKey"
		id := session.GetSession(c, cookieKey)
		if id != nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
