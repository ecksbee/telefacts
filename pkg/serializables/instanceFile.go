package serializables

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

type InstanceFile struct {
	XMLName   xml.Name   `xml:"xbrl"`
	XMLAttrs  []xml.Attr `xml:",any,attr"`
	SchemaRef []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
	} `xml:"schemaRef"`
	Context []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Entity   []struct {
			XMLName    xml.Name
			XMLAttrs   []xml.Attr `xml:",any,attr"`
			Identifier []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"identifier"`
			Segment []struct {
				XMLName        xml.Name
				XMLAttrs       []xml.Attr `xml:",any,attr"`
				ExplicitMember []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					CharData string     `xml:",chardata"`
				} `xml:"explicitMember"`
				TypedMember []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					XMLInner string     `xml:",innerxml"`
				} `xml:"typedMember"`
			} `xml:"segment"`
		} `xml:"entity"`
		Period []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			Instant  []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"instant"`
			StartDate []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"startDate"`
			EndDate []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"endDate"`
		} `xml:"period"`
		Scenario []struct {
			XMLName        xml.Name
			XMLAttrs       []xml.Attr `xml:",any,attr"`
			ExplicitMember []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				CharData string     `xml:",chardata"`
			} `xml:"explicitMember"`
			TypedMember []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				XMLInner string     `xml:",innerxml"`
			} `xml:"typedMember"`
		} `xml:"scenario"`
	} `xml:"context"`
	Unit []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Measure  []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			CharData string     `xml:",chardata"`
		} `xml:"measure"`
		Divide []struct {
			XMLName       xml.Name
			XMLAttrs      []xml.Attr `xml:",any,attr"`
			UnitNumerator []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				Measure  []struct {
					XMLName  xml.Name
					CharData string `xml:",chardata"`
				} `xml:"measure"`
			} `xml:"unitNumerator"`
			UnitDenominator []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				Measure  []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					CharData string     `xml:",chardata"`
				} `xml:"measure"`
			} `xml:"unitDenominator"`
		} `xml:"divide"`
	} `xml:"unit"`
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
