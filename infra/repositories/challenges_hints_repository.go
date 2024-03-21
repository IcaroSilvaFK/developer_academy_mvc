package repositories

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type ChallengesHintsRepository struct {
	db *gorm.DB
}

type ChallengesHintsRepositoryInterface interface {
	GetByChallengeId(challengedId string) (*models.ChallengeHintsModel, error)
	GetById(id string) (*models.ChallengeHintsModel, error)
	Create(*models.ChallengeHintsModel) error
	Delete(id string) error
}

func NewChallengesHintsRepository(
	db *gorm.DB,
) ChallengesHintsRepositoryInterface {
	return &ChallengesHintsRepository{
		db,
	}
}

func (ch *ChallengesHintsRepository) GetByChallengeId(challengedId string) (*models.ChallengeHintsModel, error) {

	var r *models.ChallengeHintsModel

	err := ch.db.Model(&models.ChallengeHintsModel{}).Find(&r, "challenge_id = ?", challengedId).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}
func (ch *ChallengesHintsRepository) GetById(id string) (*models.ChallengeHintsModel, error) {
	var r *models.ChallengeHintsModel

	err := ch.db.Model(&models.ChallengeHintsModel{}).Find(&r, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ch *ChallengesHintsRepository) Create(i *models.ChallengeHintsModel) error {

	err := ch.db.Model(&models.ChallengeHintsModel{}).Create(i).Error

	return err
}

func (ch *ChallengesHintsRepository) Delete(id string) error {

	err := ch.db.Delete(&models.ChallengeHintsModel{}, id).Error

	return err
}
