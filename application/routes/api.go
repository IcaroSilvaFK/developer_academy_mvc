package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/di"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/middlewares"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/gin-gonic/gin"
)

func NewApiRoutes(engine *gin.Engine) {

	group := engine.Group("/api/v1")
	authMiddleware := middlewares.AuthMiddleware(services.NewSessionService())

	loginController := di.NewLoginController()
	createChallengeController := di.NewCreateNewChallengeController()
	commentsChallengeController := di.NewCommentsChallengeController()
	challengeController := di.NewChallengeController()
	uController := di.NewProfileController()
	challengesCategories := di.NewChallengesCategoriesController()

	group.GET("/heath", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"check": true,
		})
	})

	group.GET("/login/:code", loginController.SignIn)

	// CHALLENGES
	group.POST("/challenges", authMiddleware, createChallengeController.Create)
	group.DELETE("/challenges/:id", authMiddleware, challengeController.Destroy)
	group.GET("/challenges/:id", challengeController.FindById)
	group.GET("/challenges", challengeController.GetAllChallenges)
	group.GET("/challenges/users/:userId", challengeController.FindUserId)

	// comments
	group.POST("/challenges/comments", authMiddleware, commentsChallengeController.Create)
	group.DELETE("/challenges/comments/:id", authMiddleware, commentsChallengeController.Destroy)
	group.GET("/challenges/comments/users/:userId", commentsChallengeController.FindUserComments)
	group.GET("/challenges/comments/:id", commentsChallengeController.FindCommentById)
	group.GET("/challenges/comments/challenge/:challengeId", commentsChallengeController.FindChallengesComments)

	//users
	group.GET("/users", uController.FindAllUsers)
	group.GET("/users/:id", uController.FindByUserId)
	group.DELETE("/users/:id", authMiddleware, uController.Delete)

	// challenges categories
	group.POST("/challenges/categories", authMiddleware, challengesCategories.Create)
	group.DELETE("/challenges/categories/:id", authMiddleware, challengesCategories.Delete)
	group.GET("/challenges/categories", challengesCategories.FindAll)
	group.GET("/challenges/categories/:id", challengesCategories.FindById)
	group.GET("/challenges/categories/users/:userId", challengesCategories.FindByUserId)
	group.PUT("/challenges/categories/:id", authMiddleware, challengesCategories.Update)
}
