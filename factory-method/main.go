package main

import (
	"factory-method/idcard"
)

func main() {
	factory := idcard.NewIDCardCreator()
	card1 := factory.Create("ひろし")
	card2 := factory.Create("とむら")
	card3 := factory.Create("はなこ")
	card1.Use()
	card2.Use()
	card3.Use()
}
