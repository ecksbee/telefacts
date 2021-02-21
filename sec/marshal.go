package sec

import (
	"ecksbee.com/telefacts/renderables"
)

func MarshalCatalog(workingDir string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	//todo get files
	filenames := []string{}
	return renderables.MarshalCatalog(h, filenames)
}

func Marshal(workingDir string, slug string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	//todo check if slug is passthrough
	hash := slug
	return renderables.MarshalRenderable(hash, h)
}

// func (h *hydratables.Hydratable) RenderCatalog(workingDir string) ([]byte, error) {
// 	schema, err := p.Schema(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	instance, err := p.Instance(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return renderables.MarshalCatalog(schema, instance)
// }

// func (p *SECProject) RenderDataGrid(workingDir string, network string, entity int, relationshipSet int) ([]byte, error) {
// 	schema, err := p.Schema(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	instance, err := p.Instance(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	presentation, err := p.PresentationLinkbase(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	calculation, err := p.CalculationLinkbase(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	definition, err := p.DefinitionLinkbase(workingDir)
// 	if err != nil {
// 		return nil, err
// 	}
// 	factFinder := p.NewFactFinder(workingDir)
// 	labelFinder := p.NewLabelFinder(workingDir)
// 	var bytes []byte
// 	switch network {
// 	case "pre":
// 		bytes, err = renderables.MarshalPGrid(entity, relationshipSet, schema, instance, presentation, factFinder, labelFinder)
// 	case "cal":
// 		bytes, err = renderables.MarshalCGrid(entity, relationshipSet, schema, instance, calculation, factFinder, labelFinder)
// 	case "def":
// 		bytes, err = renderables.MarshalDGrid(entity, relationshipSet, schema, instance, definition, factFinder, labelFinder)
// 	default:
// 		return nil, fmt.Errorf("invalid network: %s", network)
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes, nil
// }
