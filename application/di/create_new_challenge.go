package di

import "github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"

func CreateNewChallengeController() controllers.CreateNewChallengeControllerInterface {
	return controllers.NewCreateNewChallengeController()
}
