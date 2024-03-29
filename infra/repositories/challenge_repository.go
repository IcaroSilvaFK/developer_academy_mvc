package repositories

import (
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChallengeRepository struct {
	db *gorm.DB
}

type ChallengeRepositoryInterface interface {
	GetAll(page *int) ([]*models.ChallengeModel, error)
	GetById(id string) (*models.ChallengeModel, error)
	GetByUserId(id string) ([]*models.ChallengeModel, error)
	Create(*models.ChallengeModel) error
	CountChallenges() (int, error)
	Delete(id string) error
}

func NewChallengeRepository(
	db *gorm.DB,
) ChallengeRepositoryInterface {
	return &ChallengeRepository{
		db,
	}
}

// TODO add pagination
func (c *ChallengeRepository) GetAll(_ *int) ([]*models.ChallengeModel, error) {

	//if page == nil {
	//	*page = 1
	//}

	//offset := (*page - 1) * 10

	var r []*models.ChallengeModel

	//TODO implment pagination method
	tx := c.db.Model(&models.ChallengeModel{}).Find(&r)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return r, nil
}

func (c *ChallengeRepository) GetById(id string) (*models.ChallengeModel, error) {

	var r *models.ChallengeModel

	err := c.db.Table("challenges").Where("id = ?", id).Preload(clause.Associations).Find(&r).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *ChallengeRepository) GetByUserId(id string) ([]*models.ChallengeModel, error) {

	var result []*models.ChallengeModel

	err := c.db.Table("challenges").Find(&result, "user_id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ChallengeRepository) Create(cm *models.ChallengeModel) error {

	result := c.db.Create(cm)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func (c *ChallengeRepository) CountChallenges() (int, error) {

	var count int64

	if err := c.db.Table("challenges").Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (c *ChallengeRepository) Delete(id string) error {
	r := c.db.Table("challenges").Where("id = ?", id).Delete(&models.ChallengeModel{
		ID: id,
	})

	if r.Error != nil {
		return r.Error
	}

	return nil
}
