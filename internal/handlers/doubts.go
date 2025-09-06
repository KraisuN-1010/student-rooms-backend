package handlers

import (
	"errors"

	"github.com/KraisuN-1010/student-rooms-backend/services"
	"gofr.dev/pkg/gofr"
)

type DoubtHandler struct {
	service *services.DoubtService
}

func NewDoubtHandler(s *services.DoubtService) *DoubtHandler {
	return &DoubtHandler{service: s}
}

func (h *DoubtHandler) GetDoubtsByRoom(c *gofr.Context) (interface{}, error) {
	roomID := c.PathParam("roomId")
	if roomID == "" {
		return nil, errors.New("room_id is required")
	}

	doubts, err := h.service.GetDoubtsByRoom(c.Context, roomID)
	if err != nil {
		return nil, err
	}
	return doubts, nil
}

func (h *DoubtHandler) CreateDoubt(c *gofr.Context) (interface{}, error) {
	req := &services.DoubtRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	doubt, err := h.service.CreateDoubt(c.Context, req)
	if err != nil {
		return nil, err
	}

	return doubt, nil
}
