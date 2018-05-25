package main

import (
	ui "github.com/airking05/termui"
)

// TODO really don't like how manual this is
func getKeyboardMap() (result string) {
	result = `
key		command			key		command
w		turn_up			s		turn_down
a		turn_left		d		turn_right
<up>	fly_up			<down>	fly_down
<left>	fly_left		<right>	fly_right
c		start camera	C		stop camera
<enter> takeoff <t>		q		quit
`
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
