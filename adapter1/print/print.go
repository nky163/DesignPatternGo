package print

import "adapter1/banner"

type Printer interface {
	PrintWeak()
	PrintStrong()
}

type Print struct {
	*banner.Banner // 継承的な
}

func NewPrint(string_ string) *Print {
	p := &Print{banner.NewBanner(string_)}
	return p
}

func (p *Print) PrintWeak() {
	p.ShowWithParen()
}

func (p *Print) PrintStrong() {
	p.ShowWithAster()
}
