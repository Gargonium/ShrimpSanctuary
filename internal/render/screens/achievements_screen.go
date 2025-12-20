package screens

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ACSBtnFontSize   = 55 * config.ScreenCoeff
	ACSBackBtnX      = 1115 * config.ScreenCoeff
	ACSBackBtnY      = 15 * config.ScreenCoeff
	ACSBackBtnWidth  = 150 * config.ScreenCoeff
	ACSBackBtnHeight = 60 * config.ScreenCoeff
	ACSBackBtnName   = "BACK"

	ACSTrophyX           = 43 * config.ScreenCoeff
	ACSTrophyY           = 183 * config.ScreenCoeff
	ACSTrophyHorOffset   = 640 * config.ScreenCoeff
	ACSTrophyVertOffset  = 185 * config.ScreenCoeff
	ACSTrophyColumnCount = 2
	ACSTrophyRowCount    = 3
)

type Trophy struct {
}

type AchievementsScreen struct {
	Game       *game.Game
	BackButton *input.Button
	sb         *sound_bar.SoundBar
	as         *assets.AssetStorage
}

func NewAchievementsScreen(g *game.Game, sb *sound_bar.SoundBar, assetStorage *assets.AssetStorage) *AchievementsScreen {
	as := new(AchievementsScreen)
	as.Game = g
	as.sb = sb
	as.as = assetStorage

	as.BackButton = input.NewButton(
		rl.NewRectangle(
			ACSBackBtnX,
			ACSBackBtnY,
			ACSBackBtnWidth,
			ACSBackBtnHeight),
		ACSBackBtnName,
		as.HandleBackBtnClick,
		ACSBtnFontSize,
		assetStorage,
	)

	return as
}

func (as *AchievementsScreen) HandleInput() {
	btn := as.BackButton
	btn.Status = input.MouseButtonCollide(btn)
	btn.Color = config.ButtonColorFromStatus[btn.Status]
	if btn.Status == config.ClickedBtnStatus {
		btn.Action()
	}
}

func (as *AchievementsScreen) Draw() {
	rl.DrawTexture(as.as.AchievementsScreen, 0, 0, rl.White)
	as.drawButtons()
	as.drawTrophies()
}

func (as *AchievementsScreen) drawButtons() {
	btn := as.BackButton
	textVector := rl.MeasureTextEx(btn.Font, btn.Text, btn.FontSize, 2)
	rl.DrawTextEx(
		btn.Font,
		btn.Text,
		rl.Vector2{
			X: btn.Bounds.X + (btn.Bounds.Width-textVector.X)/2,
			Y: btn.Bounds.Y + (btn.Bounds.Height-textVector.Y)/2},
		btn.FontSize,
		2,
		btn.Color)
}

func (as *AchievementsScreen) drawTrophies() {
	for i := 0; i < ACSTrophyRowCount; i++ {
		for j := 0; j < ACSTrophyColumnCount; j++ {
			trophyTexture := as.as.GrayTrophy
			if as.Game.Statistics.Achievements[ACSTrophyColumnCount*i+j] {
				trophyTexture = as.as.Trophy
			}
			rl.DrawTextureV(trophyTexture,
				rl.NewVector2(
					ACSTrophyX+ACSTrophyHorOffset*float32(j),
					ACSTrophyY+ACSTrophyVertOffset*float32(i)),
				rl.White,
			)
		}
	}
}

func (as *AchievementsScreen) HandleBackBtnClick() {
	as.Game.State = config.StateAquarium
	as.sb.PlayAquariumSound()
}
