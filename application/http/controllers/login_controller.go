package controllers

import (
	"net/http"
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
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

	goContext := ctx.Request.Context()

	users, top, err := c.usvc.GetTenFirstUserAndCount(goContext)

	if err != nil {
		utils.Error("Error on request top ten users", err)
	}

	countchallenges, err := c.challengeservice.CountChallenges(goContext)

	if err != nil {
		utils.Error("Error on count challenges", err)
	}

	var r []views.UserResponseView

	for _, u := range users {
		r = append(r, *views.NewUserResponseView(u))
	}

	ctx.HTML(http.StatusOK, "index.gotmpl", gin.H{
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

	if code == "" && ctx.Request.Method == "POST" {
		input := new(dtos.LoginInputDto)

		if err := ctx.ShouldBindJSON(input); err != nil {
			errno := utils.NewBadRequestException(err.Error())
			ctx.JSON(errno.Code, errno)
			return
		}

		r, err := c.svc.LoginWithPassword(ctx.Request.Context(), input)

		if err != nil {
			ctx.JSON(err.Code, err)
			return
		}
		u := views.NewUserResponseView(r)

		c.sessionservice.Set(ctx, "user", u)

		ctx.JSON(http.StatusOK, u)
		return
	}
	goContext := ctx.Request.Context()
	u, err := c.svc.Login(goContext, code, provider)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	r := views.NewUserResponseView(u)

	c.sessionservice.Set(ctx, "user", r)

	ctx.JSON(http.StatusOK, r)
}
