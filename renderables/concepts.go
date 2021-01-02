package renderables

import (
	"strings"

	"ecks-bee.com/telefacts/xbrl"
)

func mapLocatorToHref(linkbaseSelected string, relationshipSetCurrentlyViewing string, presentation *xbrl.PresentationLinkbase, locator string) string {
	switch linkbaseSelected {
	case "pre":
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
		break
	case "def":
		break
	case "cal":
		break
	default:
		return "!!!unknown linkbase!!!"
	}
	return "#" + locator
}
