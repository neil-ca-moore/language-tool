package formats

import (
	"os"
	"path"
)

type Folder struct {
}

func (f Folder) MakeWith(str string, loc string) error {
	return os.MkdirAll(path.Join(loc, str), os.ModePerm)
}
