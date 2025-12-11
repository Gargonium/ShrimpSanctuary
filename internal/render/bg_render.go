package render

import (
	"ShrimpSanctuary/internal/config"
	_ "ShrimpSanctuary/internal/input"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawBackground() {
	rl.DrawTexture(r.bgTexture, 0, 0, rl.White)
	r.drawButtons()
}

func (r *Render) drawButtons() {
	//mouseX, mouseY := input.HandleMouseInput()
	rl.DrawTextEx(r.btnFont, "FEED", rl.Vector2{X: config.FeedBtnX, Y: config.ButtonY}, config.OtherBtnFontSize, 2, rl.Black)
	rl.DrawTextEx(r.btnFont, "CLEAN", rl.Vector2{X: config.CleanBtnX, Y: config.ButtonY}, config.CleanBtnFontSize, 2, rl.Black)
	rl.DrawTextEx(r.btnFont, "SHOP", rl.Vector2{X: config.ShopBtnX, Y: config.ButtonY}, config.OtherBtnFontSize, 2, rl.Black)
	rl.DrawTextEx(r.btnFont, "EXIT", rl.Vector2{X: config.ExitBtnX, Y: config.ButtonY}, config.OtherBtnFontSize, 2, rl.Black)
}
