//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/platforms/joystick"
)

func main() {
	joystickAdaptor := joystick.NewAdaptor()
	joystick := joystick.NewDriver(joystickAdaptor, joystick.LogitechG27)

	work := func() {
		//Pedals
		joystick.On(joystick.Event("gas"), func(data interface{}) {
			fmt.Println("gas", data)
		})

		joystick.On(joystick.Event("brake"), func(data interface{}) {
			fmt.Println("brake", data)
		})

		joystick.On(joystick.Event("clutch"), func(data interface{}) {
			fmt.Println("clutch", data)
		})

		//Wheel Body

		//Wheel
		joystick.On(joystick.Event("wheel"), func(data interface{}) {
			fmt.Println("wheel", data)
		})

		// Left Paddle
		joystick.On(joystick.Event("left_paddle_press"), func(data interface{}) {
			fmt.Println("left_paddle_press")
		})
		joystick.On(joystick.Event("left_paddle_release"), func(data interface{}) {
			fmt.Println("left_paddle_release")
		})

		// Right Paddle
		joystick.On(joystick.Event("right_paddle_press"), func(data interface{}) {
			fmt.Println("right_paddle_press")
		})
		joystick.On(joystick.Event("right_paddle_release"), func(data interface{}) {
			fmt.Println("right_paddle_release")
		})

		//Top left red wheel button
		joystick.On(joystick.Event("wheel_left_top_press"), func(data interface{}) {
			fmt.Println("wheel_left_top")
		})
		joystick.On(joystick.Event("wheel_left_top_release"), func(data interface{}) {
			fmt.Println("wheel_left_top_release")
		})

		//Middle left red wheel button
		joystick.On(joystick.Event("wheel_left_middle_press"), func(data interface{}) {
			fmt.Println("wheel_left_middle_press")
		})
		joystick.On(joystick.Event("wheel_left_middle_release"), func(data interface{}) {
			fmt.Println("wheel_left_middle_release")
		})

		//Bottom left red wheel button
		joystick.On(joystick.Event("wheel_left_bottom_press"), func(data interface{}) {
			fmt.Println("wheel_left_bottom_press")
		})
		joystick.On(joystick.Event("wheel_left_bottom_release"), func(data interface{}) {
			fmt.Println("wheel_left_bottom_release")
		})

		//Top right red wheel button
		joystick.On(joystick.Event("wheel_left_top_press"), func(data interface{}) {
			fmt.Println("wheel_left_top")
		})
		joystick.On(joystick.Event("wheel_right_top_release"), func(data interface{}) {
			fmt.Println("wheel_right_top_release")
		})

		//Middle right red wheel button
		joystick.On(joystick.Event("wheel_right_middle_press"), func(data interface{}) {
			fmt.Println("wheel_right_middle_press")
		})
		joystick.On(joystick.Event("wheel_right_middle_release"), func(data interface{}) {
			fmt.Println("wheel_right_middle_release")
		})

		//Bottom right red wheel button
		joystick.On(joystick.Event("wheel_right_bottom_press"), func(data interface{}) {
			fmt.Println("wheel_right_bottom_press")
		})
		joystick.On(joystick.Event("wheel_right_bottom_release"), func(data interface{}) {
			fmt.Println("wheel_right_bottom_release")
		})

		//GearShift

		//First gear
		joystick.On(joystick.Event("first_gear_press"), func(data interface{}) {
			fmt.Println("first_gear_press")
		})
		joystick.On(joystick.Event("first_gear_release"), func(data interface{}) {
			fmt.Println("first_gear_release")
		})

		//Second gear
		joystick.On(joystick.Event("second_gear_press"), func(data interface{}) {
			fmt.Println("second_gear_press")
		})
		joystick.On(joystick.Event("second_gear_release"), func(data interface{}) {
			fmt.Println("second_gear_release")
		})

		//Third Gear
		joystick.On(joystick.Event("third_gear_press"), func(data interface{}) {
			fmt.Println("third_gear_press")
		})
		joystick.On(joystick.Event("third_gear_release"), func(data interface{}) {
			fmt.Println("third_gear_release")
		})

		//Fourth Gear
		joystick.On(joystick.Event("fourth_gear_press"), func(data interface{}) {
			fmt.Println("fourth_gear_press")
		})
		joystick.On(joystick.Event("fourth_gear_release"), func(data interface{}) {
			fmt.Println("fourth_gear_release")
		})

		//Fifth Gear
		joystick.On(joystick.Event("fifth_gear_press"), func(data interface{}) {
			fmt.Println("fifth_gear_press")
		})
		joystick.On(joystick.Event("fifth_gear_release"), func(data interface{}) {
			fmt.Println("fifth_gear_release")
		})

		//Sixth Gear
		joystick.On(joystick.Event("sixth_gear_press"), func(data interface{}) {
			fmt.Println("sixth_gear_press")
		})
		joystick.On(joystick.Event("sixth_gear_release"), func(data interface{}) {
			fmt.Println("sixth_gear_release")
		})

		//Reverse Gear
		joystick.On(joystick.Event("reverse_gear_press"), func(data interface{}) {
			fmt.Println("reverse_gear_press")
		})
		joystick.On(joystick.Event("reverse_gear_release"), func(data interface{}) {
			fmt.Println("reverse_gear_release")
		})

		//Y
		joystick.On(joystick.Event("Y_press"), func(data interface{}) {
			fmt.Println("Y_press")
		})
		joystick.On(joystick.Event("Y_release"), func(data interface{}) {
			fmt.Println("Y_release")
		})

		//B
		joystick.On(joystick.Event("B_press"), func(data interface{}) {
			fmt.Println("B_press")
		})
		joystick.On(joystick.Event("B_release"), func(data interface{}) {
			fmt.Println("B_release")
		})

		//A
		joystick.On(joystick.Event("A_press"), func(data interface{}) {
			fmt.Println("A_press")
		})
		joystick.On(joystick.Event("A_release"), func(data interface{}) {
			fmt.Println("A_release")
		})

		//X
		joystick.On(joystick.Event("X_press"), func(data interface{}) {
			fmt.Println("X_press")
		})
		joystick.On(joystick.Event("X_release"), func(data interface{}) {
			fmt.Println("X_release")
		})

		//Red Buttons
		//Red 1
		joystick.On(joystick.Event("red_1_press"), func(data interface{}) {
			fmt.Println("red_1_press")
		})
		joystick.On(joystick.Event("red_1_release"), func(data interface{}) {
			fmt.Println("red_1_release")
		})

		//Red 2
		joystick.On(joystick.Event("red_2_press"), func(data interface{}) {
			fmt.Println("red_2_press")
		})
		joystick.On(joystick.Event("red_2_release"), func(data interface{}) {
			fmt.Println("red_2_release")
		})

		//Red 3
		joystick.On(joystick.Event("red_3_press"), func(data interface{}) {
			fmt.Println("red_3_press")
		})
		joystick.On(joystick.Event("red_3_release"), func(data interface{}) {
			fmt.Println("red_3_release")
		})

		//Red 4
		joystick.On(joystick.Event("red_4_press"), func(data interface{}) {
			fmt.Println("red_4_press")
		})
		joystick.On(joystick.Event("red_4_release"), func(data interface{}) {
			fmt.Println("red_4_release")
		})

		//DPAD

		//UP
		joystick.On(joystick.Event("up_press"), func(data interface{}) {
			fmt.Println("up_press")
		})
		joystick.On(joystick.Event("up_release"), func(data interface{}) {
			fmt.Println("up_release")
		})

		//Down
		joystick.On(joystick.Event("down_press"), func(data interface{}) {
			fmt.Println("down_press")
		})
		joystick.On(joystick.Event("down_release"), func(data interface{}) {
			fmt.Println("down_release")
		})

		//Left
		joystick.On(joystick.Event("left_press"), func(data interface{}) {
			fmt.Println("left_press")
		})
		joystick.On(joystick.Event("left_release"), func(data interface{}) {
			fmt.Println("left_release")
		})

		//Right
		joystick.On(joystick.Event("right_press"), func(data interface{}) {
			fmt.Println("right_press")
		})
		joystick.On(joystick.Event("right_release"), func(data interface{}) {
			fmt.Println("right_release")
		})

		//Diagnals
		//Up_Right
		joystick.On(joystick.Event("up_right_press"), func(data interface{}) {
			fmt.Println("up_right_press")
		})
		joystick.On(joystick.Event("up_right_release"), func(data interface{}) {
			fmt.Println("up_right_release")
		})

		//Down_Right
		joystick.On(joystick.Event("down_right_press"), func(data interface{}) {
			fmt.Println("down_right_press")
		})
		joystick.On(joystick.Event("down_right_release"), func(data interface{}) {
			fmt.Println("down_right_release")
		})

		//Down_Left
		joystick.On(joystick.Event("down_left_press"), func(data interface{}) {
			fmt.Println("down_left_press")
		})
		joystick.On(joystick.Event("down_left_release"), func(data interface{}) {
			fmt.Println("down_left_release")
		})

		//Up_Left
		joystick.On(joystick.Event("up_left_press"), func(data interface{}) {
			fmt.Println("up_left_press")
		})
		joystick.On(joystick.Event("up_left_release"), func(data interface{}) {
			fmt.Println("up_left_release")
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{joystick},
		work,
	)

	robot.Start()
}
