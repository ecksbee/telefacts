package serializables

import "encoding/xml"

type CommonContext struct {
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
}

type CommonUnit struct {
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
}
