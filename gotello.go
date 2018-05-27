package main

import (
	ui "github.com/airking05/termui"
	"time"
	"sync"
	"fmt"
	"gobot.io/x/gobot"
)

type log struct {
	sync.Mutex
	log []string
}

func (l *log) addLog(s string) {
	l.Unlock()
	l.log = append(
		l.log,
		fmt.Sprintf("%s: [%s]",
			time.Now().Format("03:04.05"), // 01/02 03:04:05PM '06 -0700
			s))
	l.Lock()
}

var l log
var d gobot.Robot

func initialize() {
	l.Lock() // not sure why this isn't the default state
	d = *initTello()
}

func getScreenSize() (lines int, columns int) {
	// try to get better height/width - TODO abstract this
	//env_lines := os.Getenv("LINES") // this might only be BSD..
	//env_lines_int, err := strconv.Atoi(env_lines)
	//env_cols := os.Getenv("COLUMNS")
	//env_cols_int, err := strconv.Atoi(env_cols)

	// TODO above, but for now..
	lines = 46
	columns = 89

	return
}

func main() {
	initialize()
	//initTello()

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	lines, columns := getScreenSize()
	l.addLog(fmt.Sprintf("found screen size to be %dx%d", lines, columns))

	// pane generation
	l.addLog("building panes..")
	c := buildConsolePane(lines, columns)
	ll := buildLogPane(lines, columns)
	s := buildStatusPane(lines, columns)
	fs := buildFlightStatusPanel(lines, columns)
	a := buildAltitudeGaugePane(lines, columns)
	g := buildBatteryGaugePane(lines, columns)
	gs := buildGroundSpeedGaugePane(lines, columns)
	rs := buildRotorSpeedGaugePane(lines, columns)

	// keyboard input handling
	l.addLog("registering keuboard handlers..")
	registerKeyboardInput() // maybe we pass the drone object here

	// render handling
	draw := func(t int) {
		ui.Render(c, ll, s, fs, a, g, gs, rs)
	}

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})
	
	// 'main()'
	l.addLog("starting the render loop..")
	ui.Loop()
}
