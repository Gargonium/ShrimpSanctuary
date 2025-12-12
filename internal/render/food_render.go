package render

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Render) DrawFood() {
	for i := range r.game.Foods {
		rl.DrawCircle(r.game.Foods[i].X, r.game.Foods[i].Y, config.FoodRadius, config.FoodColor)
		rl.DrawCircleLines(r.game.Foods[i].X, r.game.Foods[i].Y, config.FoodRadius, config.FoodBorderColor)
	}
}
