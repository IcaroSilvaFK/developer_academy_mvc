package di

import "github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"

func NewCreateNewChallengeController() controllers.CreateNewChallengeControllerInterface {
	return controllers.NewCreateNewChallengeController()
}
