package routers

import (
	"github.com/AkifhanIlgaz/dictionary-api/controllers"
	"github.com/gin-gonic/gin"
)

const BoxPath = "/box"

type BoxRouter struct {
	boxController controllers.BoxController
}

func NewBoxRouter(boxController controllers.BoxController) BoxRouter {
	return BoxRouter{
		boxController: boxController,
	}
}

func (r BoxRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(BoxPath)

	_ = router
}
