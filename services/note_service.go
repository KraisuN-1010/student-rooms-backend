package services

import (
	"context"
	"time"

	"github.com/KraisuN-1010/student-rooms-backend/db"
	"github.com/KraisuN-1010/student-rooms-backend/models"
	"github.com/google/uuid"
)

type NoteService struct{}

func NewNoteService() *NoteService {
	return &NoteService{}
}

func (s *NoteService) GetNotesByRoom(ctx context.Context, roomID string) ([]models.Note, error) {
	query := `SELECT id, room_id, title, content, content_type, file_url, file_name, file_size, file_type, is_pinned, created_by, created_at, updated_at
	          FROM posts WHERE room_id=$1`

	rows, err := db.DB.Query(ctx, query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.RoomID, &n.Title, &n.Content, &n.ContentType, &n.FileURL, &n.FileName, &n.FileSize, &n.FileType, &n.IsPinned, &n.CreatedBy, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, rows.Err()
}

func (s *NoteService) CreateNote(ctx context.Context, req *NoteRequest) (*models.Note, error) {
	note := &models.Post{
		ID:          uuid.New().String(),
		RoomID:      req.RoomID,
		Title:       req.Title,
		Content:     req.Content,
		ContentType: req.ContentType,
		FileURL:     req.FileURL,
		FileName:    req.FileName,
		FileSize:    req.FileSize,
		FileType:    req.FileType,
		IsPinned:    req.IsPinned,
		CreatedBy:   "user-id", // TODO: Get from JWT token
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO posts (id, room_id, title, content, content_type, file_url, file_name, file_size, file_type, is_pinned, created_by, created_at, updated_at)
	          VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err := db.DB.Exec(ctx, query, note.ID, note.RoomID, note.Title, note.Content, note.ContentType, note.FileURL, note.FileName, note.FileSize, note.FileType, note.IsPinned, note.CreatedBy, note.CreatedAt, note.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return note, nil
}

// CreateNoteWithRealtime creates a note and broadcasts it to the room
func (s *NoteService) CreateNoteWithRealtime(ctx context.Context, req *NoteRequest, realtimeHandler interface{}) (*models.Note, error) {
	note, err := s.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	// Broadcast the new post to the room
	if rh, ok := realtimeHandler.(interface {
		BroadcastPostCreated(roomID string, post interface{})
	}); ok {
		rh.BroadcastPostCreated(req.RoomID, note)
	}

	return note, nil
}
