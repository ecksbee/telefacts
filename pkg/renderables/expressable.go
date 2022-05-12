package renderables

import "ecksbee.com/telefacts/pkg/hydratables"

type Expressable struct {
}

func getExpressions(h *hydratables.Hydratable) (map[string]Expressable, error) {
	ret := make(map[string]Expressable)

	//todo
	return ret, nil
}
