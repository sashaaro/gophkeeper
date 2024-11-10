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
	PageSaveSecret       = "save_secret"
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

	widgetMainMenu := NewWidgetMainMenu(a.client, widgetStatus)

	a.pages.AddPage(
		PageWidgetMainMenu,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(widgetMainMenu.primitive, 10, 1, true).
			AddItem(widgetMainMenu.list, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)

	widgetRegister := NewWidgetRegister(func(f RegisterForm) {
		// On register
		if err := a.client.Register(context.Background(), f.Login, f.Password); err != nil {
			widgetStatus.Log("Fail to register: " + err.Error())
		} else {
			widgetMainMenu.UpdateMenu(BuildUserMenu(a.client.LoginName))
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
			widgetMainMenu.UpdateMenu(BuildUserMenu(a.client.LoginName))
			widgetStatus.Log("User logged in")
			widgetMainMenu.UpdateList()
		}

		a.pages.SwitchToPage(PageWidgetMainMenu)
		a.app.ForceDraw()
	}, func() {
		a.pages.SwitchToPage(PageWidgetMainMenu)
		a.app.ForceDraw()
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

	secret := struct {
		Name  string
		Value string
	}{}
	saveSecretForm := tview.NewForm()
	saveSecretForm.
		AddInputField("secret name", "", 20, nil, func(text string) {
			secret.Name = text
		}).
		AddInputField("secret value", "", 20, nil, func(text string) {
			secret.Value = text
		}).
		AddButton("Save secret data", func() {
			err := a.client.SendSecretText(secret.Name, secret.Value)
			if err != nil {
				widgetStatus.Log("error saving secret: " + err.Error())
			} else {
				widgetMainMenu.UpdateList()
				a.pages.SwitchToPage(PageWidgetMainMenu)
				widgetStatus.Log("secret saved")
			}
			a.app.ForceDraw()
		})

	a.pages.AddPage(
		PageSaveSecret,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(saveSecretForm, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)

	widgetMainMenu.primitive.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q', 'Q':
			a.app.Stop()
		case 's', 'S':
			if a.client.LoginName != "" {
				a.pages.SwitchToPage(PageSaveSecret)
				a.app.ForceDraw()
			}
		case 'r', 'R':
			if a.client.LoginName == "" {
				widgetRegister.Reset()
				a.pages.SwitchToPage(PageRegisterForm)
				a.app.ForceDraw()
			}
		case 'l', 'L':
			if a.client.LoginName == "" {
				widgetLogin.Reset()
				a.pages.SwitchToPage(PageLoginForm)
			}
		case 'e', 'E':
			if a.client.LoginName != "" {
				widgetMainMenu.UpdateMenu(BuildGuestMenu())
				widgetStatus.Log("Log out")
				a.app.ForceDraw()
			}
		case 'p', 'P':
			if err := a.client.Ping(context.Background()); err != nil {
				widgetStatus.Log("ping fails: " + err.Error())
			} else {
				widgetStatus.Log("pong")
			}
			a.app.ForceDraw()
		}
		return event
	})

	a.pages.SwitchToPage(PageWidgetMainMenu)
}

func (a *UIApp) Run() error {
	return a.app.Run()
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
