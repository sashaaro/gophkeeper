package ui

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
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
		"(t) Add secret text",
		"(c) Add secret credit card text",
		"(e) Exit. Logged as " + user,
	}
}

func (w *WidgetMainMenu) UpdateMenu(menu []string) {
	menu = append([]string{"(p) Ping"}, menu...)
	menu = append(menu, "(q) Quit")

	w.primitive.SetText(strings.Join(menu, "\n"))
}

func (w *WidgetMainMenu) UpdateList() {
	l, err := w.client.GetAll()
	if err != nil {
		w.status.Log("Fail to get all list: " + err.Error())
	} else {
		w.SetList(l)
	}
}

func (w *WidgetMainMenu) SetList(l map[string][]byte) {
	t := table.NewWriter()
	t.SetAutoIndex(true)

	t.AppendHeader(table.Row{"name", "type", "value"})

	for name, bytes := range l {
		var row string
		secretData := client.SecretData(bytes)
		if b, ok := secretData.ToBinary(); ok {
			t.AppendRow(table.Row{name, "binary", fmt.Sprintf("%v", b)})
		} else {
			row, ok = secretData.ToText()
			if ok {
				t.AppendRow(table.Row{name, "text", row})
			} else {
				creditCard, ok := secretData.ToBankCredentials()
				if ok {
					t.AppendRow(table.Row{
						name, "credit card",
						"credit card",
						fmt.Sprintf("name: %s, number: %s, code %s",
							creditCard.Name,
							creditCard.Number,
							creditCard.Code,
						),
					})
				}
			}
		}
	}

	w.list.SetText(t.Render())
}
