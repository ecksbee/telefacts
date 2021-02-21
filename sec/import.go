package sec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"regexp"
	"strings"
	"sync"

	"ecksbee.com/telefacts/actions"
)

type filingItem struct {
	LastModified string `json:"last-modified"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         string `json:"size"`
}

const xmlExt = ".xml"
const xsdExt = ".xsd"
const preExt = "_pre.xml"
const defExt = "_def.xml"
const calExt = "_cal.xml"
const labExt = "_lab.xml"
const regexSEC = "https://www.sec.gov/Archives/edgar/data/([0-9]+)/([0-9]+)"

func Import(filingURL string, workingDir string) error {
	isSEC, _ := regexp.MatchString(regexSEC, filingURL)
	if !isSEC {
		return fmt.Errorf("Not an acceptable SEC address, " + filingURL)
	}
	resp, err := http.Get(filingURL + "/index.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	filing := struct {
		Directory struct {
			Item      []filingItem `json:"item"`
			Name      string       `json:"name"`
			ParentDir string       `json:"parent-dir"`
		} `json:"directory"`
	}{}
	json.Unmarshal(body, &filing)
	items := filing.Directory.Item
	if len(items) <= 0 {
		return fmt.Errorf("Empty filing at " + filingURL)
	}
	schemaItem, err := getSchemaFromFilingItems(items)
	if err != nil {
		return err
	}
	str := schemaItem.Name
	x := strings.Index(str, "-")
	ticker := str[:x]
	if len(ticker) <= 0 {
		return fmt.Errorf("Ticker symbol not found")
	}
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		schema, err := actions.Scrape(filingURL + "/" + schemaItem.Name)
		if err != nil {
			return
		}
		dest := path.Join(workingDir, schemaItem.Name)
		err = actions.Commit(dest, schema)
	}()
	go func() {
		defer wg.Done()
		instanceItem, err := getInstanceFromFilingItems(items, ticker)
		if err != nil {
			return
		}
		instance, err := actions.Scrape(filingURL + "/" + instanceItem.Name)
		if err != nil {
			return
		}
		dest := path.Join(workingDir, instanceItem.Name)
		err = actions.Commit(dest, instance)
	}()
	go func() {
		defer wg.Done()
		preItem, err := getPresentationLinkbaseFromFilingItems(items, ticker)
		if err != nil {
			return
		}
		presentation, err := actions.Scrape(filingURL + "/" + preItem.Name)
		if err != nil {
			return
		}
		dest := path.Join(workingDir, preItem.Name)
		err = actions.Commit(dest, presentation)
	}()
	go func() {
		defer wg.Done()
		defItem, err := getDefinitionLinkbaseFromFilingItems(items, ticker)
		if err != nil {
			return
		}
		definition, err := actions.Scrape(filingURL + "/" + defItem.Name)
		if err != nil {
			return
		}
		dest := path.Join(workingDir, defItem.Name)
		err = actions.Commit(dest, definition)
	}()
	go func() {
		defer wg.Done()
		calItem, err := getCalculationLinkbaseFromFilingItems(items, ticker)
		if err != nil {
			return
		}
		calculation, err := actions.Scrape(filingURL + "/" + calItem.Name)
		if err != nil {
			return
		}
		dest := path.Join(workingDir, calItem.Name)
		err = actions.Commit(dest, calculation)
	}()
	go func() {
		defer wg.Done()
		labItem, err := getLabelLinkbaseFromFilingItems(items, ticker)
		if err != nil {
			return
		}
		label, err := actions.Scrape(filingURL + "/" + labItem.Name)
		if err != nil {
			return
		}
		dest := path.Join(workingDir, labItem.Name)
		err = actions.Commit(dest, label)
	}()
	wg.Wait()
	return err
}
