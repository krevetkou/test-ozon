package interfaces

import (
	"github.com/krevetkou/test-ozon/app/domain/repository/comment"
	"github.com/krevetkou/test-ozon/app/domain/repository/post"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostService    post.PostService
	CommentService comment.CommentService
}
