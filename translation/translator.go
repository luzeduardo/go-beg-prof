package translation

import "strings"

type StaticService struct{}

func NewStaticService() *StaticService {
	return &StaticService{}
}

func (st *StaticService) Translate(word string, language string) string {
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
	return ""
}

func sanitizeWord(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}
