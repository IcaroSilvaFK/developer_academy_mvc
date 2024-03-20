package services

import (
	"log"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type ChallengeService struct {
	repo repositories.ChallengeRepositoryInterface
}

type ChallengeServiceInterface interface {
	Create(title, description, embedUrl, userId string) error
	FindAll(page *int) ([]*models.ChallengeModel, error)
	FindById(id string) (*models.ChallengeModel, error)
	Delete(id string) error
}

func NewChallengeService(
	repo repositories.ChallengeRepositoryInterface,
) ChallengeServiceInterface {

	return &ChallengeService{
		repo,
	}
}

func (c *ChallengeService) Create(title, description, embedUrl, userId string) error {

	cm := models.NewChallengeModel(title, description, embedUrl, userId)
	log.Println("aq")
	log.Println(cm)

	return c.repo.Create(cm)
}

func (c *ChallengeService) FindAll(page *int) ([]*models.ChallengeModel, error) {

	return c.repo.GetAll(page)
}

func (c *ChallengeService) FindById(id string) (*models.ChallengeModel, error) {
	return c.repo.GetById(id)
}

func (c *ChallengeService) Delete(id string) error {
	return c.repo.Delete(id)
}
