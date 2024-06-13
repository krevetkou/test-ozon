package comment

import (
	"github.com/krevetkou/test-ozon/app/models"
)

type CommentService interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	IsCommentExists(id string) (bool, error)
	GetCommentByID(id string) (*models.Comment, error)
	UpdateComment(comment *models.Comment) (*models.Comment, error)
}
