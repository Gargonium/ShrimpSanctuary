package render

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/render/screens"
	"ShrimpSanctuary/internal/sound_bar"
)

type Render struct {
	Game           *game.Game
	sb             sound_bar.SoundBar
	AquariumScreen screens.AquariumScreen
}

func NewRender(g *game.Game, sb sound_bar.SoundBar) *Render {
	r := new(Render)
	r.AquariumScreen = screens.NewAquariumScreen(g)
	r.Game = g
	r.sb = sb

	return r
}

func (r *Render) Draw() {
	r.AquariumScreen.Draw()
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
	}
}
