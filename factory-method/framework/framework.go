package framework

type Creator interface {
	Create(owner string) ProductInterface
}

type FactoryInterface interface {
	CreateProduct(owner string) ProductInterface
	RegisterProduct(ProductInterface)
}

type ProductInterface interface {
	Use()
	GetOwner() string
}

type Factory struct {
	FactoryInterface FactoryInterface
}

func (f *Factory) Create(owner string) ProductInterface {
	p := f.FactoryInterface.CreateProduct(owner)
	f.FactoryInterface.RegisterProduct(p)
	return p
}
