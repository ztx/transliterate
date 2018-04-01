package defaults

import (
	script "github.com/ztx/transliterate/language"
	"golang.org/x/text/language"
)

var (
	Kannada = script.NewCustomScript(language.Kannada,
		script.UnicodeStartEnd[language.Kannada][0],
		script.UnicodeStartEnd[language.Kannada][1])
	Hindi = script.NewCustomScript(language.Hindi,
		script.UnicodeStartEnd[language.Hindi][0],
		script.UnicodeStartEnd[language.Hindi][1])
)
