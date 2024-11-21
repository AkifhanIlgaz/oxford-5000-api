package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/middlewares"
	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

// @title Dictionary API
// @version 1.0
// @description API for dictionary operations
// @BasePath /api/v1

const WordPath = "/word"

type WordController struct {
	wordService    services.WordService
	wordMiddleware middlewares.WordMiddleware
}

func NewWordController(wordService services.WordService, wordMiddleware middlewares.WordMiddleware) WordController {
	return WordController{
		wordService:    wordService,
		wordMiddleware: wordMiddleware,
	}
}

func (controller WordController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(WordPath)
	router.Use(controller.wordMiddleware.TrackUsage())

	router.GET("/search/:word", controller.GetWord)
}

// GetWord godoc
// @Summary Get word information
// @Description Retrieves detailed information about a word including definitions, examples, and usage
// @Tags word
// @Accept json
// @Produce json
// @Param word path string true "Word to look up"
// @Param part_of_speech query string false "Filter by part of speech (noun, verb, adjective, etc.)"
// @Success 200 {object} response.Response{data=models.WordInfo} "Word found successfully"
// @Failure 400 {object} response.Response "Invalid part of speech"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /word/{word} [get]
func (controller WordController) GetWord(ctx *gin.Context) {
	word := ctx.Param(api.WordParam)

	partOfSpeech := ctx.Query(api.PartOfSpeechParam)
	if err := models.IsValidPartOfSpeech(partOfSpeech); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	wordInfo, err := controller.wordService.GetByName(word, partOfSpeech)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.WordFound, wordInfo)
}
