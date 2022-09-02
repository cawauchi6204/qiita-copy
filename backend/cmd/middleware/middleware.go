package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() (r *gin.Engine) {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			`http://localhost:3000`,
			`http://localhost:80`,
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
