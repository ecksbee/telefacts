package hydratables

import (
	"encoding/xml"

	"ecksbee.com/telefacts/pkg/serializables"
)

type Unit struct {
	ID      string
	Measure UnitMeasure
	Divide  UnitDivide
}

type UnitMeasure struct {
	XMLName  xml.Name
	CharData string
}

type UnitDivide struct {
	UnitNumerator struct {
		Measure UnitMeasure
	}
	UnitDenominator struct {
		Measure UnitMeasure
	}
}

type Measurement struct {
	UnitID                 string
	UnitName               string
	NSUnit                 string
	ItemType               string
	NsItemType             string
	ItemTypeDate           string
	Symbol                 string
	Definition             string
	BaseStandard           string
	ConversionPresentation string
	ConversionContent      string
	Status                 string
	VersionDate            string
	NumeratorItemType      string
	NSNumeratorItemType    string
	DenominatorItemType    string
	NSDenominatorItemType  string
}

type UnitTypeRegistry map[string]map[string]Measurement

func mapMeasurements(file *serializables.UnitTypeRegistryFile) UnitTypeRegistry {
	ret := make(map[string]map[string]Measurement)
	if len(file.Units) > 0 && len(file.Units[0].Unit) > 0 {
		utr := file.Units[0]
		for _, unit := range utr.Unit {
			if len(unit.UnitID) <= 0 || len(unit.UnitName) <= 0 || len(unit.NSUnit) <= 0 {
				continue
			}
			id := unit.UnitID[0].CharData
			name := unit.UnitName[0].CharData
			ns := unit.NSUnit[0].CharData
			if id == "" || name == "" || ns == "" {
				continue
			}
			newItem := Measurement{
				UnitID:   id,
				UnitName: name,
				NSUnit:   ns,
			}
			if len(unit.Status) > 0 {
				newItem.Status = unit.Status[0].CharData
			}
			if len(unit.VersionDate) > 0 {
				newItem.VersionDate = unit.VersionDate[0].CharData
			}
			if len(unit.ItemType) > 0 {
				newItem.ItemType = unit.ItemType[0].CharData
			}
			if len(unit.NSItemType) > 0 {
				newItem.NsItemType = unit.NSItemType[0].CharData
			}
			if len(unit.ItemTypeDate) > 0 {
				newItem.ItemTypeDate = unit.ItemTypeDate[0].CharData
			}
			if len(unit.Symbol) > 0 {
				newItem.Symbol = unit.Symbol[0].CharData
			}
			if len(unit.Definition) > 0 {
				newItem.Definition = unit.Definition[0].CharData
			}
			if len(unit.BaseStandard) > 0 {
				newItem.BaseStandard = unit.BaseStandard[0].CharData
			}
			if len(unit.ConversionPresentation) > 0 {
				newItem.ConversionPresentation = unit.ConversionPresentation[0].XMLInner
			}
			if len(unit.ConversionContent) > 0 {
				newItem.ConversionContent = unit.ConversionContent[0].XMLInner
			}
			if len(unit.NumeratorItemType) > 0 {
				newItem.NumeratorItemType = unit.NumeratorItemType[0].CharData
			}
			if len(unit.NSNumeratorItemType) > 0 {
				newItem.NSNumeratorItemType = unit.NSNumeratorItemType[0].CharData
			}
			if len(unit.DenominatorItemType) > 0 {
				newItem.DenominatorItemType = unit.DenominatorItemType[0].CharData
			}
			if len(unit.NSDenominatorItemType) > 0 {
				newItem.NSDenominatorItemType = unit.NSDenominatorItemType[0].CharData
			}
			if _, found := ret[ns]; !found {
				ret[ns] = make(map[string]Measurement)
			}
			ret[ns][id] = newItem
		}
	}
	return ret
}

func (h *Hydratable) queryUTR(namespace string, id string) *Measurement {
	if units, found := h.UnitTypeRegistry[namespace]; found {
		if unit, ffound := units[id]; ffound {
			return &unit
		}
	}
	return nil
}

func (h *Hydratable) FindMeasurement(unitRef string) (*Measurement, *Measurement) {
	for _, ins := range h.Instances {
		for _, unit := range ins.Units {
			if unit.ID == unitRef {
				if unit.Measure.CharData == "" {
					if unit.Divide.UnitNumerator.Measure.CharData == "" ||
						unit.Divide.UnitDenominator.Measure.CharData == "" {
						return nil, nil
					}
					return h.queryUTR(unit.Divide.UnitNumerator.Measure.XMLName.Space,
							unit.Divide.UnitNumerator.Measure.XMLName.Local),
						h.queryUTR(unit.Divide.UnitDenominator.Measure.XMLName.Space,
							unit.Divide.UnitNumerator.Measure.XMLName.Local)
				}
				return h.queryUTR(unit.Measure.XMLName.Space, unit.Measure.XMLName.Local), nil
			}
		}
	}
	return nil, nil
}
