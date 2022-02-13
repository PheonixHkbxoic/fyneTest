package component

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TagButton struct {
	widget.Button
	panel     *fyne.Container
	ShowPanel func()
	cb func()
}

func NewTagButton(icon fyne.Resource, text string, panel *fyne.Container) *TagButton {
	tb := &TagButton{}
	tb.Icon = icon
	tb.Text = text
	tb.panel = panel
	tb.ExtendBaseWidget(tb)
	tb.OnTapped = tb.TogglePanel
	return tb
}

func (t *TagButton) TogglePanel() {
	fmt.Println("TogglePanel")
	if t.panel.Visible() {
		t.panel.Hide()
		t.panel.Refresh()
	} else {
		t.panel.Show()
		t.panel.Refresh()
	}
	if t.cb != nil {
		t.cb()
	}
}

func (t *TagButton) SetToggleCallback(cb func()) {
	t.cb = cb
}