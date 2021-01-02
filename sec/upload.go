package sec

import (
	"archive/zip"
	"mime/multipart"
	"path"
	"sync"

	"ecks-bee.com/telefacts/xbrl"
)

func (p *SECProject) Upload(zipFile multipart.File, header multipart.FileHeader, destination string, importTaxonomies bool) (int, error) {
	zipReader, err := zip.NewReader(zipFile, header.Size)
	if err != nil {
		return 0, err
	}
	unzipFiles := zipReader.File
	instance, err := getInstanceFromUnzipfiles(unzipFiles)
	if err != nil {
		return 0, err
	}
	schema, err := getSchemaFromUnzipfiles(unzipFiles)
	if err != nil {
		return 0, err
	}
	pre, err := getPresentationLinkbaseFromUnzipfiles(unzipFiles)
	if err != nil {
		return 0, err
	}
	def, err := getDefinitionLinkbaseFromUnzipfiles(unzipFiles)
	if err != nil {
		return 0, err
	}
	cal, err := getCalculationLinkbaseFromUnzipfiles(unzipFiles)
	if err != nil {
		return 0, err
	}
	lab, err := getLabelLinkbaseFromUnzipfiles(unzipFiles)
	if err != nil {
		return 0, err
	}
	return p.processClassicBundle(destination, instance, schema,
		pre, def, cal, lab, importTaxonomies)
}

func (p *SECProject) processClassicBundle(workingDir string, instance *zip.File, schema *zip.File,
	presentation *zip.File, definition *zip.File,
	calculation *zip.File, label *zip.File, importTaxonomies bool) (int, error) {
	proc := 0
	var err error
	wg := new(sync.WaitGroup)
	wg.Add(6)
	p.Lock.Lock()
	go func() {
		defer wg.Done()
		unmarshalled, err := unzipInstance(instance)
		if err != nil {
			return
		}
		err = commitInstance(path.Join(workingDir, instance.Name), unmarshalled)
		if err != nil {
			return
		}
		proc++
	}()
	go func() {
		defer wg.Done()
		unmarshalled, err := unzipSchema(schema)
		if err != nil {
			return
		}
		err = commitSchema(path.Join(workingDir, schema.Name), unmarshalled)
		if err != nil {
			return
		}
		if importTaxonomies {
			go xbrl.ImportTaxonomies(unmarshalled)
		}
		proc++
	}()
	go func() {
		defer wg.Done()
		unmarshalled, err := unzipPresentationLinkbase(presentation)
		if err != nil {
			return
		}
		err = commitPresentationLinkbase(path.Join(workingDir, presentation.Name), unmarshalled)
		if err != nil {
			return
		}
		proc++
	}()
	go func() {
		defer wg.Done()
		unmarshalled, err := unzipDefinitionLinkbase(definition)
		if err != nil {
			return
		}
		err = commitDefinitionLinkbase(path.Join(workingDir, definition.Name), unmarshalled)
		if err != nil {
			return
		}
		proc++
	}()
	go func() {
		defer wg.Done()
		unmarshalled, err := unzipCalculationLinkbase(calculation)
		if err != nil {
			return
		}
		err = commitCalculationLinkbase(path.Join(workingDir, calculation.Name), unmarshalled)
		if err != nil {
			return
		}
		proc++
	}()
	go func() {
		defer wg.Done()
		unmarshalled, err := unzipLabelLinkbase(label)
		if err != nil {
			return
		}
		err = commitLabelLinkbase(path.Join(workingDir, label.Name), unmarshalled)
		if err != nil {
			return
		}
		proc++
	}()
	wg.Wait()
	p.Lock.Unlock()
	return proc, err
}
