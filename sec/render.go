package sec

import (
	"fmt"

	"ecks-bee.com/telefacts/renderables"
)

func (p *SECProject) RenderCatalog(workingDir string) ([]byte, error) {
	schema, err := p.Schema(workingDir)
	if err != nil {
		return nil, err
	}
	instance, err := p.Instance(workingDir)
	if err != nil {
		return nil, err
	}
	return renderables.MarshalCatalog(schema, instance)
}

func (p *SECProject) RenderDataGrid(workingDir string, network string, entity int, relationshipSet int) ([]byte, error) {
	schema, err := p.Schema(workingDir)
	if err != nil {
		return nil, err
	}
	instance, err := p.Instance(workingDir)
	if err != nil {
		return nil, err
	}
	presentation, err := p.PresentationLinkbase(workingDir)
	if err != nil {
		return nil, err
	}
	calculation, err := p.CalculationLinkbase(workingDir)
	if err != nil {
		return nil, err
	}
	definition, err := p.DefinitionLinkbase(workingDir)
	if err != nil {
		return nil, err
	}
	factFinder := p.NewFactFinder(workingDir)
	var bytes []byte
	switch network {
	case "pre":
		bytes, err = renderables.MarshalPGrid(entity, relationshipSet, schema, instance, presentation, factFinder)
	case "cal":
		bytes, err = renderables.MarshalCGrid(entity, relationshipSet, schema, instance, calculation, factFinder)
	case "def":
		bytes, err = renderables.MarshalDGrid(entity, relationshipSet, schema, instance, definition, factFinder)
	default:
		return nil, fmt.Errorf("invalid network: %s", network)
	}
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
