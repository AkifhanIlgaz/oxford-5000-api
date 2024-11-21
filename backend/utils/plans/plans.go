package plans

import (
	"fmt"
)

type PlanType string

const (
	FreePlan     PlanType = "Free"
	StandardPlan PlanType = "Standard"
	ProPlan      PlanType = "Pro"
)

type Plan struct {
	Type           PlanType
	RequestsPerDay int
	Features       []string
	Price          float64
}

func GetPlan(planType PlanType) Plan {
	switch planType {
	case FreePlan:
		return Plan{
			Type:           FreePlan,
			RequestsPerDay: 100,
			Features:       []string{"Basic word lookups", "Definition search"},
			Price:          0,
		}
	case StandardPlan:
		return Plan{
			Type:           StandardPlan,
			RequestsPerDay: 1000,
			Features:       []string{"Basic word lookups", "Definition search", "Examples", "Synonyms"},
			Price:          9.99,
		}
	case ProPlan:
		return Plan{
			Type:           ProPlan,
			RequestsPerDay: 10000,
			Features:       []string{"Basic word lookups", "Definition search", "Examples", "Synonyms", "Etymology", "Advanced API features"},
			Price:          19.99,
		}
	default:
		return Plan{}
	}
}

func UpgradePlan(currentPlan, targetPlan PlanType) (Plan, error) {
	if !isValidUpgrade(currentPlan, targetPlan) {
		return Plan{}, fmt.Errorf("invalid upgrade path from %s to %s", currentPlan, targetPlan)
	}
	return GetPlan(targetPlan), nil
}

func DowngradePlan(currentPlan, targetPlan PlanType) (Plan, error) {
	if !isValidDowngrade(currentPlan, targetPlan) {
		return Plan{}, fmt.Errorf("invalid downgrade path from %s to %s", currentPlan, targetPlan)
	}
	return GetPlan(targetPlan), nil
}

func isValidUpgrade(current, target PlanType) bool {
	switch current {
	case FreePlan:
		return target == StandardPlan || target == ProPlan
	case StandardPlan:
		return target == ProPlan
	default:
		return false
	}
}

func isValidDowngrade(current, target PlanType) bool {
	switch current {
	case ProPlan:
		return target == StandardPlan || target == FreePlan
	case StandardPlan:
		return target == FreePlan
	default:
		return false
	}
}
