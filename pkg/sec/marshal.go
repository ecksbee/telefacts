package sec

import (
	"ecksbee.com/telefacts/pkg/renderables"
)

func names() map[string]map[string]string {
	//todo fetch https://www.sec.gov/files/company_tickers.json
	//todo cache
	ret := make(map[string]map[string]string)
	cik := "http://www.sec.gov/CIK"
	ret[cik] = make(map[string]string)
	ret[cik]["0001445305"] = "WORKIVA INC"
	ret[cik]["0000069891"] = "FILER DIRECT CORP"
	ret[cik]["0000843006"] = "NATIONAL BEVERAGE CORP"
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
	//todo check if slug is passthrough
	hash := slug
	return renderables.MarshalRenderable(hash, names(), h)
}
