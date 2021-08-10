package framework

type Product interface {
	Use(s string)
	createClone() Product
}
