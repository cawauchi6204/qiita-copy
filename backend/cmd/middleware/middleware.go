package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func Middleware() (r *gin.Engine) {
	r = gin.Default()
	// session設定
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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
		session := sessions.Default(c)
		loginUserJson, err := dproxy.New(session.Get("password")).String()
		if err != nil {
			c.Status(402) // 下の分岐で返されるエラーステータスが400のため、一時的にわかりやすいように402にしている
			c.Abort()
		} else {
			var loginInfo repository.User
			// Json文字列のアンマーシャル
			err := json.Unmarshal([]byte(loginUserJson), &loginInfo)
			if err != nil {
				c.Status(http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
