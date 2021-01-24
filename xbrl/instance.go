package xbrl

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html/charset"
)

const xbrli = "http://www.xbrl.org/2003/instance"

type Instance struct {
	XMLName   xml.Name   `xml:"xbrl"`
	XMLAttrs  []xml.Attr `xml:",any,attr"`
	SchemaRef []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		HRef     string     `xml:"href,attr"`
		Type     string     `xml:"type,attr"`
	} `xml:"schemaRef"`
	Context []Context `xml:"context"`
	Unit    []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		ID       string     `xml:"id,attr"`
		Measure  []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			XMLInner string     `xml:",innerxml"`
		} `xml:"measure"`
		Divide []struct {
			XMLName       xml.Name
			XMLAttrs      []xml.Attr `xml:",any,attr"`
			UnitNumerator []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				Measure  []struct {
					XMLName  xml.Name
					XMLInner string `xml:",innerxml"`
				} `xml:"measure"`
			} `xml:"unitNumerator"`
			UnitDenominator []struct {
				XMLName  xml.Name
				XMLAttrs []xml.Attr `xml:",any,attr"`
				Measure  []struct {
					XMLName  xml.Name
					XMLAttrs []xml.Attr `xml:",any,attr"`
					XMLInner string     `xml:",innerxml"`
				} `xml:"measure"`
			} `xml:"unitDenominator"`
		} `xml:"divide"`
	} `xml:"unit"`
	FootnoteLink []struct {
		XMLName  xml.Name
		Type     string     `xml:"type,attr"`
		Role     string     `xml:"role,attr"`
		Title    string     `xml:"title,attr"`
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Footnote []struct {
			XMLName  xml.Name
			CharData string     `xml:",chardata"`
			Lang     string     `xml:"lang,attr"`
			Role     string     `xml:"role,attr"`
			Type     string     `xml:"type,attr"`
			Label    string     `xml:"label,attr"`
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"footnote"`
		Loc []struct {
			Href     string     `xml:"href,attr"`
			Label    string     `xml:"label,attr"`
			Type     string     `xml:"locator,attr"`
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"loc"`
		FootnoteArc []struct {
			XMLName  xml.Name
			Arcrole  string     `xml:"arcrole,attr"`
			Type     string     `xml:"type,attr"`
			From     string     `xml:"from,attr"`
			To       string     `xml:"to,attr"`
			Title    string     `xml:"title,attr"`
			XMLAttrs []xml.Attr `xml:",any,attr"`
		} `xml:"footnoteArc"`
	} `xml:"footnoteLink"`
	Facts []Fact `xml:",any"`
}

type Fact struct {
	XMLName    xml.Name
	ID         string     `xml:"id,attr,omitempty"`
	ContextRef string     `xml:"contextRef,attr"`
	UnitRef    string     `xml:"unitRef,attr,omitempty"`
	Decimals   string     `xml:"decimals,attr,omitempty"`
	Precision  string     `xml:"precision,attr,omitempty"`
	XMLAttrs   []xml.Attr `xml:",any,attr"`
	XMLInner   string     `xml:",innerxml"`
}

type Context struct {
	XMLName  xml.Name
	ID       string     `xml:"id,attr"`
	XMLAttrs []xml.Attr `xml:",any,attr"`
	Entity   []struct {
		XMLName      xml.Name
		XMLAttrs     []xml.Attr `xml:",any,attr"`
		Identitifier []struct {
			XMLName  xml.Name
			Scheme   string `xml:"scheme,attr"`
			CharData string `xml:",chardata"`
		} `xml:"identifier"`
		Segment []struct {
			XMLName        xml.Name
			XMLAttrs       []xml.Attr `xml:",any,attr"`
			ExplicitMember []struct {
				XMLName   xml.Name
				Dimension string     `xml:"dimension,attr"`
				XMLAttrs  []xml.Attr `xml:",any,attr"`
				CharData  string     `xml:",chardata"`
			} `xml:"explicitMember"`
			TypedMember []struct {
				XMLName   xml.Name
				Dimension string     `xml:"dimension,attr"`
				XMLAttrs  []xml.Attr `xml:",any,attr"`
				XMLInner  string     `xml:",innerxml"` //todo nested elements in typedMembers
			} `xml:"typedMember"`
		} `xml:"segment"`
	} `xml:"entity"`
	Period []struct {
		XMLName  xml.Name
		XMLAttrs []xml.Attr `xml:",any,attr"`
		Instant  []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			XMLInner string     `xml:",innerxml"`
		} `xml:"instant"`
		StartDate []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			XMLInner string     `xml:",innerxml"`
		} `xml:"startDate"`
		EndDate []struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr `xml:",any,attr"`
			XMLInner string     `xml:",innerxml"`
		} `xml:"endDate"`
	} `xml:"period"`
	Scenario []struct {
		XMLName        xml.Name
		XMLAttrs       []xml.Attr `xml:",any,attr"`
		ExplicitMember []struct {
			XMLName   xml.Name
			Dimension string     `xml:"dimension,attr"`
			XMLAttrs  []xml.Attr `xml:",any,attr"`
			CharData  string     `xml:",chardata"`
		} `xml:"explicitMember"`
		TypedMember []struct {
			XMLName   xml.Name
			Dimension string     `xml:"dimension,attr"`
			XMLAttrs  []xml.Attr `xml:",any,attr"`
			XMLInner  string     `xml:",innerxml"` //todo nested elements in typedMembers
		} `xml:"typedMember"`
	} `xml:"scenario"`
}

func ReadInstance(file os.FileInfo, workingDir string) (*Instance, error) {
	filepath := path.Join(workingDir, file.Name())
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	decoded, err := DecodeInstance(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func DecodeInstance(xmlData []byte) (*Instance, error) {
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoded := Instance{}
	err := decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

func EncodeInstance(instance *Instance) ([]byte, error) {
	encoded := Instance{}
	for i, attr := range instance.XMLAttrs {
		if instance.XMLName.Space == attr.Value && attr.Value == xbrli {
			for j, aattr := range instance.XMLAttrs {
				if aattr.Name.Space == "xmlns" {
					instance.XMLAttrs[j].Name.Space = ""
					instance.XMLAttrs[j].Name.Local = "xmlns:" + aattr.Name.Local
				}
			}
			instance.XMLAttrs[i].Name.Space = ""
			instance.XMLAttrs[i].Name.Local = "xmlns"
			//todo namespace everything
			encoded.XMLAttrs = instance.XMLAttrs
			encoded.SchemaRef = instance.SchemaRef
			encoded.Context = instance.Context
			encoded.Unit = instance.Unit
			encoded.FootnoteLink = instance.FootnoteLink
			encoded.Facts = instance.Facts
			break
		}
	}
	return xml.Marshal(encoded)
}
