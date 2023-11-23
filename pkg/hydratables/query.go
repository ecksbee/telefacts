package hydratables

import (
	"fmt"
	"strings"

	"ecksbee.com/telefacts/pkg/attr"
)

func (h *Hydratable) HashQuery(query string) (string, *Concept, error) {
	i := strings.IndexRune(query, '#')
	if i < 0 {
		return "", nil, fmt.Errorf("invalid query")
	}
	base := query[:i]
	if len(base) <= 0 {
		return "", nil, fmt.Errorf("invalid base query")
	}
	fragment := query[i+1:]
	if len(fragment) <= 0 {
		return "", nil, fmt.Errorf("invalid query fragment")
	}
	var namespace string
	var concepts []Concept
	for key, value := range h.Folder.Namespaces {
		if value == base {
			namespace = key
		}
	}
	if attr.IsValidUrl(base) {
		schema, err := HydrateGlobalSchema(base)
		if err != nil {
			return namespace, nil, err
		}
		concepts = schema.Element
	} else {
		file := h.Folder.Schemas[base]
		schema, err := HydrateSchema(&file, base)
		if err != nil {
			return namespace, nil, err
		}
		concepts = schema.Element
	}
	for _, candidate := range concepts {
		if fragment == candidate.ID {
			return namespace, &candidate, nil
		}
	}
	return namespace, nil, fmt.Errorf("concept not found %s", query)
}

func (h *Hydratable) NameQuery(namespace string, localName string) (string, *Concept, error) {
	var schemaLoc string
	var concepts []Concept
	schemaLoc = h.Folder.Namespaces[namespace]
	if len(schemaLoc) <= 0 {
		return "", nil, fmt.Errorf("%s is not scoped into the folder", namespace)
	}
	if attr.IsValidUrl(schemaLoc) {
		schema, err := HydrateGlobalSchema(schemaLoc)
		if err != nil {
			return "", nil, err
		}
		concepts = schema.Element
	} else {
		file := h.Folder.Schemas[schemaLoc]
		schema, err := HydrateSchema(&file, schemaLoc)
		if err != nil {
			return "", nil, err
		}
		concepts = schema.Element
	}
	for _, candidate := range concepts {
		if localName == candidate.XMLName.Local && namespace == candidate.XMLName.Space {
			return schemaLoc + "#" + candidate.ID, &candidate, nil
		}
	}
	return "", nil, nil
}

func EntityQuery(scheme string, chardata string) (string, error) {
	names, err := HydrateEntityNames()
	if err != nil {
		return "", err
	}
	if x, found := names[scheme]; found {
		if y, foundd := x[chardata]; foundd {
			return y, nil
		}
	}
	return scheme + "/" + chardata, nil
}
