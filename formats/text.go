//Copyright (c) 2015, Neil Moore
//See LICENSE.txt for license information.

package formats

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Text struct {
}

func (f Text) MakeWith(str string, loc string) (string, error) {
	newPath := filepath.Join(loc, str+".txt")
	return newPath, ioutil.WriteFile(newPath, []byte("hello"), os.ModePerm)
}
