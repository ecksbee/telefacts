package renderables

import (
	"bufio"
	"encoding/xml"
	"io"
	"strings"

	"ecksbee.com/telefacts/pkg/hydratables"
)

func formatTypedMember(typedDomainHref string, typedMember string, h *hydratables.Hydratable) string {
	// namespace, typedDomain, err := h.HashQuery(typedDomainHref)	//todo validate typedDomain with xmlQuery https://www.golangprograms.com/dynamic-xml-parser-without-struct-in-go.html
	// if err != nil {
	// 	return ""
	// }
	reader := bufio.NewReader(strings.NewReader(typedMember))
	d := xml.NewDecoder(reader)
	ret := ""
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		switch token.(type) {
		case xml.CharData:
			charData, ok := token.(xml.CharData)
			if ok {
				str := strings.TrimSpace(string(charData))
				if str != "" {
					ret += string(charData) + ", "
				}
			}
		}
	}
	return ret
}
