package game

import (
	"ShrimpSanctuary/internal/config"
)

type Game struct {
	Shrimps    []Shrimp
	Foods      []Food
	IsRunning  bool
	IsFeeding  bool
	IsCleaning bool
}

func NewGame() Game {
	g := Game{}
	g.Shrimps = make([]Shrimp, 0)
	for i := 0; i < config.ShrimpStartCount; i++ {
		g.AddShrimpInstance(NewShrimp())
	}
	g.IsRunning = true
	g.IsFeeding = false
	g.IsCleaning = false

	return g
}

func (g *Game) Update() {
	var foodsToDelete []int
	for i := range g.Shrimps {
		g.Shrimps[i].Move()
		foodsToDelete = append(foodsToDelete, g.ShrimpFoodCollide(g.Shrimps[i])...)
	}

	for i := range g.Foods {
		if g.Foods[i].MoveAndDisappear() {
			foodsToDelete = append(foodsToDelete, i)
		}
	}
	if len(foodsToDelete) != 0 {
		g.DeleteFood(foodsToDelete)
	}
}

func (g *Game) ClickInPlayField(X, Y int32) {
	if g.IsFeeding {
		g.AddFood(X, Y)
	}
}

func (g *Game) AddFood(X, Y int32) {
	f := NewFood(X, Y)
	g.Foods = append(g.Foods, f)
}

func (g *Game) DeleteFood(foodsToDelete []int) {
	var newFoods []Food
	for i := range g.Foods {
		for j := range foodsToDelete {
			if i != j {
				newFoods = append(newFoods, g.Foods[i])
			}
		}
	}
	g.Foods = newFoods
}

func (g *Game) AddShrimpXY(X int32, Y int32) {
	s := Shrimp{X: X, Y: Y}
	g.Shrimps = append(g.Shrimps, s)
}

func (g *Game) AddShrimpInstance(shrimp Shrimp) {
	g.Shrimps = append(g.Shrimps, shrimp)
}
