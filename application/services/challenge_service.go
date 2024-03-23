package services

import (
	"errors"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type ChallengeService struct {
	repo      repositories.ChallengeRepositoryInterface
	hinrepo   ChallengeHintServiceInterface
	iaservice AIServiceInterface
}

type ChallengeServiceInterface interface {
	Create(title, description, embedUrl, userId string) error
	FindAll(page *int) ([]*models.ChallengeModel, error)
	FindById(id string) (*models.ChallengeModel, error)
	CountChallenges() (int, error)
	Delete(id string) error
}

func NewChallengeService(
	repo repositories.ChallengeRepositoryInterface,
	hinrepo ChallengeHintServiceInterface,
	iaservice AIServiceInterface,
) ChallengeServiceInterface {

	return &ChallengeService{
		repo, hinrepo, iaservice,
	}
}

func (c *ChallengeService) Create(title, description, embedUrl, userId string) error {

	if c.iaservice.VerifyIfIsValidChallenge(title) {
		return errors.New("ERROR")
	}
	cm := models.NewChallengeModel(title, description, embedUrl, userId)

	err := c.repo.Create(cm)

	if err != nil {
		return err
	}

	hint, err := c.iaservice.GenerateHintFromChallenge(title)

	if err != nil {
		return nil
	}

	err = c.hinrepo.Create(cm.ID, hint)

	if err != nil {
		utils.Error(err.Error(), err)
	}

	return nil
}

func (c *ChallengeService) FindAll(page *int) ([]*models.ChallengeModel, error) {

	return c.repo.GetAll(page)
}

func (c *ChallengeService) FindById(id string) (*models.ChallengeModel, error) {
	return c.repo.GetById(id)
}

func (c *ChallengeService) CountChallenges() (int, error) {
	return c.repo.CountChallenges()
}

func (c *ChallengeService) Delete(id string) error {
	return c.repo.Delete(id)
}
