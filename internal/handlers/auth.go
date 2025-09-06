package handlers

import (
	"github.com/KraisuN-1010/student-rooms-backend/services"
	"gofr.dev/pkg/gofr"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) SignUp(c *gofr.Context) (interface{}, error) {
	req := &services.SignUpRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	resp, err := h.service.SignUp(c.Context, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *AuthHandler) Login(c *gofr.Context) (interface{}, error) {
	req := &services.LoginRequest{}
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	token, err := h.service.Login(c.Context, req)
	if err != nil {
		return nil, err
	}

	return map[string]string{"token": token}, nil
}
