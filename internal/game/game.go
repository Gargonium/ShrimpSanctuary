package game

import (
	"ShrimpSanctuary/internal/config"
)

type Game struct {
	Shrimps   []Shrimp
	IsRunning bool
}

func NewGame() *Game {
	g := &Game{}
	g.Shrimps = make([]Shrimp, 0)
	for i := 0; i < config.ShrimpStartCount; i++ {
		g.AddShrimpInstance(NewShrimp())
	}
	g.IsRunning = true

	return g
}

func (g *Game) Update() {
	for i := range g.Shrimps {
		g.Shrimps[i].Move()
	}
}

func (g *Game) AddShrimpXY(X int32, Y int32) {
	s := Shrimp{X: X, Y: Y}
	g.Shrimps = append(g.Shrimps, s)
}

func (g *Game) AddShrimpInstance(shrimp Shrimp) {
	g.Shrimps = append(g.Shrimps, shrimp)
}
