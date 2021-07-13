package main

import "adapter1/print"

func main() {
	var printer print.Printer = print.NewPrinter("Hello")
	printer.PrintStrong()
	printer.PrintWeak()
}
