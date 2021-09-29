package serializables

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"path"
	"sync"

	"ecksbee.com/telefacts/pkg/attr"
	"golang.org/x/net/html/charset"
)

type InstanceFile struct {
	XMLName   xml.Name   `xml:"xbrl"`
	XMLAttrs  []xml.Attr `xml:",any,attr"`
	SchemaRef []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"schemaRef"`
	Context      []CommonContext `xml:"context"`
	Unit         []CommonUnit    `xml:"unit"`
	FootnoteLink []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Footnote []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			CharData string     `xml:",chardata"`
		} `xml:"footnote"`
		Loc []struct {
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"loc"`
		FootnoteArc []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"footnoteArc"`
	} `xml:"footnoteLink"`
	Facts []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		XMLInner string     `xml:",innerxml"`
	} `xml:",any"`
}

func DecodeInstanceFile(xmlData []byte) (*InstanceFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := InstanceFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadInstanceFile(filepath string) (*InstanceFile, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeInstanceFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
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
