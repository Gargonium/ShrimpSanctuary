package screens

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	StatSBtnFontSize   = 55 * config.ScreenCoeff
	StatSBackBtnX      = 1115 * config.ScreenCoeff
	StatSBackBtnY      = 15 * config.ScreenCoeff
	StatSBackBtnWidth  = 150 * config.ScreenCoeff
	StatSBackBtnHeight = 60 * config.ScreenCoeff
	StatSBackBtnName   = "BACK"

	StatSStatsFontSize = 55 * config.ScreenCoeff
	StatSStatsNumberX  = 730 * config.ScreenCoeff
	StatSStatsTextX    = 48 * config.ScreenCoeff
	StatSStatsTextY    = 154 * config.ScreenCoeff
)

type StatsScreen struct {
	Game       *game.Game
	BackButton *input.Button
	sb         *sound_bar.SoundBar
	as         *assets.AssetStorage
}

func NewStatsScreen(game *game.Game, sb *sound_bar.SoundBar, as *assets.AssetStorage) *StatsScreen {
	ss := new(StatsScreen)
	ss.Game = game
	ss.sb = sb
	ss.as = as

	ss.BackButton = input.NewButton(
		rl.NewRectangle(
			StatSBackBtnX,
			StatSBackBtnY,
			StatSBackBtnWidth,
			StatSBackBtnHeight),
		StatSBackBtnName,
		ss.HandleBackBtnClick,
		StatSBtnFontSize,
		as,
	)

	return ss
}

func (ss *StatsScreen) HandleInput() {
	btn := ss.BackButton
	btn.Status = input.MouseButtonCollide(btn)
	btn.Color = config.ButtonColorFromStatus[btn.Status]
	if btn.Status == config.ClickedBtnStatus {
		btn.Action()
	}
}

func (ss *StatsScreen) Draw() {
	rl.DrawTexture(ss.as.StatsScreen, 0, 0, rl.White)
	ss.drawButtons()
	ss.drawStats()
}

func (ss *StatsScreen) drawStats() {
	rl.DrawTextEx(
		ss.as.MolotFont,
		"Money Earned \n"+
			"Money Spent \n"+
			"Number of Shrimps \n"+
			"Shrimps Died \n"+
			"Shrimps are Fed \n"+
			"Aquarium Cleaned (times) \n"+
			"Wallpapers Bought \n"+
			"Mute Button Clicked (times) ",
		rl.NewVector2(
			StatSStatsTextX,
			StatSStatsTextY),
		StatSStatsFontSize,
		2,
		rl.White,
	)
	rl.DrawTextEx(
		ss.as.MolotFont,
		":\t\t"+strconv.Itoa(ss.Game.Statistics.MoneyEarned)+"\n"+
			":\t\t"+strconv.Itoa(ss.Game.Statistics.MoneySpent)+"\n"+
			":\t\t"+strconv.Itoa(len(ss.Game.Shrimps))+"\n"+
			":\t\t"+strconv.Itoa(ss.Game.Statistics.ShrimpDied)+"\n"+
			":\t\t"+strconv.Itoa(ss.Game.Statistics.ShrimpsFed)+"\n"+
			":\t\t"+strconv.Itoa(ss.Game.Statistics.AquariumCleaned)+"\n"+
			":\t\t"+strconv.Itoa(ss.Game.Statistics.WallpapersCount)+"\n"+
			":\t\t"+strconv.Itoa(ss.Game.Statistics.MuteBtnClicked),
		rl.NewVector2(
			StatSStatsNumberX,
			StatSStatsTextY),
		StatSStatsFontSize,
		2,
		rl.White,
	)
}

func (ss *StatsScreen) drawButtons() {
	btn := ss.BackButton
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

func (ss *StatsScreen) HandleBackBtnClick() {
	ss.Game.State = config.StateAquarium
	ss.sb.PlayAquariumSound()
}
