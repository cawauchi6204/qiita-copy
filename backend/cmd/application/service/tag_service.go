package service

import "github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"

func GetTagById(id string) (tag repository.Tag) {
	tag = repository.FindTagById(id)
	return
}

func CreateTag(id, imgUrl string) {
	tag := repository.Tag{
		ID:     id,
		ImgUrl: imgUrl,
	}
	repository.CreateTag(tag)
}
