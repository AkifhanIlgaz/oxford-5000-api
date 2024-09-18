package routers

import (
	"github.com/AkifhanIlgaz/dictionary-api/controllers"
	"github.com/gin-gonic/gin"
)

const WordsPath = "/words"

type WordRouter struct {
	wordController controllers.WordController
}

func NewWordRouter(wordController controllers.WordController) WordRouter {
	return WordRouter{
		wordController: wordController,
	}
}

func (r WordRouter) Setup(rg *gin.RouterGroup) {
	router := rg.Group(WordsPath)

	// GET /documents?field=id&value=123
	// GET /documents?field=index&value=123
	router.GET("/", r.wordController.GetWord)
}
