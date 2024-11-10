package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

const WordPath = "/word"

type WordController struct {
	wordService services.WordService
}

func NewWordController(wordService services.WordService) WordController {
	return WordController{
		wordService: wordService,
	}
}

func (controller WordController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(WordPath)

	// GET /:word?part_of_speech
	router.GET("/:word", controller.GetWord)
}

func (controller WordController) GetWord(ctx *gin.Context) {
	word := ctx.Param(api.WordParam)

	partOfSpeech := ctx.Query(api.PartOfSpeechParam)
	if err := models.IsValidPartOfSpeech(partOfSpeech); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
	}

	wordInfo, err := controller.wordService.GetByName(word, partOfSpeech)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
	}

	response.WithSuccess(ctx, http.StatusOK, message.WordFound, wordInfo)
}
