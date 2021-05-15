package sec

import (
	"archive/zip"
	"mime/multipart"
	"path"
	"sync"

	"ecksbee.com/telefacts/internal/actions"
)

func Upload(zipFile multipart.File, header multipart.FileHeader, destination string, importTaxonomies bool) error {
	zipReader, err := zip.NewReader(zipFile, header.Size)
	if err != nil {
		return err
	}
	workingDir := destination
	unzipFiles := zipReader.File
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		defer wg.Done()
		instance, err := getInstanceFromUnzipfiles(unzipFiles)
		if err != nil {
			return
		}
		unmarshalled, err := actions.Unzip(instance)
		if err != nil {
			return
		}
		err = actions.WriteFile(path.Join(workingDir, instance.Name), unmarshalled)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		schema, err := getSchemaFromUnzipfiles(unzipFiles)
		if err != nil {
			return
		}
		unmarshalled, err := actions.Unzip(schema)
		if err != nil {
			return
		}
		err = actions.WriteFile(path.Join(workingDir, schema.Name), unmarshalled)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		pre, err := getPresentationLinkbaseFromUnzipfiles(unzipFiles)
		if err != nil {
			return
		}
		unmarshalled, err := actions.Unzip(pre)
		if err != nil {
			return
		}
		err = actions.WriteFile(path.Join(workingDir, pre.Name), unmarshalled)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		def, err := getDefinitionLinkbaseFromUnzipfiles(unzipFiles)
		if err != nil {
			return
		}
		unmarshalled, err := actions.Unzip(def)
		if err != nil {
			return
		}
		err = actions.WriteFile(path.Join(workingDir, def.Name), unmarshalled)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		cal, err := getCalculationLinkbaseFromUnzipfiles(unzipFiles)
		if err != nil {
			return
		}
		unmarshalled, err := actions.Unzip(cal)
		if err != nil {
			return
		}
		err = actions.WriteFile(path.Join(workingDir, cal.Name), unmarshalled)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		lab, err := getLabelLinkbaseFromUnzipfiles(unzipFiles)
		if err != nil {
			return
		}
		unmarshalled, err := actions.Unzip(lab)
		if err != nil {
			return
		}
		err = actions.WriteFile(path.Join(workingDir, lab.Name), unmarshalled)
		if err != nil {
			return
		}
	}()
	wg.Wait()
	return err
}
