package screens

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	_ "ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	ASBtnFontSize   = 55 * config.ScreenCoeff
	ASMoneyFontSize = 30 * config.ScreenCoeff

	ASButtonsY      = 520 * config.ScreenCoeff
	ASButtonsWidth  = 140 * config.ScreenCoeff
	ASButtonsHeight = 134 * config.ScreenCoeff
	ASFeedBtnX      = 270 * config.ScreenCoeff
	ASCleanBtnX     = 470 * config.ScreenCoeff
	ASShopBtnX      = 670 * config.ScreenCoeff
	ASMenuBtnX      = 870 * config.ScreenCoeff

	ASFeedBtnName  = "FEED"
	ASCleanBtnName = "CLEAN"
	ASShopBtnName  = "SHOP"
	ASMenuBtnName  = "MENU"

	ASAchievBtnX    = 1168 * config.ScreenCoeff
	ASAchievBtnY    = 20 * config.ScreenCoeff
	ASAchievBtnSide = 86 * config.ScreenCoeff
)

type AquariumScreen struct {
	Game    *game.Game
	Buttons []*input.Button
	as      *assets.AssetStorage
	sb      *sound_bar.SoundBar
}

func NewAquariumScreen(game *game.Game, sb *sound_bar.SoundBar, assetStorage *assets.AssetStorage) *AquariumScreen {
	as := new(AquariumScreen)

	as.Game = game
	as.sb = sb
	as.as = assetStorage

	as.Buttons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				ASFeedBtnX,
				ASButtonsY,
				ASButtonsWidth,
				ASButtonsHeight,
			),
			ASFeedBtnName,
			as.HandleFeedBtnClick,
			ASBtnFontSize,
			assetStorage,
		),
		input.NewButton(
			rl.NewRectangle(
				ASCleanBtnX,
				ASButtonsY,
				ASButtonsWidth,
				ASButtonsHeight,
			),
			ASCleanBtnName,
			as.HandleCleanBtnClick,
			ASBtnFontSize,
			assetStorage,
		),
		input.NewButton(
			rl.NewRectangle(
				ASShopBtnX,
				ASButtonsY,
				ASButtonsWidth,
				ASButtonsHeight,
			),
			ASShopBtnName,
			as.HandleShopBtnClick,
			ASBtnFontSize,
			assetStorage,
		),
		input.NewButton(
			rl.NewRectangle(
				ASMenuBtnX,
				ASButtonsY,
				ASButtonsWidth,
				ASButtonsHeight,
			),
			ASMenuBtnName,
			as.HandleMenuBtnClick,
			ASBtnFontSize,
			assetStorage,
		),
		input.NewButton(
			rl.NewRectangle(
				ASAchievBtnX,
				ASAchievBtnY,
				ASAchievBtnSide,
				ASAchievBtnSide,
			),
			"",
			as.HandleAchievementsBtnClick,
			ASBtnFontSize,
			assetStorage,
		),
	}

	return as
}

func (as *AquariumScreen) Draw() {
	rl.DrawTexture(as.as.AquariumScreen, 0, 0, rl.White)
	as.drawWallpaper()
	as.drawButtons()
	as.drawShrimps()
	as.drawFood()
	as.drawPollute()
	as.drawMoney()
}

func (as *AquariumScreen) drawWallpaper() {
	switch as.Game.WallpaperState {
	case config.DefaultWallpaperState:
	case config.PvZWallpaperState:
		rl.DrawTextureV(as.as.PvZWallpaper, rl.NewVector2(config.WaterX, config.WaterY), rl.White)
	case config.CityWallpaperState:
		rl.DrawTextureV(as.as.CityWallpaper, rl.NewVector2(config.WaterX, config.WaterY), rl.White)
	case config.NiceWallpaperState:
		rl.DrawTextureV(as.as.NiceWallpaper, rl.NewVector2(config.WaterX, config.WaterY), rl.White)
	case config.GundamWallpaperState:
		rl.DrawTextureV(as.as.GundamWallpaper, rl.NewVector2(config.WaterX, config.WaterY), rl.White)
	}
	rl.DrawRectangleV(rl.NewVector2(config.WaterX, config.WaterY), rl.NewVector2(config.WaterWidth, config.WaterHeight), config.WaterColor)
}

func (as *AquariumScreen) drawButtons() {
	for _, btn := range as.Buttons {
		textVector := rl.MeasureTextEx(btn.Font, btn.Text, btn.FontSize, 2)
		rl.DrawTextEx(
			btn.Font,
			btn.Text,
			rl.Vector2{
				X: btn.Bounds.X + (btn.Bounds.Width-textVector.X)/2,
				Y: btn.Bounds.Y + textVector.Y/2},
			btn.FontSize,
			2,
			btn.Color)
	}
}

func (as *AquariumScreen) drawMoney() {
	rl.DrawTextureV(as.as.Coin, rl.NewVector2(config.MoneyX, config.MoneyY), rl.White)
	rl.DrawTextEx(as.as.MolotFont,
		strconv.Itoa(as.Game.Money),
		rl.NewVector2(
			config.MoneyX+config.StandardSquareSpriteSide+config.BorderOffset,
			config.MoneyY+(config.StandardSquareSpriteSide-ASMoneyFontSize)/2),
		ASMoneyFontSize,
		2,
		rl.White,
	)
}

