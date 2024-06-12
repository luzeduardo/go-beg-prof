package translation

import "strings"

func Translate(word string, language string) string {
	sanitizedWord := sanitizeWord(word)
	sanitizedLanguage := sanitizeWord(language)
	switch sanitizedLanguage {
	case "english":
		return sanitizedWord
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	case "dutch":
		return ""
	}
	return sanitizedWord
}

func sanitizeWord(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}
