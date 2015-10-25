package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode"

	"github.com/rjeczalik/notify"

	"bigoh.co.uk/language-tool/formats"
	"bigoh.co.uk/language-tool/replacer"
	"bigoh.co.uk/language-tool/strings"
)

func allClasses() []strings.Class {
	var res []strings.Class = make([]strings.Class, 0)
	res = append(res, strings.NewNaughty())

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

	for _, runes := range runeSets {
		res = append(res, strings.NewRandomPicker(10, runes))
	}

	return res
}

func allFormats() []formats.Format {
	return []formats.Format{
		formats.Folder{},
		formats.Text{},
	}
}

func setUpFolderListener(path string, stop <-chan bool) {
	os.RemoveAll(path)
	os.Mkdir(path, os.ModePerm)

	replacer := replacer.NewMakerRegistry(path)

	for _, format := range allFormats() {
		for _, class := range allClasses() {
			replacer.Add(format, class)
		}
	}

	c := make(chan notify.EventInfo, 1000)
	if err := notify.Watch(path, c, notify.Rename, notify.Remove); err != nil {
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

	<-stop
}

func main() {
	_ = "breakpoint"
	help := flag.Bool("help", false, "Print help")
	path := flag.String("path", "", "Path that is kept chock full of files and folders with nasty names")
	flag.Parse()

	if len(*path) != 0 {
		absPath, err := filepath.Abs(*path)
		if err != nil {
			fmt.Println("error: path ", path, " doesn't work: ", err)
		}
		stopChan := make(chan bool)
		setUpFolderListener(absPath, stopChan)
		fmt.Println("Press any key to stop")
		var anyStr string
		fmt.Scanln(anyStr)
		stopChan <- true
	} else if *help {
		flag.PrintDefaults()
	} else {
		flag.PrintDefaults()
	}
}
