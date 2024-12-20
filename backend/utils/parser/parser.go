package parser

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/PuerkitoBio/goquery"
)

// Package parser provides functionality to parse word information from Oxford Dictionary web pages.
func ParseWord(wordUrl string) (models.WordInfo, error) {
	var wordInfo models.WordInfo

	req, err := http.NewRequest(http.MethodGet, wordUrl, nil)
	if err != nil {
		return wordInfo, fmt.Errorf("client: could not create request: %s\n", err)

	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return wordInfo, fmt.Errorf("client: error making http request: %s\n", err)
	}

	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return wordInfo, fmt.Errorf("parsing word: %w", err)
	}

	if document == nil {
		return wordInfo, fmt.Errorf("parsing word: %w", err)
	}
	mainContainer := document.Find("#entryContent")

	parseHeader(mainContainer.Find(".webtop"), &wordInfo)
	parseDefinitions(mainContainer.Find("ol.senses_multiple").First().Find("li.sense"), &wordInfo)
	parseDefinitions(mainContainer.Find("ol.sense_single").First().Find("li.sense"), &wordInfo)
	parseIdioms(mainContainer.Find("div.idioms .idm-g"), &wordInfo)

	return wordInfo, nil
}

// parseHeader extracts header information from the word entry including:
// - The word itself
// - Part of speech
// - CEFR level
// - Audio pronunciations (UK and US variants)
func parseHeader(mainContainer *goquery.Selection, wordInfo *models.WordInfo) {
	// HeadingWord
	mainContainer.Find(".headword").First().Each(func(i int, s *goquery.Selection) {
		wordInfo.Word = s.Text()
	})

	// Part Of Speech
	mainContainer.Find("span.pos").Each(func(i int, s *goquery.Selection) {
		wordInfo.Header.PartOfSpeech = s.Text()
	})

	// CEFR Level
	mainContainer.Find(".symbols span").First().Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("class")
		wordInfo.Header.CEFRLevel = strings.ToUpper(strings.Split(attr, "_")[1])
	})

	// Audio
	mainContainer.Find(`span.phonetics div > div`).Each(func(i int, s *goquery.Selection) {
		audioUrl, _ := s.Attr("data-src-mp3")

		// We don't need to check `pron-us` since there is only two possibilities
		if s.HasClass("pron-uk") {
			wordInfo.Header.Audio.UK = audioUrl
		} else {
			wordInfo.Header.Audio.US = audioUrl
		}

	})

}

// parseDefinitions extracts word definitions and their corresponding examples
// from the provided DOM selection. It supports both single and multiple sense formats.
func parseDefinitions(mainContainer *goquery.Selection, wordInfo *models.WordInfo) {
	mainContainer.Each(func(i int, s *goquery.Selection) {
		var definition models.Definition

		s.Find("span.def").Each(func(i int, s *goquery.Selection) {
			definition.Meaning = s.Text()
		})

		s.Find("ul.examples > li span.x").Each(func(i int, s *goquery.Selection) {
			html, _ := s.Html()
			definition.Examples = append(definition.Examples, html)
		})
		wordInfo.Definitions = append(wordInfo.Definitions, definition)

	})

}

// parseIdioms extracts idioms related to the word, including their usage and definitions
// with examples. Each idiom can have multiple definitions.
func parseIdioms(mainContainer *goquery.Selection, wordInfo *models.WordInfo) {
	mainContainer.Each(func(i int, s *goquery.Selection) {
		var idiom models.Idiom
		idiom.Usage = s.Find("div.top-container").Text()

		s.Find(`ol[class^="sense"] li.sense`).Each(func(i int, s *goquery.Selection) {
			var definition models.Definition
			definition.Meaning = s.Find("span.def").Text()

			s.Find("ul.examples li span.x").Each(func(i int, s *goquery.Selection) {
				definition.Examples = append(definition.Examples, s.Text())
			})

			idiom.Definitions = append(idiom.Definitions, definition)
		})
		wordInfo.Idioms = append(wordInfo.Idioms, idiom)
		idiom = models.Idiom{}
	})
}
