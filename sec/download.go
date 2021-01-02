package sec

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"path"
)

func (p *SECProject) Download(workingDir string) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		filename := file.Name()
		ext := path.Ext(filename)
		if ext != ".xbrl" && ext != ".xml" && ext != ".xsd" {
			continue
		}
		filepath := path.Join(workingDir, filename)
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
		f, err := writer.Create(filename)
		if err != nil {
			return nil, err
		}
		_, err = f.Write([]byte(data))
		if err != nil {
			return nil, err
		}
	}
	err = writer.Close()
	return buf.Bytes(), err
}
