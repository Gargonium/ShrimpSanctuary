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
	BlackRoseShrimp
	IsaacShrimp
	MinecraftShrimp
	GundamShrimp
	SonicShrimp
	MiskaShrimp
	ChanelShrimp
)

var MoneyByShrimp = map[ShrimpType]int{
	CherryShrimp:    25,
	BlackRoseShrimp: 25,
	IsaacShrimp:     50,
	MinecraftShrimp: 50,
	GundamShrimp:    75,
	SonicShrimp:     75,
	MiskaShrimp:     100,
	ChanelShrimp:    100,
}

var ShrimpsTypesInShop = []ShrimpType{
	CherryShrimp,
	BlackRoseShrimp,
	IsaacShrimp,
	MinecraftShrimp,
	GundamShrimp,
	SonicShrimp,
	MiskaShrimp,
	ChanelShrimp,
}

var ShrimpCost = map[ShrimpType]int{
	CherryShrimp:    100,
	BlackRoseShrimp: 100,
	IsaacShrimp:     300,
	MinecraftShrimp: 300,
	GundamShrimp:    500,
	SonicShrimp:     500,
	MiskaShrimp:     700,
	ChanelShrimp:    700,
}

const (
	StandardSquareSpriteSide = 32
	BigSquareSpriteSide      = 64
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
	CityWallpaperState
	NiceWallpaperState
	GundamWallpaperState
)

var WallpaperTypesInShop = []WallpaperState{
	PvZWallpaperState,
	CityWallpaperState,
	NiceWallpaperState,
	GundamWallpaperState,
}

var WallpaperCost = map[WallpaperState]int{
	PvZWallpaperState:    1000,
	CityWallpaperState:   1000,
	NiceWallpaperState:   1000,
	GundamWallpaperState: 10000,
}

const (
	AutoSaveDelay = FPS * 5 * 60
)
