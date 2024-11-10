package models

import (
	"errors"
	"slices"

	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
)

var PartOfSpeeches = []string{
	"adjective",
	"adverb",
	"verb",
	"noun",
	"pronoun",
	"preposition",
	"exclamation",
	"determiner",
	"conjunction",
	"number",
	"modal verb",
	"auxiliary verb",
	"indefinite article",
}

func IsValidPartOfSpeech(partOfSpeech string) error {
	if !slices.Contains(PartOfSpeeches, partOfSpeech) {
		return errors.New(message.UnsupportedPartOfSpeech)
	}
	return nil
}

var CEFRLevels = []string{
	"A1",
	"A2",
	"B1",
	"B2",
	"C1",
}
