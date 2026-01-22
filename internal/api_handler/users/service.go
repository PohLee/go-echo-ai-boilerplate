package users

import (
	"errors"
	"time"

	"github.com/PohLee/go-echo-ai-boilerplate/internal/model"
	"github.com/PohLee/go-echo-ai-boilerplate/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req *RegisterRequest) (*UserResponse, error)
	Login(req *LoginRequest) (*LoginResponse, error)
}

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

func UserService(repo Repository, jwtService auth.JWTService) Service {
	return &service{repo: repo, jwtService: jwtService}
}

func (s *service) Register(req *RegisterRequest) (*UserResponse, error) {
	// Check if email exists
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email already registered") // Should use domain error
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Role:     "user",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *service) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User: UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}
