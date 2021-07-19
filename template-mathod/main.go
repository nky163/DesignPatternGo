package main

import (
	"template-method/display"
)

func main() {
	cd := display.NewCharDisplaier('H')
	cd.Display()
	cd2 := display.NewStringDisplaier("Hello World")
	cd2.Display()
}
