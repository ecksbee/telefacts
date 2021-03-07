package sec

import (
	"io/ioutil"

	"ecksbee.com/telefacts/pkg/serializables"
)

func Folder(workingDir string) (*serializables.Folder, error) {
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return nil, err
	}
	insFile, err := getInstanceFromOSfiles(files)
	if err != nil {
		return nil, err
	}
	entryFileName := insFile.Name()
	return serializables.Discover(workingDir, entryFileName)
}
