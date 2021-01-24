package xbrl

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html/charset"
)

const CalculationArcrole = `http://www.xbrl.org/2003/arcrole/summation-item`

type CalculationLink struct {
	XMLName xml.Name
	Role    string `xml:"role,attr"`
	Type    string `xml:"type,attr"`
	Title   string `xml:"title,attr"`
	Loc     []struct {
		Href  string `xml:"href,attr"`
		Label string `xml:"label,attr"`
		Type  string `xml:"type,attr"`
	} `xml:"loc"`
	CalculationArc []struct {
		XMLName xml.Name
		Order   string `xml:"order,attr"`
		Weight  string `xml:"weight,attr"`
		Arcrole string `xml:"arcrole,attr"`
		Type    string `xml:"type,attr"`
		From    string `xml:"from,attr"`
		To      string `xml:"to,attr"`
	} `xml:"calculationArc"`
}

type CalculationLinkbase struct {
	XMLName  xml.Name   `xml:"linkbase"`
	XMLNS    string     `xml:"xmlns,attr,omitempty"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	RoleRef  []struct {
		XMLName xml.Name
		RoleURI string `xml:"roleURI,attr"`
		Href    string `xml:"href,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"roleRef"`
	CalculationLinks []CalculationLink `xml:"calculationLink"`
}

func ReadCalculationLinkbase(file os.FileInfo, workingDir string) (*CalculationLinkbase, error) {
	filepath := path.Join(workingDir, file.Name())
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeCalculationLinkbase(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodeCalculationLinkbase(xmlData []byte) (*CalculationLinkbase, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := CalculationLinkbase{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}
