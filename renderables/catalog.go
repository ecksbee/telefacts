package renderables

import (
	"encoding/json"
	"hash/fnv"

	"ecksbee.com/telefacts/hydratables"
)

type Catalog struct {
	Files            []string
	Entities         []string
	RelationshipSets []string
	Networks         map[string]map[string]string
}

func MarshalCatalog(h *hydratables.Hydratable, filenames []string) ([]byte, error) {
	schemedEntities := sortedEntities(h)
	linkroleURIs := sortedRelationshipSets(h)
	networks := map[string]map[string]string{}
	for _, schemedEntity := range schemedEntities {
		networks[schemedEntity] = make(map[string]string)
		for _, linkroleURI := range linkroleURIs {
			hash := hash(schemedEntity, linkroleURI)
			networks[schemedEntity][linkroleURI] = hash
		}
	}
	return json.Marshal(Catalog{
		Entities:         schemedEntities,
		RelationshipSets: linkroleURIs,
		Files:            filenames,
		Networks:         networks,
	})
}

func hash(schemedEntity string, linkroleURI string) string {
	h := fnv.New128a()
	h.Write([]byte(schemedEntity + linkroleURI))
	return string(h.Sum([]byte{}))
}
