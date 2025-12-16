package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShopScreen struct {
	Game    *game.Game
	Buttons []*input.Button
	ts      *config.TextureStorage
}

func NewShopScreen(game *game.Game, ts *config.TextureStorage) *ShopScreen {
	ss := new(ShopScreen)
	ss.Game = game
	ss.ts = ts
	return ss
}

func (ss *ShopScreen) HandleInput() {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		ss.Game.State = config.StateAquarium
	}
}
func (ss *ShopScreen) Draw() {
	rl.DrawTexture(ss.ts.ShopScreen, 0, 0, rl.White)
}
func (ss *ShopScreen) drawButtons() {

}
