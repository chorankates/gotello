package main

import (
	ui "github.com/airking05/termui"
	"time"
	"sync"
	"fmt"
	"sort"
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

func getLogPaneContents(rows int) (result []string) {

	l.Unlock()
	//result = l.log[:rows] // TODO why doesn't this work?

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

func sendCommand(command string) {

	l.addLog(command)

	//switch command {
	//case "look_up":
	//	// TODO ... going to need a handle on the drone itself.. given how we're registering, probably need it to be global
	//}
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
