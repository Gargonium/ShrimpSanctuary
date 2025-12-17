package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Fonts paths
const (
	WinterFontPath = "assets/fonts/Winter.ttf"
)

// Sound paths
const (
	BgMusicPath = "assets/sounds/Background.mp3"
)

// Sprites paths
const (
	AquariumBgSprite = "assets/sprites/Screens/Aquarium.png"
	MenuBgSprite     = "assets/sprites/Screens/Menu.png"
	SettingsBgSprite = "assets/sprites/Screens/Settings.png"

	ShopBgSprite        = "assets/sprites/Screens/Shop.png"
	ShopShrimpsSprite   = "assets/sprites/Screens/ShopShrimps.png"
	ShopWallpaperSprite = "assets/sprites/Screens/ShopWallpaper.png"
	ShopDecorSprite     = "assets/sprites/Screens/ShopDecor.png"

	CherryShrimpSprite         = "assets/sprites/Shrimps/CherryShrimp.png"
	CherryShrimpReversedSprite = "assets/sprites/Shrimps/CherryShrimpReversed.png"

	PvZWallpaper    = "assets/sprites/Wallpapers/PvZWallpaper.png"
	CityWallpaper   = "assets/sprites/Wallpapers/CityWallpaper.png"
	NiceWallpaper   = "assets/sprites/Wallpapers/NiceWallpaper.png"
	GundamWallpaper = "assets/sprites/Wallpapers/GundamWallpaper.png"

	CoinSprite   = "assets/sprites/Other/Coin.png"
	MuteSprite   = "assets/sprites/Other/Mute.png"
	UnmuteSprite = "assets/sprites/Other/Unmute.png"
)

type AssetStorage struct {
	assetMgr *AssetManager

	BackgroundMusic rl.Music

	WinterFont rl.Font

	Mute   rl.Texture2D
	Unmute rl.Texture2D

	MenuScreen     rl.Texture2D
	AquariumScreen rl.Texture2D
	SettingsScreen rl.Texture2D

	ShopScreen    rl.Texture2D
	ShopShrimps   rl.Texture2D
	ShopWallpaper rl.Texture2D
	ShopDecor     rl.Texture2D

	CherryShrimp         rl.Texture2D
	CherryShrimpReversed rl.Texture2D

	PvZWallpaper    rl.Texture2D
	CityWallpaper   rl.Texture2D
	NiceWallpaper   rl.Texture2D
	GundamWallpaper rl.Texture2D

	Coin rl.Texture2D
}

func NewTextureStorage(manager *AssetManager) *AssetStorage {
	ts := new(AssetStorage)

	ts.BackgroundMusic = manager.LoadMusic(BgMusicPath)

	ts.WinterFont = manager.LoadFont(WinterFontPath)

	ts.MenuScreen = manager.LoadScreensTexture(MenuBgSprite)
	ts.AquariumScreen = manager.LoadScreensTexture(AquariumBgSprite)
	ts.SettingsScreen = manager.LoadScreensTexture(SettingsBgSprite)

	ts.ShopScreen = manager.LoadScreensTexture(ShopBgSprite)
	ts.ShopShrimps = manager.LoadScreensTexture(ShopShrimpsSprite)
	ts.ShopWallpaper = manager.LoadScreensTexture(ShopWallpaperSprite)
	ts.ShopDecor = manager.LoadScreensTexture(ShopDecorSprite)

	ts.CherryShrimp = manager.LoadShrimpsTexture(CherryShrimpSprite)
	ts.CherryShrimpReversed = manager.LoadShrimpsTexture(CherryShrimpReversedSprite)

	ts.PvZWallpaper = manager.LoadWallpapersTexture(PvZWallpaper)
	ts.CityWallpaper = manager.LoadWallpapersTexture(CityWallpaper)
	ts.NiceWallpaper = manager.LoadWallpapersTexture(NiceWallpaper)
	ts.GundamWallpaper = manager.LoadWallpapersTexture(GundamWallpaper)

	ts.Coin = manager.LoadOtherTexture(CoinSprite)
	ts.Mute = manager.LoadOtherTexture(MuteSprite)
	ts.Unmute = manager.LoadOtherTexture(UnmuteSprite)

	return ts
}
