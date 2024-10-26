package ui

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sashaaro/gophkeeper/internal/client"
)

const (
	PageLoginForm        = "login_form"
	PageRegisterForm     = "register_form"
	PageWidgetMainMenu   = "main_menu"
	PageWidgetSecretList = "secret_list"
)

type UIApp struct {
	app    *tview.Application
	client *client.Client
	pages  *tview.Pages
}

type Registerer interface {
}

func NewUIApp(client *client.Client) *UIApp {
	app := tview.NewApplication()
	pages := tview.NewPages()
	return &UIApp{
		app:    app.SetRoot(pages, true).EnableMouse(true),
		client: client,
		pages:  pages,
	}
}

func (a *UIApp) Init() {
	widgetStatus := NewWidgetStatus()

	widgetRegister := NewWidgetRegister(func(f RegisterForm) {
		// On register
		if err := a.client.Register(context.Background(), f.Login, f.Password); err != nil {
			widgetStatus.Log("Fail to register: " + err.Error())
		} else {
			widgetStatus.Log("User registered")
		}
		a.pages.SwitchToPage(PageWidgetMainMenu)
	}, func() {
		widgetStatus.Log("Cancel registration")
		a.pages.SwitchToPage(PageWidgetMainMenu)
	})
	a.pages.AddPage(
		PageRegisterForm,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(widgetRegister.form, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)

	widgetLogin := NewWidgetLogin(func(f LoginForm) {
		// On register
		if err := a.client.Login(context.Background(), f.Login, f.Password); err != nil {
			widgetStatus.Log("Fail to login: " + err.Error())
		} else {
			widgetStatus.Log("User logged in")
		}
		a.pages.SwitchToPage(PageWidgetMainMenu)
	}, func() {
		widgetStatus.Log("Cancel login")
		a.pages.SwitchToPage(PageWidgetMainMenu)
	})
	a.pages.AddPage(
		PageLoginForm,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(widgetLogin.form, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)

	widgetMainMenu := NewWidgetMainMenu()
	a.pages.AddPage(
		PageWidgetMainMenu,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(widgetMainMenu.primitive, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)
	widgetMainMenu.primitive.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			a.app.Stop()
		case 'r':
			widgetRegister.Reset()
			a.pages.SwitchToPage(PageRegisterForm)
		case 'l':
			widgetLogin.Reset()
			a.pages.SwitchToPage(PageLoginForm)
		case 'p':
			if err := a.client.Ping(context.Background()); err != nil {
				widgetStatus.Log("ping fails: " + err.Error())
			} else {
				widgetStatus.Log("pong")
			}
		}
		return event
	})
}

func (a *UIApp) Run() error {
	return a.app.Run()
}
