package handlers

import (
	"errors"

	"github.com/KraisuN-1010/student-rooms-backend/services"
	"gofr.dev/pkg/gofr"
)

type NoteHandler struct {
	service *services.NoteService
}

func NewNoteHandler(s *services.NoteService) *NoteHandler {
	return &NoteHandler{service: s}
}

func (h *NoteHandler) GetNotesByRoom(c *gofr.Context) (interface{}, error) {
	roomID := c.PathParam("roomId")
	if roomID == "" {
		return nil, errors.New("room_id is required")
	}

	notes, err := h.service.GetNotesByRoom(c.Context, roomID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (h *NoteHandler) CreateNote(c *gofr.Context) (interface{}, error) {
	req := &services.NoteRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	note, err := h.service.CreateNote(c.Context, req)
	if err != nil {
		return nil, err
	}

	return note, nil
}
