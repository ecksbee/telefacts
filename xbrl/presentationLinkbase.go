package xbrl

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html/charset"
)

const PresentationArcrole = `http://www.xbrl.org/2003/arcrole/parent-child`

type PresentationLink struct {
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
	PresentationArcs []Arc `xml:"presentationArc"`
}

type PresentationLinkbase struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName xml.Name
		RoleURI string `xml:"roleURI,attr"`
		Href    string `xml:"href,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"roleRef"`
	PresentationLinks []PresentationLink `xml:"presentationLink"`
}

func ReadPresentationLinkbase(file os.FileInfo, workingDir string) (*PresentationLinkbase, error) {
	filepath := path.Join(workingDir, file.Name())
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodePresentationLinkbase(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodePresentationLinkbase(xmlData []byte) (*PresentationLinkbase, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := PresentationLinkbase{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}
