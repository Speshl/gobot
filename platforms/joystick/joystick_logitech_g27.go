package joystick

//Mapping done on Windows 11 with logitech profiler installed

var logitechG27config = joystickConfig{
	Name: "Logitech G27",
	GUID: "", //TODO: Where do I get this?
	Axis: []pair{
		pair{
			Name: "wheel",
			ID:   0,
		},
		pair{
			Name: "gas",
			ID:   1,
		},
		pair{
			Name: "brake",
			ID:   2,
		},
		pair{
			Name: "clutch",
			ID:   4,
		},
	},
	Buttons: []pair{
		//Wheel body
		pair{
			Name: "left_paddle",
			ID:   5,
		},
		pair{
			Name: "right_paddle",
			ID:   4,
		},
		pair{
			Name: "wheel_left_top",
			ID:   7,
		},
		pair{
			Name: "wheel_left_middle",
			ID:   20,
		},
		pair{
			Name: "wheel_left_bottom",
			ID:   22,
		},
		pair{
			Name: "wheel_right_top",
			ID:   6,
		},
		pair{
			Name: "wheel_right_middle",
			ID:   19,
		},
		pair{
			Name: "wheel_right_bottom",
			ID:   21,
		},

		//Gear Shift
		pair{
			Name: "first_gear",
			ID:   8,
		},
		pair{
			Name: "second_gear",
			ID:   9,
		},
		pair{
			Name: "third_gear",
			ID:   10,
		},
		pair{
			Name: "fourth_gear",
			ID:   11,
		},
		pair{
			Name: "fifth_gear",
			ID:   12,
		},
		pair{
			Name: "sixth_gear",
			ID:   13,
		},
		pair{
			Name: "reverse_gear",
			ID:   14,
		},

		//Black Diamond
		pair{
			Name: "Y",
			ID:   15,
		},
		pair{
			Name: "B",
			ID:   18,
		},
		pair{
			Name: "A",
			ID:   17,
		},
		pair{
			Name: "X",
			ID:   16,
		},

		//Red Buttons
		pair{
			Name: "red_1",
			ID:   0,
		},
		pair{
			Name: "red_2",
			ID:   1,
		},
		pair{
			Name: "red_3",
			ID:   2,
		},
		pair{
			Name: "red_4",
			ID:   3,
		},
	},
	Hats: []hat{
		hat{
			Hat:  0,
			Name: "up",
			ID:   1,
		},
		hat{
			Hat:  0,
			Name: "down",
			ID:   4,
		},
		hat{
			Hat:  0,
			Name: "left",
			ID:   8,
		},
		hat{
			Hat:  0,
			Name: "right",
			ID:   2,
		},
		//diagnals
		hat{
			Hat:  0,
			Name: "up_right",
			ID:   3,
		},
		hat{
			Hat:  0,
			Name: "down_right",
			ID:   6,
		},
		hat{
			Hat:  0,
			Name: "down_left",
			ID:   12,
		},
		hat{
			Hat:  0,
			Name: "up_left",
			ID:   9,
		},
	},
}
