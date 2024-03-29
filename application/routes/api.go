package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/di"
	"github.com/gin-gonic/gin"
)

func NewApiRoutes(engine *gin.Engine) {

	group := engine.Group("/api/v1")

	loginController := di.NewLoginController()
	createChallengeController := di.NewCreateNewChallengeController()
	commentsChallengeController := di.NewCommentsChallengeController()
	challengeController := di.NewChallengeController()
	uController := di.NewProfileController()

	group.GET("/heath", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"check": true,
		})
	})

	group.GET("/login/:code", loginController.SignIn)

	// CHALLENGES
	group.POST("/challenges", createChallengeController.Create)
	group.DELETE("/challenges/:id", challengeController.Destroy)
	group.GET("/challenges", challengeController.GetAllChallenges)
	group.GET("/challenges/:id", challengeController.FindById)
	group.GET("/challenges/users/:userId", challengeController.FindUserId)

	// comments
	group.POST("/challenges/comments", commentsChallengeController.Create)
	group.DELETE("/challenges/comments/:id", commentsChallengeController.Destroy)
	group.GET("/challenges/comments/users/:userId", commentsChallengeController.FindUserComments)
	group.GET("/challenges/comments/:id", commentsChallengeController.FindCommentById)
	group.GET("/challenges/comments/challenge/:challengeId", commentsChallengeController.FindChallengesComments)

	//users
	group.GET("/users", uController.FindAllUsers)
	group.GET("/users/:id", uController.FindByUserId)
	group.DELETE("/users/:id", uController.Delete)
}
