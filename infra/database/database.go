package database

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=developer port=5432 sslmode=disable"

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
