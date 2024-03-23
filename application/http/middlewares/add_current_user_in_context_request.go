package middlewares

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/gin-gonic/gin"
)

func AddCurrentInContextRequest(svc services.SessionServiceInterface) gin.HandlerFunc {

	return func(c *gin.Context) {
		var u views.UserResponseView

		svc.Get(c, "user", &u)

		c.Set("user", u)

		c.Next()
	}
}
