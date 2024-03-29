package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Throllet(maxEventsPerSec int, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(ctx *gin.Context) {

		if limiter.Allow() {
			ctx.Next()
			return
		}

		ctx.Error(errors.New("Limit exceeded"))
		ctx.AbortWithStatus(http.StatusTooManyRequests)
	}
}
