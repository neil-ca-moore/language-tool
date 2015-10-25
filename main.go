package main

import (
	"log"
	"unicode"

	"github.com/rjeczalik/notify"

	"bigoh.co.uk/language-tool/formats"
	"bigoh.co.uk/language-tool/replacer"
	"bigoh.co.uk/language-tool/strings"
)

func main() {
	rootFolder := "/Users/neilmoore67/tmp/foo"

	replacer := replacer.NewMakerRegistry(rootFolder)

	runeSets := [][]rune{
		[]rune(strings.Emoji),
		strings.AllRunes(unicode.Katakana),
		strings.AllRunes(unicode.Khmer),
		strings.AllRunes(unicode.Nl),
		strings.AllRunes(unicode.Lo),
		strings.AllRunes(unicode.Sc),
		strings.AllRunes(unicode.Greek),
		strings.AllRunes(unicode.Other_Grapheme_Extend),
		strings.AllRunes(unicode.Avestan),
	}

	formatSet := []formats.Format{
		formats.Folder{},
		formats.Text{},
	}

	for _, format := range formatSet {
		replacer.Add(format, strings.NewNaughty())
		for _, runes := range runeSets {
			replacer.Add(format, strings.NewRandomPicker(10, runes))
		}
	}

	done := make(chan bool)

	c := make(chan notify.EventInfo, 1000)
	if err := notify.Watch(rootFolder, c, notify.Rename, notify.Remove); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)
	go func() {
		for {
			select {
			case event := <-c:
				replacer.Replace(event.Path())
			}
		}
	}()

	<-done
}
