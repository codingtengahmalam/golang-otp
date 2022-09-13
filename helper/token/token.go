package token

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type (
	NewTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiredAt   int64  `json:"expired_at"`
	}

	NewTokenRequest struct {
		UserID    int    `json:"user_id"`
		UserEmail string `json:"user_email"`
	}

	customClaims struct {
		UserID    int    `json:"user_id"`
		UserEmail string `json:"user_email"`
		jwt.StandardClaims
	}

	NewTokenData struct {
		UserID    int    `json:"user_id"`
		UserEmail string `json:"user_email"`
	}
)

const DurationShort = time.Minute * 5
const DurationLong = time.Minute * 180
const TypeLongSecretKey = 1
const TypeShortSecretKey = 2

var LongSecretKey = os.Getenv("LONG_TOKEN_SECRET")
var shortSecretKey = os.Getenv("SHORT_TOKEN_SECRET")

func NewCustomToken(request NewTokenRequest, duration time.Duration) (*NewTokenResponse, error) {
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

	return &NewTokenResponse{
		AccessToken: tokenString,
		TokenType:   "bearer",
		ExpiredAt:   claims.ExpiresAt,
	}, nil
}

func ExtractToken(token string, tokenType int) (*NewTokenData, error) {
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

		extract := new(NewTokenData)
		if err := json.Unmarshal(jsonBody, extract); err != nil {
			return nil, err
		}

		return extract, nil
	}

	return nil, fmt.Errorf("token not valid")
}
