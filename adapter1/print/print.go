package print

import "adapter1/banner"

type Printer interface {
	PrintWeak()
	PrintStrong()
}

type PrintBanner struct {
	*banner.Banner // 継承的な...
}

// Printerインターフェースを返すようにしてみた...
func NewPrinter(string_ string) Printer {
	p := &PrintBanner{banner.NewBanner(string_)}
	return p
}

func (p *PrintBanner) PrintWeak() {
	p.ShowWithParen()
}

func (p *PrintBanner) PrintStrong() {
	p.ShowWithAster()
}
