package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawShrimps() {
	for i := range r.game.Shrimps {
		if r.game.Shrimps[i].Vx < 0 {
			rl.DrawTexture(r.shrimpTextureReversed, r.game.Shrimps[i].X, r.game.Shrimps[i].Y, rl.White)
		} else {
			rl.DrawTexture(r.shrimpTexture, r.game.Shrimps[i].X, r.game.Shrimps[i].Y, rl.White)
		}
	}
}
