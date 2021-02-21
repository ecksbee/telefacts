package actions

import (
	"archive/zip"
	"bytes"
	"io"
)

func Unzip(unzipFile *zip.File) ([]byte, error) {
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
	return buffer.Bytes(), nil
}
