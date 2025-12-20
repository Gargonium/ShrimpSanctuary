package game

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game/entities"
	"ShrimpSanctuary/internal/sound_bar"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Statistics struct {
	ShrimpDied      int
	Achievements    []bool
	ShrimpsFed      int
	AquariumCleaned int
	ShrimpsCount    map[config.ShrimpType]int
	WallpapersCount int
	MoneyEarned     int
	MoneySpent      int
	MuteBtnClicked  int
}

type Game struct {
	State             config.GameState
	WallpaperState    config.WallpaperState
	UnlockedWallpaper []config.WallpaperState
	Shrimps           []*entities.Shrimp
	Foods             []*entities.Food
	Pollution         []*entities.Pollute
	PolluteDelay      int32
	Money             int
	IsFeeding         bool
	IsCleaning        bool
	SoundBar          *sound_bar.SoundBar
	Statistics        *Statistics
}

func NewGame(sb *sound_bar.SoundBar) *Game {
	g := new(Game)
	g.Shrimps = make([]*entities.Shrimp, 0)
	g.Foods = make([]*entities.Food, 0)
	g.Pollution = make([]*entities.Pollute, 0)
	g.Money = config.StartMoney
	g.SoundBar = sb

	g.Statistics = new(Statistics)
	g.Statistics.Achievements = make([]bool, 0)
	for i := 0; i < config.AchievementsCount; i++ {
		g.Statistics.Achievements = append(g.Statistics.Achievements, false)
	}
	g.Statistics.ShrimpsFed = 0
	g.Statistics.AquariumCleaned = 0
	g.Statistics.WallpapersCount = 0
	g.Statistics.ShrimpsCount = make(map[config.ShrimpType]int)
	for _, st := range config.ShrimpsTypesInShop {
		g.Statistics.ShrimpsCount[st] = 0
	}
	g.Statistics.MoneyEarned = config.StartMoney
	g.Statistics.ShrimpDied = 0
	g.Statistics.MoneySpent = 0
	g.Statistics.MuteBtnClicked = 0

	for i := 0; i < config.ShrimpStartCount; i++ {
		g.AddShrimpInstance(entities.NewShrimp(config.CherryShrimp))
	}

	g.PolluteDelay = 0 // config.PolluteSpawnDelay + rand.Int31n(config.PolluteSpawnDelaySpread*2) - config.PolluteSpawnDelaySpread
	g.IsFeeding = false
	g.IsCleaning = false
	g.State = config.StateMenu
	g.WallpaperState = config.DefaultWallpaperState
	g.UnlockedWallpaper = make([]config.WallpaperState, 0)

	return g
}

func (g *Game) Update() {

	if g.State == config.StateAquarium {
		for _, s := range g.Shrimps {
			s.Move()
			g.ShrimpFoodCollide(s)
			m := s.PoopMoney()
			g.Money += m
			g.Statistics.MoneyEarned += m
		}

		if g.Statistics.MoneyEarned >= config.MillionaireGoal {
			g.Statistics.Achievements[config.Millionaire] = true
		}

		for i := 0; i < len(g.Foods); i++ {
			g.Foods[i].MoveAndDisappear()
		}

		if g.PolluteDelay == 0 {
			g.AddPollute()
			g.PolluteDelay = config.PolluteSpawnDelay + rand.Int31n(config.PolluteSpawnDelaySpread*2) - config.PolluteSpawnDelaySpread
		}
		g.PolluteDelay--

		g.deleteDeadFood()
		g.deleteDeadShrimps()
	}
}

func (g *Game) deleteDeadFood() {
	var newFoods []*entities.Food
	for _, f := range g.Foods {
		if f.IsAlive {
			newFoods = append(newFoods, f)
		}
	}
	g.Foods = newFoods
}

func (g *Game) deleteDeadShrimps() {
	var newShrimps []*entities.Shrimp
	for _, s := range g.Shrimps {
		if s.IsAlive {
			newShrimps = append(newShrimps, s)
		} else {
			g.Statistics.Achievements[config.StrengthTest] = true
			g.Statistics.ShrimpDied++
		}
	}
	g.Shrimps = newShrimps
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
		for i, p := range g.Pollution {
			if rl.CheckCollisionPointRec(pos, rl.NewRectangle(p.Position.X, p.Position.Y, config.BigSquareSpriteSide, config.BigSquareSpriteSide)) {
				p.Durability--
				if p.Durability == 0 {
					g.DeletePollute(i)
					g.Statistics.AquariumCleaned++
					if g.Statistics.AquariumCleaned == config.MrPropperGoal {
						g.Statistics.Achievements[config.MrPropper] = true
					}
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
	g.Statistics.ShrimpsCount[shrimp.Type]++
	allTypes := true
	for _, s := range g.Shrimps {
		if g.Statistics.ShrimpsCount[s.Type] < config.LegendOfDepthsGoal {
			allTypes = false
			break
		}
	}
	if allTypes {
		g.Statistics.Achievements[config.LegendOfDepths] = true
	}
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

func (g *Game) ShrimpFoodCollide(s *entities.Shrimp) {
	for _, f := range g.Foods {
		if utils.CollideCircleRect(f.Position, config.FoodRadius, s.Position.X, s.Position.Y, config.StandardSquareSpriteSide, config.StandardSquareSpriteSide) {
			s.Hunger = config.ShrimpMaxHunger
			f.SelfDestruct()
			g.Statistics.ShrimpsFed++
			if g.Statistics.ShrimpsFed == config.GluttonyGoal {
				g.Statistics.Achievements[config.Gluttony] = true
			}
		}
	}
}
