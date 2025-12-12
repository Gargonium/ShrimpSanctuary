package config

import rl "github.com/gen2brain/raylib-go/raylib"

// Game constants
const (
	FPS          = 60
	ScreenWidth  = 1200
	ScreenHeight = 800
)

// Aquarium constants
const (
	AquariumX           = 100
	AquariumY           = 100
	AquariumWidth       = 1000
	AquariumHeight      = 500
	AquariumBorderWidth = 5
	SandHeight          = 25
	//SandWidth           = AquariumWidth - 2*AquariumBorderWidth
	//SandX               = AquariumX + AquariumBorderWidth
	//SandY               = AquariumY + AquariumHeight - AquariumBorderWidth - SandHeight
	//TableX              = AquariumX
	//TableY              = AquariumY + AquariumHeight
	//TableWidth          = AquariumWidth
	//TableHeight         = 30
	//TableLegWidth       = 25
)

// Dimensions of the playing field
const (
	PlayFieldX        = AquariumX + AquariumBorderWidth
	PlayFieldY        = AquariumY + AquariumBorderWidth
	PlayFieldWidth    = AquariumWidth - 2*AquariumBorderWidth
	PlayerFieldHeight = AquariumHeight - 2*AquariumBorderWidth - SandHeight
	BorderOffset      = 5
)

// Shrimps
const (
	ShrimpWidth      = 32
	ShrimpHeight     = 32
	ShrimpStartCount = 3
	ShrimpVelocity   = 1
	ShrimpMaxDelay   = 300
)

// Sprites paths
const (
	BgSprite                   = "assets/sprites/Background.png"
	CherryShrimpSprite         = "assets/sprites/CherryShrimp.png"
	CherryShrimpReversedSprite = "assets/sprites/CherryShrimpReversed.png"
)

// Fonts paths
const (
	WinterFont = "assets/fonts/Winter.ttf"
)

// Sound paths
const (
	BgMusicPath = "assets/sounds/Background.mp3"
)

// ButtonsX
const (
	ButtonTextY      = 630
	FeedBtnTextX     = 245
	CleanBtnTextX    = 440
	ShopBtnTextX     = 643
	ExitBtnTextX     = 850
	CleanBtnFontSize = 50
	OtherBtnFontSize = 55

	ButtonY      = 600
	ButtonWidth  = 132
	ButtonHeight = 136
	FeedBtnX     = 234
	CleanBtnX    = 434
	ShopBtnX     = 634
	ExitBtnX     = 834

	FeedBtnName  = "FEED"
	CleanBtnName = "CLEAN"
	ShopBtnName  = "SHOP"
	ExitBtnName  = "EXIT"

	WaitingBtnStatus = "waiting"
	HoveredBtnStatus = "hovered"
	ClickedBtnStatus = "clicked"
)

var ButtonsX = map[string]float32{
	FeedBtnName:  FeedBtnX,
	CleanBtnName: CleanBtnX,
	ShopBtnName:  ShopBtnX,
	ExitBtnName:  ExitBtnX,
}

var ButtonColorFromStatus = map[string]rl.Color{
	WaitingBtnStatus: rl.Black,
	HoveredBtnStatus: rl.Yellow,
	ClickedBtnStatus: rl.Green,
}
