package hydratables

import (
	"fmt"
	"math"
	"strconv"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
)

type LabelLinkLabel struct {
	Label    string
	Role     string
	Lang     string
	CharData string
}

type LabelArc struct {
	Order   float64
	Arcrole string
	From    string
	To      string
}

type LabelLink struct {
	Role      string
	Locs      []Loc
	Labels    []LabelLinkLabel
	LabelArcs []LabelArc
}

type LabelLinkbase struct {
	FileName  string
	RoleRefs  []RoleRef
	LabelLink []LabelLink
}

func HydrateLabelLinkbase(file *serializables.LabelLinkbaseFile, fileName string) (*LabelLinkbase, error) {
	if len(fileName) <= 0 {
		return nil, fmt.Errorf("Empty file name")
	}
	if file == nil {
		return nil, fmt.Errorf("Empty file")
	}
	ret := LabelLinkbase{}
	ret.FileName = fileName
	ret.RoleRefs = hydrateLabelLinkbaseRoleRefs(file)
	ret.LabelLink = hydrateLabelLink(file)
	return &ret, nil
}

func hydrateLabelLinkbaseRoleRefs(linkbaseFile *serializables.LabelLinkbaseFile) []RoleRef {
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

func hydrateLabelLink(linkbaseFile *serializables.LabelLinkbaseFile) []LabelLink {
	ret := make([]LabelLink, 0, len(linkbaseFile.LabelLink))
	for _, link := range linkbaseFile.LabelLink {
		typeAttr := attr.FindAttr(link.XMLAttrs, "type")
		if typeAttr == nil || typeAttr.Name.Space != attr.XLINK || typeAttr.Value != "extended" {
			continue
		}
		roleAttr := attr.FindAttr(link.XMLAttrs, "role")
		if roleAttr == nil || roleAttr.Value == "" {
			continue
		}
		newLink := LabelLink{}
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
		newLink.LabelArcs = make([]LabelArc, 0, len(link.LabelArc))
		for _, arc := range link.LabelArc {
			newArc := LabelArc{}
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
			order, err := strconv.ParseFloat(orderAttr.Value, 64)
			if err != nil {
				order = math.MaxFloat64
			}
			newArc.Arcrole = arcroleAttr.Value
			newArc.Order = order
			newArc.From = fromAttr.Value
			newArc.To = toAttr.Value
			newLink.LabelArcs = append(newLink.LabelArcs, newArc)
		}
		newLink.Labels = make([]LabelLinkLabel, 0, len(link.Label))
		for _, label := range link.Label {
			newLabel := LabelLinkLabel{}
			labelAttr := attr.FindAttr(label.XMLAttrs, "label")
			if labelAttr == nil || labelAttr.Value == "" || labelAttr.Name.Space != attr.XLINK {
				continue
			}
			roleAttr := attr.FindAttr(label.XMLAttrs, "role")
			if roleAttr == nil || roleAttr.Value == "" || roleAttr.Name.Space != attr.XLINK {
				continue
			}
			langAttr := attr.FindAttr(label.XMLAttrs, "lang")
			if langAttr == nil || langAttr.Value == "" {
				continue
			}
			newLabel.Label = labelAttr.Value
			newLabel.Role = roleAttr.Value
			newLabel.Lang = langAttr.Value
			newLabel.CharData = label.CharData
			newLink.Labels = append(newLink.Labels, newLabel)
		}
		ret = append(ret, newLink)
	}
	return ret
}
