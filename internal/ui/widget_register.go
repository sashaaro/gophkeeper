package ui

import "github.com/rivo/tview"

type RegisterForm struct {
	Login    string
	Password string
}

type OnRegister func(form RegisterForm)
type OnCancel func()

type WidgetRegisterForm struct {
	form       *tview.Form
	model      RegisterForm
	onRegister OnRegister
	onCancel   OnCancel
}

func NewWidgetRegister(onRegister OnRegister, onCancel OnCancel) *WidgetRegisterForm {
	w := WidgetRegisterForm{
		form:       tview.NewForm(),
		model:      RegisterForm{},
		onRegister: onRegister,
		onCancel:   onCancel,
	}
	w.form.AddInputField("Login", "", 20, nil, func(text string) {
		w.model.Login = text
	})
	w.form.AddPasswordField("Password", "", 20, '*', func(text string) {
		w.model.Password = text
	})
	w.form.AddButton("Register", func() {
		if w.onRegister != nil {
			w.onRegister(w.model)
		}
	})
	w.form.AddButton("Cancel", func() {
		if w.onCancel != nil {
			w.onCancel()
		}
	})
	return &w
}

func (w *WidgetRegisterForm) Reset() {
	l := w.form.GetFormItemByLabel("Login")
	if f, ok := l.(*tview.InputField); ok {
		f.SetText("")
	}
	p := w.form.GetFormItemByLabel("Password")
	if f, ok := p.(*tview.InputField); ok {
		f.SetText("")
	}
}
