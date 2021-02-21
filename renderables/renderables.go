package renderables

import (
	"encoding/json"
	"fmt"
	"sync"

	"ecksbee.com/telefacts/hydratables"
)

type Renderables struct {
	RoleURI    string
	Entity     string
	Title      string
	Lang       []Lang
	LabelRoles []LabelRole
	PGrid
	DGrid
	CGrid
}

func MarshalRenderable(slug string, h *hydratables.Hydratable) ([]byte, error) {
	schemedEntities := sortedEntities(h)
	linkroleURIs := sortedRelationshipSets(h)
	for _, schemedEntity := range schemedEntities {
		for _, linkroleURI := range linkroleURIs {
			if slug == hash(schemedEntity, linkroleURI) {
				var (
					p          PGrid
					d          DGrid
					c          CGrid
					title      string
					labelRoles []LabelRole
					langs      []Lang
					err        error
					wg         sync.WaitGroup
				)
				wg.Add(3)
				labelRoles = make([]LabelRole, 0, 20)
				langs = make([]Lang, 0, 8)
				go func(entity string, linkrole string) {
					defer wg.Done()
					localP, lr, ln, localError := pGrid(entity, linkrole, h, h)
					if localError != nil {
						err = localError
						return
					}
					p = localP
					labelRoles = append(labelRoles, lr...)
					langs = append(langs, ln...)
				}(schemedEntity, linkroleURI)
				go func(entity string, linkrole string) {
					defer wg.Done()
					localD, lr, ln, localError := dGrid(entity, linkrole, h, h)
					if localError != nil {
						err = localError
						return
					}
					d = localD
					labelRoles = append(labelRoles, lr...)
					langs = append(langs, ln...)
				}(schemedEntity, linkroleURI)
				go func(entity string, linkrole string) {
					defer wg.Done()
					localC, lr, ln, localError := cGrid(entity, linkrole, h, h)
					if localError != nil {
						err = localError
						return
					}
					c = localC
					labelRoles = append(labelRoles, lr...)
					langs = append(langs, ln...)
				}(schemedEntity, linkroleURI)

				wg.Wait()
				ret := Renderables{
					Title:      title,
					RoleURI:    linkroleURI,
					Entity:     schemedEntity,
					PGrid:      p,
					DGrid:      d,
					CGrid:      c,
					Lang:       dedupLang(langs),
					LabelRoles: dedupLabelRole(labelRoles),
				}
				if err != nil {
					return nil, err
				}
				return json.Marshal(ret)
			}
		}
	}
	return nil, fmt.Errorf("Object not found")
}
