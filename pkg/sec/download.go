package sec

import (
	"archive/zip"
	"bytes"

	"ecksbee.com/telefacts/internal/actions"
)

func Download(workingDir string) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	files, err := GetOSFiles(workingDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		data, err := actions.ReadFile(workingDir, &file)
		if err != nil {
			return nil, err
		}
		f, err := writer.Create(file.Name())
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
