package main

import (
	"prototype/framework"
)

func main() {
	m := framework.NewManager()
	m.Register("strong message", framework.NewUnderlinePen('~'))
	m.Register("warning box", framework.NewMessageBox('*'))
	m.Register("slash box", framework.NewMessageBox('/'))

	sm := m.Create("strong message")
	sm.Use("Hello wold")
	wb := m.Create("warning box")
	wb.Use("Hello World")
	sb := m.Create("slash box")
	sb.Use("Hello Wold")
}
