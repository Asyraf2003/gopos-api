package domain

import "strings"

func normalize(value string) string {
	return strings.TrimSpace(value)
}
