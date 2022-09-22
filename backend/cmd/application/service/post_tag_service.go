package service

import (
	"github.com/cawauchi6204/qiita-copy/cmd/entity"
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

func GetTagsByPostId(postId string) (postTags []entity.PostTag) {
	postTags = repository.FindTagsByPostId(postId)
	return
}

func GetPostIdsByTagId(tagId string) (postIds []string) {
	postIds = repository.FindPostsIdsByTagId(tagId)
	return
}

func CreatePostTag(postId, tagId string) entity.PostTag {
	postTag := entity.PostTag{
		PostId: postId,
		TagId:  tagId,
	}
	return repository.CreatePostTag(postTag)
}
