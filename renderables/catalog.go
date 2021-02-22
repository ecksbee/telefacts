package renderables

import (
	"encoding/hex"
	"encoding/json"
	"hash/fnv"

	"ecksbee.com/telefacts/hydratables"
)

type Catalog struct {
	Files            []string
	Subjects         []Subject
	RelationshipSets []RelationshipSet
	Networks         map[string]map[string]string
}

func MarshalCatalog(h *hydratables.Hydratable, names map[string]map[string]string,
	filenames []string) ([]byte, error) {
	schemedEntities := sortedEntities(h)
	rsets := sortedRelationshipSets(h)
	subjects := make([]Subject, 0, len(schemedEntities))
	networks := map[string]map[string]string{}
	for _, schemedEntity := range schemedEntities {
		entityStr := stringify(&schemedEntity)
		networks[entityStr] = make(map[string]string)
		for _, rset := range rsets {
			hash := hash(entityStr, rset.RoleURI)
			networks[entityStr][rset.RoleURI] = hash
		}
		subjects = append(subjects, Subject{
			Name:   names[schemedEntity.Scheme][schemedEntity.CharData],
			Entity: schemedEntity,
		})
	}
	return json.Marshal(Catalog{
		Subjects:         subjects,
		RelationshipSets: rsets,
		Files:            filenames,
		Networks:         networks,
	})
}

func hash(schemedEntity string, linkroleURI string) string {
	h := fnv.New128a()
	h.Write([]byte(schemedEntity + linkroleURI))
	return hex.EncodeToString(h.Sum([]byte{}))
}
