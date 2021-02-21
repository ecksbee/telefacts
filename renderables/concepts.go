package renderables

import (
	"strings"

	"ecksbee.com/telefacts/hydratables"
)

func mapPLocatorToHref(relationshipSetCurrentlyViewing string, presentation *hydratables.PresentationLinkbase, locator string) string {
	roleRefs := presentation.RoleRefs
	for _, roleRef := range roleRefs {
		if relationshipSetCurrentlyViewing == roleRef.RoleURI {
			presentationLinks := presentation.PresentationLinks
			for _, presentationLink := range presentationLinks {
				if presentationLink.Role == relationshipSetCurrentlyViewing {
					for _, loc := range presentationLink.Locs {
						if loc.Label == locator {
							i := strings.Index(loc.Href, "#")
							if i >= 0 {
								return loc.Href
							}
						}
					}
				}
			}
		}
	}
	return "#" + locator
}

func mapCLocatorToHref(relationshipSetCurrentlyViewing string, calculation *hydratables.CalculationLinkbase, locator string) string {
	roleRefs := calculation.RoleRefs
	for _, roleRef := range roleRefs {
		if relationshipSetCurrentlyViewing == roleRef.RoleURI {
			calculationLinks := calculation.CalculationLinks
			for _, calculationLink := range calculationLinks {
				if calculationLink.Role == relationshipSetCurrentlyViewing {
					for _, loc := range calculationLink.Locs {
						if loc.Label == locator {
							i := strings.Index(loc.Href, "#")
							if i >= 0 {
								return loc.Href
							}
						}
					}
				}
			}
		}
	}
	return "#" + locator
}

func mapDLocatorToHref(relationshipSetCurrentlyViewing string, definition *hydratables.DefinitionLinkbase, locator string) string {
	roleRefs := definition.RoleRefs
	for _, roleRef := range roleRefs {
		if relationshipSetCurrentlyViewing == roleRef.RoleURI {
			definitionLinks := definition.DefinitionLinks
			for _, definitionLink := range definitionLinks {
				if definitionLink.Role == relationshipSetCurrentlyViewing {
					for _, loc := range definitionLink.Locs {
						if loc.Label == locator {
							i := strings.Index(loc.Href, "#")
							if i >= 0 {
								return loc.Href
							}
						}
					}
				}
			}
		}
	}
	return "#" + locator
}
