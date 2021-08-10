package framework

type manager struct {
	showcase map[string]Product
}

func NewManager() *manager {
	return &manager{showcase: make(map[string]Product)}
}
func (m *manager) Register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *manager) Create(protpname string) Product {
	p := m.showcase[protpname]
	return p.createClone()
}
