package ui

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/sashaaro/gophkeeper/internal/client"
	"strings"
)

type WidgetMainMenu struct {
	primitive *tview.TextView
	list      *tview.TextView
	client    *client.Client
	status    *WidgetStatus
}

func NewWidgetMainMenu(client *client.Client, status *WidgetStatus) *WidgetMainMenu {
	o := &WidgetMainMenu{
		primitive: tview.NewTextView(),
		list:      tview.NewTextView(),
		client:    client,
		status:    status,
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

func (w *WidgetMainMenu) UpdateList() {
	l, err := w.client.GetAll()
	if err != nil {
		w.status.Log("Fail to get all list: " + err.Error())
	} else {
		text := "----------\n"
		for name, bytes := range l {
			var row string
			secretData := client.SecretData(bytes)
			if b, ok := secretData.ToBinary(); ok {
				row = "binary = " + string(b)
			} else {
				row, ok = secretData.ToText()
				if ok {
					row = "text - " + row
				} else {
					creditCard, ok := secretData.ToBankCredentials()
					if ok {
						row = fmt.Sprintf("credit card - name: %s, number: %s, code %s", creditCard.Name, creditCard.Number, creditCard.Code)
					}
				}
			}
			text += fmt.Sprintf("%s %v\n", name, row)
		}
		w.list.SetText(text)
	}
}
