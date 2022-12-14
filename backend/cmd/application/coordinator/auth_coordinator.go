package coordinator

import (
	"net/http"

	"github.com/cawauchi6204/qiita-copy/cmd/application/service"
	"github.com/cawauchi6204/qiita-copy/cmd/common/crypto"
	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/session"
	"github.com/cawauchi6204/qiita-copy/cmd/presentation/request"
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
		service.CreateUser(request.Id, request.Email, encryptedPassword)
		cookieKey := "loginUserIdKey"
		session.NewSession(c, cookieKey, request.Email)
		c.Redirect(http.StatusFound, "/")
	}
}

func Login(c *gin.Context) {
	var request request.EmailLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		user := service.GetUserByEmail(request.Email)
		err = crypto.CompareHashAndPassword(user.Password, request.Password)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			cookieKey := "loginUserIdKey"
			session.NewSession(c, cookieKey, request.Email)
			c.Redirect(http.StatusFound, "/")
		}
	}
}

func GetMyInfo(c *gin.Context) (user entity.User) {
	cookieKey := "loginUserIdKey"
	email := session.GetSession(c, cookieKey)
	user = service.GetUserByEmail(email)
	return
}

func Logout(c *gin.Context) {
	cookieKey := "loginUserIdKey"
	session.DeleteSession(c, cookieKey)
	c.Redirect(http.StatusFound, "/")
}
