package main

import "fmt"

// interface
type iAlpha interface {
	work()
	common()
}

// abstract concreat struct
type alpha struct {
	name string
}

func (a *alpha) common() {
	fmt.Println("common called")
}

// implementor struct
type beta struct {
	*alpha
}

// bata struct is able to access the default method "common" and field "name"
func (b *beta) work() {
	fmt.Println("work called")
	fmt.Printf("name is %s\n", b.name)
	b.common()
}

// iAlphaインターフェイスはalphaから直接作れない alpha はcommon()しか持っていないので
// Alpha はabstract concrete
// ※問題はbのcommonの中でworkメソッドを使えない
func main() {
	a := &alpha{name: "test"}
	b := &beta{alpha: a}
	b.work()
	b.common()
}
