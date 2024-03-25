package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	apputils "github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"

	"github.com/gin-gonic/gin"
)

type ChallengeController struct {
	svc services.ChallengeServiceInterface
}

type ChallengeControllerInterface interface {
	Index(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

func NewChallengeController(
	svc services.ChallengeServiceInterface,
) ChallengeControllerInterface {
	return &ChallengeController{
		svc,
	}
}

func (cc *ChallengeController) Index(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" || !utils.IsValidId(id) {
		//TODO implment redirect to error page
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	c, err := cc.svc.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	d, _ := c.CreatedAt.UTC().MarshalText()
	u := apputils.GetCurrentUserInRequestContext(ctx)

	countComments := len(c.Comments)

	ctx.HTML(http.StatusOK, "challenge.gotmpl", gin.H{
		"challenge":        c,
		"created":          string(d),
		"quantityComments": countComments,
		"comments":         views.NewCommentChallengeListOutputView(c.Comments),
		"id":               id,
		"user":             u,
	})
}

func (cc *ChallengeController) Create(ctx *gin.Context) {

	ctx.JSON(http.StatusCreated, nil)
}

func (cc *ChallengeController) Destroy(ctx *gin.Context) {

	id := ctx.Param("id")

	if !utils.IsValidId(id) {

		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := cc.svc.Delete(id); err != nil {

		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
