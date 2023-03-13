package utils

import (
	"strings"
)

func DefaultToTrue(s string) bool {
	pred := strings.ToLower(s)
	return pred != "f" && pred != "F" && pred != "false" || pred == "n"
}

func DefaultToFalse(s string) bool {
	pred := strings.ToLower(s)
	return pred == "t" || pred == "T" || pred == "true" || pred == "y"
}
