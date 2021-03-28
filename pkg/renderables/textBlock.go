package renderables

import (
	gohtml "html"
	"io"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/hydratables"

	"golang.org/x/net/html"
)

func renderTextBlock(fact *hydratables.Fact, cf ConceptFinder, mf MeasurementFinder, labelRole LabelRole) *FactExpression {
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
	preview := ""
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()
		err := tokenizer.Err()
		if err == io.EOF || len(preview) > 40 {
			break
		}

		switch tt {
		case html.TextToken:
			data := strings.TrimSpace(token.Data)

			if len(data) > 0 {
				preview += data + " "
			}
		}
	}

	return &FactExpression{
		TextPreview: preview,
		TextBlock:   fact.XMLInner,
	}
}
