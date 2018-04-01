package language

func isKannda(c rune) bool {
	if c >= '\u0C80' && c <= '\u0CF2' {
		return true
	}
	return false
}

func isDevanagari(c rune) bool {
	if c >= '\u0900' && c <= '\u097F' {
		return true
	}
	return false
}
