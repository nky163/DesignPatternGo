package print

import "adapter1/banner"

// Target
type Printer interface {
	PrintWeak()
	PrintStrong()
}

// Adapter
type PrintBanner struct {
	// Adaptee
	*banner.Banner // 継承的な...
}

// Printerインターフェースを返すようにしてみた...
// 使う側はインターフェースに依存する
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
