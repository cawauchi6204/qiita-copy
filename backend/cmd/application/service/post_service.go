package service

import "github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"

func GetAllPosts() (posts []repository.Post) {
	posts = repository.FindPostsAll()
	return
}

func GetAllPostsByUserId(userId string) (posts []repository.Post) {
	posts = repository.FindPostsAllByUserId(userId)
	return
}
