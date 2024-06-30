package controllers

import (
	"net/http"
	"strconv"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	svc        services.ChallengeServiceInterface
	catService services.ChallengesCategoriesServiceInterface
}

type HomeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewHomeController(
	svc services.ChallengeServiceInterface,
	catService services.ChallengesCategoriesServiceInterface,
) HomeControllerInterface {
	return &HomeController{
		svc, catService,
	}
}

func (hc *HomeController) Index(ctx *gin.Context) {
	q := ctx.Param("page")

	v, err := strconv.Atoi(q)

	if err != nil {
		v = 1
	}

	challenges, restErr := hc.svc.FindAll(ctx.Request.Context(), &v)

	if restErr != nil {
		ctx.Redirect(http.StatusPermanentRedirect, "/error")
		return
	}
	categories, restErr := hc.catService.GetAll(ctx.Request.Context(), "")

	if restErr != nil {
		utils.Error("Error on get all categories", err)
	}

	u := utils.GetCurrentUserInRequestContext(ctx)

	r := views.NewChallengeResponseOutputList(challenges)

	ctx.HTML(200, "home.gotmpl", gin.H{
		"challenges": r,
		"user":       u,
		"categories": categories,
	})
}
