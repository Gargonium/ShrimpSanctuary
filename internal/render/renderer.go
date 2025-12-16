package render

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/render/screens"
	"ShrimpSanctuary/internal/sound_bar"
)

type Render struct {
	Game           *game.Game
	sb             *sound_bar.SoundBar
	ts             *config.TextureStorage
	AquariumScreen *screens.AquariumScreen
	MenuScreen     *screens.MenuScreen
	SettingsScreen *screens.SettingsScreen
	ShopScreen     *screens.ShopScreen
}

func NewRender(g *game.Game, sb *sound_bar.SoundBar, ts *config.TextureStorage) *Render {
	r := new(Render)
	r.AquariumScreen = screens.NewAquariumScreen(g, ts)
	r.MenuScreen = screens.NewMenuScreen(g, ts)
	r.SettingsScreen = screens.NewSettingsScreen(g, sb, ts)
	r.ShopScreen = screens.NewShopScreen(g, ts)
	r.Game = g
	r.sb = sb
	r.ts = ts

	return r
}

func (r *Render) Draw() {
	switch r.Game.State {
	case config.StateAquarium:
		r.AquariumScreen.Draw()
	case config.StateMenu:
		r.MenuScreen.Draw()
	case config.StateSettings:
		r.SettingsScreen.Draw()
	case config.StateShop:
		r.ShopScreen.Draw()
	case config.StateQuit:
	}
}

func (r *Render) Update() {
	r.HandleInput()
	r.Game.Update()
	r.sb.Update()
}

func (r *Render) HandleInput() {
	switch r.Game.State {
	case config.StateAquarium:
		r.AquariumScreen.HandleInput()
	case config.StateMenu:
		r.MenuScreen.HandleInput()
	case config.StateSettings:
		r.SettingsScreen.HandleInput()
	case config.StateShop:
		r.ShopScreen.HandleInput()
	case config.StateQuit:
	}
}
