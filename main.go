package main

import (
	"fmt"
	"unicode"

	"bigoh.co.uk/language-tool/formats"
	"bigoh.co.uk/language-tool/strings"
)

func main() {
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Katakana)))
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Khmer)))
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Imperial_Aramaic)))
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Lao)))
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Cc)))
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Nl)))
	fmt.Println(strings.RunesToString(strings.AllRunes(unicode.Other_Grapheme_Extend)))

	folder := formats.Folder{}
	fmt.Println(folder.MakeWith(strings.NewRandomPicker(10, strings.AllRunes(unicode.Katakana)).Make(), "/Users/neilmoore67/tmp"))

	text := formats.Text{}
	fmt.Println(text.MakeWith(strings.NewRandomPicker(10, strings.AllRunes(unicode.Nl)).Make(), "/Users/neilmoore67/tmp"))
	fmt.Println(text.MakeWith(strings.NewRandomPicker(10, []rune(strings.Emoji)).Make(), "/Users/neilmoore67/tmp"))
}
