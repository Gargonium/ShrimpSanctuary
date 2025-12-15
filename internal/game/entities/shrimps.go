package entities

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Shrimp struct {
	Position rl.Vector2
	Vx, Vy   float32
	Delay    int32
}

func NewShrimp() *Shrimp {
	shrimp := new(Shrimp)
	shrimp.Position.X = (rand.Float32() * config.PlayFieldWidth) + config.PlayFieldX
	shrimp.Position.Y = (rand.Float32() * config.PlayerFieldHeight) + config.PlayFieldY
	shrimp.Vx, shrimp.Vy = config.ShrimpMaxVelocity, config.ShrimpMaxVelocity
	shrimp.Delay = config.ShrimpBehaviourMaxDelay
	shrimp.ShrimpWallCollide()
	return shrimp
}

func (s *Shrimp) ShrimpWallCollide() {

	minX := float32(config.PlayFieldX + config.BorderOffset)
	maxX := float32(config.PlayFieldX + config.PlayFieldWidth - config.ShrimpWidth - config.BorderOffset)
	minY := float32(config.PlayFieldY + config.BorderOffset)
	maxY := float32(config.PlayFieldY + config.PlayerFieldHeight - config.ShrimpHeight - config.BorderOffset)

	s.Position.X, s.Vx = utils.ClampAndBounce(s.Position.X, minX, maxX, s.Vx)
	s.Position.Y, s.Vy = utils.ClampAndBounce(s.Position.Y, minY, maxY, s.Vy)
}

// Move TODO Переделать
func (s *Shrimp) Move() {
	s.Delay--
	if s.Delay == 0 {

		s.Vx = rand.Float32()*2*config.ShrimpMaxVelocity - config.ShrimpMaxVelocity
		s.Vy = rand.Float32()*2*config.ShrimpMaxVelocity - config.ShrimpMaxVelocity

		s.Delay = rand.Int31()%config.ShrimpBehaviourMaxDelay + config.FPS
	}
	s.Position.X += s.Vx
	s.Position.Y += s.Vy
	s.ShrimpWallCollide()
}
