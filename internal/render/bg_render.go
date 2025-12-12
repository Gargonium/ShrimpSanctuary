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
	rl.DrawTextEx(r.btnFont, config.FeedBtnName, rl.Vector2{X: config.FeedBtnTextX, Y: config.ButtonTextY}, config.OtherBtnFontSize, 2, r.buttonsColor["FEED"])
	rl.DrawTextEx(r.btnFont, config.CleanBtnName, rl.Vector2{X: config.CleanBtnTextX, Y: config.ButtonTextY}, config.CleanBtnFontSize, 2, r.buttonsColor["CLEAN"])
	rl.DrawTextEx(r.btnFont, config.ShopBtnName, rl.Vector2{X: config.ShopBtnTextX, Y: config.ButtonTextY}, config.OtherBtnFontSize, 2, r.buttonsColor["SHOP"])
	rl.DrawTextEx(r.btnFont, config.ExitBtnName, rl.Vector2{X: config.ExitBtnTextX, Y: config.ButtonTextY}, config.OtherBtnFontSize, 2, r.buttonsColor["EXIT"])
}
