package entities

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	Position rl.Vector2
	lifeTime int
}

func NewFood(pos rl.Vector2) *Food {
	f := new(Food)
	f.Position = pos

	f.lifeTime = config.FoodLifeTime
	return f
}

func (f *Food) MoveAndDisappear() bool {
	if f.lifeTime != 0 && f.Position.Y < config.PlayFieldY+config.PlayerFieldHeight {
		f.Position.Y += config.FoodVelocity
		f.lifeTime--
	} else if f.lifeTime == 0 {
		return true
	}
	return false
}
