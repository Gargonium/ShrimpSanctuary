package config

import (
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
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

type TextureStorage struct {
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

func NewTextureStorage() *TextureStorage {
	ts := new(TextureStorage)

	ts.MenuScreen = utils.SpriteToTexture(MenuBgSprite)
	ts.AquariumScreen = utils.SpriteToTexture(AquariumBgSprite)
	ts.SettingsScreen = utils.SpriteToTexture(SettingsBgSprite)

	ts.ShopScreen = utils.SpriteToTexture(ShopBgSprite)
	ts.ShopShrimps = utils.SpriteToTexture(ShopShrimpsSprite)
	ts.ShopWallpaper = utils.SpriteToTexture(ShopWallpaperSprite)
	ts.ShopDecor = utils.SpriteToTexture(ShopDecorSprite)

	ts.CherryShrimp = utils.SpriteToTexture(CherryShrimpSprite)
	ts.CherryShrimpReversed = utils.SpriteToTexture(CherryShrimpReversedSprite)

	ts.PvZWallpaper = utils.SpriteToTexture(PvZWallpaper)
	ts.CityWallpaper = utils.SpriteToTexture(CityWallpaper)
	ts.NiceWallpaper = utils.SpriteToTexture(NiceWallpaper)
	ts.GundamWallpaper = utils.SpriteToTexture(GundamWallpaper)

	ts.Coin = utils.SpriteToTexture(CoinSprite)
	ts.Mute = utils.SpriteToTexture(MuteSprite)
	ts.Unmute = utils.SpriteToTexture(UnmuteSprite)

	return ts
}
