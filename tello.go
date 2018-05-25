package main

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot"
	"time"
	"fmt"
)

func initTello() {
	// TODO what params should we take here?
	//


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

	robot.Start()
}