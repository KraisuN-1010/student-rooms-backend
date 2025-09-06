package services

import (
	"context"
	"time"

	"github.com/KraisuN-1010/student-rooms-backend/db"
	"github.com/KraisuN-1010/student-rooms-backend/models"
	"github.com/google/uuid"
)

type DoubtService struct{}

func NewDoubtService() *DoubtService {
	return &DoubtService{}
}

func (s *DoubtService) GetDoubtsByRoom(ctx context.Context, roomID string) ([]models.Doubt, error) {
	query := `SELECT id, room_id, question_text, description, status, is_urgent, created_by, created_at, updated_at
	          FROM doubts WHERE room_id=$1`

	rows, err := db.DB.Query(ctx, query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doubts []models.Doubt
	for rows.Next() {
		var d models.Doubt
		if err := rows.Scan(&d.ID, &d.RoomID, &d.QuestionText, &d.Description, &d.Status, &d.IsUrgent, &d.CreatedBy, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		doubts = append(doubts, d)
	}
	return doubts, rows.Err()
}

func (s *DoubtService) CreateDoubt(ctx context.Context, req *DoubtRequest) (*models.Doubt, error) {
	doubt := &models.Doubt{
		ID:           uuid.New().String(),
		RoomID:       req.RoomID,
		QuestionText: req.QuestionText,
		Description:  req.Description,
		Status:       req.Status,
		IsUrgent:     req.IsUrgent,
		CreatedBy:    "user-id", // TODO: Get from JWT token
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO doubts (id, room_id, question_text, description, status, is_urgent, created_by, created_at, updated_at)
	          VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := db.DB.Exec(ctx, query, doubt.ID, doubt.RoomID, doubt.QuestionText, doubt.Description, doubt.Status, doubt.IsUrgent, doubt.CreatedBy, doubt.CreatedAt, doubt.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return doubt, nil
}
