package models

import "time"

type DoubtStatus string

const (
	OpenDoubt     DoubtStatus = "open"
	AnsweredDoubt DoubtStatus = "answered"
	ClosedDoubt   DoubtStatus = "closed"
)

type Doubt struct {
	ID           string      `json:"id"`
	RoomID       string      `json:"room_id"`
	QuestionText string      `json:"question_text"`
	Description  string      `json:"description,omitempty"`
	Status       DoubtStatus `json:"status"`
	IsUrgent     bool        `json:"is_urgent"`
	CreatedBy    string      `json:"created_by"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
