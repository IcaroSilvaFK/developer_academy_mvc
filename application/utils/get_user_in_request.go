package utils

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/gin-gonic/gin"
)

func GetCurrentUserInRequestContext(c *gin.Context) views.UserResponseView {

	u, _ := c.Get("user")

	return u.(views.UserResponseView)
}
