package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/gin-gonic/gin"
)

type WordController struct {
	wordService services.WordService
}

func NewWordController(wordService services.WordService) WordController {
	return WordController{
		wordService: wordService,
	}
}

func (controller WordController) SetupRoutes(rg *gin.RouterGroup) {

}

// TODO: Response functions
func (controller WordController) GetWord(ctx *gin.Context) {
	searchBy := ctx.Query(api.SearchByQueryParam)
	value := ctx.Query(api.ValueQueryParam)
	switch searchBy {
	case api.SearchById:
		word, err := controller.wordService.GetById(value)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, &word)

	case api.SearchByIndex:
		index, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		word, err := controller.wordService.GetByIndex(index)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, &word)
	default:
		// TODO: Error constant
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("unsupported search value"))
	}

}
