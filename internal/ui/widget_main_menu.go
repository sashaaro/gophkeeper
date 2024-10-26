package ui

import "github.com/rivo/tview"

type WidgetMainMenu struct {
	primitive *tview.TextView
}

func NewWidgetMainMenu() *WidgetMainMenu {
	v := tview.NewTextView()
	v.SetText(`(r) Register
(p) Ping
(l) Login
(q) Quit
`)

	return &WidgetMainMenu{
		primitive: v,
	}
}
