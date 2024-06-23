package repositories

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type ChallengesHintsRepository struct {
	db *gorm.DB
}

type ChallengesHintsRepositoryInterface interface {
	GetByChallengeId(ctx context.Context, challengedId string) (*models.ChallengeHintsModel, error)
	GetById(ctx context.Context, id string) (*models.ChallengeHintsModel, error)
	Create(ctx context.Context, ch *models.ChallengeHintsModel) error
	Delete(ctx context.Context, id string) error
}

func NewChallengesHintsRepository(
	db *gorm.DB,
) ChallengesHintsRepositoryInterface {
	return &ChallengesHintsRepository{
		db,
	}
}

func (ch *ChallengesHintsRepository) GetByChallengeId(ctx context.Context, challengedId string) (*models.ChallengeHintsModel, error) {

	var r *models.ChallengeHintsModel

	err := ch.db.WithContext(ctx).Model(&models.ChallengeHintsModel{}).Find(&r, "challenge_id = ?", challengedId).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}
func (ch *ChallengesHintsRepository) GetById(ctx context.Context, id string) (*models.ChallengeHintsModel, error) {
	var r *models.ChallengeHintsModel

	err := ch.db.WithContext(ctx).Model(&models.ChallengeHintsModel{}).Find(&r, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ch *ChallengesHintsRepository) Create(ctx context.Context, i *models.ChallengeHintsModel) error {

	err := ch.db.WithContext(ctx).Model(&models.ChallengeHintsModel{}).Create(i).Error

	return err
}

func (ch *ChallengesHintsRepository) Delete(ctx context.Context, id string) error {

	err := ch.db.WithContext(ctx).Delete(&models.ChallengeHintsModel{}, id).Error

	return err
}
