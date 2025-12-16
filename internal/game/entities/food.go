package entities

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	Position rl.Vector2
	lifeTime int
	IsAlive  bool
}

func NewFood(pos rl.Vector2) *Food {
	f := new(Food)
	f.Position = pos
	f.IsAlive = true

	f.lifeTime = config.FoodLifeTime
	return f
}

func (f *Food) SelfDestruct() {
	f.IsAlive = false
}

func (f *Food) MoveAndDisappear() {
	if !f.IsAlive {
		return
	}
	if f.lifeTime != 0 && f.Position.Y < config.PlayFieldY+config.PlayerFieldHeight {
		f.Position.Y += config.FoodVelocity
	} else if f.lifeTime == 0 {
		f.SelfDestruct()
		return
	}
	f.lifeTime--
}
