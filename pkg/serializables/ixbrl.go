package serializables

import (
	"fmt"
	"io/ioutil"
)

type IxbrlFile struct {
	//todo add all sort of stuff
}

func ReadIxbrlFile(filepath string) (*IxbrlFile, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("//todo decode ixbrl file")
	}
	var decoded *IxbrlFile
	return decoded, nil
}

func (folder *Folder) inlineSchemaRef(file *IxbrlFile) {
	fmt.Println("//todo do something")
}
