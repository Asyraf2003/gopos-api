package domain

import "strings"

func normalizeCode(code string) *string {
	trimmed := strings.TrimSpace(code)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

func normalizeDisplayText(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func normalizeSearchText(value string) string {
	return strings.ToLower(normalizeDisplayText(value))
}

func validateThreshold(reorderPointQty, criticalThresholdQty *int) error {
	if (reorderPointQty == nil) != (criticalThresholdQty == nil) {
		return ErrProductThresholdPairRequired
	}

	if reorderPointQty == nil {
		return nil
	}

	if *reorderPointQty < 0 || *criticalThresholdQty < 0 {
		return ErrProductThresholdNegative
	}

	if *criticalThresholdQty > *reorderPointQty {
		return ErrProductCriticalAboveReorder
	}

	return nil
}

func copyIntPtr(value *int) *int {
	if value == nil {
		return nil
	}
	copied := *value
	return &copied
}
