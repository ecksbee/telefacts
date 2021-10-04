package renderables

import (
	"ecksbee.com/telefacts/pkg/hydratables"
)

type DGrid struct {
	RootDomains []RootDomain
	DRS         DRS
}

func dGrid(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder) (DGrid, []LabelRole, []Lang, error) {
	rootDomains, labelRoles, langs := getRootDomains(schemedEntity, linkroleURI, h, factFinder, conceptFinder, measurementFinder)
	return DGrid{
		RootDomains: rootDomains,
		DRS:         getDRS(schemedEntity, linkroleURI, h),
	}, labelRoles, langs, nil
}
