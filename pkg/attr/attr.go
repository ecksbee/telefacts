package attr

import (
	"encoding/xml"
	"strings"

	"github.com/antchfx/xmlquery"
)

func FindAttr(attrs []xml.Attr, attr string) *xml.Attr {
	for _, a := range attrs {
		if a.Name.Local == attr {
			return &a
		}
	}
	return nil
}

func FindXpathAttr(attrs []xmlquery.Attr, attr string) *xmlquery.Attr {
	for _, a := range attrs {
		if a.Name.Local == attr {
			return &a
		}
	}
	return nil
}

func Xmlns(attrs []xml.Attr, prefixedName string) xml.Name {
	i := strings.IndexRune(prefixedName, ':')
	if i < 0 {
		xmlnsAttr := FindAttr(attrs, `xmlns`)
		space := ""
		if xmlnsAttr != nil {
			space = xmlnsAttr.Value
		}
		return xml.Name{
			Space: space,
			Local: prefixedName,
		}
	}
	prefix := prefixedName[:i]
	local := prefixedName[i+1:]
	if len(local) <= 0 {
		return xml.Name{}
	}
	space := Ns(attrs, prefix)
	if space == "" {
		defaultAttr := FindAttr(attrs, `xmlns`)
		if defaultAttr != nil {
			space = defaultAttr.Value
		}
	}
	return xml.Name{
		Space: space,
		Local: local,
	}
}

func Ns(attrs []xml.Attr, prefix string) string {
	for _, attr := range attrs {
		if attr.Name.Space == "xmlns" && attr.Name.Local == prefix {
			return attr.Value
		}
	}
	return ""
}
