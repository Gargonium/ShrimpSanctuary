package config

import rl "github.com/gen2brain/raylib-go/raylib"

// Game constants
const (
	FPS          = 60
	ScreenWidth  = 1280
	ScreenHeight = 720
	StartMoney   = 100000
)

// Aquarium constants
const (
	AquariumX           = 140
	AquariumY           = 20
	AquariumWidth       = 1000
	AquariumHeight      = 500
	AquariumBorderWidth = 5
	SandHeight          = 25
	MoneyX              = 10
	MoneyY              = 25
	WaterX              = 145
	WaterY              = 25
	WaterWidth          = 990
	WaterHeight         = 465
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
	BorderOffset      = Offset5
)

// Offsets
const (
	Offset5 = 5
)

type ButtonStatus string

// Buttons
const (
	WaitingBtnStatus ButtonStatus = "waiting"
	HoveredBtnStatus ButtonStatus = "hovered"
	ClickedBtnStatus ButtonStatus = "clicked"
)

var PlayFieldBounds = rl.Rectangle{X: PlayFieldX, Y: PlayFieldY, Width: PlayFieldWidth, Height: PlayerFieldHeight}

// Shrimps
const (
	ShrimpStartCount        = 3
	ShrimpMaxVelocity       = 0.5
	ShrimpBehaviourMaxDelay = FPS * 5
	ShrimpMoneyDelay        = FPS * 60
	ShrimpMaxHunger         = FPS * 60 * 2
)

type ShrimpType int

const (
	CherryShrimp ShrimpType = iota
)

var MoneyByShrimp = map[ShrimpType]int{
	CherryShrimp: 25,
}

var ShrimpsTypesInShop = []ShrimpType{
	CherryShrimp,
	CherryShrimp,
	CherryShrimp,
	CherryShrimp,
	CherryShrimp,
	CherryShrimp,
	CherryShrimp,
	CherryShrimp,
}

var ShrimpCost = map[ShrimpType]int{
	CherryShrimp: 100,
}

const (
	StandardSquareSpriteSide = 32
	BigSquareSpriteSide      = 64
)

// Fonts paths
const (
	WinterFont = "assets/fonts/Winter.ttf"
)

// Sound paths
const (
	BgMusicPath = "assets/sounds/Background.mp3"
)

// Food
const (
	FoodRadius   = 5
	FoodVelocity = 1
	FoodLifeTime = FPS * 5
)

// Pollute
const (
	PolluteSpawnDelay       = FPS * 120
	PolluteSpawnDelaySpread = FPS * 30
	PolluteMaxDurability    = 4
	PolluteRadius           = 10
)

// GameState
type GameState int

const (
	StateMenu GameState = iota
	StateAquarium
	StateSettings
	StateShop
	StateQuit
)

type ShopState int

const (
	ShopStateShrimps ShopState = iota
	ShopStateWallpaper
	ShopStateDecor
)

type WallpaperState int

const (
	DefaultWallpaperState WallpaperState = iota
	PvZWallpaperState
)

var WallpaperTypesInShop = []WallpaperState{
	PvZWallpaperState,
	PvZWallpaperState,
	PvZWallpaperState,
	PvZWallpaperState,
}

var WallpaperCost = map[WallpaperState]int{
	PvZWallpaperState: 1000,
}

const (
	AutoSaveDelay = FPS * 5 * 60
)
