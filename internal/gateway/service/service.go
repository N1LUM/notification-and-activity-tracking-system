package service

import (
	"context"
	"notify-activity-tracking-system/internal/gateway/auth"
	"notify-activity-tracking-system/internal/gateway/dto/user_context"
	user_context2 "notify-activity-tracking-system/internal/gateway/models"
	"notify-activity-tracking-system/internal/gateway/repository"

	"github.com/google/uuid"
)

type User interface {
	CreateUser(input user_context.CreateUserInputDTO) (*user_context2.User, error)
	GetAllUsers() ([]user_context2.User, error)
	GetByID(id uuid.UUID) (*user_context2.User, error)
	GetUserByUsername(username string) (*user_context2.User, error)
	Update(id uuid.UUID, input user_context.UpdateUserInputDTO) (*user_context2.User, error)
	DeleteUser(id uuid.UUID, ctx context.Context) error
}

type Auth interface {
	Login(username, password string, ctx context.Context) (*user_context2.User, string, string, error)
	Refresh(userID string, token string, ctx context.Context) (string, string, error)
}

type Service struct {
	User
	Auth
}

func NewService(repos *repository.Repository, jwtManager *auth.JWTManager) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Auth: NewAuthService(repos.Auth, jwtManager),
	}
}
