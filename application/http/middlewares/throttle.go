package middlewares

import (
	"errors"
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Throttle(maxEventsPerSec int, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(ctx *gin.Context) {

		if limiter.Allow() {
			ctx.Next()
			return
		}

		ctx.Error(errors.New(utils.LIMIT_EXCEEDED))
		ctx.AbortWithStatus(http.StatusTooManyRequests)
	}
}
