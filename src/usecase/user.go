package usecase

import (
	"context"
	"golang-otp/config"
	"golang-otp/helper"
	"golang-otp/src/model"
	"golang-otp/src/request"
	"golang-otp/src/response"
	"golang-otp/thirdparty"
)

type userUsecase struct {
	cfg      config.Config
	userRepo model.UserRepository
}

func NewUserUsecase(cfg config.Config, user model.UserRepository) model.UserUsecase {
	return &userUsecase{
		cfg:      cfg,
		userRepo: user,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, request request.CreateUserRequest) (*model.User, error) {
	hashPassword, _ := helper.NewPassword(request.Password)

	user, err := u.userRepo.Create(ctx, &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword,
	})
	//send email
	go u.cfg.GoMail().SendEmail(ctx, thirdparty.SendEmailRequest{
		To:      user.Email,
		Body:    "Berikut ini adalah OTP Anda " + request.Password,
		Subject: "Selamat Datang di Senja Labs",
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) RequestOtp(ctx context.Context, request request.GetOtpRequest) error {
	// get email
	user, err := u.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	otp := helper.RandNumeric(5)
	cacheKey := "otp:" + user.Email
	err = u.cfg.Redis().Set(ctx, cacheKey, otp)
	if err != nil {
		return err
	}

	//send email
	go u.cfg.GoMail().SendEmail(ctx, thirdparty.SendEmailRequest{
		To:      user.Email,
		Body:    "Berikut ini adalah OTP Anda " + otp,
		Subject: "OTP dari Senja Labs",
	})

	return nil

}

func (u *userUsecase) LoginWithOTP(ctx context.Context, request request.LoginWithOTPRequest) (*response.LoginWithOtpResponse, error) {
	//TODO implement me
	panic("implement me")
}
