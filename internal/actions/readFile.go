package actions

import (
	"io/fs"
	"io/ioutil"
	"path"
)

func ReadFile(src string, file *fs.FileInfo) ([]byte, error) {
	filename := (*file).Name()
	filepath := path.Join(src, filename)
	return ioutil.ReadFile(filepath)
}
