package renderables

import (
	"encoding/hex"
	"encoding/json"
	"hash/fnv"

	"ecksbee.com/telefacts/pkg/hydratables"
)

type Catalog struct {
	Subjects         []Subject
	RelationshipSets []RelationshipSet
	Networks         map[string]map[string]string
	Expressions      map[string]Expressable
}

func MarshalCatalog(h *hydratables.Hydratable) ([]byte, error) {
	schemedEntities := sortedEntities(h)
	rsets := sortedRelationshipSets(h)
	subjects := make([]Subject, 0, len(schemedEntities))
	networks := map[string]map[string]string{}
	expressions, err := getExpressions(h, h)
	if err != nil {
		return nil, err
	}
	for _, schemedEntity := range schemedEntities {
		entityStr := stringify(&schemedEntity)
		networks[entityStr] = make(map[string]string)
		for _, rset := range rsets {
			hash := hash(entityStr, rset.RoleURI)
			networks[entityStr][rset.RoleURI] = hash
		}
		name := schemedEntity.Scheme + "/" + schemedEntity.CharData
		hydratedName, err := hydratables.EntityQuery(schemedEntity.Scheme,
			schemedEntity.CharData)
		if err == nil {
			name = hydratedName
		}
		subjects = append(subjects, Subject{
			Name:   name,
			Entity: schemedEntity,
		})
	}
	return json.Marshal(Catalog{
		Subjects:         subjects,
		RelationshipSets: rsets,
		Networks:         networks,
		Expressions:      expressions,
	})
}

func hash(schemedEntity string, linkroleURI string) string {
	h := fnv.New128a()
	h.Write([]byte(schemedEntity + linkroleURI))
	return hex.EncodeToString(h.Sum([]byte{}))
}
