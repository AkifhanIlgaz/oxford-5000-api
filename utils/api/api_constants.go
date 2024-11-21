package api

const (
	WordParam         string = "word"
	PartOfSpeechParam string = "part_of_speech"
	UidParam          string = "uid"
	NameParam         string = "name"
	ApiKeyParam       string = "apikey"
)

const (
	AuthHeader string = "Authorization"
)

const (
	ContextUid string = "uid"
)

const (
	FreePlanUsage int = 10
	ProPlanUsage  int = 1000
)

var PlanUsageLimits = map[string]int{
	"free": FreePlanUsage,
	"pro":  ProPlanUsage,
}
