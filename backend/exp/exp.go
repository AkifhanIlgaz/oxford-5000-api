package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/dictionary-api/utils/parser"
)

func main() {
	wordInfo, err := parser.ParseWord("https://www.oxfordlearnersdictionaries.com/definition/english/abandon_1")
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.MarshalIndent(wordInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}
