package sec

import (
	"io/ioutil"
	"sync"

	"ecks-bee.com/telefacts/xbrl"
	gocache "github.com/patrickmn/go-cache"
)

type SECProject struct {
	ID       string
	AppCache *gocache.Cache
	Lock     sync.RWMutex
}

func (p *SECProject) Schema(workingDir string) (*xbrl.Schema, error) {
	p.Lock.RLock()
	if x, found := p.AppCache.Get(p.ID + "/schema"); found {
		p.Lock.RUnlock()
		ret := x.(xbrl.Schema)
		return &ret, nil
	}
	files, err := ioutil.ReadDir(workingDir)
	schemaFile, err := getSchemaFromOSfiles(files)
	p.Lock.RUnlock()
	if err != nil {
		return nil, err
	}
	schema, err := xbrl.ReadSchema(schemaFile, workingDir)
	if err != nil {
		return nil, err
	}
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		p.AppCache.Set(p.ID+"/schema", *schema, gocache.DefaultExpiration)
	}()
	return schema, nil
}

func (p *SECProject) SaveSchema(destination string, unmarshalled *xbrl.Schema) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	go func() {
		p.AppCache.Set(p.ID+"/schema", *unmarshalled, gocache.DefaultExpiration)
	}()
	return commitSchema(destination, unmarshalled)
}

func (p *SECProject) Instance(workingDir string) (*xbrl.Instance, error) {
	p.Lock.RLock()
	if x, found := p.AppCache.Get(p.ID + "/instance"); found {
		p.Lock.RUnlock()
		ret := x.(xbrl.Instance)
		return &ret, nil
	}
	files, err := ioutil.ReadDir(workingDir)
	insFile, err := getInstanceFromOSfiles(files)
	p.Lock.RUnlock()
	if err != nil {
		return nil, err
	}
	ins, err := xbrl.ReadInstance(insFile, workingDir)
	if err != nil {
		return nil, err
	}
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		p.AppCache.Set(p.ID+"/instance", *ins, gocache.DefaultExpiration)
	}()
	return ins, nil
}

func (p *SECProject) SaveInstance(destination string, unmarshalled *xbrl.Instance) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	go func() {
		p.AppCache.Set(p.ID+"/instance", &unmarshalled, gocache.DefaultExpiration)
	}()
	return commitInstance(destination, unmarshalled)
}

func (p *SECProject) PresentationLinkbase(workingDir string) (*xbrl.PresentationLinkbase, error) {
	p.Lock.RLock()
	if x, found := p.AppCache.Get(p.ID + "/presentation"); found {
		p.Lock.RUnlock()
		ret := x.(xbrl.PresentationLinkbase)
		return &ret, nil
	}
	files, err := ioutil.ReadDir(workingDir)
	preFile, err := getPresentationLinkbaseFromOSfiles(files)
	p.Lock.RUnlock()
	if err != nil {
		return nil, err
	}
	pre, err := xbrl.ReadPresentationLinkbase(preFile, workingDir)
	if err != nil {
		return nil, err
	}
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		p.AppCache.Set(p.ID+"/presentation", *pre, gocache.DefaultExpiration)
	}()
	return pre, nil
}

func (p *SECProject) SavePresentationLinkbase(destination string, unmarshalled *xbrl.PresentationLinkbase) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	go func() {
		p.AppCache.Set(p.ID+"/presentation", &unmarshalled, gocache.DefaultExpiration)
	}()
	return commitPresentationLinkbase(destination, unmarshalled)
}

func (p *SECProject) DefinitionLinkbase(workingDir string) (*xbrl.DefinitionLinkbase, error) {
	p.Lock.RLock()
	if x, found := p.AppCache.Get(p.ID + "/definition"); found {
		p.Lock.RUnlock()
		ret := x.(xbrl.DefinitionLinkbase)
		return &ret, nil
	}
	files, err := ioutil.ReadDir(workingDir)
	defFile, err := getDefinitionLinkbaseFromOSfiles(files)
	p.Lock.RUnlock()
	if err != nil {
		return nil, err
	}
	def, err := xbrl.ReadDefinitionLinkbase(defFile, workingDir)
	if err != nil {
		return nil, err
	}
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		p.AppCache.Set(p.ID+"/definition", *def, gocache.DefaultExpiration)
	}()
	return def, nil
}

func (p *SECProject) SaveDefinitionLinkbase(destination string, unmarshalled *xbrl.DefinitionLinkbase) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	go func() {
		p.AppCache.Set(p.ID+"/definition", &unmarshalled, gocache.DefaultExpiration)
	}()
	return commitDefinitionLinkbase(destination, unmarshalled)
}

func (p *SECProject) CalculationLinkbase(workingDir string) (*xbrl.CalculationLinkbase, error) {
	p.Lock.RLock()
	if x, found := p.AppCache.Get(p.ID + "/calculation"); found {
		p.Lock.RUnlock()
		ret := x.(xbrl.CalculationLinkbase)
		return &ret, nil
	}
	files, err := ioutil.ReadDir(workingDir)
	calcFile, err := getCalculationLinkbaseFromOSfiles(files)
	p.Lock.RUnlock()
	if err != nil {
		return nil, err
	}
	cal, err := xbrl.ReadCalculationLinkbase(calcFile, workingDir)
	if err != nil {
		return nil, err
	}
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		p.AppCache.Set(p.ID+"/calculation", *cal, gocache.DefaultExpiration)
	}()
	return cal, nil
}

func (p *SECProject) SaveCalculationLinkbase(destination string, unmarshalled *xbrl.CalculationLinkbase) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	go func() {
		p.AppCache.Set(p.ID+"/definition", &unmarshalled, gocache.DefaultExpiration)
	}()
	return commitCalculationLinkbase(destination, unmarshalled)
}

func (p *SECProject) LabelLinkbase(workingDir string) (*xbrl.LabelLinkbase, error) {
	if x, found := p.AppCache.Get(p.ID + "/label"); found {
		ret := x.(xbrl.LabelLinkbase)
		return &ret, nil
	}
	files, err := ioutil.ReadDir(workingDir)
	labelFile, err := getLabelLinkbaseFromOSfiles(files)
	if err != nil {
		return nil, err
	}
	label, err := xbrl.ReadLabelLinkbase(labelFile, workingDir)
	if err != nil {
		return nil, err
	}
	go func() {
		p.Lock.Lock()
		defer p.Lock.Unlock()
		p.AppCache.Set(p.ID+"/label", *label, gocache.DefaultExpiration)
	}()
	return label, nil
}

func (p *SECProject) SaveLabelLinkbase(destination string, unmarshalled *xbrl.LabelLinkbase) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	go func() {
		p.AppCache.Set(p.ID+"/label", &unmarshalled, gocache.DefaultExpiration)
	}()
	return commitLabelLinkbase(destination, unmarshalled)
}
