package renderables

import (
	"encoding/json"
	"fmt"
	"sync"

	"ecksbee.com/telefacts/pkg/hydratables"
)

type RelationshipSet struct {
	Title   string
	RoleURI string
}

type Subject struct {
	Name   string
	Entity Entity
}

type Entity struct {
	Scheme   string
	CharData string
}

type Renderable struct {
	RelationshipSet RelationshipSet
	Subject         Subject
	Lang            []Lang
	LabelRoles      []LabelRole
	PGrid           PGrid
	DGrid           DGrid
	CGrid           CGrid
}

func MarshalRenderable(slug string, names map[string]map[string]string, h *hydratables.Hydratable) ([]byte, error) {
	schemedEntities := sortedEntities(h)
	rsets := sortedRelationshipSets(h)
	for _, schemedEntity := range schemedEntities {
		for _, rset := range rsets {
			if slug == hash(stringify(&schemedEntity), rset.RoleURI) {
				var (
					p          PGrid
					d          DGrid
					c          CGrid
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
					localP, lr, ln, localError := pGrid(entity, linkrole, h, h, h)
					if localError != nil {
						err = localError
						return
					}
					p = localP
					labelRoles = append(labelRoles, lr...)
					langs = append(langs, ln...)
				}(stringify(&schemedEntity), rset.RoleURI)
				go func(entity string, linkrole string) {
					defer wg.Done()
					localD, lr, ln, localError := dGrid(entity, linkrole, h, h, h)
					if localError != nil {
						err = localError
						return
					}
					d = localD
					labelRoles = append(labelRoles, lr...)
					langs = append(langs, ln...)
				}(stringify(&schemedEntity), rset.RoleURI)
				go func(entity string, linkrole string) {
					defer wg.Done()
					localC, lr, ln, localError := cGrid(entity, linkrole, h, h, h)
					if localError != nil {
						err = localError
						return
					}
					c = localC
					labelRoles = append(labelRoles, lr...)
					langs = append(langs, ln...)
				}(stringify(&schemedEntity), rset.RoleURI)
				wg.Wait()
				langs = dedupLang(langs)
				labelRoles = dedupLabelRole(labelRoles)
				p, d, c = formatPeriod(p, d, c, labelRoles, langs)
				ret := Renderable{
					Subject: Subject{
						Name: names[schemedEntity.Scheme][schemedEntity.CharData],
						Entity: struct {
							Scheme   string
							CharData string
						}{
							Scheme:   schemedEntity.Scheme,
							CharData: schemedEntity.CharData,
						},
					},
					RelationshipSet: RelationshipSet{
						Title:   rset.Title,
						RoleURI: rset.RoleURI,
					},
					PGrid:      p,
					DGrid:      d,
					CGrid:      c,
					Lang:       langs,
					LabelRoles: labelRoles,
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
