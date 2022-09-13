package repository

import (
	"context"
	"golang-otp/config"
	"golang-otp/src/model"
)

type userRepository struct {
	Cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository {
	return &userRepository{Cfg: cfg}
}

func (u *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := u.Cfg.Database().WithContext(ctx).
		Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
