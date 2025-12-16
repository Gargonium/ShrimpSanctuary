package config

import (
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureStorage struct {
	MenuScreen     rl.Texture2D
	AquariumScreen rl.Texture2D
	SettingsScreen rl.Texture2D
	ShopScreen     rl.Texture2D

	CherryShrimp         rl.Texture2D
	CherryShrimpReversed rl.Texture2D

	Coin rl.Texture2D
}

func NewTextureStorage() *TextureStorage {
	ts := new(TextureStorage)
	ts.MenuScreen = utils.SpriteToTexture(MenuBgSprite)
	ts.AquariumScreen = utils.SpriteToTexture(AquariumBgSprite)
	ts.SettingsScreen = utils.SpriteToTexture(SettingsBgSprite)
	ts.ShopScreen = utils.SpriteToTexture(ShopBgSprite)

	ts.CherryShrimp = utils.SpriteToTexture(CherryShrimpSprite)
	ts.CherryShrimpReversed = utils.SpriteToTexture(CherryShrimpReversedSprite)

	ts.Coin = utils.SpriteToTexture(CoinSprite)

	return ts
}
