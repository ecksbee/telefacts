package serializables

import (
	"encoding/xml"
	"path"
	"sync"

	"ecksbee.com/telefacts/pkg/attr"
)

type Folder struct {
	wLock                 sync.Mutex
	Dir                   string
	Document              *Document
	Namespaces            map[string]string
	Instances             map[string]InstanceFile
	Schemas               map[string]SchemaFile
	LabelLinkbases        map[string]LabelLinkbaseFile
	PresentationLinkbases map[string]PresentationLinkbaseFile
	DefinitionLinkbases   map[string]DefinitionLinkbaseFile
	CalculationLinkbases  map[string]CalculationLinkbaseFile
}

func Discover(id string) (*Folder, error) {
	entryFileName, err := GetEntryFileName(id)
	if err != nil {
		return nil, err
	}
	workingDir := path.Join(VolumePath, "folders", id)
	ret := &Folder{
		Dir:                   workingDir,
		Namespaces:            make(map[string]string),
		Instances:             make(map[string]InstanceFile),
		Schemas:               make(map[string]SchemaFile),
		LabelLinkbases:        make(map[string]LabelLinkbaseFile),
		PresentationLinkbases: make(map[string]PresentationLinkbaseFile),
		DefinitionLinkbases:   make(map[string]DefinitionLinkbaseFile),
		CalculationLinkbases:  make(map[string]CalculationLinkbaseFile),
	}
	filepath := path.Join(workingDir, entryFileName)
	instanceFile, err := ReadInstanceFile(filepath)
	if err != nil {
		return nil, err
	}
	ret.schemaRef(instanceFile)
	ret.wLock.Lock()
	defer ret.wLock.Unlock()
	ret.Instances[entryFileName] = *instanceFile
	doc, err := GetDocument(id)
	if err != nil {
		return nil, err
	}
	ret.Document = doc
	return ret, nil
}

func (folder *Folder) schemaRef(file *InstanceFile) {
	if file == nil {
		return
	}
	schemaRefs := file.SchemaRef
	var wg sync.WaitGroup
	wg.Add(len(schemaRefs))
	for _, iitem := range schemaRefs {
		go func(item struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr "xml:\",any,attr\""
		}) {
			defer wg.Done()
			if item.XMLName.Space != attr.LINK {
				return
			}
			hrefAttr := attr.FindAttr(item.XMLAttrs, "href")
			if hrefAttr == nil || hrefAttr.Value == "" || hrefAttr.Name.Space != attr.XLINK {
				return
			}
			if attr.IsValidUrl(hrefAttr.Value) {
				go DiscoverGlobalSchema(hrefAttr.Value)
				return
			}
			filepath := path.Join(folder.Dir, hrefAttr.Value)
			discoveredSchema, err := ReadSchemaFile(filepath)
			if err != nil {
				return
			}
			targetNS := attr.FindAttr(discoveredSchema.XMLAttrs, "targetNamespace")
			if targetNS == nil || targetNS.Value == "" {
				return
			}
			folder.wLock.Lock()
			folder.Namespaces[targetNS.Value] = hrefAttr.Value
			folder.wLock.Unlock()
			var wwg sync.WaitGroup
			wwg.Add(3)
			go func() {
				defer wwg.Done()
				folder.importSchema(discoveredSchema)
			}()
			go func() {
				defer wwg.Done()
				folder.includeSchema(discoveredSchema)
			}()
			go func() {
				defer wwg.Done()
				folder.linkbaseRefSchema(discoveredSchema)
			}()
			wwg.Wait()
			folder.wLock.Lock()
			folder.Schemas[hrefAttr.Value] = *discoveredSchema
			folder.wLock.Unlock()
		}(iitem)
	}
	wg.Wait()
}

func (folder *Folder) includeSchema(file *SchemaFile) {
	if file == nil {
		return
	}
	includes := file.Include
	var wg sync.WaitGroup
	wg.Add(len(includes))
	for _, iitem := range includes {
		go func(item struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr "xml:\",any,attr\""
		}) {
			defer wg.Done()
			if item.XMLName.Space != attr.XSD {
				return
			}
			schemaLocationAttr := attr.FindAttr(item.XMLAttrs, "schemaLocation")
			if schemaLocationAttr == nil || schemaLocationAttr.Value == "" {
				return
			}
			if attr.IsValidUrl(schemaLocationAttr.Value) {
				go DiscoverGlobalSchema(schemaLocationAttr.Value)
				return
			}
			filepath := path.Join(folder.Dir, schemaLocationAttr.Value)
			discoveredSchema, err := ReadSchemaFile(filepath)
			if err != nil {
				return
			}
			targetNS := attr.FindAttr(discoveredSchema.XMLAttrs, "targetNamespace")
			if targetNS == nil || targetNS.Value == "" {
				return
			}
			folder.wLock.Lock()
			folder.Namespaces[targetNS.Value] = schemaLocationAttr.Value
			folder.wLock.Unlock()
			var wwg sync.WaitGroup
			wwg.Add(3)
			go func() {
				defer wwg.Done()
				folder.importSchema(discoveredSchema)
			}()
			go func() {
				defer wwg.Done()
				folder.includeSchema(discoveredSchema)
			}()
			go func() {
				defer wwg.Done()
				folder.linkbaseRefSchema(discoveredSchema)
			}()
			wwg.Wait()
			folder.wLock.Lock()
			folder.Schemas[schemaLocationAttr.Value] = *discoveredSchema
			folder.wLock.Unlock()
		}(iitem)
	}
	wg.Wait()
}

