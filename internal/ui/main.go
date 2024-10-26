package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	PageLoginForm        = "login_form"
	PageRegisterForm     = "register_form"
	PageWidgetMainMenu   = "main_menu"
	PageWidgetSecretList = "secret_list"
)

type UIApp struct {
	app   *tview.Application
	pages *tview.Pages
}

func NewUIApp() *UIApp {
	app := tview.NewApplication()
	pages := tview.NewPages()
	return &UIApp{
		app:   app.SetRoot(pages, true).EnableMouse(true),
		pages: pages,
	}
}

func (a *UIApp) Run() error {
	return a.app.Run()
}

func (a *UIApp) Init() {
	widgetRegister := NewWidgetRegister(func() {
		// On register
		a.pages.SwitchToPage(PageWidgetMainMenu)
	}, func() {
		a.pages.SwitchToPage(PageWidgetMainMenu)
	})
	a.pages.AddPage(PageRegisterForm, widgetRegister.form, true, true)

	widgetMainMenu := NewWidgetMainMenu()
	a.pages.AddPage(PageWidgetMainMenu, widgetMainMenu.primitive, true, true)
	widgetMainMenu.primitive.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			a.app.Stop()
		}
		if event.Rune() == 'r' {
			widgetRegister.Reset()
			a.pages.SwitchToPage(PageRegisterForm)
		}
		return event
	})
}
