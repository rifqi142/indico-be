package services

import (
	"fmt"
	"time"

	"github.com/rifqi142/indico-be/internal/dto"
	"github.com/rifqi142/indico-be/internal/utils"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	jwtSecret     string
	jwtExpiration time.Duration
}

func NewAuthService(jwtSecret string, jwtExpiration time.Duration) AuthService {
	return &authService{
		jwtSecret:     jwtSecret,
		jwtExpiration: jwtExpiration,
	}
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {	
	token, err := utils.GenerateToken(req.Username, s.jwtSecret, s.jwtExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token: token,
		Type:  "Bearer",
	}, nil
}
