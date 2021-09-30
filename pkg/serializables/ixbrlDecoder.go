package serializables

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sync"

	"ecksbee.com/telefacts/pkg/attr"
)

func DecodeIxbrl(byteArray []byte) (*IxbrlFile, error) {
	var ixReferences *IxReferences
	var errRef error
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		ixReferences, errRef = DecodeIxReferences(byteArray)
	}()
	var ixResources *IxResources
	var errRes error
	go func() {
		defer wg.Done()
		ixResources, errRes = DecodeIxResources(byteArray)
	}()
	var ixHiddenFacts *IxHiddenFacts
	var ixRenderedFacts *IxRenderedFacts
	var errFacts error
	go func() {
		defer wg.Done()
		ixHiddenFacts, ixRenderedFacts, errFacts = DecodeInlineFacts(byteArray)
	}()
	wg.Wait()
	for _, err := range []error{
		errRef,
		errRes,
		errFacts,
	} {
		if err != nil {
			return nil, err
		}
	}
	decoded := IxbrlFile{}
	decoded.Header.References = *ixReferences
	decoded.Header.Resources = *ixResources
	decoded.Header.Hidden = *ixHiddenFacts
	decoded.RenderedFacts = *ixRenderedFacts
	return &decoded, nil
}

func DecodeIxReferences(byteArray []byte) (*IxReferences, error) {
	reader := bytes.NewReader(byteArray)
	d := xml.NewDecoder(reader)
	decoded := IxReferences{}
	decoded.SchemaRef = make([]string, 0)
	decoded.LinkbaseRef = make([]string, 0)
	isInBody := false
	ixHeader := 0
	ixReferences := 0
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		switch token.(type) {
		case xml.StartElement:
			startElem, ok := token.(xml.StartElement)
			if ok {
				if isInBody {
					switch startElem.Name.Space {
					case attr.IX:
						switch startElem.Name.Local {
						case "header":
							ixHeader++
						case "references":
							if ixHeader <= 0 {
								return nil, fmt.Errorf("malformed ixbrl file: misplaced references element")
							}
							ixReferences++
						}
					case attr.LINK:
						switch startElem.Name.Local {
						case "schemaRef":
							if ixHeader <= 0 || ixReferences <= 0 {
								return nil, fmt.Errorf("malformed ixbrl file: misplaced schemaRef element")
							}
							hrefAttr := attr.FindAttr(startElem.Attr, "href")
							if hrefAttr.Value != "" && hrefAttr.Name.Space == attr.XLINK {
								decoded.SchemaRef = append(decoded.SchemaRef, hrefAttr.Value)
							}
						case "linkbaseRef":
							if ixHeader <= 0 || ixReferences <= 0 {
								return nil, fmt.Errorf("malformed ixbrl file: misplaced linkbaseRef element")
							}
							hrefAttr := attr.FindAttr(startElem.Attr, "href")
							if hrefAttr.Value != "" && hrefAttr.Name.Space == attr.XLINK {
								decoded.LinkbaseRef = append(decoded.LinkbaseRef, hrefAttr.Value)
							}
						}
					}
				} else {
					if startElem.Name.Local == "body" &&
						startElem.Name.Space == "http://www.w3.org/1999/xhtml" {
						isInBody = true
					}
				}
			}
		case xml.EndElement:
			endElem, ok := token.(xml.EndElement)
			if ok {
				if isInBody {
					switch endElem.Name.Space {
					case "http://www.w3.org/1999/xhtml":
						if endElem.Name.Local == "body" {
							isInBody = false
						}
					case attr.IX:
						switch endElem.Name.Local {
						case "header":
							ixHeader--
						case "references":
							ixReferences--
						}
					}
				}
			}
		}
	}
	return &decoded, nil
}
