package models

// WordInfo represents detailed information about a word
// @Description Detailed information about a word
type WordInfo struct {
	Index  int    `json:"index"`
	Source string `json:"source"`
	Word   string `json:"word"`
	Header
	Definitions []Definition `json:"definitions"`
	Idioms      []Idiom      `json:"idioms"`
}

type Header struct {
	Audio struct {
		UK string `json:"UK" bson:"UK"`
		US string `json:"US" bson:"US"`
	} `json:"audio" bson:"audio"`
	PartOfSpeech string `json:"partOfSpeech" bson:"partOfSpeech"`
	CEFRLevel    string `json:"CEFRLevel" bson:"CEFRLevel"`
}

// Definition represents a single definition of a word
// @Description A single definition of a word
type Definition struct {
	// PartOfSpeech indicates the grammatical category
	// @Description Grammatical category (noun, verb, etc.)
	PartOfSpeech string `json:"partOfSpeech"`
	
	// Meaning contains the actual definition
	// @Description The actual definition of the word
	Meaning  string   `json:"meaning"`
	Examples []string `json:"examples"`
}

type Idiom struct {
	Usage       string       `json:"usage"`
	Definitions []Definition `json:"definition"`
}
