package main

import "fmt"

type iAlpha interface {
	work()
	common(iAlpha)
}

type alpha struct {
	name string
}

func (a *alpha) common(i iAlpha) {
	fmt.Println("common called")
	i.work()
}

type beta struct {
	alpha
}

func (b *beta) work() {
	fmt.Println("work called")
	fmt.Printf("Name is %s\n", b.name)
}

func main() {
	a := alpha{name: "test"}
	b := &beta{alpha: a}
	b.common(b)
}
