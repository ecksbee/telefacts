package hydratables

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"ecksbee.com/telefacts/attr"
	"ecksbee.com/telefacts/serializables"
)

const COMMON_SCHEMA = `http://www.xbrl.org/2003/xbrl-instance-2003-12-31.XSD`

type Concept struct {
	XMLName           xml.Name
	ID                string
	Name              string
	Type              xml.Name
	SubstitutionGroup xml.Name
	Nillable          bool
	PeriodType        string
	Balance           string
	Abstract          bool
}

type Schema struct {
	FileName string
	Annotation
	Element []Concept
}

type Annotation struct {
	Appinfo
}

type Appinfo struct {
	RoleTypes []RoleType
}

type RoleType struct {
	RoleURI    string
	ID         string
	Definition string
	UsedOn     []struct {
		CharData string
	}
}

func HydrateSchema(file *serializables.SchemaFile, fileName string) (*Schema, error) {
	if len(fileName) <= 0 {
		return nil, fmt.Errorf("Empty file name")
	}
	if file == nil {
		return nil, fmt.Errorf("Empty file")
	}
	ret := Schema{}
	ret.Annotation = hydrateAnnotation(file)
	ret.Element = hydrateConcepts(file)
	return &ret, nil
}

func hydrateTargetNamespace(file *serializables.SchemaFile) string {
	attr := attr.FindAttr(file.XMLAttrs, "targetNamespace")
	if attr == nil {
		return ""
	}
	return attr.Value
}

func hydrateElementFormDefault(file *serializables.SchemaFile) string {
	attr := attr.FindAttr(file.XMLAttrs, "elementFormDefault")
	if attr == nil {
		return ""
	}
	return attr.Value
}

func hydrateAttributeFormDefault(file *serializables.SchemaFile) string {
	attr := attr.FindAttr(file.XMLAttrs, "attributeFormDefault")
	if attr == nil {
		return ""
	}
	return attr.Value
}

func hydrateAnnotation(file *serializables.SchemaFile) Annotation {
	ret := Annotation{}
	if len(file.Annotation) <= 0 {
		return ret
	}
	if file.Annotation[0].XMLName.Space != attr.XSD {
		return ret
	}
	if len(file.Annotation[0].Appinfo) <= 0 {
		return ret
	}
	if file.Annotation[0].Appinfo[0].XMLName.Space != attr.XSD {
		return ret
	}
	roleTypes := make([]RoleType, 0, len(file.Annotation[0].Appinfo[0].RoleType))
	for _, roleType := range file.Annotation[0].Appinfo[0].RoleType {
		if roleType.XMLName.Space != attr.LINK {
			continue
		}
		newRoleType := RoleType{}
		idAttr := attr.FindAttr(roleType.XMLAttrs, "id")
		if idAttr == nil || idAttr.Value == "" {
			continue
		}
		newRoleType.ID = idAttr.Value
		roleuriAttr := attr.FindAttr(roleType.XMLAttrs, "roleURI")
		if roleuriAttr == nil || roleuriAttr.Value == "" {
			continue
		}
		newRoleType.RoleURI = roleuriAttr.Value
		if len(roleType.Definition) > 0 {
			newRoleType.Definition = roleType.Definition[0].CharData
		}
		if len(roleType.UsedOn) > 0 {
			newRoleType.UsedOn = make([]struct{ CharData string }, 0, len(roleType.UsedOn))
			for _, usedOn := range roleType.UsedOn {
				if usedOn.XMLName.Space != attr.LINK {
					continue
				}
				newRoleType.UsedOn = append(newRoleType.UsedOn, struct{ CharData string }{
					CharData: usedOn.CharData,
				})
			}
		}
		roleTypes = append(roleTypes, newRoleType)
	}
	return Annotation{
		Appinfo: Appinfo{
			RoleTypes: roleTypes,
		},
	}
}

func hydrateConcepts(file *serializables.SchemaFile) []Concept {
	ret := make([]Concept, 0, len(file.Element))
	tlAttrs := file.XMLAttrs
	targetNS := hydrateTargetNamespace(file)
	for _, element := range file.Element {
		if element.XMLName.Space != attr.XSD {
			continue
		}
		idAttr := attr.FindAttr(element.XMLAttrs, "id")
		if idAttr == nil || idAttr.Value == "" {
			continue
		}
		nameAttr := attr.FindAttr(element.XMLAttrs, "name")
		if nameAttr == nil || nameAttr.Value == "" {
			continue
		}
		substitutionGroupAttr := attr.FindAttr(element.XMLAttrs, "substitutionGroup")
		if substitutionGroupAttr == nil || substitutionGroupAttr.Value == "" {
			continue
		}
		typeAttr := attr.FindAttr(element.XMLAttrs, "type")
		if typeAttr == nil || typeAttr.Value == "" {
			continue
		}
		isAbstract := false
		abstractAttr := attr.FindAttr(element.XMLAttrs, "abstract")
		if abstractAttr != nil {
			v, err := strconv.ParseBool(abstractAttr.Value)
			if err != nil {
				isAbstract = v
			}
		}
		balance := ""
		balanceAttr := attr.FindAttr(element.XMLAttrs, "balance")
		if balanceAttr != nil && balanceAttr.Name.Space == attr.XBRLI {
			balance = balanceAttr.Value
		}
		isNillable := false
		nillableAttr := attr.FindAttr(element.XMLAttrs, "nillable")
		if nillableAttr != nil {
			v, err := strconv.ParseBool(nillableAttr.Value)
			if err != nil {
				isNillable = v
			}
		}
		periodType := ""
		periodTypeAttr := attr.FindAttr(element.XMLAttrs, "periodType")
		if periodTypeAttr != nil && periodTypeAttr.Name.Space == attr.XBRLI {
			periodType = periodTypeAttr.Value
		}
		ret = append(ret, Concept{
			XMLName: xml.Name{
				Space: targetNS,
				Local: nameAttr.Value,
			},
			ID:                idAttr.Value,
			Type:              attr.Xmlns(tlAttrs, nameAttr.Value),
			SubstitutionGroup: attr.Xmlns(tlAttrs, substitutionGroupAttr.Value),
			Abstract:          isAbstract,
			Balance:           balance,
			Nillable:          isNillable,
			PeriodType:        periodType,
		})
	}
	return ret
}
