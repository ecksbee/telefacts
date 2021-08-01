package sec

import (
	"encoding/json"
	"io/ioutil"

	"ecksbee.com/telefacts/pkg/renderables"
)

func names() map[string]map[string]string {
	names := make(map[string]map[string]string)
	b, err := ioutil.ReadFile("/names.json")
	if err != nil {
		return nil
	}
	err = json.Unmarshal(b, &names)
	if err != nil {
		return nil
	}
	return names
}

func MarshalCatalog(workingDir string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	filenames := []string{}
	return renderables.MarshalCatalog(h, names(), filenames)
}

func Marshal(workingDir string, slug string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	hash := slug
	return renderables.MarshalRenderable(hash, names(), h)
}
