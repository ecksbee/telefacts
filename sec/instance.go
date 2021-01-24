package sec

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"ecks-bee.com/telefacts/xbrl"
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

func unzipInstance(unzipFile *zip.File) (*xbrl.Instance, error) { //todo move to xbrl
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
	decoded, err := xbrl.DecodeInstance(buffer.Bytes())
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func commitInstance(dest string, instance *xbrl.Instance) error {
	data, err := xbrl.EncodeInstance(instance)
	if err != nil {
		return err
	}
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

func scrapeInstanceFromSEC(filingURL string, filingItem *filingItem) (*xbrl.Instance, error) {
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
	return xbrl.DecodeInstance(buffer.Bytes())
}
