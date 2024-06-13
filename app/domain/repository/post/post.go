package post

import (
	"github.com/krevetkou/test-ozon/app/models"
)

type PostService interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
}
