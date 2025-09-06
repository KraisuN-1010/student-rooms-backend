package services

import "github.com/KraisuN-1010/student-rooms-backend/models"

// Room request/response types
type RoomRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
}

type RoomResponse struct {
	*models.Room
}

// Auth request/response types
type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Bio      string `json:"bio"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

// Post request/response types (posts table)
type PostRequest struct {
	RoomID      string                `json:"room_id" validate:"required"`
	Title       string                `json:"title" validate:"required"`
	Content     string                `json:"content"`
	ContentType models.ContentType    `json:"content_type" validate:"required"`
	FileURL     string                `json:"file_url"`
	FileName    string                `json:"file_name"`
	FileSize    int64                 `json:"file_size"`
	FileType    string                `json:"file_type"`
	IsPinned    bool                  `json:"is_pinned"`
}

type PostResponse struct {
	*models.Post
}

// Note request/response types (alias for backward compatibility)
type NoteRequest = PostRequest
type NoteResponse = PostResponse

// Comment request/response types
type CommentRequest struct {
	ParentID   string                     `json:"parent_id" validate:"required"`
	ParentType models.CommentParentType   `json:"parent_type" validate:"required"`
	RoomID     string                     `json:"room_id" validate:"required"`
	Content    string                     `json:"content" validate:"required"`
	FileURL    string                     `json:"file_url"`
	FileName   string                     `json:"file_name"`
	IsSolution bool                       `json:"is_solution"`
}

type CommentResponse struct {
	*models.Comment
}

// Doubt request/response types
type DoubtRequest struct {
	RoomID       string              `json:"room_id" validate:"required"`
	QuestionText string              `json:"question_text" validate:"required"`
	Description  string              `json:"description"`
	Status       models.DoubtStatus  `json:"status"`
	IsUrgent     bool                `json:"is_urgent"`
}

type DoubtResponse struct {
	*models.Doubt
}
