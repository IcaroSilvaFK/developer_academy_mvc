package controllers

import (
	"net/http"
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	svc              services.LoginServiceInterface
	usvc             services.UserServiceInterface
	challengeservice services.ChallengeServiceInterface
	sessionservice   services.SessionServiceInterface
}

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

func NewLoginController(
	svc services.LoginServiceInterface,
	usvc services.UserServiceInterface,
	challengeservice services.ChallengeServiceInterface,
	sessionservice services.SessionServiceInterface,
) LoginControllerInterface {
	return &LoginController{
		svc, usvc, challengeservice, sessionservice,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {

	users, top, err := c.usvc.GetTenFirstUserAndCount(ctx.Request.Context())

	if err != nil {
		utils.Error("Error on request top ten users", err)
	}

	countchallenges, err := c.challengeservice.CountChallenges(
		ctx.Request.Context(),
	)

	if err != nil {
		utils.Error("Error on count challenges", err)
	}

	var r []views.UserResponseView

	for _, u := range users {
		r = append(r, *views.NewUserResponseView(u))
	}

	ctx.HTML(http.StatusOK, "login.gotmpl", gin.H{
		"users":            r,
		"quantity":         top,
		"error":            err,
		"challenges":       countchallenges,
		"client_id":        os.Getenv(utils.GITHUB_CLIENT_ID),
		"gitlab_client_id": os.Getenv(utils.GITLAB_APP_ID),
	})
}

// @Summary		Signin with code provides the provider metho to signin now using [github,gitlab]
// @Description	Sigin with platform
// @Tags			signin
// @Param			code		query	string	false "Code return on execute signIn with github or gitlab"
// @Param			provider	query	string	false "github or gitlab or empty"
// @Produce		json
// @Success		200 {object} 	views.UserResponseView
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/login [get]
func (c *LoginController) SignIn(ctx *gin.Context) {

	code := ctx.Param("code")
	provider := ctx.Query("provider")

	if code == "" {
		err := utils.NewBadRequestException("Missing a param code in request")
		ctx.JSON(err.Code, err)
		return
	}

	u, err := c.svc.Login(ctx.Request.Context(), code, provider)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	r := views.NewUserResponseView(u)

	c.sessionservice.Set(ctx, "user", r)

	ctx.JSON(http.StatusOK, r)
}