func (folder *Folder) importSchema(file *SchemaFile) {
	if file == nil {
		return
	}
	imports := file.Import
	var wg sync.WaitGroup
	wg.Add(len(imports))
	for _, iitem := range imports {
		go func(item struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr "xml:\",any,attr\""
		}) {
			defer wg.Done()
			if item.XMLName.Space != attr.XSD {
				return
			}
			namespaceAttr := attr.FindAttr(item.XMLAttrs, "namespace")
			if namespaceAttr == nil || namespaceAttr.Value == "" {
				return
			}
			schemaLocationAttr := attr.FindAttr(item.XMLAttrs, "schemaLocation")
			if schemaLocationAttr == nil || schemaLocationAttr.Value == "" {
				return
			}
			folder.wLock.Lock()
			folder.Namespaces[namespaceAttr.Value] = schemaLocationAttr.Value
			folder.wLock.Unlock()
			if attr.IsValidUrl(schemaLocationAttr.Value) {
				go DiscoverGlobalSchema(schemaLocationAttr.Value)
				return
			}
			filepath := path.Join(folder.Dir, schemaLocationAttr.Value)
			discoveredSchema, err := ReadSchemaFile(filepath)
			if err != nil {
				return
			}
			var wwg sync.WaitGroup
			wwg.Add(3)
			go func() {
				defer wwg.Done()
				folder.importSchema(discoveredSchema)
			}()
			go func() {
				defer wwg.Done()
				folder.includeSchema(discoveredSchema)
			}()
			go func() {
				defer wwg.Done()
				folder.linkbaseRefSchema(discoveredSchema)
			}()
			wwg.Wait()
			folder.wLock.Lock()
			folder.Schemas[schemaLocationAttr.Value] = *discoveredSchema
			folder.wLock.Unlock()
		}(iitem)
	}
	wg.Wait()
}

func (folder *Folder) linkbaseRefSchema(file *SchemaFile) {
	if file == nil {
		return
	}
	var wg sync.WaitGroup
	for _, annotation := range file.Annotation {
		if annotation.XMLName.Space != attr.XSD {
			continue
		}
		for _, appinfo := range annotation.Appinfo {
			if appinfo.XMLName.Space != attr.XSD {
				continue
			}
			for _, iitem := range appinfo.LinkbaseRef {
				wg.Add(1)
				go func(item struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr "xml:\",any,attr\""
				}) {
					defer wg.Done()
					if item.XMLName.Space != attr.LINK {
						return
					}
					arcroleAttr := attr.FindAttr(item.XMLAttrs, "arcrole")
					if arcroleAttr == nil || arcroleAttr.Name.Space != attr.XLINK || arcroleAttr.Value != attr.LINKARCROLE {
						return
					}
					typeAttr := attr.FindAttr(item.XMLAttrs, "type")
					if typeAttr == nil || typeAttr.Name.Space != attr.XLINK || typeAttr.Value != "simple" {
						return
					}
					roleAttr := attr.FindAttr(item.XMLAttrs, "role")
					if roleAttr == nil || roleAttr.Name.Space != attr.XLINK || roleAttr.Value == "" {
						return
					}
					hrefAttr := attr.FindAttr(item.XMLAttrs, "href")
					if hrefAttr == nil || hrefAttr.Name.Space != attr.XLINK || hrefAttr.Value == "" {
						return
					}
					if attr.IsValidUrl(hrefAttr.Value) {
						go DiscoverGlobalFile(hrefAttr.Value)
						return
					}
					filepath := path.Join(folder.Dir, hrefAttr.Value)
					switch roleAttr.Value {
					case attr.PresentationLinkbaseRef:
						discoveredPre, err := ReadPresentationLinkbaseFile(filepath)
						if err != nil {
							return
						}
						folder.wLock.Lock()
						folder.PresentationLinkbases[hrefAttr.Value] = *discoveredPre
						folder.wLock.Unlock()
						break
					case attr.DefinitionLinkbaseRef:
						discoveredDef, err := ReadDefinitionLinkbaseFile(filepath)
						if err != nil {
							return
						}
						folder.wLock.Lock()
						folder.DefinitionLinkbases[hrefAttr.Value] = *discoveredDef
						folder.wLock.Unlock()
						break
					case attr.CalculationLinkbaseRef:
						discoveredCal, err := ReadCalculationLinkbaseFile(filepath)
						if err != nil {
							return
						}
						folder.wLock.Lock()
						folder.CalculationLinkbases[hrefAttr.Value] = *discoveredCal
						folder.wLock.Unlock()
						break
					case attr.LabelLinkbaseRef:
						discoveredLab, err := ReadLabelLinkbaseFile(filepath)
						if err != nil {
							return
						}
						folder.wLock.Lock()
						folder.LabelLinkbases[hrefAttr.Value] = *discoveredLab
						folder.wLock.Unlock()
						break
					default:
						break
					}
				}(iitem)
			}
		}
	}
	wg.Wait()
}
