package transliterate

import (
	"github.com/ztx/transliterate/language"
	"github.com/ztx/transliterate/mapp"
)

func NewTransliterater(fromLang, toLang language.Script) mapp.Transliterater {
	t := mapp.Transliterater{}
	t.CodePointDiff = fromLang.UnicodeStart - toLang.UnicodeStart
	return t
}
