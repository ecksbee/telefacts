package renderables

import (
	gohtml "html"
	"io"
	"strconv"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"

	"golang.org/x/net/html"
)

func findHtmlAttr(attrs []html.Attribute, attr string) *html.Attribute {
	for _, a := range attrs {
		if a.Key == attr {
			return &a
		}
	}
	return nil
}

func colSpanHelper(attrs []html.Attribute, cb func(int)) {
	a := findHtmlAttr(attrs, "colspan")
	if a != nil {
		colspan, err := strconv.Atoi(a.Val)
		if err == nil {
			for i := 0; i < colspan-1; i++ {
				cb(colspan)
			}
		}
	}
}

func colSpan2Helper(stack []int, cb func(int)) {
	if len(stack) > 0 {
		colspan := stack[len(stack)-1]
		for i := 0; i < colspan-1; i++ {
			cb(colspan)
		}
	}
}

func renderTextBlock(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder) *FactExpression {
	_, concept, err := cf.HashQuery(fact.Href)
	if err != nil {
		return &FactExpression{
			Core: "error",
		}
	}
	isTextBlock := concept.Type.Local == attr.TextBlockItemType
	// concept.Type.Space == attr.NONNUM || concept.Type.Space == "http://www.xbrl.org/dtr/type/2020-01-21" //todo
	if !isTextBlock {
		return nil
	}
	r := strings.NewReader(gohtml.UnescapeString(fact.XMLInner))
	tokenizer := html.NewTokenizer(r)
	tokenized := ""
	disableBr := 0
	thcolspans := make([]int, 0)
	tdcolspans := make([]int, 0)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()
		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.SelfClosingTagToken, html.StartTagToken:
			switch token.Data {
			case "table":
				disableBr++
				tokenized += "<table>"
			case "tr":
				tokenized += "<tr>"
			case "th":
				tokenized += "<th>"
				colSpanHelper(token.Attr, func(colspan int) {
					thcolspans = append(thcolspans, colspan)
					tokenized += "<th>"
				})
			case "td":
				tokenized += "<td>"
				colSpanHelper(token.Attr, func(colspan int) {
					tdcolspans = append(tdcolspans, colspan)
					tokenized += "<td>"
				})
			}
			if disableBr <= 0 {
				blockElems := []string{
					"div",
					"p",
					"h1",
					"h2",
					"h3",
					"h4",
					"h5",
					"h6",
				}
				for _, elem := range blockElems {
					if elem == token.Data {
						tokenized += "<br /><br />"
						break
					}
				}
			}
		case html.TextToken:
			data := strings.TrimSpace(token.Data)
			if len(data) > 0 {
				tokenized += data + " "
			}
		case html.EndTagToken:
			switch token.Data {
			case "table":
				disableBr--
				tokenized += "</table>"
			case "tr":
				tokenized += "</tr>"
			case "th":
				tokenized += "</th>"
				colSpan2Helper(thcolspans, func(colspan int) {
					thcolspans = thcolspans[:len(thcolspans)-1]
					tokenized += "</th>"
				})
			case "td":
				tokenized += "</td>"
				colSpan2Helper(tdcolspans, func(colspan int) {
					tdcolspans = tdcolspans[:len(tdcolspans)-1]
					tokenized += "</td>"
				})
			}
		}
	}

	return &FactExpression{
		InnerHtml: tokenized,
	}
}
