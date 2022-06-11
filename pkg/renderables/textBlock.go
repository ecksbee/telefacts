package renderables

import (
	gohtml "html"
	"io"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"

	"golang.org/x/net/html"
)

func renderTextBlock(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder) *FactExpression {
	_, concept, err := cf.HashQuery(fact.Href)
	if err != nil {
		return &FactExpression{
			Core: "error",
		}
	}
	isTextBlock := concept.Type.Space == attr.NONNUM &&
		concept.Type.Local == attr.TextBlockItemType
	if !isTextBlock {
		return nil
	}
	r := strings.NewReader(gohtml.UnescapeString(fact.XMLInner))
	tokenizer := html.NewTokenizer(r)
	tokenized := ""
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()
		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.SelfClosingTagToken, html.StartTagToken:
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
		case html.TextToken:
			data := strings.TrimSpace(token.Data)
			if len(data) > 0 {
				tokenized += data
			}
		}
	}

	return &FactExpression{
		InnerHtml: tokenized,
	}
}
