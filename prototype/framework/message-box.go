package framework

import (
	"fmt"
)

type messageBox struct {
	decochar rune
}

func NewMessageBox(c rune) Product {
	return &messageBox{decochar: c}
}

func (u *messageBox) Use(s string) {
	for i := 0; i-2 < len(s); i++ {
		fmt.Printf("%c", u.decochar)
	}

	fmt.Println("")
	fmt.Printf("%c%s%c\n", u.decochar, s, u.decochar)

	for i := 0; i-2 < len(s); i++ {
		fmt.Printf("%c", u.decochar)
	}
	fmt.Println("")
}

func (u *messageBox) createClone() Product {
	return NewMessageBox(u.decochar)
}
