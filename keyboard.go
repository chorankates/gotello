package main

import (
	ui "github.com/airking05/termui"
	"text/tabwriter"
	"io"
	"fmt"
)

// TODO really don't like how manual this is
func getKeyboardMap() (result string) {
	result = `key		command			key		command
w		turn_up			s		turn_down
a		turn_left		d		turn_right
<up>	fly_up			<down>	fly_down
<left>	fly_left		<right>	fly_right
c		start camera	C		stop camera
+		zoom_in			-		zoom_out
<enter> takeoff <t>		q		quit
`
	return
}

func getKeyboardMapBetter() (result string) {
	buffer := new(io.Writer)
	w := tabwriter.NewWriter(*buffer, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug) // boilerplate

	fmt.Fprintf(w, "key\tcommand\tkey\tcommand")
	//fmt.Fprintf(w, "w\tturn_up\ts\tturn_down") // not actually supported on the tello
	fmt.Fprintf(w, "a\tturn_left\td\tturn_right")
	fmt.Fprintf(w, "<up>\tfly_up\t<down>\tfly_down")
	fmt.Fprintf(w, "<left>\tfly_left\t<right>\tfly_right")
	fmt.Fprintf(w, "c\tstart_camera\tC\tstop_camera")
	fmt.Fprintf(w, "+\tzoom_in\t-\tzoom_out")
	fmt.Fprintf(w, "<enter>\ttakeoff<t>\tq\tquit")
	//w.Flush()

	// TODO how do we get the contents of w into result

	return
}

func registerKeyboardInput() {

	// flight controls
	ui.Handle("/sys/kbd/a", func(ui.Event) {
		sendCommand("turn_left")
	})

	ui.Handle("/sys/kbd/d", func(ui.Event) {
		sendCommand("turn_right")
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		sendCommand("fly_up")
	})

	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		sendCommand("fly_down")
	})

	ui.Handle("/sys/kbd/<left>", func(ui.Event) {
		sendCommand("fly_left")
	})

	ui.Handle("/sys/kbd/<right>", func(ui.Event) {
		sendCommand("fly_right")
	})

	ui.Handle("/sys/kbd/+", func(ui.Event) {
		sendCommand("zoom_in")
	})

	ui.Handle("/sys/kbd/-", func(ui.Event) {
		sendCommand("zoom_out")
	})

	// TODO or should this be a toggle as well?
	ui.Handle("/sys/kbd/c", func(ui.Event) {
		sendCommand("start_camera")
	})

	ui.Handle("/sys/kbd/C", func(ui.Event) {
		sendCommand("stop_camera")
	})

	// TODO not sure that this is the right key or whether or not this should be a toggle
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		sendCommand("takeoff_or_land")
	})

	// navigation / otherwise
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

}
