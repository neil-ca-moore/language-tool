package strings

import (
	"bytes"
	"unicode"
)

func AllRunes(rt *unicode.RangeTable) []rune {
	var runes []rune
	for _, aRange := range rt.R16 {
		for i := aRange.Lo; i < aRange.Hi; i += aRange.Stride {
			runes = append(runes, rune(i))
		}
	}
	for _, aRange := range rt.R32 {
		for i := aRange.Lo; i < aRange.Hi; i += aRange.Stride {
			runes = append(runes, rune(i))
		}
	}
	return runes
}

func RunesToString(runes []rune) string {
	var buffer bytes.Buffer
	for _, rune := range runes {
		buffer.WriteRune(rune)
	}
	return buffer.String()
}
