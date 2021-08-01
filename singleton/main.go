package main

import (
	"fmt"
	"singleton/singleton"
)

func main() {
	for i := 0; i < 100; i++ {
		go singleton.GetInstance()
	}
	fmt.Scanln()
}
