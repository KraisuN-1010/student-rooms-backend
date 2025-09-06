package services

import (
	"context"
	"time"

	"github.com/KraisuN-1010/student-rooms-backend/db"
	"github.com/KraisuN-1010/student-rooms-backend/models"
	"github.com/google/uuid"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s *CommentService) GetCommentsByParent(ctx context.Context, parentID string) ([]models.Comment, error) {
	query := `SELECT id, parent_id, parent_type, room_id, content, file_url, file_name, is_solution, created_by, created_at, updated_at
	          FROM comments WHERE parent_id=$1`

	rows, err := db.DB.Query(ctx, query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(
			&c.ID, &c.ParentID, &c.ParentType, &c.RoomID, &c.Content, &c.FileURL,
			&c.FileName, &c.IsSolution, &c.CreatedBy, &c.CreatedAt, &c.UpdatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}

func (s *CommentService) CreateComment(ctx context.Context, req *CommentRequest) (*models.Comment, error) {
	comment := &models.Comment{
		ID:         uuid.New().String(),
		ParentID:   req.ParentID,
		ParentType: req.ParentType,
		RoomID:     req.RoomID,
		Content:    req.Content,
		FileURL:    req.FileURL,
		FileName:   req.FileName,
		IsSolution: req.IsSolution,
		CreatedBy:  "user-id", // TODO: Get from JWT token
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	query := `INSERT INTO comments 
	(id, parent_id, parent_type, room_id, content, file_url, file_name, is_solution, created_by, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := db.DB.Exec(ctx, query,
		comment.ID, comment.ParentID, comment.ParentType, comment.RoomID, comment.Content,
		comment.FileURL, comment.FileName, comment.IsSolution, comment.CreatedBy, comment.CreatedAt, comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
