package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/middlewares"
	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

const boxPath string = "/box"

type BoxController struct {
	boxService services.BoxService
}

func NewBoxController(boxService services.BoxService) BoxController {
	return BoxController{
		boxService: boxService,
	}
}

func (controller BoxController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(boxPath)

	router.POST("/action", controller.Action)
}

func (controller BoxController) Action(ctx *gin.Context) {

	user, err := middlewares.GetUserFromContext(ctx)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(200, &user)
	var boxActionRequest models.BoxActionRequest

	if err := ctx.ShouldBindJSON(&boxActionRequest); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = controller.boxService.Action(boxActionRequest)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, "success", nil)
}
