package symbols

type Env interface {
	Put(s string, sym Symbol)
	Get(s string) Symbol
}

func NewEnv(p Env) Env {
	return &envImpl{
		table: make(map[string]Symbol),
		prev:  p,
	}
}

type envImpl struct {
	table map[string]Symbol
	prev  Env
}

func (e *envImpl) Put(s string, sym Symbol) {
	e.table[s] = sym
}

func (e *envImpl) Get(s string) Symbol {
	for env := e; env != nil; env = e.prev.(*envImpl) {
		if sym, ok := e.table[s]; ok {
			return sym
		}
	}
	return nil
}
