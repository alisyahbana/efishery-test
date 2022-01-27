package user

import (
	"errors"
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/helper"
	"github.com/alisyahbana/efishery-test/pkg/common/key"
	"github.com/alisyahbana/efishery-test/pkg/service/user/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/karirdotcom/qframework/pkg/common/qerror"
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
			Role:      payload.Role,
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
		"username":  userData.Username,
		"phone":     userData.Phone,
		"role":      userData.Role,
		"timestamp": time.Now(),
	})

	tokenString, err := token.SignedString([]byte(key.GetConfig().SignatureJwt))
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to sign JWT token")
		return nil, err
	}

	return &tokenString, nil
}

func (s UserService) ValidateToken(tokenAuhorization string) (*data.UserClaims, error) {
	token, err := jwt.Parse(tokenAuhorization, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key.GetConfig().SignatureJwt), nil
	})
	if err != nil {
		return nil, err
	}

	if token != nil && err == nil {
		claim := token.Claims.(jwt.MapClaims)
		if claim.VerifyExpiresAt(time.Now().Unix(), false) == false {
			return nil, qerror.NewValidationError("ACCESS_TOKEN_EXPIRED")
		}

		resp := data.UserClaims{
			Username:  claim["username"].(string),
			Phone:     claim["phone"].(string),
			Role:      claim["role"].(string),
			Timestamp: claim["timestamp"].(string),
		}
		return &resp, nil
	}

	return nil, nil
}
