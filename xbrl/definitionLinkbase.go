package xbrl

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html/charset"
)

type DefinitionLinkbase struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName xml.Name
		RoleURI string `xml:"roleURI,attr"`
		Href    string `xml:"href,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"roleRef"`
	DefinitionLink []struct {
		XMLName xml.Name
		Role    string `xml:"role,attr"`
		Type    string `xml:"type,attr"`
		Title   string `xml:"title,attr"`
		Loc     []struct {
			Href  string `xml:"href,attr"`
			Label string `xml:"label,attr"`
			Type  string `xml:"locator,attr"`
		} `xml:"loc"`
		ArcroleRef []struct {
			Href       string `xml:"href,attr"`
			Type       string `xml:"locator,attr"`
			ArcroleURI string `xml:"arcroleURI,attr"`
		} `xml:"arcroleRef"`
		DefinitionArc []struct {
			XMLName xml.Name
			Order   string `xml:"order,attr"`
			Arcrole string `xml:"arcrole,attr"`
			Type    string `xml:"type,attr"`
			From    string `xml:"from,attr"`
			To      string `xml:"to,attr"`
		} `xml:"definitionArc"`
	} `xml:"definitionLink"`
}

func ReadDefinitionLinkbase(file os.FileInfo, workingDir string) (*DefinitionLinkbase, error) {
	filepath := path.Join(workingDir, file.Name())
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeDefinitionLinkbase(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodeDefinitionLinkbase(xmlData []byte) (*DefinitionLinkbase, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := DefinitionLinkbase{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}
