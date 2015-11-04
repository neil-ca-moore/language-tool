//Copyright (c) 2015, Neil Moore
//See LICENSE.txt for license information.

package formats

import (
	"os"
	"path"
)

type Folder struct {
}

func (f Folder) MakeWith(str string, loc string) (string, error) {
	newPath := path.Join(loc, str)
	return newPath, os.MkdirAll(newPath, os.ModePerm)
}
