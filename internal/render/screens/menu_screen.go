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
	MSBtnFontSize = 55 * config.ScreenCoeff

	MSPlayBtnName     = "PLAY"
	MSSettingsBtnName = "SETTINGS"
	MSExitBtnName     = "EXIT"

	MSButtonsX      = 490 * config.ScreenCoeff
	MSButtonsWidth  = 300 * config.ScreenCoeff
	MSButtonsHeight = 60 * config.ScreenCoeff
	MSPlayBtnY      = 280 * config.ScreenCoeff
	MSSettingsBtnY  = 370 * config.ScreenCoeff
	MSExitBtnY      = 460 * config.ScreenCoeff

	MSMuteX       = 15 * config.ScreenCoeff
	MSMuteY       = 15 * config.ScreenCoeff
	MSMuteBtnName = "MUTE"
)

type MenuScreen struct {
	Game    *game.Game
	Buttons []*input.Button
	sb      *sound_bar.SoundBar
	ts      *assets.AssetStorage
}

func NewMenuScreen(game *game.Game, sb *sound_bar.SoundBar, as *assets.AssetStorage) *MenuScreen {
	ms := new(MenuScreen)
	ms.Game = game
	ms.sb = sb
	ms.ts = as

	ms.Buttons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				MSButtonsX,
				MSPlayBtnY,
				MSButtonsWidth,
				MSButtonsHeight,
			),
			MSPlayBtnName,
			ms.HandlePlayBtnClick,
			MSBtnFontSize,
			as,
		),
		input.NewButton(
			rl.NewRectangle(
				MSButtonsX,
				MSSettingsBtnY,
				MSButtonsWidth,
				MSButtonsHeight,
			),
			MSSettingsBtnName,
			ms.HandleSettingsBtnClick,
			MSBtnFontSize,
			as,
		),
		input.NewButton(
			rl.NewRectangle(
				MSButtonsX,
				MSExitBtnY,
				MSButtonsWidth,
				MSButtonsHeight,
			),
			MSExitBtnName,
			ms.HandleExitBtnClick,
			MSBtnFontSize,
			as,
		),
		input.NewButton(
			rl.NewRectangle(
				MSMuteX,
				MSMuteY,
				config.BigSquareSpriteSide,
				config.BigSquareSpriteSide,
			),
			MSMuteBtnName,
			ms.HandleMuteBtnClick,
			MSBtnFontSize,
			as,
		),
	}

	return ms
}

func (ms *MenuScreen) Draw() {
	rl.DrawTexture(ms.ts.MenuScreen, 0, 0, rl.White)
	ms.drawButtons()
}

func (ms *MenuScreen) drawButtons() {
	for _, btn := range ms.Buttons {
		textVector := rl.MeasureTextEx(btn.Font, btn.Text, btn.FontSize, 2)
		if btn.Text == MSMuteBtnName {
			if ms.sb.IsMuted() {
				rl.DrawTexture(ms.ts.Mute, 0, 0, rl.White)
			} else {
				rl.DrawTexture(ms.ts.Unmute, 0, 0, rl.White)
			}
			continue
		}
		rl.DrawTextEx(
			btn.Font,
			btn.Text,
			rl.Vector2{
				X: btn.Bounds.X + (btn.Bounds.Width-textVector.X)/2,
				Y: btn.Bounds.Y + 6},
			btn.FontSize,
			2,
			btn.Color)
	}
}

func (ms *MenuScreen) HandleInput() {
	for _, btn := range ms.Buttons {
		btnStatus := input.MouseButtonCollide(btn)
		btn.Color = config.ButtonColorFromStatus[btnStatus]
		if btnStatus == config.ClickedBtnStatus {
			btn.Action()
		}
	}
}

func (ms *MenuScreen) HandlePlayBtnClick() {
	ms.Game.State = config.StateAquarium
	ms.sb.PlayAquariumSound()
}

func (ms *MenuScreen) HandleSettingsBtnClick() {
	ms.Game.State = config.StateSettings
}

func (ms *MenuScreen) HandleExitBtnClick() {
	ms.Game.State = config.StateQuit
}

func (ms *MenuScreen) HandleMuteBtnClick() {
	ms.Game.Statistics.MuteBtnClicked++
	ms.sb.Mute()
}
