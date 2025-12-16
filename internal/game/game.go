package game

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game/entities"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Game struct {
	State        config.GameState
	Shrimps      []*entities.Shrimp
	Foods        []*entities.Food
	Pollution    []*entities.Pollute
	PolluteDelay int32
	Money        int
	IsFeeding    bool
	IsCleaning   bool
}

func NewGame() *Game {
	g := new(Game)
	g.Shrimps = make([]*entities.Shrimp, 0)
	g.Foods = make([]*entities.Food, 0)
	g.Pollution = make([]*entities.Pollute, 0)
	g.Money = config.StartMoney

	for i := 0; i < config.ShrimpStartCount; i++ {
		g.AddShrimpInstance(entities.NewShrimp(config.CherryShrimp))
	}

	g.PolluteDelay = 0 //config.PolluteSpawnDelay + rand.Int31n(config.PolluteSpawnDelaySpread * 2) - config.PolluteSpawnDelaySpread
	g.IsFeeding = false
	g.IsCleaning = false
	g.State = config.StateMenu

	return g
}

func (g *Game) Update() {

	if g.State == config.StateAquarium {
		var foodsToDelete []int
		for _, s := range g.Shrimps {
			s.Move()
			foodsToDelete = append(foodsToDelete, g.ShrimpFoodCollide(s)...)
			g.Money += s.PoopMoney()
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

}

func (g *Game) AddPollute() {
	p := entities.NewPollute()
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
	f := entities.NewFood(pos)
	g.Foods = append(g.Foods, f)
}

func (g *Game) DeleteFood(foodsToDelete []int) {
	var newFoods []*entities.Food
	for i := range g.Foods {
		for j := range foodsToDelete {
			if i != j {
				newFoods = append(newFoods, g.Foods[i])
			}
		}
	}
	g.Foods = newFoods
}

func (g *Game) AddShrimpInstance(shrimp *entities.Shrimp) {
	g.Shrimps = append(g.Shrimps, shrimp)
}

func (g *Game) DeletePollute(toDel int) {
	var newPollution []*entities.Pollute
	for i := range g.Pollution {
		if i != toDel {
			newPollution = append(newPollution, g.Pollution[i])
		}
	}
	g.Pollution = newPollution
}

func (g *Game) ShrimpFoodCollide(s *entities.Shrimp) []int {
	var foodCollide []int
	for i := range g.Foods {
		f := g.Foods[i]
		if utils.CollideCircleRect(f.Position, config.FoodRadius, s.Position.X, s.Position.Y, config.StandardSquareSpriteSide, config.StandardSquareSpriteSide) {
			foodCollide = append(foodCollide, i)
		}
	}
	return foodCollide
}
