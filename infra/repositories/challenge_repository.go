package repositories

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type ChallengeRepository struct {
	db *gorm.DB
}

type ChallengeRepositoryInterface interface {
	GetAll(page *int) ([]*models.ChallengeModel, error)
	GetById(id string) (*models.ChallengeModel, error)
	Create(title, description, embedUrl, userId string) error
	Delete(id string) error
}

func NewChallengeRepository(
	db *gorm.DB,
) ChallengeRepositoryInterface {
	return &ChallengeRepository{
		db,
	}
}

func (c *ChallengeRepository) GetAll(page *int) ([]*models.ChallengeModel, error) {

	if page == nil {
		*page = 1
	}

	offset := (*page - 1) * 10

	var r []*models.ChallengeModel

	tx := c.db.Model(&models.ChallengeModel{}).Offset(offset).Limit(utils.PAGE_SIZE).Find(&r)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return r, nil
}

func (c *ChallengeRepository) GetBydId(id string) (*models.ChallengeModel, error) {

	var r *models.ChallengeModel

	err := c.db.Model(&models.ChallengeModel{}).Select(&r, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *ChallengeRepository) Create(title, description, embedUrl, userId string) error {
	result := c.db.Create(&models.ChallengeModel{
		Title:       title,
		Description: description,
		EmbedUrl:    embedUrl,
		UserId:      userId,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *ChallengeRepository) Delete(id string) error {

	r := c.db.Model(&models.ChallengeModel{}).Delete(&id)

	if r.Error != nil {
		return r.Error
	}

	return nil
}
