package idcard

import (
	"factory-method/framework"
	"fmt"
)

// ConcreteProduct
type IDCard struct {
	owner string
}

func (i *IDCard) GetOwner() string {
	return i.owner
}

// implement ProductInterface
func (i *IDCard) Use() {
	fmt.Printf("use %s's card.\n", i.GetOwner())
}

// 外部に公開しない（Newさせる）
type iDCardFactory struct {
	owners            []string
	framework.Factory // FactoryがもつfactoryInterfaceを満たすこと、、、
}

// implement factoryInterface
func (cf *iDCardFactory) CreateProduct(owner string) framework.ProductInterface {
	return &IDCard{owner: owner}
}

// implement factoryInterface
func (cf *iDCardFactory) RegisterProduct(framework.ProductInterface) {
	// owner := framework.ProductInterface.GetOwner()
	// cf.owners = append(cf.owners, owner)
}

// ConcreteFactory
func NewIDCardCreator() framework.Creator {
	f := framework.Factory{}
	cf := &iDCardFactory{
		Factory: f,
		owners:  []string{},
	}
	cf.Factory.FactoryInterface = cf
	return cf
}
