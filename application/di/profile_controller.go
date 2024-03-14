package di

import "github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"

func NewProfileController() controllers.ProfileControllerInterface {

	return controllers.NewProfileController()
}
