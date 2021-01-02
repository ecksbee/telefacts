package renderables

import (
	"encoding/json"

	"ecks-bee.com/telefacts/xbrl"
)

type Catalog struct {
	Entities         []string
	RelationshipSets []string
}

func MarshalCatalog(schema *xbrl.Schema, instance *xbrl.Instance) ([]byte, error) {
	schemedEntities := sortedEntities(instance)
	linkroleURIs := sortedRelationshipSets(schema)
	return json.Marshal(Catalog{
		Entities:         schemedEntities,
		RelationshipSets: linkroleURIs,
	})
}
