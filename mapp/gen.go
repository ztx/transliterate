package mapp

import "unicode/utf8"

// import "fmt"

func RuneReader(runes ...rune) <-chan rune {
	runeChan := make(chan rune)
	go func() {
		for _, r := range runes {
			runeChan <- r
		}
		close(runeChan)
	}()

	return runeChan
}

type RuneMapper interface {
	To(rune) <-chan rune
}

type Transliterater struct {
	//special single rune in devanagari which result in
	//multiple runes in kannda
	//eg: ॐ
	composites map[rune]string

	//diff b/w normal code points
	codePointDiff rune
}

var DevanagariToKannada = Transliterater{
	composites:    map[rune]string{'ॐ': "ಓಂ"},
	codePointDiff: 896}

var KannadaToDevanagari = Transliterater{
	composites:    make(map[rune]string),
	codePointDiff: -896}

func (t Transliterater) To(c rune) <-chan rune {
	//if a single rune results in a string after transliterate
	if key, ok := t.composites[c]; ok {
		return RuneReader([]rune(key)...)
	}

	var r rune
	r = c + t.codePointDiff
	if !utf8.ValidRune(r) {
		r = c
	}

	return RuneReader(r)
}
