package main

import (
	ui "github.com/airking05/termui"
	"fmt"
	"math/rand"
)

// TODO if this just holds control information, shouldn't it be 'help' and toggle-able
func buildConsolePane(lines int, columns int) *ui.Par {
	l.addLog("building Console pane")
	c := ui.NewPar("gotello 0.1 - basic control for Tello")
	c.Height = int(float64(lines) * 0.2)
	c.Width = columns

	c.TextFgColor = ui.ColorWhite
	c.BorderLabel = "console"
	c.BorderFg = ui.ColorWhite

	c.Y = lines - c.Height // align to the bottom

	c.Text = getKeyboardMap()
	//c.Text = getKeyboardMapBetter()

	return c
}

func buildLogPane(lines int, columns int) *ui.List {
	l.addLog("building Log pane")
	
	// TODO this should be a real timestamp
	ll := ui.NewList()
	ll.Height = int(float64(lines) * 0.8) + 1
	ll.Width = columns / 2
	ll.BorderLabel = "log"
	ll.BorderFg = ui.ColorWhite

	ll.Handle("/timer/1s", func(e ui.Event) {
		ll.Items = getLogPaneContents(ll.Height)
	})

	return ll
}

func buildStatusPane(lines int, columns int) *ui.List {
	l.addLog("building Status pane")

	s := ui.NewList()
	//s.Height = int(float64(lines) * 0.8)
	s.Height = int(float64(lines) * 0.5)
	s.Width = columns / 2
	s.BorderLabel = "status"
	s.BorderFg = ui.ColorWhite
	s.X = s.Width // align to the center

	s.Handle("/timer/1s", func(e ui.Event) {
		s.Items = getStatusPaneContents()
	})

	return s
}

func buildAltitudeGaugePane(lines int, columns int) *ui.Gauge {
	l.addLog("building Altitude gauge")

	a := ui.NewGauge()
	a.Height = int(float64(lines) * 0.1)
	a.Width = columns / 2
	a.BorderLabel = "altitude"
	a.BorderFg = ui.ColorWhite
	a.X = a.Width
	a.Y = lines - a.Height * 7 + 3
	
	a.Handle("/timer/1s", func(e ui.Event) {
		a.Percent = getAltitudePercentage()
	})

	return a
}

func getAltitudePercentage() (result int) {
	result = rand.Intn(100)
	return
}

func buildBatteryGaugePane(lines int, columns int) *ui.Gauge {
	l.addLog("building Battery gauge")

	g := ui.NewGauge()
	g.Height = int(float64(lines) * 0.1)
	g.Width = columns / 2
	g.BorderLabel = "battery"
	g.BorderFg = ui.ColorWhite
	g.X = g.Width // align to center
	g.Y = lines - g.Height * 6 + 3

	g.Percent = 0

	g.Handle("/timer/1s", func(e ui.Event) {
		pct := getBatteryPercentage()
		g.Percent = pct

		if (pct > 66) && (pct <= 100) {
			g.BarColor = ui.ColorGreen
		} else if (pct > 33)  && pct <= 66 {
			g.BarColor = ui.ColorYellow
		} else {
			g.BarColor = ui.ColorRed
		}
	})

	return g
}

func buildGroundSpeedGaugePane(lines int, columns int) *ui.Gauge {
	l.addLog("building GroundSpeed gauge")

	gs := ui.NewGauge()
	gs.Height = int(float64(lines) * 0.1)
	gs.Width = columns / 2
	gs.BorderLabel = "ground speed"
	gs.BorderFg = ui.ColorWhite
	gs.X = gs.Width
	gs.Y = lines - gs.Height * 4 + 3

	gs.Percent = 0

	gs.Handle("/timer/1s", func(e ui.Event) {
		pct := getGroundSpeedPercentage()
		gs.Percent = pct

		if (pct > 66) && (pct <= 100) {
			gs.BarColor = ui.ColorGreen
		} else if (pct > 33)  && pct <= 66 {
			gs.BarColor = ui.ColorYellow
		} else {
			gs.BarColor = ui.ColorRed
		}

	})

	return gs
}

func buildFlightStatusPanel(lines int, columns int) *ui.Par {
	l.addLog("building FlightStatus panel")

	fs := ui.NewPar("flight status")
	fs.Height = int(float64(lines) * 0.1)
	fs.Width = columns / 2
	fs.BorderLabel = "flight status"
	fs.BorderFg = ui.ColorWhite
	fs.X = fs.Width
	fs.Y = lines - fs.Height * 8 + 3

	fs.Text = "disconnected"

	fs.Handle("/timer/1s", func(e ui.Event) {
		status := getFlightStatus()

		fs.Text = status

		switch status {
		case "disconnected":
			fs.TextFgColor = ui.ColorWhite
		case "connected":
			fs.TextFgColor = ui.ColorYellow
		case "flying":
			fs.TextFgColor = ui.ColorGreen
		}

	})

	return fs

}

func getFlightStatus() (result string) {
	statuses := []string{
		"disconnected",
		"connected",
		"flying",
	}

	return statuses[rand.Intn(len(statuses))]
}

func getGroundSpeedPercentage() (result int) {
	// TODO use a constant to define the max speed
	result = rand.Intn(100)
	return
}

func buildRotorSpeedGaugePane(lines int, columns int) *ui.Gauge {
	l.addLog("building RotorSpeed gauge")

	rs := ui.NewGauge()
	rs.Height = int(float64(lines) * 0.1)
	rs.Width = columns / 2
	rs.BorderLabel = "rotor speed"
	rs.BorderFg = ui.ColorWhite
	rs.X = rs.Width
	rs.Y = lines - rs.Height * 5 + 3

	rs.Percent = 0

	rs.Handle("/timer/1s", func(e ui.Event) {
		pct := getRotorSpeedPerentage()

		rs.Percent = pct

		if (pct > 66) && (pct <= 100) {
			rs.BarColor = ui.ColorGreen
		} else if (pct > 33)  && pct <= 66 {
			rs.BarColor = ui.ColorYellow
		} else {
			rs.BarColor = ui.ColorRed
		}
	})

	return rs
}

func getRotorSpeedPerentage() (result int) {
	// TODO use a constant to define the max speed
	result = rand.Intn(100)
	return
}

func getBatteryPercentage() (result int) {
	//  TODO come up with a way to share this data with `getStatusPaneContents()`
	// maybe do some sort of achine in the GetStatus function
	//status := getStatus()
	//
	//result = status.BatteryPercent

	result = rand.Intn(100)
	return
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

	status := getStatus()

	result = []string{
		fmt.Sprintf("status:       [%v]", status.Status),
		fmt.Sprintf("altitude:     [%v]", status.Altitude),
		fmt.Sprintf("heading:      [%v]", status.Heading),
		fmt.Sprintf("ground speed: [%v]", status.GroundSpeed),
		fmt.Sprintf("rotor speed:  [%v]", status.RotorSpeed),
		fmt.Sprintf("battery %%:    [%v]", status.BatteryPercent),
		fmt.Sprintf("battery left: [%v]", status.BatteryLeft),
	}

	return
}

