package main

import (
	ui "github.com/airking05/termui"
)

func buildConsolePane(lines int, columns int) *ui.Par {
	c := ui.NewPar("gotello 0.1 - basic control for Tello")
	c.Height = int(float64(lines) * 0.2)
	c.Width = columns

	c.TextFgColor = ui.ColorGreen
	c.BorderLabel = "console"
	c.BorderFg = ui.ColorWhite
	//c.Handle("/timer/1s", func(e ui.Event) {
	//	cnt := e.Data.(ui.EvtTimer)
	//	if cnt.Count%2 == 0 {
	//		c.TextFgColor = ui.ColorRed
	//	} else {
	//		c.TextFgColor = ui.ColorWhite
	//	}
	//})
	c.Y = lines - c.Height // align to the bottom

	c.Text = getKeyboardMap()

	return c
}

func buildLogPane(lines int, columns int) *ui.List {
	// TODO this should be a real timestamp
	l := ui.NewList()
	l.Height = int(float64(lines) * 0.8)
	l.Width = columns / 2
	l.BorderLabel = "log"
	l.BorderFg = ui.ColorWhite


	l.Handle("/timer/1s", func(e ui.Event) {
		l.Items = getLogPaneContents(l.Height)
	})

	return l
}

// TODO what should this actually hold? key presses -- we'll use the log window for responses from the drone
func buildStatusPane(lines int, columns int) *ui.List {
	// TODO come up with some better default text (?)
	s := ui.NewList()
	s.Height = int(float64(lines) * 0.8)
	s.Width = columns / 2
	s.BorderLabel = "status"
	s.BorderFg = ui.ColorWhite
	s.X = s.Width // aligh to the center

	s.Handle("/timer/1s", func(e ui.Event) {
		s.Items = getStatusPaneContents()
	})

	return s
}
