package ui

import "github.com/rivo/tview"

type WidgetMainMenu struct {
	primitive *tview.TextView
}

func NewWidgetMainMenu() *WidgetMainMenu {
	v := tview.NewTextView()
	v.SetText(`(r) Register`)

	return &WidgetMainMenu{
		primitive: v,
	}
}
