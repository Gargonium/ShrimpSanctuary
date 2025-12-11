package game

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	"math/rand"
)

type Shrimp struct {
	X, Y      int32
	Vx, Vy    int32
	Delay     int32
	Behaviour int32
}

func NewShrimp() *Shrimp {
	shrimp := new(Shrimp)
	shrimp.X = (rand.Int31() % config.PlayFieldWidth) + config.PlayFieldX
	shrimp.Y = (rand.Int31() % config.PlayerFieldHeight) + config.PlayFieldY
	shrimp.Vx, shrimp.Vy = config.ShrimpVelocity, config.ShrimpVelocity
	shrimp.Delay = config.ShrimpMaxDelay
	shrimp.Behaviour = rand.Int31() % 3
	shrimp.ShrimpWallCollide()
	return shrimp
}

func (s *Shrimp) ShrimpWallCollide() {

	minX := int32(config.PlayFieldX + config.BorderOffset)
	maxX := int32(config.PlayFieldX + config.PlayFieldWidth - config.ShrimpWidth - config.BorderOffset)
	minY := int32(config.PlayFieldY + config.BorderOffset)
	maxY := int32(config.PlayFieldY + config.PlayerFieldHeight - config.ShrimpHeight - config.BorderOffset)

	s.X, s.Vx = utils.ClampAndBounce(s.X, minX, maxX, s.Vx)
	s.Y, s.Vy = utils.ClampAndBounce(s.Y, minY, maxY, s.Vy)
}

func (s *Shrimp) Move() {
	s.Delay--
	if s.Delay == 0 {
		s.Behaviour = rand.Int31() % 3
		if s.Behaviour != 0 {
			rvx := rand.Int31() % 3
			rvy := rand.Int31() % 3
			switch rvx {
			case 0:
				s.Vx = 0
			case 1:
				s.Vx = config.ShrimpVelocity
			case 2:
				s.Vx = -config.ShrimpVelocity
			}
			switch rvy {
			case 0:
				s.Vy = 0
			case 1:
				s.Vy = config.ShrimpVelocity
			case 2:
				s.Vy = -config.ShrimpVelocity
			}
		}
		s.Delay = rand.Int31()%config.ShrimpMaxDelay + config.FPS
	}
	if s.Behaviour != 0 {
		s.X += s.Vx
		s.Y += s.Vy
		s.ShrimpWallCollide()
	}
}
