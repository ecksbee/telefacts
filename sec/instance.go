package sec

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getInstanceFromOSfiles(files []os.FileInfo) (os.FileInfo, error) {
	var xmls []os.FileInfo
	for _, file := range files {
		s := file.Name()
		if file.IsDir() || filepath.Ext(s) != xmlExt {
			continue
		}
		if s[len(s)-len(preExt):len(s)] == preExt {
			continue
		}
		if s[len(s)-len(defExt):len(s)] == defExt {
			continue
		}
		if s[len(s)-len(calExt):len(s)] == calExt {
			continue
		}
		if s[len(s)-len(labExt):len(s)] == labExt {
			continue
		}
		xmls = append(xmls, file)
	}
	if len(xmls) <= 0 {
		return nil, fmt.Errorf("No instance found")
	}
	return xmls[0], nil
}

func getInstanceFromUnzipfiles(unzipFiles []*zip.File) (*zip.File, error) {
	var xmls []*zip.File
	for _, unzipFile := range unzipFiles {
		if filepath.Ext(unzipFile.Name) == xmlExt {
			s := unzipFile.Name
			if s[len(s)-len(preExt):len(s)] == preExt {
				continue
			}
			if s[len(s)-len(defExt):len(s)] == defExt {
				continue
			}
			if s[len(s)-len(calExt):len(s)] == calExt {
				continue
			}
			if s[len(s)-len(labExt):len(s)] == labExt {
				continue
			}
			xmls = append(xmls, unzipFile)
		}
	}
	if len(xmls) <= 0 {
		return nil, fmt.Errorf("No instance found")
	}
	return xmls[0], nil
}

func getInstanceFromFilingItems(filingItems []filingItem, ticker string) (*filingItem, error) {
	for _, f := range filingItems {
		s := f.Name
		ext := filepath.Ext(s)
		a := (ext == xmlExt && strings.Index(s, ticker) == 0)
		b := len(s) >= 8
		if b {
			longExt := s[len(s)-8:]
			b = longExt != preExt && longExt != defExt && longExt != calExt && longExt != labExt
		}
		if a && b {
			return &f, nil
		}
	}
	return nil, fmt.Errorf("Cannot identify a single instance")
}
