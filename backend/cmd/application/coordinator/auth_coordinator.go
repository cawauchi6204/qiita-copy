package coordinator

import (
	"encoding/json"
	"net/http"

	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/common/crypto"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var request request.EmailSignUpRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		encryptedPassword, err := crypto.PasswordEncrypt(request.Password)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		user := service.CreateUser(request.Name, request.Email, encryptedPassword)
		// TODO: 作成されたユーザー情報をreturnしたい
		session := sessions.Default(c)
		// セッションに格納する為にユーザ情報をJson化
		loginUser, err := json.Marshal(user)
		if err == nil {
			session.Set("loginUser", string(loginUser))
			session.Save()
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}
}

func Login(c *gin.Context) {
	var request request.EmailLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		// メールアドレスでDBからユーザ取得（詳細は割愛）
		user := service.GetUserByEmail(request.Email)
		// ハッシュ値でのパスワード比較
		err = crypto.CompareHashAndPassword(user.Password, request.Password)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			session := sessions.Default(c)
			// セッションに格納する為にユーザ情報をJson化
			loginUser, err := json.Marshal(user)
			if err == nil {
				session.Set("loginUser", string(loginUser))
				session.Save()
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusInternalServerError)
			}
		}
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.String(http.StatusOK, "ログアウトしました")
}
