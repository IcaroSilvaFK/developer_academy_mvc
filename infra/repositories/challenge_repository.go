package repositories

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChallengeRepository struct {
	db *gorm.DB
}

type ChallengeRepositoryInterface interface {
	GetAll(ctx context.Context, page *int) ([]*models.ChallengeModel, error)
	GetById(ctx context.Context, id string) (*models.ChallengeModel, error)
	GetByUserId(ctx context.Context, id string) ([]*models.ChallengeModel, error)
	Create(ctx context.Context, u *models.ChallengeModel) error
	CountChallenges(ctx context.Context) (int, error)
	Delete(ctx context.Context, id string) error
}

func NewChallengeRepository(
	db *gorm.DB,
) ChallengeRepositoryInterface {
	return &ChallengeRepository{
		db,
	}
}

// TODO add pagination
func (c *ChallengeRepository) GetAll(ctx context.Context, _ *int) ([]*models.ChallengeModel, error) {

	//if page == nil {
	//	*page = 1
	//}

	//offset := (*page - 1) * 10

	var r []*models.ChallengeModel

	//TODO implment pagination method
	tx := c.db.WithContext(ctx).Preload(clause.Associations).Model(&models.ChallengeModel{}).Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: true}).Find(&r)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return r, nil
}

func (c *ChallengeRepository) GetById(ctx context.Context, id string) (*models.ChallengeModel, error) {

	var r *models.ChallengeModel

	err := c.db.WithContext(ctx).Table("challenges").Where("id = ?", id).Preload(clause.Associations).Find(&r).Error

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *ChallengeRepository) GetByUserId(ctx context.Context, id string) ([]*models.ChallengeModel, error) {

	var result []*models.ChallengeModel

	err := c.db.WithContext(ctx).Table("challenges").Find(&result, "user_id = ?", id).Preload(clause.Associations).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ChallengeRepository) Create(ctx context.Context, cm *models.ChallengeModel) error {

	result := c.db.WithContext(ctx).Create(cm)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func (c *ChallengeRepository) CountChallenges(ctx context.Context) (int, error) {

	var count int64

	if err := c.db.WithContext(ctx).Table("challenges").Where("deleted_at IS NULL").Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (c *ChallengeRepository) Delete(ctx context.Context, id string) error {
	r := c.db.WithContext(ctx).Table("challenges").Where("id = ?", id).Delete(&models.ChallengeModel{
		ID: id,
	})

	if r.Error != nil {
		return r.Error
	}

	return nil
}
