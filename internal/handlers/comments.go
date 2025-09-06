package handlers

import (
	"errors"

	"github.com/KraisuN-1010/student-rooms-backend/services"
	"gofr.dev/pkg/gofr"
)

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{service: s}
}

func (h *CommentHandler) GetCommentsByParent(c *gofr.Context) (interface{}, error) {
	parentID := c.PathParam("parentId")
	if parentID == "" {
		return nil, errors.New("parent_id is required")
	}

	comments, err := h.service.GetCommentsByParent(c.Context, parentID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (h *CommentHandler) CreateComment(c *gofr.Context) (interface{}, error) {
	req := &services.CommentRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	comment, err := h.service.CreateComment(c.Context, req)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
