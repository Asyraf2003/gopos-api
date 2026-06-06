package domain

type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

func (r RiskLevel) Valid() bool {
	return r == RiskLow || r == RiskMedium || r == RiskHigh
}
