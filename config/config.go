package config

import (
	"golang-otp/config/cache"
	"golang-otp/config/postgres"
	"golang-otp/thirdparty"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type (
	config struct {
	}

	Config interface {
		GoMail() thirdparty.EmailSmtp
		ServiceName() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *gorm.DB
		Redis() cache.Redis
	}
)

func (c *config) GoMail() thirdparty.EmailSmtp {
	return thirdparty.NewEmailSmtp()
}

func (c *config) Redis() cache.Redis {
	return cache.InitRedis()
}

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) ServicePort() int {
	v := os.Getenv("PORT")
	port, _ := strconv.Atoi(v)

	return port
}

func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}
