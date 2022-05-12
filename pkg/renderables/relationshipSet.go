package renderables

import (
	"sort"

	"ecksbee.com/telefacts/pkg/hydratables"
)

func sortedRelationshipSets(h *hydratables.Hydratable) []RelationshipSet {
	rsets := dedupRelationshipSets(h)
	sort.SliceStable(rsets, func(i, j int) bool {
		if rsets[i].Title == rsets[j].Title {
			return rsets[i].RoleURI < rsets[j].RoleURI
		}
		return rsets[i].Title < rsets[j].Title
	})
	return rsets
}

func dedupRelationshipSets(h *hydratables.Hydratable) []RelationshipSet {
	rsets := []RelationshipSet{}
	for _, schema := range h.Schemas {
		if len(schema.Annotation.Appinfo.RoleTypes) <= 0 {
			continue
		}
		for _, e := range schema.Annotation.Appinfo.RoleTypes {
			if len(e.RoleURI) <= 0 {
				continue
			}
			rsets = append(rsets, RelationshipSet{
				RoleURI: e.RoleURI,
				Title:   e.Definition,
			})
		}
	}
	uniques := func(arr []RelationshipSet) []RelationshipSet {
		occured := map[RelationshipSet]bool{}
		u := []RelationshipSet{}
		for e := range arr {
			if occured[arr[e]] != true {
				occured[arr[e]] = true
				u = append(u, arr[e])
			}
		}
		return u
	}(rsets)
	return uniques
}
