package serializables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/google/uuid"
)

type Underscore struct {
	Entry string
	Note  string
}

func GetEntryFileName(id string) (string, error) {
	underscore := Underscore{}
	b, err := ioutil.ReadFile(path.Join(WorkingDirectoryPath, "folders", id, "_"))
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(b, &underscore)
	return underscore.Entry, err
}

func NewFolder(key string, underscore Underscore) (string, error) {
	if WorkingDirectoryPath == "" {
		return "", fmt.Errorf("empty WorkingDirectoryPath")
	}
	telefactsId := func() uuid.UUID {
		if key == "" {
			return uuid.New()
		}
		var NilUuid uuid.UUID
		bytes := []byte(key)
		return uuid.NewMD5(NilUuid, bytes)
	}
	workingDir := path.Join(WorkingDirectoryPath, "folders")
	id := telefactsId()
	pathStr := path.Join(workingDir, id.String())
	_, err := os.Stat(pathStr)
	for os.IsExist(err) {
		if key != "" {
			return id.String(), err
		}
		id = telefactsId()
		pathStr = path.Join(workingDir, id.String())
		_, err = os.Stat(pathStr)
	}
	err = os.Mkdir(pathStr, 0755)
	if err != nil {
		return "", err
	}
	meta := path.Join(pathStr, "_")
	file, _ := os.OpenFile(meta, os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(underscore)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
