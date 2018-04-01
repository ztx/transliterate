package language

import (
	//"fmt"
	 "github.com/ztx/transliterate/mapp"
	lan "golang.org/x/text/language"
)

type Script struct {
	language     lan.Tag
	UnicodeStart rune
	UnicodeEnd   rune

	//map of functions providing transliteration to other languages
	transliterationFuncs map[lan.Tag]mapp.RuneMapper
	// standardizer functions
	StandardizerFuncs []Standardizer
}

//type unicodeStartEnd map[lan.Tag][]rune

type Standardizer interface {
	Standardize(string) string
}

func NewScript(lang lan.Tag) Script {
	return Script{language: lang,
			transliterationFuncs: make(map[lan.Tag]mapp.RuneMapper),
			UnicodeStart:UnicodeStartEnd[lang][0],
			UnicodeEnd:UnicodeStartEnd[lang][1],
	}
}

func NewCustomScript(lang lan.Tag, unicodeStart, unicodeEnd rune) Script {
	return Script{language: lang,
		UnicodeStart:         unicodeStart,
		UnicodeEnd:           unicodeEnd,
		transliterationFuncs: make(map[lan.Tag]mapp.RuneMapper)}
}

//ValidRune validates a character to its script
//ValidRune returns true if
//the rune is between unicodeStart and unicodeEnd
func (s *Script) ValidRune(c rune) bool {
	if c >= s.UnicodeStart && c <= s.UnicodeEnd {
		return true
	}
	return false
}

func (s *Script) RegisterRuneMapper(language lan.Tag, runeMapper mapp.RuneMapper) {
	s.transliterationFuncs[language] = runeMapper
}

func (s *Script) TransliterateRune(toLanguage lan.Tag, c rune) <-chan rune {
	runeMapper := s.transliterationFuncs[toLanguage]
	return runeMapper.To(c)
}

func (s *Script) TransliterateString(toLanguage lan.Tag, str string) string {
	runeMapper := s.transliterationFuncs[toLanguage]
	transliteratedStr := ""
	//i := 0
	runes := make([]rune, 0)
	for _, r := range str {
		if s.ValidRune(r) {

			for c := range runeMapper.To(r) {

				runes = append(runes, c)
			}
		} else {
			runes = append(runes, r)
		}
	}
	transliteratedStr = string(runes)
	return transliteratedStr
}

//After transliteration not so well known characters may show up which
//need to be replaced by well known characters
//Eg: in normal-Hindi there is no short e character which present in Kannada (ಎ)
func (s *Script) Standardize(str string) string {
	result := ""
	for _, f := range s.StandardizerFuncs {
		result = f.Standardize(result)
	}
	return result
}
