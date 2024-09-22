package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

const WordsPath = "/words"

type WordController struct {
	wordService services.WordService
}

func NewWordController(wordService services.WordService) WordController {
	return WordController{
		wordService: wordService,
	}
}

func (controller WordController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(WordsPath)

	// GET /documents?field=id&value=123
	// GET /documents?field=index&value=123
	router.GET("/", controller.GetWord)
}

// TODO: Response functions
func (controller WordController) GetWord(ctx *gin.Context) {
	searchBy := ctx.Query(api.SearchByQueryParam)
	value := ctx.Query(api.ValueQueryParam)
	switch searchBy {
	case api.SearchById:
		word, err := controller.wordService.GetById(value)
		if err != nil {
			log.Println(err.Error())
			response.WithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		response.WithSuccess(ctx, http.StatusOK, message.WordFound, word)
	case api.SearchByIndex:
		index, err := strconv.Atoi(value)
		if err != nil {
			log.Println(err.Error())
			response.WithError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		word, err := controller.wordService.GetByIndex(index)
		if err != nil {
			log.Println(err.Error())
			response.WithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		response.WithSuccess(ctx, http.StatusOK, message.WordFound, word)
	default:
		log.Println(message.UnsupportedSearchValue)
		response.WithError(ctx, http.StatusInternalServerError, message.UnsupportedSearchValue)
	}

}
