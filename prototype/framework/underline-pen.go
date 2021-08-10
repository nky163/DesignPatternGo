package framework

import (
	"fmt"
)

type underlinePen struct {
	ulchar rune
}

func NewUnderlinePen(c rune) Product {
	return &underlinePen{ulchar: c}
}

func (u *underlinePen) Use(s string) {
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", u.ulchar)
	}
	fmt.Println("")
}

func (u *underlinePen) createClone() Product {
	return NewUnderlinePen(u.ulchar)
}
