package services

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type ChallengeHintService struct {
	repo repositories.ChallengesHintsRepositoryInterface
}

type ChallengeHintServiceInterface interface {
	Create(challengeId, text string) error
	FindByChallengeId(challengeId string) (*models.ChallengeHintsModel, error)
	FindById(id string) (*models.ChallengeHintsModel, error)
	Delete(id string) error
}

func NewChallengeHintService(
	repo repositories.ChallengesHintsRepositoryInterface,
) ChallengeHintServiceInterface {

	return &ChallengeHintService{
		repo,
	}
}

func (ch *ChallengeHintService) Create(challengeId, text string) error {

	h := models.NewChallengeHintsModel(challengeId, text)

	err := ch.repo.Create(h)

	return err
}

func (ch *ChallengeHintService) FindByChallengeId(challengeId string) (*models.ChallengeHintsModel, error) {

	h, err := ch.repo.GetByChallengeId(challengeId)

	if err != nil {
		return nil, err
	}

	return h, nil
}

func (ch *ChallengeHintService) FindById(id string) (*models.ChallengeHintsModel, error) {

	h, err := ch.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return h, nil
}

func (ch *ChallengeHintService) Delete(id string) error {
	return ch.repo.Delete(id)
}
