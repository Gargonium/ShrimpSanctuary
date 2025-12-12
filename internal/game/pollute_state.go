package game

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Pollute struct {
	Position   rl.Vector2
	Durability int32
}

func NewPollute() Pollute {
	p := Pollute{}
	p.Position.X = (rand.Float32() * config.PlayFieldWidth) + config.PlayFieldX
	p.Position.Y = (rand.Float32() * config.PlayerFieldHeight) + config.PlayFieldY
	p.Durability = config.PolluteMaxDurability - 1
	return p
}
