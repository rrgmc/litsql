package jen

import (
	"unicode"
)

func InitialIsLower(s string) bool {
	for _, r := range s {
		return r == unicode.ToLower(r)
	}
	return false
}

func InitialIsUpper(s string) bool {
	for _, r := range s {
		return r == unicode.ToUpper(r)
	}
	return false
}

// InitialToLower converts initial to lower.
func InitialToLower(s string) string {
	for _, r := range s {
		u := string(unicode.ToLower(r))
		return u + s[len(u):]
	}

	return s
}

// InitialToUpper converts initial to upper.
func InitialToUpper(s string) string {
	for _, r := range s {
		u := string(unicode.ToUpper(r))
		return u + s[len(u):]
	}

	return ""
}
