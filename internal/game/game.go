package game

import (
	"ShrimpSanctuary/internal/config"
)

type Game struct {
	Shrimps []*Shrimp
}

func NewGame() *Game {
	g := &Game{}
	g.Shrimps = make([]*Shrimp, 0)

	for i := 0; i < config.ShrimpStartCount; i++ {
		g.AddShrimpInstance(NewShrimp())
	}

	return g
}

func (g *Game) Update() {
	for i := range g.Shrimps {
		g.Shrimps[i].Move()
	}

}

func (g *Game) AddShrimpXY(X int32, Y int32) {
	shrimp := Shrimp{X: X, Y: Y}
	g.Shrimps = append(g.Shrimps, &shrimp)
}

func (g *Game) AddShrimpInstance(shrimp *Shrimp) {
	g.Shrimps = append(g.Shrimps, shrimp)
}
