package translation

func Translate(word string, language string) string {
	switch language {
	case "english":
		return word
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	case "dutch":
		return ""
	}
	return word
}
