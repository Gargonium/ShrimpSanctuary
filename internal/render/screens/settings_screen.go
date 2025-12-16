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
	SeSBtnFontSize = 55

	SeSBackBtnName  = "BACK"
	SeSApplyBtnName = "APPLY"

	SeSButtonsY      = 610
	SeSButtonsWidth  = 300
	SeSButtonsHeight = 60
	SeSBackBtnX      = 300
	SeSApplyBtnX     = 680

	SeSSliderMinX       = 700
	SeSSliderMaxX       = 1000
	SeSSliderSquareSide = 25
	SeSMusicSliderY     = 385
	SeSEffectsSliderY   = 505

	SeSSliderStickX        = SeSSliderMinX
	SeSMusicSliderStickY   = SeSMusicSliderY
	SeSEffectsSliderStickY = SeSEffectsSliderY
	SeSSliderStickWidth    = SeSSliderMaxX - SeSSliderMinX + SeSSliderSquareSide
	SeSSliderStickHeight   = SeSSliderSquareSide
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
	ss.MusicSliderX = SeSSliderMinX + (SeSSliderMaxX-SeSSliderMinX)*sb.GetMusicVolume()
	ss.EffectsSliderX = SeSSliderMinX + (SeSSliderMaxX-SeSSliderMinX)*sb.GetEffectsVolume()

	ss.Buttons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				SeSBackBtnX,
				SeSButtonsY,
				SeSButtonsWidth,
				SeSButtonsHeight,
			),
			SeSBackBtnName,
			ss.HandleBackBtnClick,
			SeSBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				SeSApplyBtnX,
				SeSButtonsY,
				SeSButtonsWidth,
				SeSButtonsHeight,
			),
			SeSApplyBtnName,
			ss.HandleApplyBtnClick,
			SeSBtnFontSize,
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

	musicSliderStatus := input.MouseSliderCollide(rl.NewRectangle(SeSSliderStickX, SeSMusicSliderStickY, SeSSliderStickWidth, SeSSliderStickHeight))
	if musicSliderStatus == config.ClickedBtnStatus {
		ss.MusicSliderX = utils.Clamp(mousePosX-SeSSliderSquareSide/2, SeSSliderMinX, SeSSliderMaxX)
	}

	effectsSliderStatus := input.MouseSliderCollide(rl.NewRectangle(SeSSliderStickX, SeSEffectsSliderStickY, SeSSliderStickWidth, SeSSliderStickHeight))
	if effectsSliderStatus == config.ClickedBtnStatus {
		ss.EffectsSliderX = utils.Clamp(mousePosX-SeSSliderSquareSide/2, SeSSliderMinX, SeSSliderMaxX)
	}
}

func (ss *SettingsScreen) Draw() {
	rl.DrawTexture(ss.ts.SettingsScreen, 0, 0, rl.White)
	ss.drawVolumeSlider()
	ss.drawButtons()
}

func (ss *SettingsScreen) drawVolumeSlider() {
	rl.DrawRectangleRounded(rl.NewRectangle(ss.MusicSliderX, SeSMusicSliderY, SeSSliderSquareSide, SeSSliderSquareSide), 10, 1, config.VolumeSliderColor)
	rl.DrawRectangleRounded(rl.NewRectangle(ss.EffectsSliderX, SeSEffectsSliderY, SeSSliderSquareSide, SeSSliderSquareSide), 10, 1, config.VolumeSliderColor)
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
	ss.SoundBar.ChangeMusicVolume((ss.MusicSliderX - SeSSliderMinX) / (SeSSliderMaxX - SeSSliderMinX))
	ss.SoundBar.ChangeEffectsVolume((ss.EffectsSliderX - SeSSliderMinX) / (SeSSliderMaxX - SeSSliderMinX))
}
