package entities

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Shrimp struct {
	Position       rl.Vector2
	Vx, Vy         float32
	BehaviourDelay int32
	MoneyDelay     int32
	Type           config.ShrimpType
	Hunger         int
	IsAlive        bool
}

func NewShrimp(t config.ShrimpType) *Shrimp {
	shrimp := new(Shrimp)
	shrimp.Position.X = (rand.Float32() * config.PlayFieldWidth) + config.PlayFieldX
	shrimp.Position.Y = (rand.Float32() * config.PlayerFieldHeight) + config.PlayFieldY
	shrimp.Vx = rand.Float32()*2*config.ShrimpMaxVelocity - config.ShrimpMaxVelocity
	shrimp.Vy = rand.Float32()*2*config.ShrimpMaxVelocity - config.ShrimpMaxVelocity
	shrimp.BehaviourDelay = rand.Int31()%config.ShrimpBehaviourMaxDelay + config.FPS
	shrimp.Type = t
	shrimp.MoneyDelay = config.ShrimpMoneyDelay
	shrimp.ShrimpWallCollide()
	shrimp.Hunger = config.ShrimpMaxHunger
	shrimp.IsAlive = true
	return shrimp
}

func (s *Shrimp) ShrimpWallCollide() {

	minX := float32(config.PlayFieldX + config.BorderOffset)
	maxX := float32(config.PlayFieldX + config.PlayFieldWidth - config.StandardSquareSpriteSide - config.BorderOffset)
	minY := float32(config.PlayFieldY + config.BorderOffset)
	maxY := float32(config.PlayFieldY + config.PlayerFieldHeight - config.StandardSquareSpriteSide - config.BorderOffset)

	s.Position.X, s.Vx = utils.ClampAndBounce(s.Position.X, minX, maxX, s.Vx)
	s.Position.Y, s.Vy = utils.ClampAndBounce(s.Position.Y, minY, maxY, s.Vy)
}

func (s *Shrimp) Move() {
	s.BehaviourDelay--
	if s.BehaviourDelay == 0 {

		s.Vx = rand.Float32()*2*config.ShrimpMaxVelocity - config.ShrimpMaxVelocity
		s.Vy = rand.Float32()*2*config.ShrimpMaxVelocity - config.ShrimpMaxVelocity

		s.BehaviourDelay = rand.Int31()%config.ShrimpBehaviourMaxDelay + config.FPS
	}
	s.Position.X += s.Vx
	s.Position.Y += s.Vy
	s.ShrimpWallCollide()

	s.Starve()
}

func (s *Shrimp) Starve() {
	s.Hunger--
	if s.Hunger == 0 {
		s.Die()
	}
}

func (s *Shrimp) Die() {
	s.IsAlive = false
}

func (s *Shrimp) PoopMoney() int {
	if !s.IsAlive {
		s.MoneyDelay--
		if s.MoneyDelay == 0 {
			s.MoneyDelay = config.ShrimpMoneyDelay
			return config.MoneyByShrimp[s.Type]
		}
	}
	return 0
}
