package ui

import "github.com/rivo/tview"

type LoginForm struct {
	Login    string
	Password string
}

type OnLogin func(form LoginForm)

type WidgetLoginForm struct {
	form     *tview.Form
	model    LoginForm
	onLogin  OnLogin
	onCancel OnCancel
}

func NewWidgetLogin(onLogin OnLogin, onCancel OnCancel) *WidgetLoginForm {
	w := WidgetLoginForm{
		form:     tview.NewForm(),
		model:    LoginForm{},
		onLogin:  onLogin,
		onCancel: onCancel,
	}
	w.form.AddInputField("Login", "", 20, nil, func(text string) {
		w.model.Login = text
	})
	w.form.AddPasswordField("Password", "", 20, '*', func(text string) {
		w.model.Password = text
	})
	w.form.AddButton("Login", func() {
		if w.onLogin != nil {
			w.onLogin(w.model)
		}
	})
	w.form.AddButton("Cancel", func() {
		if w.onCancel != nil {
			w.onCancel()
		}
	})
	return &w
}

func (w *WidgetLoginForm) Reset() {
	l := w.form.GetFormItemByLabel("Login")
	if f, ok := l.(*tview.InputField); ok {
		f.SetText("")
	}
	p := w.form.GetFormItemByLabel("Password")
	if f, ok := p.(*tview.InputField); ok {
		f.SetText("")
	}
}
