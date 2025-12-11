package render

import (
	"ShrimpSanctuary/internal/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawShrimps(g *game.Game) {
	for i := range g.Shrimps {
		if g.Shrimps[i].Vx < 0 {
			rl.DrawTexture(r.shrimpTextureReversed, g.Shrimps[i].X, g.Shrimps[i].Y, rl.White)
		} else {
			rl.DrawTexture(r.shrimpTexture, g.Shrimps[i].X, g.Shrimps[i].Y, rl.White)
		}
	}
}
