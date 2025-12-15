package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShopScreen struct {
	Game      *game.Game
	Buttons   []*input.Button
	bgTexture rl.Texture2D
}

func NewShopScreen(game *game.Game) *ShopScreen {
	ss := new(ShopScreen)
	ss.Game = game
	ss.bgTexture = utils.SpriteToTexture(config.ShopBgSprite)
	return ss
}

func (ss *ShopScreen) HandleInput() {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		ss.Game.State = config.StateAquarium
	}
}
func (ss *ShopScreen) Draw() {
	rl.DrawTexture(ss.bgTexture, 0, 0, rl.White)
}
func (ss *ShopScreen) drawButtons() {

}
