package services

import (
	"context"
	"time"

	"github.com/KraisuN-1010/student-rooms-backend/db"
	"github.com/KraisuN-1010/student-rooms-backend/models"
	"github.com/google/uuid"
)

type RoomService struct{}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (s *RoomService) GetRooms(ctx context.Context) ([]models.Room, error) {
	query := `SELECT id, name, description, is_private, invite_code, created_by, created_at, updated_at FROM rooms`

	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var r models.Room
		if err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.IsPrivate, &r.InviteCode, &r.CreatedBy, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}
	return rooms, rows.Err()
}

func (s *RoomService) CreateRoom(ctx context.Context, req *RoomRequest) (*models.Room, error) {
	room := &models.Room{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
		InviteCode:  generateInviteCode(),
		CreatedBy:   "user-id", // TODO: Get from JWT token
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO rooms (id, name, description, is_private, invite_code, created_by, created_at, updated_at)
	          VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := db.DB.Exec(ctx, query, room.ID, room.Name, room.Description, room.IsPrivate, room.InviteCode, room.CreatedBy, room.CreatedAt, room.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func generateInviteCode() string {
	return uuid.New().String()[:8]
}
