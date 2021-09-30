package serializables

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type IxbrlNonfraction struct {
	XMLName  xml.Name
	XMLAttrs []xml.Attr `xml:",any,attr"`
	XMLInner string     `xml:",innerxml"`
}

type IxbrlNonnumeric struct {
	XMLName  xml.Name
	XMLAttrs []xml.Attr `xml:",any,attr"`
	XMLInner string     `xml:",innerxml"`
}

type IxbrlFootnote struct {
	XMLName  xml.Name
	XMLAttrs []xml.Attr `xml:",any,attr"`
	XMLInner string     `xml:",innerxml"`
}

type IxReferences struct {
	SchemaRef   []string
	LinkbaseRef []string
}

type IxResources struct {
	Contexts []CommonContext
	Units    []CommonUnit
}

type IxHiddenFacts struct {
	Nonfractions []IxbrlNonfraction
	Nonnumerics  []IxbrlNonnumeric
	Footnotes    []IxbrlFootnote
}

type IxRenderedFacts struct {
	Nonfractions []IxbrlNonfraction
	Nonnumerics  []IxbrlNonnumeric
	Footnotes    []IxbrlFootnote
}

type IxbrlHeader struct {
	References IxReferences
	Resources  IxResources
	Hidden     IxHiddenFacts
}

type IxbrlFile struct {
	Header        IxbrlHeader
	RenderedFacts IxRenderedFacts
}

func ReadIxbrlFile(filepath string) (*IxbrlFile, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, fmt.Errorf("reading ixbrl file failed")
	}
	return DecodeIxbrl(data)
}

func (folder *Folder) inlineSchemaRef(file *IxbrlFile) {
	fmt.Println("//todo do something")
}
