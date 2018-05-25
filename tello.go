package main

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot"
	"time"
	"fmt"
)

func sendCommand(command string) {

	l.addLog(command)

	//switch command {
	//case "look_up":
	//	// TODO ... going to need a handle on the drone itself.. given how we're registering, probably need it to be global
	//}
}

func initTello() {
	// TODO what params should we take here?

	drone := tello.NewDriver("8888")

	work := func() {
		drone.TakeOff()

		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	fmt.Println(robot.Name)

	// TODO we need a way to dynamically add to 'work' while in-flight
	//robot.Start()
}