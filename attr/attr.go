package attr

import (
	"encoding/xml"
	"strings"
)

func FindAttr(attrs []xml.Attr, attr string) *xml.Attr {
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
		return xml.Name{}
	}
	prefix := prefixedName[:i]
	local := prefixedName[i+1:]
	if len(local) <= 0 {
		return xml.Name{}
	}
	return xml.Name{
		Space: Ns(attrs, prefix),
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
