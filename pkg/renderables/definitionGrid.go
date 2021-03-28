package renderables

import (
	"ecksbee.com/telefacts/pkg/hydratables"
)

type DGrid struct {
	RootDomains []RootDomain
	//todo XBRL v1 definition arc roles
}

func dGrid(schemedEntity string, linkroleURI string, h *hydratables.Hydratable,
	factFinder FactFinder, conceptFinder ConceptFinder, measurementFinder MeasurementFinder) (DGrid, []LabelRole, []Lang, error) {
	rootDomains, labelRoles, langs := getRootDomains(schemedEntity, linkroleURI, h, factFinder, conceptFinder, measurementFinder)
	return DGrid{
		RootDomains: rootDomains,
	}, labelRoles, langs, nil
}
