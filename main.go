package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"golang.org/x/text/unicode/runenames"
)

var userPrefs = []language.Tag{
	language.Make("gsw"), // Swiss German
	language.Make("fr"),  // French
}

var serverLangs = []language.Tag{
	language.AmericanEnglish, // en-US fallback
	language.German,          // de
}

var matcher = language.NewMatcher(serverLangs)

func main() {
	tag, index, confidence := matcher.Match(userPrefs...)

	fmt.Printf("best match: %s (%s) index=%d confidence=%v\n",
		display.English.Tags().Name(tag),
		display.Self.Name(tag),
		index, confidence)
	// best match: German (Deutsch) index=1 confidence=High
	fmt.Println(language.Hindi, language.Kannada)

	r := '\U00000C90'
	hr := '\u0910'

	fmt.Printf("%08x %q\n", r, runenames.Name(r))
	fmt.Printf("%c %c %d", r, hr, r-hr)
	fmt.Printf("\n%c %c %d", hr+898, r, r-hr)

}
