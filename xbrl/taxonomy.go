package xbrl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

type Concept struct {
	XMLName           xml.Name
	ID                string `xml:"id,attr"`
	Name              string `xml:"name,attr"`
	Nillable          bool   `xml:"nillable,attr"`
	PeriodType        string `xml:"periodType,attr"`
	Balance           string `xml:"balance,attr"`
	Type              string `xml:"type,attr"`
	SubstitutionGroup string `xml:"substitutionGroup,attr"`
	Abstract          bool   `xml:"abstract,attr"`
}

var (
	lock     sync.RWMutex
	appcache *gocache.Cache
)

func InjectCache(c *gocache.Cache) {
	appcache = c
}

func ImportTaxonomy(urlStr string) error {
	if appcache == nil {
		return fmt.Errorf("No accessible cache")
	}
	dest, err := urlToFilename(urlStr)
	if err != nil {
		return err
	}
	dirString, _ := path.Split(dest)
	_, err = os.Stat(dirString)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirString, 0755)
		if err != nil {
			return err
		}
	}
	resp, err := http.Get(urlStr)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	byteArray := buffer.Bytes()
	schema, err := DecodeSchema(byteArray)
	if err != nil {
		return err
	}
	appcache.Set(dest, *schema, gocache.DefaultExpiration)
	_, err = file.Write(byteArray)
	if err != nil {
		return err
	}
	return err
}

func urlToFilename(urlStr string) (string, error) {
	urlPath, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	if len(urlPath.Scheme) <= 0 {
		return "", fmt.Errorf("Empty scheme")
	}
	dest := urlPath.Scheme
	hostname := urlPath.Hostname() //todo only import "trusted hostnames"
	if len(hostname) <= 0 {
		return "", fmt.Errorf("Empty hostname")
	}
	dest = path.Join(dest, hostname)
	var splits = strings.Split(urlPath.Path, "/")
	for _, split := range splits {
		dest = path.Join(dest, split)
	}
	return path.Join(".", "taxonomies", dest), nil
}

func GetGlobalSchema(filePath string) (*Schema, error) {
	lock.RLock()
	if x, found := appcache.Get(filePath); found {
		lock.RUnlock()
		ret := x.(Schema)
		return &ret, nil
	}
	osFile, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, err
	}
	if osFile.IsDir() {
		return nil, fmt.Errorf("Invalid schemaLoc format")
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	schema, err := DecodeSchema(data)
	if err != nil {
		return nil, err
	}
	lock.RUnlock()
	lock.Lock()
	defer lock.Unlock()
	appcache.Set(filePath, *schema, gocache.DefaultExpiration)
	return schema, nil
}
