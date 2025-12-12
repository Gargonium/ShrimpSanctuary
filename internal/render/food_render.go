package render

import (
	"ShrimpSanctuary/internal/colors"
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawFood() {
	for i := range r.game.Foods {
		rl.DrawCircleV(r.game.Foods[i].Position, config.FoodRadius, colors.FoodColor)
		rl.DrawCircleLinesV(r.game.Foods[i].Position, config.FoodRadius, colors.FoodBorderColor)
	}
}
