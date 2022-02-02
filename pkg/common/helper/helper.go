package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alisyahbana/efishery-test/pkg/common/key"
	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type CommonTokenPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type CommonTokenClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	jwt.StandardClaims
}

type ProfileResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
}

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func CreateCommonToken(payload CommonTokenPayload) (string, error) {
	expireToken := time.Now().AddDate(0, 1, 0).Unix()
	claims := CommonTokenClaims{
		Username: payload.Username,
		Email:    payload.Email,
		Address:  payload.Address,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key.GetConfig().SignatureJwt))
	if err != nil {
		return "", err
	}

	newToken := "Bearer " + tokenString

	return newToken, nil
}

func AuthorizeUser(r *http.Request) (*ProfileResponse, error) {
	Token := r.Header.Get("authorization")

	if len(Token) == 0 {
		return nil, nil
	}

	var body CommonTokenPayload

	newToken := strings.TrimPrefix(Token, "Bearer ")

	token, err := jwt.Parse(newToken, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key.GetConfig().SignatureJwt), nil
	})

	if token != nil && err == nil {
		claim := token.Claims.(jwt.MapClaims)
		if claim.VerifyExpiresAt(time.Now().Unix(), false) == false {
			return nil, fmt.Errorf("ACCESS_TOKEN_EXPIRED")
		}

		body = CommonTokenPayload{
			Username: claim["username"].(string),
			Email:    claim["email"].(string),
			Address:  claim["address"].(string),
		}

		return &ProfileResponse{
			Username: body.Username,
			Email:    body.Email,
			Address:  body.Address,
			Token:    Token,
		}, nil
	} else {
		return nil, fmt.Errorf("invalid access token")
	}
}

type ErrorResponse struct {
	Error string `json:"message"`
}

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorReturn(writer http.ResponseWriter, statusCode int, err error) {
	errorResponse := ErrorResponse{}
	writer.WriteHeader(statusCode)
	writer.Header().Set("Content-Type", "application/json")
	if err == nil {
		return
	}
	errorResponse.Error = err.Error()
	json.NewEncoder(writer).Encode(errorResponse)
	return
}

func SuccessReturn(writer http.ResponseWriter, response interface{}) {
	message := MessageResponse{
		Message: "Success",
		Data:    response,
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(message)
	return
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type CurrencyRate struct {
	Ratio float64 `json:"IDR_USD"`
}

func GetRatioUSD() (float64, error) {
	// get from cache
	cacheKey := "rate-usd-idr"
	var rate *CurrencyRate
	valueCache, found := GetCache(cacheKey)
	if found {
		byteData, _ := json.Marshal(valueCache)
		err := json.Unmarshal(byteData, &rate)
		if err != nil {
			return 0, err
		}
		fmt.Println("success get ratio from cache")
	} else {
		// save to cache
		apiKey := "a89dabc7c704a030831d"

		url := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=%s", apiKey)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			err = errors.New("Failed to get conversion rate")
			return 0, err
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &rate)
		if err != nil {
			log.Println(err)
			err = errors.New("Failed to parse conversion rate data")
			return 0, err
		} else if rate.Ratio == 0 {
			log.Println("Conversion Rate is 0!")
			err = errors.New("Failed to get conversion rate data")
			return 0, err
		}
		SaveCache(cacheKey, rate)
	}
	return rate.Ratio, nil
}

func SaveCache(key string, valude interface{}) bool {
	Cache.Set(key, valude, 1*time.Hour)
	return true
}

func GetCache(key string) (interface{}, bool) {
	var value interface{}
	dataCache, found := Cache.Get(key)
	if found {
		value = dataCache.(interface{})
	}
	return value, found
}
