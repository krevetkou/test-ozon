package persistence

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krevetkou/test-ozon/app/domain/repository/post"
	"github.com/krevetkou/test-ozon/app/models"
	"strings"
)

type postService struct {
	db *gorm.DB
}

func NewPost(db *gorm.DB) *postService {
	return &postService{
		db,
	}
}

var _ post.PostService = &postService{}

func (s *postService) CreatePost(post *models.Post) (*models.Post, error) {
	err := s.db.Create(&post).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("post title already taken")
		}
		return nil, err
	}

	return post, nil

}

func (s *postService) GetPostByID(id string) (*models.Post, error) {
	p := &models.Post{}

	err := s.db.Where("id = ?", id).Take(&p).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *postService) GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post

	err := s.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *postService) UpdatePost(post *models.Post) (*models.Post, error) {
	fmt.Print(post)
	err := s.db.Save(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}
