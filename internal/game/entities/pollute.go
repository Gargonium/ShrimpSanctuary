package entities

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Pollute struct {
	Position   rl.Vector2
	Durability int32
}

func NewPollute() *Pollute {
	p := new(Pollute)
	p.Position.X = utils.Clamp(
		(rand.Float32()*config.PlayFieldWidth)+config.PlayFieldX,
		config.PlayFieldX,
		config.PlayFieldX+config.PlayFieldWidth-config.BigSquareSpriteSide)
	p.Position.Y = utils.Clamp(
		(rand.Float32()*config.PlayerFieldHeight)+config.PlayFieldY,
		config.PlayFieldY,
		config.PlayFieldY+config.PlayerFieldHeight-config.BigSquareSpriteSide)
	p.Durability = config.PolluteMaxDurability
	return p
}
