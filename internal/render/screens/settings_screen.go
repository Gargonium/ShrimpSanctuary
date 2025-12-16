package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SetBtnFontSize = 55

	SetBackBtnName  = "BACK"
	SetApplyBtnName = "APPLY"

	SetButtonsY      = 610
	SetButtonsWidth  = 300
	SetButtonsHeight = 60
	SetBackBtnX      = 300
	SetApplyBtnX     = 680

	SetSliderMinX       = 700
	SetSliderMaxX       = 1000
	SetSliderSquareSide = 25
	SetMusicSliderY     = 385
	SetEffectsSliderY   = 505
)

type SettingsScreen struct {
	Game           *game.Game
	SoundBar       *sound_bar.SoundBar
	Buttons        []*input.Button
	ts             *config.TextureStorage
	MusicSliderX   float32
	EffectsSliderX float32
}

func NewSettingsScreen(game *game.Game, sb *sound_bar.SoundBar, ts *config.TextureStorage) *SettingsScreen {
	ss := new(SettingsScreen)
	ss.Game = game
	ss.SoundBar = sb
	ss.ts = ts
	ss.MusicSliderX = SetSliderMinX + (SetSliderMaxX-SetSliderMinX)/2
	ss.EffectsSliderX = SetSliderMinX + (SetSliderMaxX-SetSliderMinX)/2

	ss.Buttons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				SetBackBtnX,
				SetButtonsY,
				SetButtonsWidth,
				SetButtonsHeight,
			),
			SetBackBtnName,
			ss.HandleBackBtnClick,
			SetBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				SetApplyBtnX,
				SetButtonsY,
				SetButtonsWidth,
				SetButtonsHeight,
			),
			SetApplyBtnName,
			ss.HandleApplyBtnClick,
			SetBtnFontSize,
		),
	}

	return ss
}

func (ss *SettingsScreen) HandleInput() {
	for _, btn := range ss.Buttons {
		btnStatus := input.MouseButtonCollide(btn)
		btn.Color = config.ButtonColorFromStatus[btnStatus]
		if btnStatus == config.ClickedBtnStatus {
			btn.Action()
		}
	}
	mousePosX := rl.GetMousePosition().X

	musicSliderStatus := input.MouseSliderCollide(rl.NewRectangle(ss.MusicSliderX, SetMusicSliderY, SetSliderSquareSide, SetSliderSquareSide))
	if musicSliderStatus == config.ClickedBtnStatus {
		ss.MusicSliderX = utils.Clamp(mousePosX-SetSliderSquareSide/2, SetSliderMinX, SetSliderMaxX)
	}

	effectsSliderStatus := input.MouseSliderCollide(rl.NewRectangle(ss.EffectsSliderX, SetEffectsSliderY, SetSliderSquareSide, SetSliderSquareSide))
	if effectsSliderStatus == config.ClickedBtnStatus {
		ss.EffectsSliderX = utils.Clamp(mousePosX-SetSliderSquareSide/2, SetSliderMinX, SetSliderMaxX)
	}
}

func (ss *SettingsScreen) Draw() {
	rl.DrawTexture(ss.ts.SettingsScreen, 0, 0, rl.White)
	ss.drawVolumeSlider()
	ss.drawButtons()
}

func (ss *SettingsScreen) drawVolumeSlider() {
	rl.DrawRectangleRounded(rl.NewRectangle(ss.MusicSliderX, SetMusicSliderY, SetSliderSquareSide, SetSliderSquareSide), 10, 1, config.VolumeSliderColor)
	rl.DrawRectangleRounded(rl.NewRectangle(ss.EffectsSliderX, SetEffectsSliderY, SetSliderSquareSide, SetSliderSquareSide), 10, 1, config.VolumeSliderColor)
}

func (ss *SettingsScreen) drawButtons() {
	for _, btn := range ss.Buttons {
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

func (ss *SettingsScreen) HandleBackBtnClick() {
	ss.Game.State = config.StateMenu
}

func (ss *SettingsScreen) HandleApplyBtnClick() {
	ss.ApplyVolume()
}

func (ss *SettingsScreen) ApplyVolume() {
	ss.SoundBar.ChangeMusicVolume((ss.MusicSliderX - SetSliderMinX) / (SetSliderMaxX - SetSliderMinX))
	ss.SoundBar.ChangeEffectsVolume((ss.EffectsSliderX - SetSliderMinX) / (SetSliderMaxX - SetSliderMinX))
}
