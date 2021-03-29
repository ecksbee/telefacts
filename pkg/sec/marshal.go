package sec

import (
	"encoding/json"
	"fmt"

	"ecksbee.com/telefacts/internal/actions"
	"ecksbee.com/telefacts/pkg/renderables"
)

func names() map[string]map[string]string {
	ret := make(map[string]map[string]string)
	b, err := actions.Scrape(`https://www.sec.gov/files/company_tickers.json`)
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
