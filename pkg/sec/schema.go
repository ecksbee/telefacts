package sec

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
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
