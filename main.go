package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode"

	"github.com/rjeczalik/notify"

	"github.com/neil-ca-moore/language-toollanguage-tool/formats"
	"github.com/neil-ca-moore/language-toollanguage-tool/replacer"
	"github.com/neil-ca-moore/language-toollanguage-tool/strings"
)

func allClasses() []strings.Class {
	var res []strings.Class = make([]strings.Class, 0)
	res = append(res, strings.NewNaughty())

	runeSets := map[string][]rune{
		"emoji":    []rune(strings.Emoji),
		"NFC":      []rune(strings.AccentedNFC()),
		"NFD":      []rune(strings.AccentedNFD()),
		"NFKC":     []rune(strings.AccentedNFKC()),
		"NFKD":     []rune(strings.AccentedNFKD()),
		"katakana": strings.AllRunes(unicode.Katakana),
		"khmer":    strings.AllRunes(unicode.Khmer),
		"Nl":       strings.AllRunes(unicode.Nl),
		"Lo":       strings.AllRunes(unicode.Lo),
		"Sc":       strings.AllRunes(unicode.Sc),
		"Greek":    strings.AllRunes(unicode.Greek),
		"Linear-B": strings.AllRunes(unicode.Linear_B),
		"Arabic":   strings.AllRunes(unicode.Arabic),
		"Symb":     strings.AllRunes(unicode.Symbol),
	}

	for tag, runes := range runeSets {
		res = append(res, strings.NewRandomPicker(tag, 10, runes))
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
		go setUpFolderListener(absPath, stopChan)
		fmt.Println("Press return key to stop")
		var s string
		fmt.Scanf("%s\n", s)
		stopChan <- true
	} else if *help {
		flag.PrintDefaults()
	} else {
		flag.PrintDefaults()
	}
}
