package ui

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sashaaro/gophkeeper/internal/client"
	"github.com/sashaaro/gophkeeper/internal/entity"
)

const (
	PageLoginForm      = "login_form"
	PageRegisterForm   = "register_form"
	PageTextForm       = "text_form"
	PageCardForm       = "card_form"
	PageWidgetMainMenu = "main_menu"
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

	textSecretForm := createTextForm(a, widgetStatus, widgetMainMenu)

	a.pages.AddPage(
		PageTextForm,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(textSecretForm, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)

	creditCardForm := createCardForm(a, widgetStatus, widgetMainMenu)

	a.pages.AddPage(
		PageCardForm,
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(creditCardForm, 10, 1, true).
			AddItem(widgetStatus.primitive, 10, 1, true),
		true,
		true,
	)

	widgetMainMenu.primitive.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q', 'Q':
			a.app.Stop()
		case 't', 'T':
			if a.client.LoginName != "" {
				resetTextFrom(textSecretForm)
				a.pages.SwitchToPage(PageTextForm)
				a.app.ForceDraw()
			}
		case 'c', 'C':
			if a.client.LoginName != "" {
				resetTextFrom(creditCardForm)
				a.pages.SwitchToPage(PageCardForm)
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
				a.client.Logout()
				widgetMainMenu.UpdateMenu(BuildGuestMenu())
				widgetStatus.Log("Log out")
				widgetMainMenu.SetList(nil)
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

func createTextForm(a *UIApp, widgetStatus *WidgetStatus, widgetMainMenu *WidgetMainMenu) *tview.Form {
	textSecret := struct {
		Name  string
		Value string
	}{}
	textSecretForm := tview.NewForm()
	textSecretForm.
		AddInputField("secret name", "", 20, nil, func(text string) {
			textSecret.Name = text
		}).
		AddInputField("secret text", "", 20, nil, func(text string) {
			textSecret.Value = text
		}).
		AddButton("Save text secret", func() {
			err := a.client.SendSecretText(textSecret.Name, textSecret.Value)
			if err != nil {
				widgetStatus.Log("error saving secret: " + err.Error())
			} else {
				widgetMainMenu.UpdateList()
				a.pages.SwitchToPage(PageWidgetMainMenu)
				widgetStatus.Log("secret saved")
			}
			a.app.ForceDraw()
		})
	return textSecretForm
}

func createCardForm(a *UIApp, widgetStatus *WidgetStatus, widgetMainMenu *WidgetMainMenu) *tview.Form {
	creditCard := entity.CreditCard{
		Number: "",
		Date:   "",
		Name:   "",
		Code:   "",
	}
	creditCardForm := tview.NewForm()

	var secretName string
	creditCardForm.
		AddInputField("secret name", "", 20, nil, func(text string) {
			secretName = text
		}).
		AddInputField("credit cart number", "", 20, nil, func(text string) {
			creditCard.Number = text
		}).
		AddInputField("date", "", 20, nil, func(text string) {
			creditCard.Date = text
		}).
		AddInputField("name", "", 20, nil, func(text string) {
			creditCard.Name = text
		}).
		AddInputField("code", "", 20, nil, func(text string) {
			creditCard.Code = text
		}).
		AddButton("Save credit card secret", func() {
			err := a.client.SendSecretCreditCard(secretName, creditCard)
			if err != nil {
				widgetStatus.Log("error saving secret: " + err.Error())
			} else {
				widgetMainMenu.UpdateList()
				a.pages.SwitchToPage(PageWidgetMainMenu)
				widgetStatus.Log("secret saved")
			}
			a.app.ForceDraw()
		})
	return creditCardForm
}

func (a *UIApp) Run() error {
	return a.app.Run()
}

func resetTextFrom(f *tview.Form) {
	for i := 0; i < f.GetFormItemCount(); i++ {
		f.GetFormItem(i).(*tview.InputField).SetText("")
	}
}
