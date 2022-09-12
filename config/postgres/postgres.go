package postgres

import (
	"github.com/rs/zerolog/log"
	"golang-otp/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitGorm() *gorm.DB {

	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}
	db.AutoMigrate(&model.User{})

	return db

}
