package middlewares

import (
	"errors"
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(
	svc services.SessionServiceInterface,
) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var u views.UserResponseView

		svc.Get(ctx, "user", &u)

		if u.ID == "" {
			ctx.Error(errors.New("THE USER IS NOT AUTHENTICATED"))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
