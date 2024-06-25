package utils

import (
	"github.com/gin-gonic/gin"
)

func GetCurrentUserInRequestContext(c *gin.Context) interface{} {

	u, _ := c.Get("user")

	return u
}