func (as *AquariumScreen) drawShrimps() {
	for _, s := range as.Game.Shrimps {
		if !s.IsAlive {
			continue
		}

		var ShrimpTexture, ReversedShrimpTexture rl.Texture2D
		switch s.Type {
		case config.CherryShrimp:
			ShrimpTexture = as.as.CherryShrimp
			ReversedShrimpTexture = as.as.CherryShrimpReversed
		case config.GundamShrimp:
			ShrimpTexture = as.as.GundamShrimp
			ReversedShrimpTexture = as.as.GundamShrimpReversed
		case config.IsaacShrimp:
			ShrimpTexture = as.as.IsaacShrimp
			ReversedShrimpTexture = as.as.IsaacShrimpReversed
		case config.MinecraftShrimp:
			ShrimpTexture = as.as.MinecraftShrimp
			ReversedShrimpTexture = as.as.MinecraftShrimpReversed
		case config.MiskaShrimp:
			ShrimpTexture = as.as.MiskaShrimp
			ReversedShrimpTexture = as.as.MiskaShrimpReversed
		case config.ChanelShrimp:
			ShrimpTexture = as.as.ChanelShrimp
			ReversedShrimpTexture = as.as.ChanelShrimpReversed
		case config.BlackRoseShrimp:
			ShrimpTexture = as.as.BlackRoseShrimp
			ReversedShrimpTexture = as.as.BlackRoseShrimpReversed
		case config.SonicShrimp:
			ShrimpTexture = as.as.SonicShrimp
			ReversedShrimpTexture = as.as.SonicShrimpReversed
		}

		shrimpConditionTint := rl.NewColor(
			255,
			uint8(float32(s.Hunger)/float32(config.ShrimpMaxHunger)*255),
			uint8(float32(s.Hunger)/float32(config.ShrimpMaxHunger)*255),
			255,
		)
		if s.Vx >= 0 {
			rl.DrawTextureV(ShrimpTexture, s.Position, shrimpConditionTint)
		} else {
			rl.DrawTextureV(ReversedShrimpTexture, s.Position, shrimpConditionTint)
		}
	}
}

func (as *AquariumScreen) drawPollute() {
	for _, p := range as.Game.Pollution {

		polluteDurabilityTint := rl.NewColor(
			255,
			255,
			255,
			uint8(float32(p.Durability)/float32(config.PolluteMaxDurability)*255),
		)

		rl.DrawTextureV(as.as.Pollute, p.Position, polluteDurabilityTint)
	}
}

func (as *AquariumScreen) drawFood() {
	for _, f := range as.Game.Foods {
		if !f.IsAlive {
			continue
		}

		foodColor := config.FoodColor
		foodBorderColor := config.FoodBorderColor

		foodCondition := float32(f.GetLifeTime()) / float32(config.FoodLifeTime)

		if foodCondition < 0.5 {
			foodCondition *= 2
			foodColor.A = uint8(float32(foodColor.A) * foodCondition)
			foodBorderColor.A = uint8(float32(foodBorderColor.A) * foodCondition)
		}

		rl.DrawCircleV(f.Position, config.FoodRadius, foodColor)
		rl.DrawCircleLinesV(f.Position, config.FoodRadius, foodBorderColor)
	}
}

func (as *AquariumScreen) HandleInput() {
	if input.MousePlayFieldClick() {
		as.Game.ClickInPlayField(input.GetMouseVector())
		if as.Game.IsFeeding {
			as.sb.PlayFoodDropSound()
		}
		if as.Game.IsCleaning {
			as.sb.PlayCleanSound()
		}
	}

	for _, btn := range as.Buttons {
		btnStatus := input.MouseButtonCollide(btn)

		allowColorUpdate := true

		switch btn.Text {
		case ASFeedBtnName:
			if as.Game.IsFeeding && !as.Game.IsCleaning {
				allowColorUpdate = false
			}
		case ASCleanBtnName:
			if as.Game.IsCleaning && !as.Game.IsFeeding {
				allowColorUpdate = false
			}
		}

		if allowColorUpdate {
			btn.Color = config.ButtonColorFromStatus[btnStatus]
		}

		if btnStatus == config.ClickedBtnStatus {
			btn.Action()
		}
	}
}

func (as *AquariumScreen) HandleFeedBtnClick() {
	as.Game.IsFeeding = !as.Game.IsFeeding
	as.Game.IsCleaning = false
}

func (as *AquariumScreen) HandleCleanBtnClick() {
	as.Game.IsCleaning = !as.Game.IsCleaning
	as.Game.IsFeeding = false
}

func (as *AquariumScreen) HandleShopBtnClick() {
	as.Game.State = config.StateShop
	as.sb.StopAquariumSound()
}

func (as *AquariumScreen) HandleMenuBtnClick() {
	as.Game.State = config.StateMenu
	as.sb.StopAquariumSound()
}

func (as *AquariumScreen) HandleAchievementsBtnClick() {
	as.Game.State = config.StateAchievements
	as.sb.StopAquariumSound()
}
