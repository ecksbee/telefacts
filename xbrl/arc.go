package xbrl

import (
	"encoding/xml"
)

type Arc struct {
	XMLName        xml.Name
	Order          string `xml:"order,attr"`
	Arcrole        string `xml:"arcrole,attr"`
	Type           string `xml:"type,attr"`
	From           string `xml:"from,attr"`
	To             string `xml:"to,attr"`
	Closed         bool   `xml:"closed,attr,omitempty"`
	ContextElement string `xml:"contextElement,attr,omitempty"`
	TargetRole     string `xml:"targetRole,attr,omitempty"`
}
