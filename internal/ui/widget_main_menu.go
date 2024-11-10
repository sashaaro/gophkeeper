package ui

import (
	"github.com/rivo/tview"
	"strings"
)

type WidgetMainMenu struct {
	primitive *tview.TextView
}

func NewWidgetMainMenu() *WidgetMainMenu {
	v := tview.NewTextView()

	o := &WidgetMainMenu{
		primitive: v,
	}

	o.UpdateMenu("")

	return o
}

func (w *WidgetMainMenu) UpdateMenu(user string) {
	menu := []string{"(r) Register", "(p) Ping"}
	if user == "" {
		menu = append(menu, "(l) Login")
	} else {
		menu = append(menu, "(e) Exit. Logged as "+user)
	}
	menu = append(menu, "(q) Quit")

	w.primitive.SetText(strings.Join(menu, "\n"))
}
