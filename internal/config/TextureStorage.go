package config

import (
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
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

	Coin rl.Texture2D
}

func NewTextureStorage() *TextureStorage {
	ts := new(TextureStorage)

	ts.Mute = utils.SpriteToTexture(MuteSprite)
	ts.Unmute = utils.SpriteToTexture(UnmuteSprite)

	ts.MenuScreen = utils.SpriteToTexture(MenuBgSprite)
	ts.AquariumScreen = utils.SpriteToTexture(AquariumBgSprite)
	ts.SettingsScreen = utils.SpriteToTexture(SettingsBgSprite)

	ts.ShopScreen = utils.SpriteToTexture(ShopBgSprite)
	ts.ShopShrimps = utils.SpriteToTexture(ShopShrimpsSprite)
	ts.ShopWallpaper = utils.SpriteToTexture(ShopWallpaperSprite)
	ts.ShopDecor = utils.SpriteToTexture(ShopDecorSprite)

	ts.CherryShrimp = utils.SpriteToTexture(CherryShrimpSprite)
	ts.CherryShrimpReversed = utils.SpriteToTexture(CherryShrimpReversedSprite)

	ts.Coin = utils.SpriteToTexture(CoinSprite)

	return ts
}
