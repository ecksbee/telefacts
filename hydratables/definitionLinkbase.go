package hydratables

import (
	"fmt"
	"math"
	"strconv"

	"ecksbee.com/telefacts/attr"
	"ecksbee.com/telefacts/serializables"
)

type DefinitionArc struct {
	Order          float64
	Arcrole        string
	From           string
	To             string
	Closed         bool
	ContextElement string
	TargetRole     string
}

type DefinitionLink struct {
	Role           string
	Locs           []Loc
	DefinitionArcs []DefinitionArc
}

type DefinitionLinkbase struct {
	FileName        string
	RoleRefs        []RoleRef
	DefinitionLinks []DefinitionLink
}

func HydrateDefinitionLinkbase(file *serializables.DefinitionLinkbaseFile, fileName string) (*DefinitionLinkbase, error) {
	if len(fileName) <= 0 {
		return nil, fmt.Errorf("Empty file name")
	}
	if file == nil {
		return nil, fmt.Errorf("Empty file")
	}
	ret := DefinitionLinkbase{}
	ret.FileName = fileName
	ret.RoleRefs = hydrateDefinitionLinkbaseRoleRefs(file)
	ret.DefinitionLinks = hydrateDefinitionLink(file)
	return &ret, nil
}

func hydrateDefinitionLinkbaseRoleRefs(linkbaseFile *serializables.DefinitionLinkbaseFile) []RoleRef {
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

func hydrateDefinitionLink(linkbaseFile *serializables.DefinitionLinkbaseFile) []DefinitionLink {
	ret := make([]DefinitionLink, 0, len(linkbaseFile.DefinitionLink))
	for _, link := range linkbaseFile.DefinitionLink {
		typeAttr := attr.FindAttr(link.XMLAttrs, "type")
		if typeAttr == nil || typeAttr.Name.Space != attr.XLINK || typeAttr.Value != "extended" {
			continue
		}
		roleAttr := attr.FindAttr(link.XMLAttrs, "role")
		if roleAttr == nil || roleAttr.Value == "" {
			continue
		}
		newLink := DefinitionLink{}
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
		newLink.DefinitionArcs = make([]DefinitionArc, 0, len(link.DefinitionArc))
		for _, arc := range link.DefinitionArc {
			newArc := DefinitionArc{}
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
			newArc.From = fromAttr.Value
			newArc.To = toAttr.Value
			closedAttr := attr.FindAttr(arc.XMLAttrs, "closed")
			if closedAttr != nil {
				closed, err := strconv.ParseBool(closedAttr.Value)
				if err == nil {
					closed = false
				}
				newArc.Closed = closed
			}
			contextElementAttr := attr.FindAttr(arc.XMLAttrs, "contextElement")
			if contextElementAttr != nil {
				newArc.ContextElement = contextElementAttr.Value
			}
			targetRoleAttr := attr.FindAttr(arc.XMLAttrs, "targetRole")
			if targetRoleAttr != nil {
				newArc.TargetRole = targetRoleAttr.Value
			}
			newLink.DefinitionArcs = append(newLink.DefinitionArcs, newArc)
		}
		ret = append(ret, newLink)
	}
	return ret
}
