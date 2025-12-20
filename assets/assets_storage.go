package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Fonts paths
const (
	WinterFontPath = "assets/fonts/Winter.ttf"
	MolotFontPath  = "assets/fonts/Molot.ttf"
)

// Sound paths
const (
	BgMusicPath       = "assets/sounds/Background.mp3"
	AquariumSoundPath = "assets/sounds/AquariumBackground.mp3"
	CleanSoundPath    = "assets/sounds/CleanSound.mp3"
	FoodDropSoundPath = "assets/sounds/FoodDropSound.mp3"
)

// Sprites paths
const (
	AquariumBgSprite     = "assets/sprites/Screens/Aquarium.png"
	MenuBgSprite         = "assets/sprites/Screens/Menu.png"
	SettingsBgSprite     = "assets/sprites/Screens/Settings.png"
	AchievementsBgSprite = "assets/sprites/Screens/Achievements.png"
	StatsBgSprite        = "assets/sprites/Screens/Stats.png"

	ShopBgSprite        = "assets/sprites/Screens/Shop.png"
	ShopShrimpsSprite   = "assets/sprites/Screens/ShopShrimps.png"
	ShopWallpaperSprite = "assets/sprites/Screens/ShopWallpaper.png"

	CherryShrimpSprite            = "assets/sprites/Shrimps/CherryShrimp.png"
	CherryShrimpReversedSprite    = "assets/sprites/Shrimps/CherryShrimpReversed.png"
	GundamShrimpSprite            = "assets/sprites/Shrimps/GundamShrimp.png"
	GundamShrimpReversedSprite    = "assets/sprites/Shrimps/GundamShrimpReversed.png"
	IsaacShrimpSprite             = "assets/sprites/Shrimps/IsaacShrimp.png"
	IsaacShrimpReversedSprite     = "assets/sprites/Shrimps/IsaacShrimpReversed.png"
	MinecraftShrimpSprite         = "assets/sprites/Shrimps/MinecraftShrimp.png"
	MinecraftShrimpReversedSprite = "assets/sprites/Shrimps/MinecraftShrimpReversed.png"
	MiskaShrimpSprite             = "assets/sprites/Shrimps/MiskaShrimp.png"
	MiskaShrimpReversedSprite     = "assets/sprites/Shrimps/MiskaShrimpReversed.png"
	ChanelShrimpSprite            = "assets/sprites/Shrimps/ChanelShrimp.png"
	ChanelShrimpReversedSprite    = "assets/sprites/Shrimps/ChanelShrimpReversed.png"
	BlackRoseShrimpSprite         = "assets/sprites/Shrimps/BlackRoseShrimp.png"
	BlackRoseShrimpReversedSprite = "assets/sprites/Shrimps/BlackRoseShrimpReversed.png"
	SonicShrimpSprite             = "assets/sprites/Shrimps/SonicShrimp.png"
	SonicShrimpReversedSprite     = "assets/sprites/Shrimps/SonicShrimpReversed.png"

	PvZWallpaper    = "assets/sprites/Wallpapers/PvZWallpaper.png"
	CityWallpaper   = "assets/sprites/Wallpapers/CityWallpaper.png"
	NiceWallpaper   = "assets/sprites/Wallpapers/NiceWallpaper.png"
	GundamWallpaper = "assets/sprites/Wallpapers/GundamWallpaper.png"

	CoinSprite       = "assets/sprites/Other/Coin.png"
	MuteSprite       = "assets/sprites/Other/Mute.png"
	UnmuteSprite     = "assets/sprites/Other/Unmute.png"
	PolluteSprite    = "assets/sprites/Other/Pollute.png"
	TrophySprite     = "assets/sprites/Other/Trophy.png"
	GrayTrophySprite = "assets/sprites/Other/GrayTrophy.png"
)

type AssetStorage struct {
	assetMgr *AssetManager

	WinterFont rl.Font
	MolotFont  rl.Font

	BackgroundMusic rl.Music
	AquariumSound   rl.Music
	CleanSound      rl.Music
	FoodDropSound   rl.Music

	MenuScreen         rl.Texture2D
	AquariumScreen     rl.Texture2D
	SettingsScreen     rl.Texture2D
	AchievementsScreen rl.Texture2D
	StatsScreen        rl.Texture2D

	ShopScreen    rl.Texture2D
	ShopShrimps   rl.Texture2D
	ShopWallpaper rl.Texture2D

	CherryShrimp            rl.Texture2D
	CherryShrimpReversed    rl.Texture2D
	GundamShrimp            rl.Texture2D
	GundamShrimpReversed    rl.Texture2D
	IsaacShrimp             rl.Texture2D
	IsaacShrimpReversed     rl.Texture2D
	MinecraftShrimp         rl.Texture2D
	MinecraftShrimpReversed rl.Texture2D
	MiskaShrimp             rl.Texture2D
	MiskaShrimpReversed     rl.Texture2D
	ChanelShrimp            rl.Texture2D
	ChanelShrimpReversed    rl.Texture2D
	BlackRoseShrimp         rl.Texture2D
	BlackRoseShrimpReversed rl.Texture2D
	SonicShrimp             rl.Texture2D
	SonicShrimpReversed     rl.Texture2D

	PvZWallpaper    rl.Texture2D
	CityWallpaper   rl.Texture2D
	NiceWallpaper   rl.Texture2D
	GundamWallpaper rl.Texture2D

	Coin       rl.Texture2D
	Mute       rl.Texture2D
	Unmute     rl.Texture2D
	Pollute    rl.Texture2D
	Trophy     rl.Texture2D
	GrayTrophy rl.Texture2D
}

