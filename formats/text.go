package formats

import (
	"io/ioutil"
	"os"
	"path"
)

type Text struct {
}

func (f Text) MakeWith(str string, loc string) error {
	return ioutil.WriteFile(path.Join(loc, str+".txt"), []byte("hello"), os.ModePerm)
}
