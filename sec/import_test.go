package sec_test

import (
	"os"
	"path"
	"testing"

	"ecks-bee.com/telefacts/sec"
	"ecks-bee.com/telefacts/server"
	"github.com/google/uuid"
)

func TestImport(t *testing.T) {
	workingDir := path.Join("..", "projects")
	id := uuid.New()
	pathStr := path.Join(workingDir, "test_"+id.String())
	_, err := os.Stat(pathStr)
	for os.IsExist(err) {
		id = uuid.New()
		pathStr = path.Join(workingDir, "test_"+id.String())
		_, err = os.Stat(pathStr)
	}
	err = os.Mkdir(pathStr, 0755)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	appcache := server.NewCache()
	secProject := sec.SECProject{
		ID:       id.String(),
		AppCache: appcache,
	}
	err = secProject.Import("https://www.sec.gov/Archives/edgar/data/843006/000165495420001999", pathStr, true)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}

func TestImport_Large(t *testing.T) {
	workingDir := path.Join("..", "projects")
	id := uuid.New()
	pathStr := path.Join(workingDir, "test_"+id.String())
	_, err := os.Stat(pathStr)
	for os.IsExist(err) {
		id = uuid.New()
		pathStr = path.Join(workingDir, "test_"+id.String())
		_, err = os.Stat(pathStr)
	}
	err = os.Mkdir(pathStr, 0755)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	appcache := server.NewCache()
	secProject := sec.SECProject{
		ID:       id.String(),
		AppCache: appcache,
	}
	err = secProject.Import("https://www.sec.gov/Archives/edgar/data/69891/000143774920014395", pathStr, true)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}

func TestImport_Gold(t *testing.T) {
	workingDir := path.Join("..", "projects")
	id := uuid.New()
	pathStr := path.Join(workingDir, "test_"+id.String())
	_, err := os.Stat(pathStr)
	for os.IsExist(err) {
		id = uuid.New()
		pathStr = path.Join(workingDir, "test_"+id.String())
		_, err = os.Stat(pathStr)
	}
	err = os.Mkdir(pathStr, 0755)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
	appcache := server.NewCache()
	secProject := sec.SECProject{
		ID:       id.String(),
		AppCache: appcache,
	}
	err = secProject.Import("https://www.sec.gov/Archives/edgar/data/1445305/000144530520000124", pathStr, true)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}
