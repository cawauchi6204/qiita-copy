package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/cawauchi6204/qiita-copy/cmd/common/crypto"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

func GetAllUsers() (users []repository.User) {
	users = repository.FindUsersAll()
	return
}

func GetUserById(userId int) (user repository.User) {
	user = repository.FindUserById(userId)
	return
}

func createUser(id, email, password string) {
	// TODO: これはEntityに持たせたい
	u := repository.User{
		ID:             id,
		Email:          email,
		Password:       password,
		Description:    "",
		HpUrl:          "",
		Location:       "",
		GithubAccount:  "",
		OrganizationId: "",
		IsDeleted:      0,
		CreatedAt:      time.Now(),
	}
	repository.CreateUser(u)
}

func Signup(id, email, password string) (*repository.User, error) {
	user := repository.User{}
	repository.FindUserByEmail(email)
	if user.ID != "" {
		err := errors.New("同一名のメールアドレスが既に登録されています。")
		fmt.Println(err)
		return nil, err
	}

	encryptPw, err := crypto.PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return nil, err
	}
	createUser(id, email, encryptPw)
	return &user, nil
}

func Login(email, password string) (*repository.User, error) {
	user := repository.FindUserByEmail(email)
	if user.ID == "" {
		err := errors.New("メールアドレスが一致するユーザーが存在しません")
		fmt.Println(err)
		return nil, err
	}

	err := crypto.CompareHashAndPassword(user.Password, password)
	if err != nil {
		fmt.Println("パスワードが一致しませんでした。:", err)
		return nil, err
	}

	return &user, nil
}
