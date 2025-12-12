package game

import (
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Game struct {
	Shrimps      []Shrimp
	Foods        []Food
	Pollution    []Pollute
	PolluteDelay int32
	IsRunning    bool
	IsFeeding    bool
	IsCleaning   bool
}

func NewGame() Game {
	g := Game{}
	g.Shrimps = make([]Shrimp, 0)
	g.Foods = make([]Food, 0)
	g.Pollution = make([]Pollute, 0)

	for i := 0; i < config.ShrimpStartCount; i++ {
		g.AddShrimpInstance(NewShrimp())
	}

	g.PolluteDelay = 0 //config.PolluteSpawnDelay + rand.Int31n(config.PolluteSpawnDelaySpread * 2) - config.PolluteSpawnDelaySpread
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

	if g.PolluteDelay == 0 {
		g.AddPollute()
		g.PolluteDelay = config.PolluteSpawnDelay + rand.Int31n(config.PolluteSpawnDelaySpread*2) - config.PolluteSpawnDelaySpread
	}
	g.PolluteDelay--

}

func (g *Game) AddPollute() {
	p := NewPollute()
	g.Pollution = append(g.Pollution, p)
}

func (g *Game) ClickInPlayField(pos rl.Vector2) {
	if g.IsFeeding {
		g.AddFood(pos)
	}
	if g.IsCleaning {
		for i := range g.Pollution {
			if rl.CheckCollisionPointCircle(pos, g.Pollution[i].Position, config.PolluteRadius) {
				g.Pollution[i].Durability--
				if g.Pollution[i].Durability == 0 {
					g.DeletePollute(i)
				}
				break
			}
		}
	}
}

func (g *Game) AddFood(pos rl.Vector2) {
	f := NewFood(pos)
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

func (g *Game) AddShrimpXY(X, Y float32) {
	s := Shrimp{}
	s.Position.X = X
	s.Position.Y = Y
	g.Shrimps = append(g.Shrimps, s)
}

func (g *Game) AddShrimpInstance(shrimp Shrimp) {
	g.Shrimps = append(g.Shrimps, shrimp)
}

func (g *Game) DeletePollute(toDel int) {
	var newPollution []Pollute
	for i := range g.Pollution {
		if i != toDel {
			newPollution = append(newPollution, g.Pollution[i])
		}
	}
	g.Pollution = newPollution
}
