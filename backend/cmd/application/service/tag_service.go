package service

import (
	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

func GetTagById(id string) (tag entity.Tag) {
	tag = repository.FindTagById(id)
	return
}

func CreateTag(id, imgUrl string) {
	tag := entity.Tag{
		ID:     id,
		ImgUrl: imgUrl,
	}
	repository.CreateTag(tag)
}
