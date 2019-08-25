package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	App            *tview.Application
	MainLog, Stats *tview.TextView
	grid           *tview.Grid
)

func newPrimitive(text string, app *tview.Application, borders bool) *tview.TextView {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	tv.Box.SetBorder(borders)
	return tv
}

func SetupTUI(interactEvtHandler func(event *tcell.EventKey) *tcell.EventKey) error {
	App = tview.NewApplication()
	MainLog = newPrimitive("LOGS\n", App, false)
	Stats = newPrimitive("STATS\n", App, true)
	App.SetInputCapture(interactEvtHandler)
	grid = tview.NewGrid().
		SetRows(0, 4).
		AddItem(MainLog, 0, 0, 1, 3, 0, 0, true)

	grid.AddItem(Stats, 1, 0, 1, 3, 0, 0, false)

	return nil
}

func StartTUI() {
	if err := App.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
