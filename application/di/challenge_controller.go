package di

import "github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"

func NewChallengeController() controllers.ChallengeControllerInterface {
	return controllers.NewChallengeController()
}
