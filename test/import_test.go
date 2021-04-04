package telefacts_test

import (
	"math/rand"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	"ecksbee.com/telefacts/pkg/sec"
	"github.com/google/uuid"
)

var (
	lock sync.Mutex
)

func TestImport(t *testing.T) {
	workingDir := path.Join(".", "data")
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
	defer os.RemoveAll(pathStr)
	loop := true
	dur := (time.Second * 1) + (time.Second * time.Duration(rand.Int31n(4)))
	for loop {
		select {
		case <-time.After(dur):
			lock.Lock()
			loop = false
		}
	}
	defer lock.Unlock()
	<-time.After(time.Second * 1)
	err = sec.Import("https://www.sec.gov/Archives/edgar/data/843006/000165495420001999", pathStr)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}

func TestImport_Large(t *testing.T) {
	workingDir := path.Join(".", "data")
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
	defer os.RemoveAll(pathStr)
	loop := true
	dur := (time.Second * 1) + (time.Second * time.Duration(rand.Int31n(4)))
	for loop {
		select {
		case <-time.After(dur):
			lock.Lock()
			loop = false
		}
	}
	defer lock.Unlock()
	<-time.After(time.Second * 1)
	err = sec.Import("https://www.sec.gov/Archives/edgar/data/69891/000143774920014395", pathStr)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}

func TestImport_Gold(t *testing.T) {
	workingDir := path.Join(".", "data")
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
	defer os.RemoveAll(pathStr)
	loop := true
	dur := (time.Second * 1) + (time.Second * time.Duration(rand.Int31n(4)))
	for loop {
		select {
		case <-time.After(dur):
			lock.Lock()
			loop = false
		}
	}
	defer lock.Unlock()
	<-time.After(time.Second * 1)
	err = sec.Import("https://www.sec.gov/Archives/edgar/data/1445305/000144530520000124", pathStr)
	if err != nil {
		t.Fatalf("Error: " + err.Error())
	}
}
