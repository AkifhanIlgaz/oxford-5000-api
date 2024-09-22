package controllers

import (
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/gin-gonic/gin"
)

type BoxController struct {
	boxService services.BoxService
}

func NewBoxController(boxService services.BoxService) BoxController {
	return BoxController{
		boxService: boxService,
	}
}

func (controller BoxController) SetupRoutes(rg *gin.RouterGroup) {

}
