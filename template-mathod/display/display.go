package display

import "fmt"

// method-set of template
type printer interface {
	open()
	print()
	close()
}

// interface
type Displaier interface {
	Display()
}

// template-class
type abstractDisplay struct {
	printer printer
}

// template-method
func (ad *abstractDisplay) Display() {
	if ad.printer == nil {
		fmt.Println("ad.printer is nil")
		return
	}
	ad.printer.open()
	for i := 0; i < 5; i++ {
		ad.printer.print()
	}
	ad.printer.close()
}

type charDisplay struct {
	abstractDisplay
	ch byte
}

func NewCharDisplaier(ch byte) Displaier {
	ad := abstractDisplay{}
	cd := &charDisplay{
		abstractDisplay: ad,
		ch:              ch,
	}
	cd.printer = cd
	return cd
}

func (cd *charDisplay) open() {
	fmt.Print("<<")
}

func (cd *charDisplay) print() {
	fmt.Printf("%c", cd.ch)
}

func (cd *charDisplay) close() {
	fmt.Printf(">>\n")
}

type stringDisplay struct {
	abstractDisplay
	str string
}

func NewStringDisplaier(str string) Displaier {
	ad := abstractDisplay{}
	sd := &stringDisplay{
		abstractDisplay: ad,
		str:             str,
	}
	sd.printer = sd
	return sd
}

func (sd *stringDisplay) open() {
	sd.printLine()
}

func (sd *stringDisplay) print() {
	fmt.Printf("%s%s%s\n", "|", sd.str, "|")
}

func (sd *stringDisplay) close() {
	sd.printLine()
}

func (sd *stringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(sd.str); i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")
}
