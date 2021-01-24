package xbrl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html/charset"
)

const xsd = "http://www.w3.org/2001/XMLSchema"
const link = "http://www.xbrl.org/2003/linkbase"

type Schema struct {
	XMLName         xml.Name   `xml:"schema"`
	XMLNS           string     `xml:"xmlns,attr,omitempty"`
	TargetNamespace string     `xml:"targetNamespace,attr"`
	XMLAttrs        []xml.Attr `xml:",any,attr"`
	Annotation      []struct {
		XMLName xml.Name
		Appinfo struct {
			XMLName  xml.Name
			RoleType []struct {
				XMLName    xml.Name
				RoleURI    string `xml:"roleURI,attr"`
				ID         string `xml:"id,attr"`
				Definition string `xml:"definition"`
				UsedOn     []struct {
					XMLName xml.Name
					Text    string `xml:",chardata"`
				} `xml:"usedOn"`
			} `xml:"roleType"`
		} `xml:"appinfo"`
	} `xml:"annotation"`
	Import []struct {
		XMLName        xml.Name
		Namespace      string `xml:"namespace,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"import"`
	Element []Concept `xml:"element"`
}

func ImportTaxonomies(schema *Schema) {
	imports := schema.Import
	for _, importItem := range imports {
		if len(importItem.Namespace) <= 0 {
			continue
		}
		if len(importItem.SchemaLocation) <= 0 {
			continue
		}
		ImportTaxonomy(importItem.SchemaLocation)
	}
}

func ReadSchema(file os.FileInfo, workingDir string) (*Schema, error) {
	filepath := path.Join(workingDir, file.Name())
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeSchema(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodeSchema(xmlData []byte) (*Schema, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := Schema{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func HashQuery(schema *Schema, query string) (string, *Concept, error) {
	i := strings.IndexRune(query, '#')
	if i < 0 {
		return "", nil, fmt.Errorf("Invalid query")
	}
	base := query[:i]
	if len(base) <= 0 {
		return "", nil, fmt.Errorf("Invalid base query")
	}
	fragment := query[i+1 : len(query)]
	if len(base) <= 0 {
		return "", nil, fmt.Errorf("Invalid query fragment")
	}
	var namespace string
	var concepts []Concept
	if strings.HasPrefix(base, "http") {
		filePath, err := urlToFilename(base)
		if err != nil {
			return "", nil, err
		}
		globalSchema, err := GetGlobalSchema(filePath)
		if err != nil {
			return "", nil, err
		}
		concepts = globalSchema.Element
		namespace = globalSchema.TargetNamespace
	} else {
		concepts = schema.Element
		namespace = schema.TargetNamespace
	}
	for _, candidate := range concepts {
		if fragment == candidate.ID {
			return namespace, &candidate, nil
		}
	}
	return "", nil, nil
}

func NameQuery(schema *Schema, namespace string, localName string) (string, *Concept, error) {
	var schemaLoc string
	var concepts []Concept
	if namespace == schema.TargetNamespace {
		concepts = schema.Element
		schemaLoc = ""
	} else {
		for _, importTag := range schema.Import {
			if importTag.Namespace == namespace {
				schemaLoc = importTag.SchemaLocation
				break
			}
		}
		if len(schemaLoc) <= 0 {
			return "", nil, fmt.Errorf("Invalid namespace")
		}
		filePath, err := urlToFilename(schemaLoc)
		if err != nil {
			return "", nil, err
		}
		globalSchema, err := GetGlobalSchema(filePath)
		if err != nil {
			return "", nil, err
		}
		concepts = globalSchema.Element
	}
	for _, candidate := range concepts {
		if localName == candidate.Name {
			return schemaLoc + "#" + candidate.ID, &candidate, nil
		}
	}
	return "", nil, nil
}
