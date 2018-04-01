package language

import "golang.org/x/text/language"

var (
	UnicodeStartEnd = map[language.Tag][]rune{
		language.Kannada: []rune{'\u0C80', '\u0CF2'},
		language.Hindi:   []rune{'\u0900', '\u097F'},
	}
)
