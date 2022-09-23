package usecase

import (
	"fmt"
	"sync"
)

type Usecase struct {
	name      string
	devices   []string // devices required
	policy    string   // default policy
	endpoints []string // the backends' urls
}

type UsecaseSet struct {
	Usecases map[string]Usecase
	lock     sync.Mutex
}

type Request struct {
	Name    string
	Devices string // requested device priority
	Policy  string // requested policy
	Uri     string // data source uri
}

func NewRequest(name string, devices string, policy string, uri string) *Request {
	r := &Request{
		Name:    name,
		Devices: devices,
		Policy:  policy,
		Uri:     uri,
	}

	return r
}

func NewUsecase(name string, devices []string, policy string, endpoints []string) *Usecase {
	return &Usecase{
		name:      name,
		devices:   devices,
		policy:    policy,
		endpoints: endpoints,
	}
}

func NewUsecaseSet() *UsecaseSet {
	return &UsecaseSet{Usecases: make(map[string]Usecase)}
}

func (us *UsecaseSet) Delete(name string) {
	us.lock.Lock()
	delete(us.Usecases, name)
	us.lock.Unlock()
}

func (us *UsecaseSet) Update(u *Usecase) {
	us.lock.Lock()
	us.Usecases[u.name] = *u
	us.lock.Unlock()
}

func (us *UsecaseSet) LookUp(name string) (Usecase, bool) {
	us.lock.Lock()
	u, ok := us.Usecases[name]
	us.lock.Unlock()
	return u, ok
}

func (us *UsecaseSet) List() {
	us.lock.Lock()
	for k, v := range us.Usecases {
		fmt.Printf("usecase %s: %v\n", k, v)
	}
	us.lock.Unlock()
}