func NewTextureStorage(manager *AssetManager) *AssetStorage {
	ts := new(AssetStorage)

	ts.WinterFont = manager.LoadFont(WinterFontPath)
	ts.MolotFont = manager.LoadFont(MolotFontPath)

	ts.BackgroundMusic = manager.LoadMusic(BgMusicPath)
	ts.AquariumSound = manager.LoadMusic(AquariumSoundPath)
	ts.CleanSound = manager.LoadMusic(CleanSoundPath)
	ts.FoodDropSound = manager.LoadMusic(FoodDropSoundPath)

	ts.MenuScreen = manager.LoadScreensTexture(MenuBgSprite)
	ts.AquariumScreen = manager.LoadScreensTexture(AquariumBgSprite)
	ts.SettingsScreen = manager.LoadScreensTexture(SettingsBgSprite)
	ts.AchievementsScreen = manager.LoadScreensTexture(AchievementsBgSprite)
	ts.StatsScreen = manager.LoadScreensTexture(StatsBgSprite)

	ts.ShopScreen = manager.LoadScreensTexture(ShopBgSprite)
	ts.ShopShrimps = manager.LoadScreensTexture(ShopShrimpsSprite)
	ts.ShopWallpaper = manager.LoadScreensTexture(ShopWallpaperSprite)

	ts.CherryShrimp = manager.LoadShrimpsTexture(CherryShrimpSprite)
	ts.CherryShrimpReversed = manager.LoadShrimpsTexture(CherryShrimpReversedSprite)
	ts.GundamShrimp = manager.LoadShrimpsTexture(GundamShrimpSprite)
	ts.GundamShrimpReversed = manager.LoadShrimpsTexture(GundamShrimpReversedSprite)
	ts.IsaacShrimp = manager.LoadShrimpsTexture(IsaacShrimpSprite)
	ts.IsaacShrimpReversed = manager.LoadShrimpsTexture(IsaacShrimpReversedSprite)
	ts.MinecraftShrimp = manager.LoadShrimpsTexture(MinecraftShrimpSprite)
	ts.MinecraftShrimpReversed = manager.LoadShrimpsTexture(MinecraftShrimpReversedSprite)
	ts.MiskaShrimp = manager.LoadShrimpsTexture(MiskaShrimpSprite)
	ts.MiskaShrimpReversed = manager.LoadShrimpsTexture(MiskaShrimpReversedSprite)
	ts.ChanelShrimp = manager.LoadShrimpsTexture(ChanelShrimpSprite)
	ts.ChanelShrimpReversed = manager.LoadShrimpsTexture(ChanelShrimpReversedSprite)
	ts.BlackRoseShrimp = manager.LoadShrimpsTexture(BlackRoseShrimpSprite)
	ts.BlackRoseShrimpReversed = manager.LoadShrimpsTexture(BlackRoseShrimpReversedSprite)
	ts.SonicShrimp = manager.LoadShrimpsTexture(SonicShrimpSprite)
	ts.SonicShrimpReversed = manager.LoadShrimpsTexture(SonicShrimpReversedSprite)

	ts.PvZWallpaper = manager.LoadWallpapersTexture(PvZWallpaper)
	ts.CityWallpaper = manager.LoadWallpapersTexture(CityWallpaper)
	ts.NiceWallpaper = manager.LoadWallpapersTexture(NiceWallpaper)
	ts.GundamWallpaper = manager.LoadWallpapersTexture(GundamWallpaper)

	ts.Coin = manager.LoadOtherTexture(CoinSprite)
	ts.Mute = manager.LoadOtherTexture(MuteSprite)
	ts.Unmute = manager.LoadOtherTexture(UnmuteSprite)
	ts.Pollute = manager.LoadOtherTexture(PolluteSprite)
	ts.Trophy = manager.LoadOtherTexture(TrophySprite)
	ts.GrayTrophy = manager.LoadOtherTexture(GrayTrophySprite)

	return ts
}
