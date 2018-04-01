package script

import (
	//"fmt"
	"github.com/ztx/transliterate/mapp"
	"golang.org/x/text/language"
)

type Script struct {
	language     language.Tag
	unicodeStart rune
	unicodeEnd   rune

	//map of functions providing transliteration to other languages
	transliterationFuncs map[language.Tag]mapp.RuneMapper
	// standardizer functions
	StandardizerFuncs []Standardizer
}

type unicodeStartEnd map[language.Tag][]rune

type Standardizer interface {
	Standardize(string) string
}

func NewScript(lang language.Tag) Script {
	return Script{language: lang, transliterationFuncs: make(map[language.Tag]mapp.RuneMapper)}
}

func NewCustomScript(lang language.Tag, unicodeStart, unicodeEnd rune) Script {
	return Script{language: lang,
		unicodeStart:         unicodeStart,
		unicodeEnd:           unicodeEnd,
		transliterationFuncs: make(map[language.Tag]mapp.RuneMapper)}
}

//ValidRune validates a character to its script
//ValidRune returns true if
//the rune is between unicodeStart and unicodeEnd
func (s *Script) ValidRune(c rune) bool {
	if c >= s.unicodeStart && c <= s.unicodeEnd {
		return true
	}
	return false
}

func (s *Script) RegisterRuneMapper(language language.Tag, runeMapper mapp.RuneMapper) {
	s.transliterationFuncs[language] = runeMapper
}

func (s *Script) TransliterateRune(toLanguage language.Tag, c rune) <-chan rune {
	runeMapper := s.transliterationFuncs[toLanguage]
	return runeMapper.To(c)
}

func (s *Script) TransliterateString(toLanguage language.Tag, str string) string {
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
//Eg: in general-Hindi there is no short e character which present in Kannada (ಎ)
func (s *Script) Standardize(str string) string {
	result := ""
	for _, f := range s.StandardizerFuncs {
		result = f.Standardize(result)
	}
	return result
}
