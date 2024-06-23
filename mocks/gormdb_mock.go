package mocks_test

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewGormDbMock() *gorm.DB {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

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

func DownDatabase(db *gorm.DB) {
	d, _ := db.DB()

	d.Close()
}
