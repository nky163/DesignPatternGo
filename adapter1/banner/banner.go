package banner

import "fmt"

type Banner struct {
	string_ string
}

func NewBanner(string_ string) *Banner {
	return &Banner{string_: string_}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.string_)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.string_)
}
