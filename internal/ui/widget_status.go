package ui

import (
	"strings"

	"github.com/rivo/tview"
)

const StatusBufferSize = 5

type WidgetStatus struct {
	primitive *tview.TextView
	buf       [StatusBufferSize]string
}

func NewWidgetStatus() *WidgetStatus {
	return &WidgetStatus{
		primitive: tview.NewTextView().SetText("Hi"),
	}
}

func (w *WidgetStatus) Log(s string) {
	for i := 0; i < StatusBufferSize-1; i++ {
		w.buf[i] = w.buf[i+1]
	}
	w.buf[StatusBufferSize-1] = s
	w.primitive.SetText(strings.Join(w.buf[:], "\n"))
}
