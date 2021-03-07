package hydratables

import (
	"ecksbee.com/telefacts/pkg/serializables"
)

type Hydratable struct {
	Folder *serializables.Folder
	UnitTypeRegistry
	Instances             map[string]Instance
	Schemas               map[string]Schema
	LabelLinkbases        map[string]LabelLinkbase
	PresentationLinkbases map[string]PresentationLinkbase
	DefinitionLinkbases   map[string]DefinitionLinkbase
	CalculationLinkbases  map[string]CalculationLinkbase
}

func Hydrate(folder *serializables.Folder) (*Hydratable, error) {
	ret := &Hydratable{
		Folder:                folder,
		Instances:             make(map[string]Instance),
		Schemas:               make(map[string]Schema),
		LabelLinkbases:        make(map[string]LabelLinkbase),
		PresentationLinkbases: make(map[string]PresentationLinkbase),
		DefinitionLinkbases:   make(map[string]DefinitionLinkbase),
		CalculationLinkbases:  make(map[string]CalculationLinkbase),
	}
	for filename, file := range folder.Schemas {
		entry, err := HydrateSchema(&file, filename)
		if err != nil {
			return nil, err
		}
		ret.Schemas[filename] = *entry
	}
	for filename, file := range folder.PresentationLinkbases {
		entry, err := HydratePresentationLinkbase(&file, filename)
		if err != nil {
			return nil, err
		}
		ret.PresentationLinkbases[filename] = *entry
	}
	for filename, file := range folder.DefinitionLinkbases {
		entry, err := HydrateDefinitionLinkbase(&file, filename)
		if err != nil {
			return nil, err
		}
		ret.DefinitionLinkbases[filename] = *entry
	}
	for filename, file := range folder.CalculationLinkbases {
		entry, err := HydrateCalculationLinkbase(&file, filename)
		if err != nil {
			return nil, err
		}
		ret.CalculationLinkbases[filename] = *entry
	}
	for filename, file := range folder.LabelLinkbases {
		entry, err := HydrateLabelLinkbase(&file, filename)
		if err != nil {
			return nil, err
		}
		ret.LabelLinkbases[filename] = *entry
	}
	for filename, file := range folder.Instances {
		entry, err := HydrateInstance(&file, filename, ret)
		if err != nil {
			return nil, err
		}
		ret.Instances[filename] = *entry
	}
	utr, err := HydrateUnitTypeRegistry()
	if err != nil {
		return nil, err
	}
	ret.UnitTypeRegistry = *utr
	return ret, nil
}
