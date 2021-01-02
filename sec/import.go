package sec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"

	"ecks-bee.com/telefacts/xbrl"
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

func (p *SECProject) Import(filingURL string, workingDir string, importTaxonomies bool) error {
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
	schema, err := scrapeSchemaFromSEC(filingURL, schemaItem)
	if err != nil {
		return err
	}
	if importTaxonomies {
		go xbrl.ImportTaxonomies(schema)
	}
	instanceItem, err := getInstanceFromFilingItems(items, ticker)
	if err != nil {
		return err
	}
	instance, err := scrapeInstanceFromSEC(filingURL, instanceItem)
	if err != nil {
		return err
	}
	preItem, err := getPresentationLinkbaseFromFilingItems(items, ticker)
	if err != nil {
		return err
	}
	presentation, err := scrapePresentationLinkbaseFromSEC(filingURL, preItem)
	if err != nil {
		return err
	}
	defItem, err := getDefinitionLinkbaseFromFilingItems(items, ticker)
	if err != nil {
		return err
	}
	definition, err := scrapeDefinitionLinkbaseFromSEC(filingURL, defItem)
	if err != nil {
		return err
	}
	calItem, err := getCalculationLinkbaseFromFilingItems(items, ticker)
	if err != nil {
		return err
	}
	calculation, err := scrapeCalculationLinkbaseFromSEC(filingURL, calItem)
	if err != nil {
		return err
	}
	labItem, err := getLabelLinkbaseFromFilingItems(items, ticker)
	if err != nil {
		return err
	}
	label, err := scrapeLabelLinkbaseFromSEC(filingURL, labItem)
	if err != nil {
		return err
	}
	wg := new(sync.WaitGroup)
	wg.Add(6)
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		defer wg.Done()
		dest := path.Join(workingDir, schemaItem.Name)
		err = commitSchema(dest, schema)
	}()
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		defer wg.Done()
		dest := path.Join(workingDir, instanceItem.Name)
		err = commitInstance(dest, instance)
	}()
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		defer wg.Done()
		dest := path.Join(workingDir, preItem.Name)
		err = commitPresentationLinkbase(dest, presentation)
	}()
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		defer wg.Done()
		dest := path.Join(workingDir, defItem.Name)
		err = commitDefinitionLinkbase(dest, definition)
	}()
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		defer wg.Done()
		dest := path.Join(workingDir, calItem.Name)
		err = commitCalculationLinkbase(dest, calculation)
	}()
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		defer wg.Done()
		dest := path.Join(workingDir, labItem.Name)
		err = commitLabelLinkbase(dest, label)
	}()
	wg.Wait()
	if err != nil {
		return err
	}

	meta := path.Join(workingDir, "_")
	file, _ := os.OpenFile(meta, os.O_CREATE, 0755)
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(struct {
		SEC string
	}{
		SEC: filingURL,
	})
}
