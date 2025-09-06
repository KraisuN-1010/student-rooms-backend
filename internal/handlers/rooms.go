package handlers

import (
	"github.com/KraisuN-1010/student-rooms-backend/services"
	"gofr.dev/pkg/gofr"
)

type RoomHandler struct {
	service *services.RoomService
}

func NewRoomHandler(s *services.RoomService) *RoomHandler {
	return &RoomHandler{service: s}
}

func (h *RoomHandler) GetRooms(c *gofr.Context) (interface{}, error) {
	rooms, err := h.service.GetRooms(c.Context)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (h *RoomHandler) CreateRoom(c *gofr.Context) (interface{}, error) {
	req := &services.RoomRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	room, err := h.service.CreateRoom(c.Context, req)
	if err != nil {
		return nil, err
	}

	return room, nil
}
