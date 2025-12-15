package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// BackGround
var (
	BackgroundColor = rl.NewColor(20, 25, 35, 255)
	//TableColor          = rl.NewColor(101, 79, 48, 255)
	//AquariumBorderColor = rl.Black
	//SandColor           = rl.NewColor(200, 144, 65, 255)
	//WaterColor          = rl.NewColor(62, 85, 130, 255)
)

var ButtonColorFromStatus = map[string]rl.Color{
	WaitingBtnStatus: rl.Black,
	HoveredBtnStatus: rl.Yellow,
	ClickedBtnStatus: rl.Green,
}

// Food
var FoodColor = rl.Brown
var FoodBorderColor = rl.DarkBrown

// Pollute
var PolluteColor = rl.NewColor(0, 0, 0, 200)
