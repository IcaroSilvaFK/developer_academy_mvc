package middlewares

import (
	"time"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		init := time.Now()
		utils.Info(
			"Request initilizade",
			zap.String("client", ctx.ClientIP()),
			zap.String("route", ctx.FullPath()),
		)

		ctx.Next()

		elapsed := time.Since(init)

		utils.Info(
			"Request finalized",
			zap.String("client", ctx.ClientIP()),
			zap.String("route", ctx.FullPath()),
			zap.Any("duration", elapsed),
		)
	}
}
