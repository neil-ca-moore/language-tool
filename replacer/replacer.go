package replacer

import (
	"path/filepath"

	"github.com/neil-ca-moore/language-toollanguage-tool/formats"
	"github.com/neil-ca-moore/language-toollanguage-tool/strings"
)

type FormatAndClass struct {
	format formats.Format
	class  strings.Class
}

type MakerRegistry struct {
	existing   map[string]FormatAndClass
	rootFolder string
}

func NewMakerRegistry(rootFolder string) MakerRegistry {
	return MakerRegistry{
		existing:   make(map[string]FormatAndClass),
		rootFolder: rootFolder,
	}
}

func (m MakerRegistry) makeNew(maker FormatAndClass) {
	newPath, err := maker.format.MakeWith(maker.class.Make(), filepath.Join(m.rootFolder))
	if err == nil {
		m.existing[newPath] = maker
	}
}

func (m MakerRegistry) Replace(path string) {
	maker, ok := m.existing[path]
	if ok {
		delete(m.existing, path)
		m.makeNew(maker)
	}
}

func (m MakerRegistry) Add(format formats.Format, class strings.Class) {
	m.makeNew(FormatAndClass{format, class})
}
