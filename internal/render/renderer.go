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
	AquariumScreen *screens.AquariumScreen
	MenuScreen     *screens.MenuScreen
	SettingsScreen *screens.SettingsScreen
}

func NewRender(g *game.Game, sb *sound_bar.SoundBar) *Render {
	r := new(Render)
	r.AquariumScreen = screens.NewAquariumScreen(g)
	r.MenuScreen = screens.NewMenuScreen(g)
	r.SettingsScreen = screens.NewSettingsScreen(g, sb)
	r.Game = g
	r.sb = sb

	return r
}

func (r *Render) Draw() {
	switch r.Game.State {
	case config.StatePlaying:
		r.AquariumScreen.Draw()
	case config.StateMenu:
		r.MenuScreen.Draw()
	case config.StateSettings:
		r.SettingsScreen.Draw()
	case config.StateShop:
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
	case config.StatePlaying:
		r.AquariumScreen.HandleInput()
	case config.StateMenu:
		r.MenuScreen.HandleInput()
	case config.StateSettings:
		r.SettingsScreen.HandleInput()
	case config.StateShop:
	case config.StateQuit:
	}
}
