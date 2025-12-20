package screens

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SeSBtnFontSize = 55 * config.ScreenCoeff

	SeSBackBtnName  = "BACK"
	SeSApplyBtnName = "APPLY"

	SeSButtonsY      = 610 * config.ScreenCoeff
	SeSButtonsWidth  = 300 * config.ScreenCoeff
	SeSButtonsHeight = 60 * config.ScreenCoeff
	SeSBackBtnX      = 300 * config.ScreenCoeff
	SeSApplyBtnX     = 680 * config.ScreenCoeff

	SeSSliderMinX       = 700 * config.ScreenCoeff
	SeSSliderMaxX       = 1000 * config.ScreenCoeff
	SeSSliderSquareSide = 25 * config.ScreenCoeff
	SeSMusicSliderY     = 385 * config.ScreenCoeff
	SeSEffectsSliderY   = 505 * config.ScreenCoeff

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
	ts             *assets.AssetStorage
	MusicSliderX   float32
	EffectsSliderX float32
}

func NewSettingsScreen(game *game.Game, sb *sound_bar.SoundBar, as *assets.AssetStorage) *SettingsScreen {
	ss := new(SettingsScreen)
	ss.Game = game
	ss.SoundBar = sb
	ss.ts = as
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
			as,
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
			as,
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
