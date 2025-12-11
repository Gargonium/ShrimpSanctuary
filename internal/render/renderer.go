package render

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Render struct {
	bgTexture             rl.Texture2D
	shrimpTexture         rl.Texture2D
	shrimpTextureReversed rl.Texture2D
	btnFont               rl.Font
}

func NewRender() *Render {
	r := Render{}
	r.bgTexture = spriteToTexture(config.BgSprite)
	r.shrimpTexture = spriteToTexture(config.CherryShrimpSprite)
	r.shrimpTextureReversed = spriteToTexture(config.CherryShrimpReversedSprite)
	r.btnFont = loadFont(config.WinterFont)
	return &r
}

func loadFont(fontPath string) rl.Font {
	fontTtf := rl.LoadFont(fontPath) //rl.LoadFontEx(config.WinterFont, 8, 0, 250)
	return fontTtf
}

func spriteToTexture(spritePath string) rl.Texture2D {
	image := rl.LoadImage(spritePath)
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	return texture
}

func (r *Render) Draw(g *game.Game) {
	r.DrawBackground()
	r.DrawShrimps(g)
}
