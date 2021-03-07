package serializables

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

type PresentationLinkbaseFile struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"roleRef"`
	PresentationLink []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Loc      []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"loc"`
		PresentationArc []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"presentationArc"`
	} `xml:"presentationLink"`
}

func DecodePresentationLinkbaseFile(xmlData []byte) (*PresentationLinkbaseFile, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := PresentationLinkbaseFile{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func ReadPresentationLinkbaseFile(filepath string) (*PresentationLinkbaseFile, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodePresentationLinkbaseFile(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
