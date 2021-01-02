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
	factFinder := p.NewFactFinder(workingDir)
	var bytes []byte
	switch network {
	case "pre":
		bytes, err = renderables.MarshalPGrid(entity, relationshipSet, schema, instance, presentation, factFinder)
	default:
		return nil, fmt.Errorf("invalid network: %s", network)
	}
	return bytes, nil
}
