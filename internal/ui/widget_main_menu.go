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

	o.UpdateMenu(BuildGuestMenu())

	return o
}

func BuildGuestMenu() []string {
	return []string{
		"(r) Register",
		"(l) Login",
	}
}

func BuildUserMenu(user string) []string {
	return []string{
		"(s) Save data",
		"(e) Exit. Logged as " + user,
	}
}

func (w *WidgetMainMenu) UpdateMenu(menu []string) {
	baseMenu := []string{"(p) Ping"}
	baseMenu = append(baseMenu, menu...)
	menu = append(menu, "(q) Quit")

	w.primitive.SetText(strings.Join(menu, "\n"))
}
