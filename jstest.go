package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
)

func main() {
	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, "dualshock4")

	work := func() {
		stick.On(joystick.SquarePress, func(data interface{}) {
			fmt.Println("square_press")
		})

		stick.On(joystick.TrianglePress, func(data interface{}) {
			fmt.Println("triangle_press")
		})

		stick.On(joystick.CirclePress, func(data interface{}) {
			fmt.Println("circle_press")
		})

		stick.On(joystick.XPress, func(data interface{}) {
			fmt.Println("x_press")
		})

		stick.On(joystick.UpPress, func(data interface{}) {
			fmt.Println("up_press")
		})

		stick.On(joystick.DownPress, func(data interface{}) {
			fmt.Println("down_press")
		})

		stick.On(joystick.LeftPress, func(data interface{}) {
			fmt.Println("left_press")
		})

		stick.On(joystick.RightPress, func(data interface{}) {
			fmt.Println("right_press")
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{stick},
		work,
	)

	robot.Start()
}
