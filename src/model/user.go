package model

import (
	"context"
	"golang-otp/helper/token"
	"golang-otp/src/request"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase interface {
	RequestOtp(ctx context.Context, request request.GetOtpRequest) error
	CreateUser(ctx context.Context, request request.CreateUserRequest) (*User, error)
	LoginWithOTP(ctx context.Context, request request.LoginWithOTPRequest) (*token.NewTokenResponse, error)
}
