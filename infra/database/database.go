package database

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	sing   sync.Once
	dbConn *gorm.DB
)

func GetConnection() *gorm.DB {

	sing.Do(func() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  false,
			},
		)

		dsn := os.Getenv(utils.DATABASE_URL)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			panic(err)
		}

		db.AutoMigrate(
			&models.UserModel{},
			&models.ChallengeModel{},
			&models.ChallengeCommentModel{},
			&models.ChallengeHintsModel{},
			&models.ChallengesCategoriesModel{},
			&models.ChallengeCategoriesUsersModel{},
		)

		dbConn = db
	})

	return dbConn
}
