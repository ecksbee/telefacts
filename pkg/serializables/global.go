package serializables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
)

var (
	WorkingDirectoryPath  string
	GlobalTaxonomySetPath string
)

func DiscoverFundamentalSchema() (*SchemaFile, error) {
	data, err := DiscoverGlobalFile(attr.LRR)
	if err != nil {
		return nil, err
	}
	return DecodeSchemaFile(data)
}

func UrlToFilename(urlStr string) (string, error) {
	urlPath, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	if len(urlPath.Scheme) <= 0 {
		return "", fmt.Errorf("empty scheme")
	}
	dest := urlPath.Scheme
	hostname := urlPath.Hostname()
	if len(hostname) <= 0 {
		return "", fmt.Errorf("empty hostname")
	}
	dest = path.Join(dest, hostname)
	var splits = strings.Split(urlPath.Path, "/")
	for _, split := range splits {
		dest = path.Join(dest, split)
	}
	return path.Join(GlobalTaxonomySetPath, "concepts", dest), nil
}

func DiscoverGlobalFile(urlStr string) ([]byte, error) {
	dest, err := UrlToFilename(urlStr)
	if err != nil {
		return nil, err
	}
	var ret []byte
	if file, err := os.Stat(dest); err == nil {
		dirString, _ := path.Split(dest)
		filename := file.Name()
		filepath := path.Join(dirString, filename)
		ret, err = ioutil.ReadFile(filepath)
		if err == nil {
			return ret, nil
		}
	}
	return nil, err
}

func DiscoverGlobalSchema(urlStr string) (*SchemaFile, error) {
	bytes, err := DiscoverGlobalFile(urlStr)
	if err != nil {
		return nil, err
	}
	return DecodeSchemaFile(bytes)
}

func DiscoverEntityNames() (map[string]map[string]string, error) {
	filename := path.Join(WorkingDirectoryPath, "names.json")
	names := make(map[string]map[string]string)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return names, err
	}
	err = json.Unmarshal(b, &names)
	return names, err
}
