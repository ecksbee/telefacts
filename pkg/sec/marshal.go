package sec

import (
	"encoding/json"
	"fmt"

	"ecksbee.com/telefacts/internal/actions"
	"ecksbee.com/telefacts/pkg/renderables"
	gocache "github.com/patrickmn/go-cache"
)

func names() map[string]map[string]string {
	tickerURL := `https://www.sec.gov/files/company_tickers.json`
	ret := make(map[string]map[string]string)
	if x, found := appcache.Get(tickerURL); found {
		lock.RLock()
		defer lock.RUnlock()
		ret := x.(map[string]map[string]string)
		return ret
	}
	b, err := actions.Scrape(tickerURL)
	if err != nil {
		return ret
	}
	type SECTickers map[string]struct {
		CIK    int    `json:"cik_str"`
		Ticker string `json:"ticker"`
		Title  string `json:"title"`
	}
	var f SECTickers
	err = json.Unmarshal(b, &f)
	if err != nil {
		return ret
	}
	cik := "http://www.sec.gov/CIK"
	ret[cik] = make(map[string]string)
	for _, obj := range f {
		ret[cik][fmt.Sprintf("%010d", obj.CIK)] = obj.Title
	}
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appcache.Set(tickerURL, ret, gocache.DefaultExpiration)
	}()
	return ret
}

func MarshalCatalog(workingDir string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	//todo get files
	filenames := []string{}
	return renderables.MarshalCatalog(h, names(), filenames)
}

func Marshal(workingDir string, slug string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	//todo check if slug is passthrough to a file
	hash := slug
	return renderables.MarshalRenderable(hash, names(), h)
}
