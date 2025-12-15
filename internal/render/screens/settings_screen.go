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
	SSBtnFontSize = 55

	SSBackBtnName  = "BACK"
	SSApplyBtnName = "APPLY"

	SSButtonsY      = 610
	SSButtonsWidth  = 300
	SSButtonsHeight = 60
	SSBackBtnX      = 300
	SSApplyBtnX     = 680

	SSSliderMinX       = 700
	SSSliderMaxX       = 1000
	SSSliderSquareSide = 25
	SSMusicSliderY     = 385
	SSEffectsSliderY   = 505
)

type SettingsScreen struct {
	Game           *game.Game
	SoundBar       *sound_bar.SoundBar
	Buttons        []*input.Button
	bgTexture      rl.Texture2D
	MusicSliderX   float32
	EffectsSliderX float32
}

func NewSettingsScreen(game *game.Game, sb *sound_bar.SoundBar) *SettingsScreen {
	ss := new(SettingsScreen)
	ss.bgTexture = utils.SpriteToTexture(config.SettingsBgSprite)
	ss.Game = game
	ss.SoundBar = sb
	ss.MusicSliderX = SSSliderMinX + (SSSliderMaxX-SSSliderMinX)/2
	ss.EffectsSliderX = SSSliderMinX + (SSSliderMaxX-SSSliderMinX)/2

	ss.Buttons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				SSBackBtnX,
				SSButtonsY,
				SSButtonsWidth,
				SSButtonsHeight,
			),
			SSBackBtnName,
			ss.HandleBackBtnClick,
			SSBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				SSApplyBtnX,
				SSButtonsY,
				SSButtonsWidth,
				SSButtonsHeight,
			),
			SSApplyBtnName,
			ss.HandleApplyBtnClick,
			SSBtnFontSize,
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

	musicSliderStatus := input.MouseSliderCollide(rl.NewRectangle(ss.MusicSliderX, SSMusicSliderY, SSSliderSquareSide, SSSliderSquareSide))
	if musicSliderStatus == config.ClickedBtnStatus {
		ss.MusicSliderX = utils.Clamp(mousePosX-SSSliderSquareSide/2, SSSliderMinX, SSSliderMaxX)
	}

	effectsSliderStatus := input.MouseSliderCollide(rl.NewRectangle(ss.EffectsSliderX, SSEffectsSliderY, SSSliderSquareSide, SSSliderSquareSide))
	if effectsSliderStatus == config.ClickedBtnStatus {
		ss.EffectsSliderX = utils.Clamp(mousePosX-SSSliderSquareSide/2, SSSliderMinX, SSSliderMaxX)
	}
}

func (ss *SettingsScreen) Draw() {
	rl.DrawTexture(ss.bgTexture, 0, 0, rl.White)
	ss.drawVolumeSlider()
	ss.drawButtons()
}

func (ss *SettingsScreen) drawVolumeSlider() {
	rl.DrawRectangleRounded(rl.NewRectangle(ss.MusicSliderX, SSMusicSliderY, SSSliderSquareSide, SSSliderSquareSide), 10, 1, config.VolumeSliderColor)
	rl.DrawRectangleRounded(rl.NewRectangle(ss.EffectsSliderX, SSEffectsSliderY, SSSliderSquareSide, SSSliderSquareSide), 10, 1, config.VolumeSliderColor)
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
	ss.SoundBar.ChangeMusicVolume((ss.MusicSliderX - SSSliderMinX) / (SSSliderMaxX - SSSliderMinX))
	ss.SoundBar.ChangeEffectsVolume((ss.EffectsSliderX - SSSliderMinX) / (SSSliderMaxX - SSSliderMinX))
}
