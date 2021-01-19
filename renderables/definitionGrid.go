package renderables

import (
	"encoding/json"
	"fmt"

	"ecks-bee.com/telefacts/xbrl"
)

type DGrid struct {
	RootDomains []RootDomain
	//todo XBRL v1 definition arc roles
}

func MarshalDGrid(entityIndex int, relationshipSetIndex int, schema *xbrl.Schema,
	instance *xbrl.Instance, definition *xbrl.DefinitionLinkbase,
	factFinder FactFinder) ([]byte, error) {
	schemedEntities := sortedEntities(instance)
	if entityIndex > len(schemedEntities)-1 {
		return nil, fmt.Errorf("invalid entity index")
	}
	linkroleURIs := sortedRelationshipSets(schema)
	if relationshipSetIndex > len(linkroleURIs)-1 {
		return nil, fmt.Errorf("invalid relationship set index")
	}
	linkroleURI := linkroleURIs[relationshipSetIndex]
	schemedEntity := schemedEntities[entityIndex]
	rootDomains := getRootDomains(schemedEntity, linkroleURI, schema, instance,
		definition, factFinder)
	return json.Marshal(DGrid{
		RootDomains: rootDomains,
	})
}
