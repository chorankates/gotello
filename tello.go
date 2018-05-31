package main

import (
	"gobot.io/x/gobot/platforms/dji/tello"
	//"time"
	//"fmt"
	"gobot.io/x/gobot"
)

type Drone struct {
	// TODO need to include the actual gobot drone object here too
	// and we also need to redirect the consoloe output it's doing through it's logger to ours

	Status string
	Signal float32
	Altitude float32
	Heading float32
	GroundSpeed float32
	RotorSpeed float32
	BatteryPercent int // though we probably will get it as a float..
	BatteryLeft string // this needs to be computed by us

}

func getStatus() (status Drone) {

	// TODO fake it until you make it
	status = *new(Drone)

	status.Status = "disconnected"
	status.Signal = 0
	status.Altitude = 0
	status.Heading = 0
	status.GroundSpeed = 0
	status.RotorSpeed = 0
	status.BatteryPercent = 0
	status.BatteryLeft = "0 minutes"

	return
}



func sendCommand(command string) {

	l.addLog(command)

	switch command {
	case "turn_left":

	case "turn_right":
	case "fly_up":
	case "fly_down":
	case "fly_forward":
		//d.Forward(20)
	case "fly_backward":
	case "fly_left":
	case "fly_right":
	case "takeoff_or_land":
	}
}

func initTello() *gobot.Robot {
	// TODO what params should we take here?

	drone := tello.NewDriver("8888")

	//work := func() {
	//	drone.TakeOff()
	//
	//	gobot.After(5*time.Second, func() {
	//		drone.Land()
	//	})
	//}


	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		//work, // this needs to be a queue we pop on/off to/from
	)

	//fmt.Println(robot.Name)

	// TODO we need a way to dynamically add to 'work' while in-flight
	//robot.Start()

	return robot
}