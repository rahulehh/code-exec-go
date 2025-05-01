package executor

var supportedLangs = map[string]struct{}{
	"python": {},
	"c":      {},
	"cpp":    {},
	"go":     {},
}

// isLanguageSupported checks if a language is supported
func isLanguageSupported(lang string) bool {
	_, ok := supportedLangs[lang]
	return ok
}
