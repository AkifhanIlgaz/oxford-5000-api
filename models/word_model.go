package models

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

type Definition struct {
	Meaning  string   `json:"meaning"`
	Examples []string `json:"examples"`
}

type Idiom struct {
	Usage       string       `json:"usage"`
	Definitions []Definition `json:"definition"`
}
