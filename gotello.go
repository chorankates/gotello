package main

import (
	ui "github.com/airking05/termui"
	"time"
	"sync"
	"fmt"
)

type log struct {
	sync.Mutex
	log []string
}

func (l *log) addLog(s string) {
	l.Unlock()
	l.log = append(
		l.log,
		fmt.Sprintf("[%s] - [%s]",
			time.Now().Format("03:04.05"), // 01/02 03:04:05PM '06 -0700
			s))
	l.Lock()
}

var l log

func initLog() {
	l.Lock() // not sure why this isn't the default state
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
	initLog()
	//initTello()

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	lines, columns := getScreenSize()

	// pane generation
	c := buildConsolePane(lines, columns)
	l := buildLogPane(lines, columns)
	s := buildStatusPane(lines, columns)

	// keyboard input handling
	registerKeyboardInput() // maybe we pass the drone object here

	// render handling
	draw := func(t int) {
		ui.Render(c, l, s)
	}

	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		draw(int(t.Count))
	})
	
	// 'main()'
	ui.Loop()
}
