package comment

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Service - our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines comment structure
type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(id uint32) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(id uint32, newComment Comment) (Comment, error)
	DeleteComment(id uint32) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(id uint32) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, id); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(id uint32, newComment Comment) (Comment, error) {
	oldComment, err := s.GetComment(id)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&oldComment).Updates(newComment); result.Error != nil {
		return Comment{}, err
	}
	return oldComment, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(id uint32) error {
	if result := s.DB.Delete(&Comment{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - retrieves all comments from database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
