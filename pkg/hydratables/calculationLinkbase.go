package hydratables

import (
	"fmt"
	"math"
	"strconv"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
)

type CalculationArc struct {
	Order   float64
	Arcrole string
	From    string
	To      string
	Weight  float64
}

type CalculationLink struct {
	Role            string
	Locs            []Loc
	CalculationArcs []CalculationArc
}

type CalculationLinkbase struct {
	FileName         string
	RoleRefs         []RoleRef
	CalculationLinks []CalculationLink
}

func HydrateCalculationLinkbase(file *serializables.CalculationLinkbaseFile, fileName string) (*CalculationLinkbase, error) {
	if len(fileName) <= 0 {
		return nil, fmt.Errorf("empty file name")
	}
	if file == nil {
		return nil, fmt.Errorf("empty file")
	}
	ret := CalculationLinkbase{}
	ret.FileName = fileName
	ret.RoleRefs = hydrateCalculationLinkbaseRoleRefs(file)
	ret.CalculationLinks = hydrateCalculationLink(file)
	return &ret, nil
}

func hydrateCalculationLinkbaseRoleRefs(linkbaseFile *serializables.CalculationLinkbaseFile) []RoleRef {
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

func hydrateCalculationLink(linkbaseFile *serializables.CalculationLinkbaseFile) []CalculationLink {
	ret := make([]CalculationLink, 0, len(linkbaseFile.CalculationLink))
	for _, link := range linkbaseFile.CalculationLink {
		typeAttr := attr.FindAttr(link.XMLAttrs, "type")
		if typeAttr == nil || typeAttr.Name.Space != attr.XLINK || typeAttr.Value != "extended" {
			continue
		}
		roleAttr := attr.FindAttr(link.XMLAttrs, "role")
		if roleAttr == nil || roleAttr.Value == "" {
			continue
		}
		newLink := CalculationLink{}
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
		newLink.CalculationArcs = make([]CalculationArc, 0, len(link.CalculationArc))
		for _, arc := range link.CalculationArc {
			newArc := CalculationArc{}
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
			weightAttr := attr.FindAttr(arc.XMLAttrs, "weight")
			if weightAttr == nil || weightAttr.Value == "" {
				continue
			}
			order, err := strconv.ParseFloat(orderAttr.Value, 64)
			if err != nil {
				order = math.MaxFloat64
			}
			weight, err := strconv.ParseFloat(weightAttr.Value, 64)
			if err != nil {
				weight = 0.0
			}
			newArc.Arcrole = arcroleAttr.Value
			newArc.Order = order
			newArc.From = fromAttr.Value
			newArc.To = toAttr.Value
			newArc.Weight = weight
			newLink.CalculationArcs = append(newLink.CalculationArcs, newArc)
		}
		ret = append(ret, newLink)
	}
	return ret
}
