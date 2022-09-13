package service

import "github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"

func CreateTag(id, imgUrl string) {
	tag := repository.Tag{
		ID:     id,
		ImgUrl: imgUrl,
	}
	repository.CreateTag(tag)
}
