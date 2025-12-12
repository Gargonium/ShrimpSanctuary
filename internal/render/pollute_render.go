package render

import (
	"ShrimpSanctuary/internal/colors"
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawPollute() {
	for i := range r.game.Pollution {
		polCol := colors.PolluteColor
		polCol.A = uint8(float32(polCol.A) * float32(r.game.Pollution[i].Durability) / float32(config.PolluteMaxDurability))
		rl.DrawCircleV(r.game.Pollution[i].Position, config.PolluteRadius, polCol)
	}
}
