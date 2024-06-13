package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/krevetkou/test-ozon/app/domain/repository/comment"
	"github.com/krevetkou/test-ozon/app/models"
)

type commentService struct {
	db *gorm.DB
}

func NewAnswer(db *gorm.DB) *commentService {
	return &commentService{
		db,
	}
}

// We implement the interface defined in the domain
var _ comment.CommentService = &commentService{}

func (s *commentService) CreateComment(comment *models.Comment) (*models.Comment, error) {
	err := s.db.Create(&comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) IsCommentExists(id string) (bool, error) {
	c := &models.Comment{}
	err := s.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *commentService) GetCommentByID(id string) (*models.Comment, error) {
	c := &models.Comment{}

	err := s.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *commentService) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	err := s.db.Save(&comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}
