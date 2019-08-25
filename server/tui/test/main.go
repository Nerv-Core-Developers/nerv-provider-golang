package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

var (
	app            *tview.Application
	mainLog, stats *tview.TextView
	grid           *tview.Grid
)

const corporate = `Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.

Bring to the table win-win survival strategies to ensure proactive domination. At the end of the day, going forward, a new normal that has evolved from generation X is on the runway heading towards a streamlined cloud solution. User generated content in real-time will have multiple touchpoints for offshoring.

Capitalize on low hanging fruit to identify a ballpark value added activity to beta test. Override the digital divide with additional clickthroughs from DevOps. Nanotechnology immersion along the information highway will close the loop on focusing solely on the bottom line.

[yellow]Press Enter, then Tab/Backtab for word selections Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.`

func newPrimitive(text string, app *tview.Application, borders bool) *tview.TextView {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	tv.Box.SetTitleAlign(tview.AlignLeft).SetBorder(borders)
	return tv
}

func SetupTUI() error {
	app = tview.NewApplication()
	mainLog = newPrimitive("LOGS\n", app, false)
	stats = newPrimitive("STATS\n", app, true)

	grid = tview.NewGrid().
		SetRows(0, 4).
		AddItem(mainLog, 0, 0, 1, 3, 0, 0, true)

	grid.AddItem(stats, 1, 0, 1, 3, 0, 0, false)

	return nil
}

func main() {
	SetupTUI()

	go func() {
		numSelections := 0
		for _, word := range strings.Split(corporate, " ") {
			if word == "the" {
				word = "[red]the[white]"
			}
			if word == "to" {
				word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
				numSelections++
			}
			fmt.Fprintf(mainLog, "%s ", word)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
