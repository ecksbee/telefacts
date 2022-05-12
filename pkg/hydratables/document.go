package hydratables

import "ecksbee.com/telefacts/pkg/serializables"

type Expressable struct {
	Href       string
	ID         string
	ContextRef string
	UnitRef    string
	Precision  Precision
	IsNil      bool
	XMLInner   string
}

type Document struct {
	Expressions []Expressable
}

func HydrateDocument(folder *serializables.Folder) (*Document, error) {
	//todo
	return nil, nil
}
