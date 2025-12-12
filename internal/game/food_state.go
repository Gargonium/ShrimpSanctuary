package game

import "ShrimpSanctuary/internal/config"

type Food struct {
	X, Y     int32
	lifeTime int
}

func NewFood(X, Y int32) Food {
	f := Food{}
	f.X = X
	f.Y = Y
	f.lifeTime = config.FoodLifeTime
	return f
}

func (f *Food) MoveAndDisappear() bool {
	if f.lifeTime != 0 && f.Y < config.PlayFieldY+config.PlayerFieldHeight {
		f.Y += config.FoodVelocity
		f.lifeTime--
	} else if f.lifeTime == 0 {
		return true
	}
	return false
}
