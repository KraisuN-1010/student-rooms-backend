package models

import "time"

type CommentParentType string

const (
	PostCommentParent    CommentParentType = "post"
	DoubtCommentParent   CommentParentType = "doubt"
	CommentReplyParent   CommentParentType = "comment"
)

type Comment struct {
	ID         string            `json:"id"`
	ParentID   string            `json:"parent_id"`
	ParentType CommentParentType `json:"parent_type"`
	RoomID     string            `json:"room_id"`
	Content    string            `json:"content"`
	FileURL    string            `json:"file_url,omitempty"`
	FileName   string            `json:"file_name,omitempty"`
	IsSolution bool              `json:"is_solution"`
	CreatedBy  string            `json:"created_by"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}
