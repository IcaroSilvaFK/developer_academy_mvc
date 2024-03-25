package database

import (
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {

	dsn := os.Getenv(utils.DATABASE_URL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.UserModel{},
		&models.ChallengeModel{},
		&models.ChallengeCommentModel{},
		&models.ChallengeHintsModel{},
	)

	return db
}
