package main

import "adapter1/print"

func main() {
	var printer print.Printer = print.NewPrint("Hello")
	printer.PrintStrong()
	printer.PrintWeak()
}
