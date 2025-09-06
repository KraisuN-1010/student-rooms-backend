package services

import (
	"context"
	"errors"
	"time"

	"github.com/KraisuN-1010/student-rooms-backend/db"
	"github.com/KraisuN-1010/student-rooms-backend/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignUp(ctx context.Context, req *SignUpRequest) (*AuthResponse, error) {
	// Check if user already exists
	existingUser, _ := s.getUserByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	user := &models.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Bio:       req.Bio,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users (id, name, email, password_hash, avatar_url, bio, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.DB.Exec(ctx, query, user.ID, user.Name, user.Email, hashedPwd, user.AvatarURL, user.Bio, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (string, error) {
	user, err := s.authenticateUser(ctx, req.Email, req.Password)
	if err != nil {
		return "", err
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) authenticateUser(ctx context.Context, email, password string) (*models.User, error) {
	user := &models.User{}
	var hashedPwd []byte

	query := `SELECT id, name, email, password_hash, avatar_url, bio FROM users WHERE email=$1`
	row := db.DB.QueryRow(ctx, query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &hashedPwd, &user.AvatarURL, &user.Bio)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(hashedPwd, []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *AuthService) getUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, avatar_url, bio FROM users WHERE email=$1`
	row := db.DB.QueryRow(ctx, query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.AvatarURL, &user.Bio)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (s *AuthService) generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key")) // TODO: Use environment variable
}

