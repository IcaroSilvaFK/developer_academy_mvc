package di

import "github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"

func NewHomeController() controllers.HomeControllerInterface {

	return controllers.NewHomeController()
}
