package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MSBtnFontSize = 55

	MSPlayBtnName     = "PLAY"
	MSSettingsBtnName = "SETTINGS"
	MSExitBtnName     = "EXIT"

	MSButtonsX      = 490
	MSButtonsWidth  = 300
	MSButtonsHeight = 60
	MSPlayBtnY      = 250
	MSSettingsBtnY  = 340
	MSExitBtnY      = 430
)

type MenuScreen struct {
	Game    *game.Game
	Buttons []*input.Button
	ts      *config.TextureStorage
}

func NewMenuScreen(game *game.Game, ts *config.TextureStorage) *MenuScreen {
	ms := new(MenuScreen)
	ms.Game = game
	ms.ts = ts

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
}

func (ms *MenuScreen) HandleSettingsBtnClick() {
	ms.Game.State = config.StateSettings
}

func (ms *MenuScreen) HandleExitBtnClick() {
	ms.Game.State = config.StateQuit
}
