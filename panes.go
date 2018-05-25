package main

import (
	ui "github.com/airking05/termui"
	"sort"
	"fmt"
)

// TODO if this just holds control information, shouldn't it be 'help' and toggle-able
func buildConsolePane(lines int, columns int) *ui.Par {
	c := ui.NewPar("gotello 0.1 - basic control for Tello")
	c.Height = int(float64(lines) * 0.2)
	c.Width = columns

	c.TextFgColor = ui.ColorWhite
	c.BorderLabel = "console"
	c.BorderFg = ui.ColorWhite

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

func buildStatusPane(lines int, columns int) *ui.List {
	s := ui.NewList()
	s.Height = int(float64(lines) * 0.8)
	s.Width = columns / 2
	s.BorderLabel = "status"
	s.BorderFg = ui.ColorWhite
	s.X = s.Width // align to the center

	s.Handle("/timer/1s", func(e ui.Event) {
		s.Items = getStatusPaneContents()
	})

	return s
}

func getLogPaneContents(rows int) (result []string) {

	l.Unlock()

	total := len(l.log) - 1

	for i := 0; i < rows && i < total; i++ {
		result = append(result, l.log[total - i])
	}
	l.Lock()

	return
}

func getStatusPaneContents() (result []string) {
	m := map[string]string {
		"status": "disconnected", // (disconnected, connected/landed, flying)
		"ground speed": "0.0",
		"rotor speed": "0.0",
		"battery %": "0.00%", // want this to be a bar graph eventually
		"battery left": "0s",
	}

	// i do not like this. i would very much like to just sort on the keys during the appending
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		result = append(result, fmt.Sprintf("%s: [%s]", key, m[key]))
	}

	return
}

