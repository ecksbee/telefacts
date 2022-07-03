package telefacts_test

import (
	zipPkg "archive/zip"
	bytesPkg "bytes"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"ecksbee.com/telefacts/pkg/serializables"
)

func TestDecode(t *testing.T) {
	workingDir := path.Join(".", "data", "folders", "test_ix_extract")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		os.MkdirAll(workingDir, fs.FileMode(0700))
	}
	defer func() {
		os.RemoveAll(workingDir)
	}()
	zipFile := path.Join(".", "data", "test_ix_extract.zip")
	err = unZipTestData(workingDir, zipFile)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
		return
	}
	sourceFilePath := "cmg-20200331x10q.htm"
	filepath := path.Join(workingDir, sourceFilePath)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
		return
	}
	doc := serializables.DecodeIxbrlFile(data)
	if doc == nil {
		t.Fatalf("Error: failed to decode IXBRL source document")
		return
	}
	if len(doc.SchemaRefs) != 1 {
		t.Fatalf("expected 1 schemaRef; outcome %d;\n", len(doc.SchemaRefs))
	}
}

func bencmarkExtract(targetPath string, doc serializables.Document, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := doc.Extract(targetPath)
		if err != nil {
			b.Fatalf("Error: " + err.Error())
		}
	}
}

func BenchmarkExtract_Ix_Narrative(b *testing.B) {
	serializables.INDENT = true
	workingDir := path.Join(".", "data", "folders", "test_ix_extract")
	serializables.XSLT_NRTV = path.Join(".", "XSLT_NRTV.xslt")
	_, err := os.Stat(workingDir)
	if os.IsNotExist(err) {
		os.MkdirAll(workingDir, fs.FileMode(0700))
	}
	defer func() {
		os.RemoveAll(workingDir)
	}()
	zipFile := path.Join(".", "data", "test_ix_extract.zip")
	err = unZipTestData(workingDir, zipFile)
	if err != nil {
		b.Fatalf("Error: " + err.Error())
		return
	}
	sourceFilePath := "cmg-20200331x10q.htm"
	filepath := path.Join(workingDir, sourceFilePath)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		b.Fatalf("Error: " + err.Error())
		return
	}
	doc := serializables.DecodeIxbrlFile(data)
	if doc == nil {
		b.Fatalf("Error: failed to decode IXBRL source document")
		return
	}
	targetpath := path.Join(workingDir, "test_ix_extracted.xml")
	bencmarkExtract(targetpath, *doc, b)
}

func unzip(bytes []byte) ([]*zipPkg.File, error) {
	bytesReader := bytesPkg.NewReader(bytes)
	zipReader, err := zipPkg.NewReader(bytesReader, bytesReader.Size())
	if err != nil {
		return nil, err
	}
	return zipReader.File, nil
}

func unzipFile(unzipFile *zipPkg.File) ([]byte, error) {
	rc, err := unzipFile.Open()
	defer rc.Close()
	if err != nil {
		return nil, err
	}
	var buffer bytesPkg.Buffer
	_, err = io.Copy(&buffer, rc)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func writeFile(dest string, data []byte) error {
	dirString, _ := path.Split(dest)
	_, err := os.Stat(dirString)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirString, 0755)
		if err != nil {
			return err
		}
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

func unZipTestData(workingDir string, zipFile string) error {
	zipBytes, err := ioutil.ReadFile(zipFile)
	if err != nil {
		return err
	}
	unzippeds, err := unzip(zipBytes)
	if err != nil {
		return err
	}
	for _, unzipped := range unzippeds {
		unzipBytes, err := unzipFile(unzipped)
		if err != nil {
			return err
		}
		dest := path.Join(workingDir, unzipped.Name)
		err = writeFile(dest, unzipBytes)
		if err != nil {
			return err
		}
	}
	return nil
}
