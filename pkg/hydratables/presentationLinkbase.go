package hydratables

import (
	"fmt"
	"math"
	"strconv"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
)

type PresentationArc struct {
	Order          float64
	Arcrole        string
	From           string
	To             string
	PreferredLabel string
}

type PresentationLink struct {
	Role             string
	Locs             []Loc
	PresentationArcs []PresentationArc
}

type PresentationLinkbase struct {
	FileName          string
	RoleRefs          []RoleRef
	PresentationLinks []PresentationLink
}

func HydratePresentationLinkbase(file *serializables.PresentationLinkbaseFile, fileName string) (*PresentationLinkbase, error) {
	if len(fileName) <= 0 {
		return nil, fmt.Errorf("empty file name")
	}
	if file == nil {
		return nil, fmt.Errorf("empty file")
	}
	ret := PresentationLinkbase{}
	ret.FileName = fileName
	ret.RoleRefs = hydratePresentationLinkbaseRoleRefs(file)
	ret.PresentationLinks = hydratePresentationLink(file)
	return &ret, nil
}

func hydratePresentationLinkbaseRoleRefs(linkbaseFile *serializables.PresentationLinkbaseFile) []RoleRef {
	ret := make([]RoleRef, 0, len(linkbaseFile.RoleRef))
	for _, roleRef := range linkbaseFile.RoleRef {
		if roleRef.XMLName.Space != attr.LINK {
			continue
		}
		roleURIAttr := attr.FindAttr(roleRef.XMLAttrs, "roleURI")
		if roleURIAttr == nil || roleURIAttr.Value == "" {
			continue
		}
		hrefAttr := attr.FindAttr(roleRef.XMLAttrs, "href")
		if hrefAttr == nil || hrefAttr.Value == "" {
			continue
		}
		if hrefAttr.Name.Space != attr.XLINK {
			continue
		}
		newRoleRef := RoleRef{
			RoleURI: roleURIAttr.Value,
			Href:    hrefAttr.Value,
		}
		ret = append(ret, newRoleRef)
	}
	return ret
}

func hydratePresentationLink(linkbaseFile *serializables.PresentationLinkbaseFile) []PresentationLink {
	ret := make([]PresentationLink, 0, len(linkbaseFile.PresentationLink))
	for _, link := range linkbaseFile.PresentationLink {
		typeAttr := attr.FindAttr(link.XMLAttrs, "type")
		if typeAttr == nil || typeAttr.Name.Space != attr.XLINK || typeAttr.Value != "extended" {
			continue
		}
		roleAttr := attr.FindAttr(link.XMLAttrs, "role")
		if roleAttr == nil || roleAttr.Value == "" {
			continue
		}
		newLink := PresentationLink{}
		newLink.Role = roleAttr.Value
		newLink.Locs = make([]Loc, 0, len(link.Loc))
		for _, loc := range link.Loc {
			newLoc := Loc{}
			ttypeAttr := attr.FindAttr(loc.XMLAttrs, "type")
			if ttypeAttr == nil || ttypeAttr.Name.Space != attr.XLINK || ttypeAttr.Value != "locator" {
				continue
			}
			labelAttr := attr.FindAttr(loc.XMLAttrs, "label")
			if labelAttr == nil || labelAttr.Name.Space != attr.XLINK || labelAttr.Value == "" {
				continue
			}
			hrefAttr := attr.FindAttr(loc.XMLAttrs, "href")
			if hrefAttr == nil || hrefAttr.Value == "" {
				continue
			}
			newLoc.Href = hrefAttr.Value
			newLoc.Label = labelAttr.Value
			newLink.Locs = append(newLink.Locs, newLoc)
		}
		newLink.PresentationArcs = make([]PresentationArc, 0, len(link.PresentationArc))
		for _, arc := range link.PresentationArc {
			newArc := PresentationArc{}
			ttypeAttr := attr.FindAttr(arc.XMLAttrs, "type")
			if ttypeAttr == nil || ttypeAttr.Name.Space != attr.XLINK || ttypeAttr.Value != "arc" {
				continue
			}
			orderAttr := attr.FindAttr(arc.XMLAttrs, "order")
			if orderAttr == nil || orderAttr.Value == "" {
				continue
			}
			arcroleAttr := attr.FindAttr(arc.XMLAttrs, "arcrole")
			if arcroleAttr == nil || arcroleAttr.Name.Space != attr.XLINK || arcroleAttr.Value == "" {
				continue
			}
			fromAttr := attr.FindAttr(arc.XMLAttrs, "from")
			if fromAttr == nil || fromAttr.Name.Space != attr.XLINK || fromAttr.Value == "" {
				continue
			}
			toAttr := attr.FindAttr(arc.XMLAttrs, "to")
			if toAttr == nil || toAttr.Name.Space != attr.XLINK || toAttr.Value == "" {
				continue
			}
			newArc.Arcrole = arcroleAttr.Value
			order, err := strconv.ParseFloat(orderAttr.Value, 64)
			if err != nil {
				order = math.MaxFloat64
			}
			newArc.Order = order
			preferredLabelAttr := attr.FindAttr(arc.XMLAttrs, "preferredLabel")
			if preferredLabelAttr != nil {
				newArc.PreferredLabel = preferredLabelAttr.Value
			}
			newArc.From = fromAttr.Value
			newArc.To = toAttr.Value
			newLink.PresentationArcs = append(newLink.PresentationArcs, newArc)
		}
		ret = append(ret, newLink)
	}
	return dedupPresentationLink(ret)
}

func dedupPresentationLink(links []PresentationLink) []PresentationLink {
	occured := map[string]PresentationLink{}
	for _, link := range links {
		if merged, found := occured[link.Role]; found {
			merged.Locs = append(merged.Locs, link.Locs...)
			merged.PresentationArcs = append(merged.PresentationArcs, link.PresentationArcs...)
			occured[link.Role] = merged
		} else {
			occured[link.Role] = link
		}
	}
	ret := make([]PresentationLink, 0)
	for _, r := range occured {
		ret = append(ret, r)
	}
	return ret
}
