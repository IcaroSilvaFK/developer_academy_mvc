package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	svc              services.LoginServiceInterface
	usvc             services.UserServiceInterface
	challengeservice services.ChallengeServiceInterface
}

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

func NewLoginController(
	svc services.LoginServiceInterface,
	usvc services.UserServiceInterface,
	challengeservice services.ChallengeServiceInterface,
) LoginControllerInterface {
	return &LoginController{
		svc, usvc, challengeservice,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {

	users, top, err := c.usvc.GetTenFirstUserAndCount()

	if err != nil {
		fmt.Println(err)
	}

	countchallenges, err := c.challengeservice.CountChallenges()

	if err != nil {
		log.Println(err)
	}

	var r []views.UserResponseView

	for _, u := range users {
		r = append(r, *views.NewUserResponseView(u))
	}

	ctx.HTML(200, "login.gotmpl", gin.H{
		"users":      r,
		"quantity":   top,
		"error":      err,
		"challenges": countchallenges,
		"client_id":  os.Getenv(utils.GITHUB_CLIENT_ID),
	})
}

func (c *LoginController) SignIn(ctx *gin.Context) {
	code := ctx.Param("code")
	session := sessions.Default(ctx)

	if code == "" {
		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "Bad request missing code",
		})
		return
	}

	u, err := c.svc.Login(code)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := views.NewUserResponseView(u)

	bt, _ := json.Marshal(r)

	// TODO define in constant this key
	session.Set("user", string(bt))
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"user": r,
	})
}
