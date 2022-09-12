package helper

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type (
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiredAt   int64  `json:"expired_at"`
	}

	TokenRequest struct {
		UserID    string `json:"user_id"`
		UserEmail string `json:"user_email"`
	}

	customClaims struct {
		UserID    string `json:"user_id"`
		UserEmail string `json:"user_email"`
		jwt.StandardClaims
	}

	TokenData struct {
		UserID    string `json:"user_id"`
		UserEmail string `json:"user_email"`
	}
)

const DurationShort = time.Minute * 5
const DurationLong = time.Minute * 180
const TypeLongSecretKey = 1
const TypeShortSecretKey = 2

var LongSecretKey = os.Getenv("LONG_TOKEN_SECRET")
var shortSecretKey = os.Getenv("SHORT_TOKEN_SECRET")

func NewCustomToken(request TokenRequest, duration time.Duration) (*TokenResponse, error) {
	var secretKey string

	claims := &customClaims{
		request.UserID,
		request.UserEmail,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	switch duration {
	case DurationLong:
		secretKey = LongSecretKey
	case DurationShort:
		secretKey = shortSecretKey
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken: tokenString,
		TokenType:   "bearer",
		ExpiredAt:   claims.ExpiresAt,
	}, nil
}

func ExtractToken(token string, tokenType int) (*TokenData, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		switch tokenType {
		case TypeLongSecretKey:
			return LongSecretKey, nil
		case TypeShortSecretKey:
			return shortSecretKey, nil
		}

		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		jsonBody, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}

		extract := new(TokenData)
		if err := json.Unmarshal(jsonBody, extract); err != nil {
			return nil, err
		}

		return extract, nil
	}

	return nil, fmt.Errorf("token not valid")
}
