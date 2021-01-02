package sec

import (
	"ecks-bee.com/telefacts/xbrl"
)

type FactFinder struct {
	Dir     string
	Project *SECProject
}

func (p *SECProject) NewFactFinder(workingDir string) *FactFinder {
	return &FactFinder{
		Dir:     workingDir,
		Project: p,
	}
}

func (f *FactFinder) FindFact(href string, contextRef string) *xbrl.Fact { //todo return err
	schema, err := f.Project.Schema(f.Dir)
	if err != nil {
		return nil
	}
	namespace, concept, err := xbrl.HashQuery(schema, href)
	if err != nil {
		return nil
	}
	ins, err := f.Project.Instance(f.Dir)
	facts := ins.Facts
	for _, fact := range facts {
		if namespace == fact.XMLName.Space && concept.Name == fact.XMLName.Local {
			return &fact
		}
	}
	return nil
}
