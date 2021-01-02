package xbrl

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html/charset"
)

type LabelLinkbase struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName xml.Name
		RoleURI string `xml:"roleURI,attr"`
		Href    string `xml:"href,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"roleRef"`
	LabelLink []struct {
		XMLName xml.Name
		Role    string `xml:"role,attr"`
		Type    string `xml:"type,attr"`

		Loc []struct {
			Href  string `xml:"href,attr"`
			Label string `xml:"label,attr"`
			Type  string `xml:"locator,attr"`
		} `xml:"loc"`
		ArcroleRef []struct {
			Href       string `xml:"href,attr"`
			Type       string `xml:"locator,attr"`
			ArcroleURI string `xml:"arcroleURI,attr"`
		} `xml:"arcroleRef"`
		LabelArc []struct {
			XMLName xml.Name
			Arcrole string `xml:"arcrole,attr"`
			Type    string `xml:"type,attr"`
			From    string `xml:"from,attr"`
			To      string `xml:"to,attr"`
		} `xml:"labelArc"`
	} `xml:"labelLink"`
}

func ReadLabelLinkbase(file os.FileInfo, workingDir string) (*LabelLinkbase, error) {
	filepath := path.Join(workingDir, file.Name())
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeLabelLinkbase(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodeLabelLinkbase(xmlData []byte) (*LabelLinkbase, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := LabelLinkbase{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}
