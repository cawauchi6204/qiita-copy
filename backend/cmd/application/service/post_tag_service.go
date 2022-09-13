package service

import (
	"github.com/cawauchi6204/qiita-copy/cmd/infrastructure/repository"
)

func CreatePostTag(postId, tagId string) repository.PostTag {
	postTag := repository.PostTag{
		PostId: postId,
		TagId:  tagId,
	}
	return repository.CreatePostTag(postTag)
}
