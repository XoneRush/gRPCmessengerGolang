package main

import (
	"Client/forms"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	cli := forms.Client{
		App:       tview.NewApplication(),
		IndexForm: tview.NewForm(),
	}

	if err := cli.App.SetRoot(tview.NewBox(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	cli.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 123 {
			cli.App.Stop()
		}
		return event
	})

}
