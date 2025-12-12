package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawShrimps() {
	for i := range r.game.Shrimps {
		if r.game.Shrimps[i].Vx < 0 {
			rl.DrawTextureV(r.shrimpTextureReversed, r.game.Shrimps[i].Position, rl.White)
		} else {
			rl.DrawTextureV(r.shrimpTexture, r.game.Shrimps[i].Position, rl.White)
		}
	}
}
