package hydratables

func (h *Hydratable) FindFact(href string, contextRef string) *Fact { //todo return err
	for _, ins := range h.Instances {
		for _, fact := range ins.Facts {
			if fact.Href == href && fact.ContextRef == contextRef {
				return &fact
			}
		}
	}
	return nil
}
