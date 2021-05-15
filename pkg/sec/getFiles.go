package sec

import (
	"io/fs"
	"io/ioutil"
	"sync"
)

func GetOSFiles(workingDir string) ([]fs.FileInfo, error) {
	ret := make([]fs.FileInfo, 0)
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return ret, err
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(6)
	go func() {
		defer wg.Done()
		instance, err := getInstanceFromOSfiles(files)
		if err != nil {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		files = append(files, instance)
	}()
	go func() {
		defer wg.Done()
		schema, err := getSchemaFromOSfiles(files)
		if err != nil {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		files = append(files, schema)
	}()
	go func() {
		defer wg.Done()
		pre, err := getPresentationLinkbaseFromOSfiles(files)
		if err != nil {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		files = append(files, pre)
	}()
	go func() {
		defer wg.Done()
		def, err := getDefinitionLinkbaseFromOSfiles(files)
		if err != nil {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		files = append(files, def)
	}()
	go func() {
		defer wg.Done()
		cal, err := getCalculationLinkbaseFromOSfiles(files)
		if err != nil {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		files = append(files, cal)
	}()
	go func() {
		defer wg.Done()
		lab, err := getLabelLinkbaseFromOSfiles(files)
		if err != nil {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		files = append(files, lab)
	}()
	wg.Wait()
	return files, err
}
