package service

import "github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"

func GetAllPostsByUserId(userId int) (posts []repository.Post) {
	posts = repository.FindPostsAllByUserId(userId)
	return
}
