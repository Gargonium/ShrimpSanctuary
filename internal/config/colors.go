package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ButtonColorFromStatus = map[string]rl.Color{
	WaitingBtnStatus: rl.Black,
	HoveredBtnStatus: rl.Gray,
	ClickedBtnStatus: rl.Green,
}

var VolumeSliderColor = rl.NewColor(44, 55, 76, 255)

// Food
var FoodColor = rl.Brown
var FoodBorderColor = rl.DarkBrown

// Pollute
var PolluteColor = rl.NewColor(0, 0, 0, 200)
