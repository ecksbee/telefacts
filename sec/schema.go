package sec

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"ecks-bee.com/telefacts/xbrl"
)

func getSchemaFromOSfiles(files []os.FileInfo) (os.FileInfo, error) {
	var xmls []os.FileInfo
	for _, file := range files {
		s := file.Name()
		if file.IsDir() || filepath.Ext(s) == "" {
			continue
		}
		if filepath.Ext(s) == xsdExt {
			xmls = append(xmls, file)
		}
	}
	if len(xmls) <= 0 {
		return nil, fmt.Errorf("No presentation linkbase found")
	}
	return xmls[0], nil
}

func getSchemaFromUnzipfiles(unzipFiles []*zip.File) (*zip.File, error) {
	var xsds []*zip.File
	for _, unzipFile := range unzipFiles {
		if filepath.Ext(unzipFile.Name) == xsdExt {
			xsds = append(xsds, unzipFile)
		}
	}
	if len(xsds) <= 0 {
		return nil, fmt.Errorf("No schema found")
	}
	return xsds[0], nil
}

func unzipSchema(unzipFile *zip.File) (*xbrl.Schema, error) {
	rc, err := unzipFile.Open()
	defer rc.Close()
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, rc)
	if err != nil {
		return nil, err
	}
	decoded, err := xbrl.DecodeSchema(buffer.Bytes())
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func commitSchema(dest string, schema *xbrl.Schema) error {
	data, _ := xml.Marshal(schema)
	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func getSchemaFromFilingItems(filingItems []filingItem) (*filingItem, error) {
	var candidates []filingItem
	for _, f := range filingItems {
		s := f.Name
		ext := filepath.Ext(s)
		if ext == xsdExt {
			candidates = append(candidates, f)
		}
	}
	if len(candidates) > 1 {
		return nil, fmt.Errorf("Cannot identify a single schema")
	}
	if len(candidates) <= 0 {
		return nil, fmt.Errorf("No schema found")
	}
	return &(candidates[0]), nil
}

func scrapeSchemaFromSEC(filingURL string, filingItem *filingItem) (*xbrl.Schema, error) {
	resp, err := http.Get(filingURL + "/" + filingItem.Name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		return nil, err
	}
	return xbrl.DecodeSchema(buffer.Bytes())
}
