package serializables

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"ecksbee.com/telefacts/pkg/attr"
)

func DecodeInlineFacts(byteArray []byte) (*IxHiddenFacts, *IxLexicalFacts, error) {
	reader := bytes.NewReader(byteArray)
	d := xml.NewDecoder(reader)
	ixHidden := IxHiddenFacts{
		Nonfractions: make([]IxbrlNonfraction, 0),
		Nonnumeric:   make([]IxbrlNonnumeric, 0),
		Footnotes:    make([]IxbrlFootnote, 0),
	}
	ixLexicalFacts := IxLexicalFacts{
		Nonfractions: make([]IxbrlNonfraction, 0),
		Nonnumeric:   make([]IxbrlNonnumeric, 0),
		Footnotes:    make([]IxbrlFootnote, 0),
	}
	isInBody := false
	ixHeader := 0
	ixHiddenCount := 0
	factualTags := make([]xml.StartElement, 0)
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
						case "hidden":
							if ixHeader <= 0 {
								return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced hidden element")
							}
							ixHiddenCount++
						case "nonFraction", "nonfraction":
							if ixHeader > 0 {
								if ixHiddenCount <= 0 {
									return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced nonFraction element")
								}
							}
							factualTags = append(factualTags, startElem)
						case "nonNumeric", "nonnumeric":
							if ixHeader > 0 {
								if ixHiddenCount <= 0 {
									return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced nonNumeric element")
								}
							}
							factualTags = append(factualTags, startElem)
						case "footnote":
							if ixHeader > 0 {
								if ixHiddenCount <= 0 {
									return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced footnote element")
								}
							}
							factualTags = append(factualTags, startElem)
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
						case "hidden":
							ixHiddenCount--
						case "nonFraction", "nonfraction":
							if len(factualTags) <= 0 {
								return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced end nonFraction element")
							}
							factualTags = factualTags[:len(factualTags)-1]
						case "nonNumeric", "nonnumeric":
							if len(factualTags) <= 0 {
								return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced end nonNumeric element")
							}
							factualTags = factualTags[:len(factualTags)-1]
						case "footnote":
							if len(factualTags) <= 0 {
								return nil, nil, fmt.Errorf("malformed ixbrl file: misplaced end footnote element")
							}
							factualTags = factualTags[:len(factualTags)-1]
						}
					}
				}
			}
		case xml.CharData:
			_, ok := token.(xml.CharData)
			if ok {
				if isInBody {
					//todo do something
					// fmt.Printf("%v", charData)
				}
			}
		}
	}
	return &ixHidden, &ixLexicalFacts, nil
}
