package models

import "time"

type ContentType string

const (
	NoteContent         ContentType = "note"
	TopicContent        ContentType = "topic"
	AnnouncementContent ContentType = "announcement"
)

// Post represents the posts table (equivalent to notes)
type Post struct {
	ID          string      `json:"id"`
	RoomID      string      `json:"room_id"`
	Title       string      `json:"title"`
	Content     string      `json:"content,omitempty"`
	ContentType ContentType `json:"content_type"`
	FileURL     string      `json:"file_url,omitempty"`
	FileName    string      `json:"file_name,omitempty"`
	FileSize    int64       `json:"file_size,omitempty"`
	FileType    string      `json:"file_type,omitempty"`
	IsPinned    bool        `json:"is_pinned"`
	CreatedBy   string      `json:"created_by"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// Note is an alias for Post for backward compatibility
type Note = Post
