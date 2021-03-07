package serializables

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"ecksbee.com/telefacts/internal/actions"
	"ecksbee.com/telefacts/pkg/attr"
	gocache "github.com/patrickmn/go-cache"
)

var (
	lock      sync.RWMutex
	appcache  *gocache.Cache
	globalDir string
)

func InjectCache(c *gocache.Cache) {
	appcache = c
}

func SetGlobalDir(dir string) {
	globalDir = dir
}

func DiscoverUnitTypeRegistry() (*UnitTypeRegistryFile, error) {
	data, err := DiscoverGlobalFile(attr.UTR)
	if err != nil {
		return nil, err
	}
	return DecodeUnitTypeRegistry(data)
}

func DiscoverFundamentalSchema() (*SchemaFile, error) {
	data, err := DiscoverGlobalFile(attr.LRR)
	if err != nil {
		return nil, err
	}
	return DecodeSchemaFile(data)
}

func ImportSchema(file *SchemaFile) map[string]string {
	ret := make(map[string]string)
	if file == nil {
		return ret
	}
	imports := file.Import
	var wg sync.WaitGroup
	for _, iitem := range imports {
		go func(item struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr "xml:\",any,attr\""
		}) {
			wg.Add(1)
			defer wg.Done()
			if item.XMLName.Space != attr.XSD {
				return
			}
			namespaceAttr := attr.FindAttr(item.XMLAttrs, "namespace")
			if namespaceAttr == nil || namespaceAttr.Value == "" {
				return
			}
			schemaLocationAttr := attr.FindAttr(item.XMLAttrs, "schemaLocation")
			if schemaLocationAttr == nil || schemaLocationAttr.Value == "" {
				return
			}
			go DiscoverGlobalFile(schemaLocationAttr.Value)
			ret[namespaceAttr.Value] = schemaLocationAttr.Value
		}(iitem)
	}
	wg.Wait()
	return ret
}

func urlToFilename(urlStr string) (string, error) {
	if globalDir == "" {
		return "", fmt.Errorf("No directory set")
	}
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
	return path.Join(globalDir, dest), nil
}

func DiscoverGlobalFile(urlStr string) ([]byte, error) {
	if appcache == nil {
		return nil, fmt.Errorf("No accessible cache")
	}
	lock.RLock()
	if x, found := appcache.Get(urlStr); found {
		ret := x.([]byte)
		lock.RUnlock()
		return ret, nil
	}
	lock.RUnlock()
	dest, err := urlToFilename(urlStr)
	if err != nil {
		return nil, err
	}
	dirString, _ := path.Split(dest)
	_, err = os.Stat(dirString)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirString, 0755)
		if err != nil {
			return nil, err
		}
	}
	ret, err := actions.Scrape(urlStr)
	if err != nil {
		return nil, err
	}
	err = actions.Commit(dest, ret)
	if err != nil {
		return nil, err
	}
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appcache.Set(urlStr, ret, gocache.DefaultExpiration)
	}()
	return ret, err
}

func DiscoverGlobalSchema(urlStr string) (*SchemaFile, error) {
	bytes, err := DiscoverGlobalFile(urlStr)
	if err != nil {
		return nil, err
	}
	return DecodeSchemaFile(bytes)
}
