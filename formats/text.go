package formats

import (
	"io/ioutil"
	"os"
	"path"
)

type Text struct {
}

func (f Text) MakeWith(str string, loc string) (string, error) {
	newPath := path.Join(loc, str+".txt")
	return newPath, ioutil.WriteFile(newPath, []byte("hello"), os.ModePerm)
}
