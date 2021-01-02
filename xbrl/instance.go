package xbrl

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html/charset"
)

type Instance struct {
	XMLName   xml.Name   `xml:"xbrl"`
	XMLAttrs  []xml.Attr `xml:",any,attr"`
	SchemaRef []struct {
		XMLName xml.Name
		HRef    string `xml:"href,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"schemaRef"`
	Context []Context `xml:"context"`
	Unit    []struct {
		XMLName xml.Name
		Measure string `xml:"measure"`
		Divide  []struct {
			UnitNumerator []struct {
				Measure string `xml:"measure"`
			} `xml:"unitNumerator"`
			UnitDenominator []struct {
				Measure string `xml:"measure"`
			} `xml:"unitDenominator"`
		} `xml:"divide"`
	} `xml:"unit"`
	FootnoteLink []struct {
		XMLName  xml.Name
		Type     string `xml:"type,attr"`
		Role     string `xml:"role,attr"`
		Title    string `xml:"title,attr"`
		Footnote []struct {
			XMLName xml.Name
			Text    string `xml:",chardata"`
			Lang    string `xml:"lang,attr"`
			Role    string `xml:"role,attr"`
			Type    string `xml:"type,attr"`
			Label   string `xml:"label,attr"`
		} `xml:"footnote"`
		Loc []struct {
			Href  string `xml:"href,attr"`
			Label string `xml:"label,attr"`
			Type  string `xml:"locator,attr"`
		} `xml:"loc"`
		FootnoteArc []struct {
			XMLName xml.Name
			Arcrole string `xml:"arcrole,attr"`
			Type    string `xml:"type,attr"`
			From    string `xml:"from,attr"`
			To      string `xml:"to,attr"`
			Title   string `xml:"title,attr"`
		} `xml:"footnoteArc"`
	} `xml:"footnoteLink"`
	Facts []Fact `xml:",any"`
}

type Fact struct {
	XMLName    xml.Name
	ID         string `xml:"id,attr"`
	ContextRef string `xml:"contextRef,attr"`
	UnitRef    string `xml:"unitRef,attr"`
	Decimals   string `xml:"decimals,attr"`
	Precision  string `xml:"precision,attr"`
	Text       string `xml:",chardata"`
	XMLInner   string `xml:",innerxml"` //todo tuples
}

type Context struct {
	XMLName xml.Name
	ID      string `xml:"id,attr"`
	Entity  struct {
		Identitifier []struct {
			XMLName xml.Name
			Scheme  string `xml:"scheme,attr"`
			Text    string `xml:",chardata"`
		} `xml:"identifier"`
		Segment struct {
			XMLName        xml.Name
			ExplicitMember []struct {
				XMLName   xml.Name
				Dimension string `xml:"dimension,attr"`
				Text      string `xml:",chardata"`
			} `xml:"explicitMember"`
			TypedMember []struct {
				XMLName   xml.Name
				Dimension string `xml:"dimension,attr"`
				XMLInner  string `xml:",innerxml"` //todo nested elements in typedMembers
			} `xml:"typedMember"`
			SegmentConcepts []struct {
				XMLName  xml.Name
				XMLInner string `xml:",innerxml"` //todo nested elements in segments
			} `xml:",any"`
		} `xml:"segment"`
		Scenario struct {
			XMLName        xml.Name
			XMLAttrs       []xml.Attr `xml:",any,attr"`
			ExplicitMember []struct {
				XMLName   xml.Name
				Dimension string `xml:"dimension,attr"`
				Text      string `xml:",chardata"`
			} `xml:"explicitMember"`
			TypedMember []struct {
				XMLName   xml.Name
				Dimension string `xml:"dimension,attr"`
				XMLInner  string `xml:",innerxml"` //todo nested elements in typedMembers
			} `xml:"typedMember"`
			ScenarioConcepts []struct {
				XMLName  xml.Name
				XMLInner string `xml:",innerxml"` //todo nested elements in scenarios
			} `xml:",any"`
		} `xml:"scenario"`
	} `xml:"entity"`
	Period struct {
		XMLName   xml.Name
		Instant   string `xml:"instant"`
		StartDate string `xml:"startDate"`
		EndDate   string `xml:"endDate"`
	} `xml:"period"`
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
