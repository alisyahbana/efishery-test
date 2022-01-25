package user

import (
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/service/user/data"
	"time"
)

type UserService struct {
	data data.MysqlUserData
}

func New() UserService {
	return UserService{
		data: data.MysqlUserData{},
	}
}

type RegisterPayload struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	GeneratedPassword string `json:"generated_password"`
}

func (s UserService) Register(payload RegisterPayload) (*RegisterResponse, error) {
	if payload.Username == "" {
		return nil, fmt.Errorf("user name required")
	}
	userExisting, err := s.data.GetUser(payload.Username)
	if err != nil {
		return nil, err
	}

	if userExisting != nil {
		return nil, fmt.Errorf("user already exist")
	} else {
		generatedPassword := helper.RandStringBytes(4)
		// create user
		_, err := s.data.CreateUser(data.User{
			Username:  payload.Username,
			Password:  generatedPassword,
			Phone:     payload.Phone,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		})
		if err != nil {
			return nil, err
		}

		return &RegisterResponse{GeneratedPassword: generatedPassword}, nil
	}
}
