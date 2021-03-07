package hydratables

type Fact struct {
	Href       string
	ID         string
	ContextRef string
	UnitRef    string
	Decimals   string
	Precision  string
	IsNil      bool
	XMLInner   string
}

func (h *Hydratable) FindFact(href string, contextRef string) *Fact {
	for _, ins := range h.Instances {
		for _, fact := range ins.Facts {
			if fact.Href == href && fact.ContextRef == contextRef {
				return &fact
			}
		}
	}
	return nil
}
