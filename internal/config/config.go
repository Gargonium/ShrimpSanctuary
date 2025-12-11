package config

// Game constants
const (
	FPS          = 20
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
	SandWidth           = AquariumWidth - 2*AquariumBorderWidth
	SandX               = AquariumX + AquariumBorderWidth
	SandY               = AquariumY + AquariumHeight - AquariumBorderWidth - SandHeight
	TableX              = AquariumX
	TableY              = AquariumY + AquariumHeight
	TableWidth          = AquariumWidth
	TableHeight         = 30
	TableLegWidth       = 25
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

// Buttons
const (
	ButtonY          = 630
	FeedBtnX         = 245
	CleanBtnX        = 440
	ShopBtnX         = 643
	ExitBtnX         = 850
	CleanBtnFontSize = 50
	OtherBtnFontSize = 55
)
