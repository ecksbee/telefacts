package renderables

import (
	"strings"

	"ecks-bee.com/telefacts/xbrl"
)

func mapPLocatorToHref(relationshipSetCurrentlyViewing string, presentation *xbrl.PresentationLinkbase, locator string) string {
	roleRefs := presentation.RoleRef
	for _, roleRef := range roleRefs {
		if relationshipSetCurrentlyViewing == roleRef.RoleURI {
			presentationLinks := presentation.PresentationLinks
			for _, presentationLink := range presentationLinks {
				if presentationLink.Role == relationshipSetCurrentlyViewing {
					for _, loc := range presentationLink.Loc {
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

func mapCLocatorToHref(relationshipSetCurrentlyViewing string, calculation *xbrl.CalculationLinkbase, locator string) string {
	roleRefs := calculation.RoleRef
	for _, roleRef := range roleRefs {
		if relationshipSetCurrentlyViewing == roleRef.RoleURI {
			calculationLinks := calculation.CalculationLinks
			for _, calculationLink := range calculationLinks {
				if calculationLink.Role == relationshipSetCurrentlyViewing {
					for _, loc := range calculationLink.Loc {
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
