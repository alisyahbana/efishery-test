package user

import (
	"errors"
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/common/key"
	"github.com/alisyahbana/efishery-test/pkg/service/user/data"
	"github.com/dgrijalva/jwt-go"
	"log"
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

type LoginPayload struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (s UserService) Login(payload LoginPayload) (*string, error) {
	userData, err := s.data.GetUserByPhoneAndPassword(payload.Phone, payload.Password)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, fmt.Errorf("user not found")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  userData.Username,
		"phone": userData.Phone,
		//"role":      userData.Role,
		"timestamp": userData.CreatedAt,
	})

	tokenString, err := token.SignedString([]byte(key.GetConfig().SignatureJwt))
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to sign JWT token")
		return nil, err
	}

	return &tokenString, nil
}
