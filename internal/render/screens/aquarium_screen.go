package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	_ "ShrimpSanctuary/internal/input"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	ASBtnFontSize   = 55
	ASMoneyFontSize = 30

	ASButtonsY      = 520
	ASButtonsWidth  = 140
	ASButtonsHeight = 134
	ASFeedBtnX      = 270
	ASCleanBtnX     = 470
	ASShopBtnX      = 670
	ASMenuBtnX      = 870

	ASFeedBtnName  = "FEED"
	ASCleanBtnName = "CLEAN"
	ASShopBtnName  = "SHOP"
	ASMenuBtnName  = "MENU"
)

type AquariumScreen struct {
	Game    *game.Game
	Buttons []*input.Button
	ts      *config.TextureStorage
}

func NewAquariumScreen(game *game.Game, ts *config.TextureStorage) *AquariumScreen {
	as := new(AquariumScreen)

	as.Game = game
	as.ts = ts

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
		),
	}

	return as
}

func (as *AquariumScreen) Draw() {
	rl.DrawTexture(as.ts.AquariumScreen, 0, 0, rl.White)
	as.drawButtons()
	as.drawMoney()
	as.drawShrimps()
	as.drawFood()
	as.drawPollute()
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
	rl.DrawTexture(as.ts.Coin, config.MoneyX, config.MoneyY, rl.White)
	rl.DrawText(strconv.Itoa(as.Game.Money),
		config.MoneyX+config.StandardSquareSpriteSide+config.BorderOffset,
		config.MoneyY+(config.StandardSquareSpriteSide-ASMoneyFontSize)/2,
		ASMoneyFontSize, rl.White)
}

func (as *AquariumScreen) drawShrimps() {
	for i := range as.Game.Shrimps {
		if as.Game.Shrimps[i].Vx < 0 {
			rl.DrawTextureV(as.ts.CherryShrimpReversed, as.Game.Shrimps[i].Position, rl.White)
		} else {
			rl.DrawTextureV(as.ts.CherryShrimp, as.Game.Shrimps[i].Position, rl.White)
		}
	}
}

func (as *AquariumScreen) drawPollute() {
	for i := range as.Game.Pollution {
		polCol := config.PolluteColor
		polCol.A = uint8(float32(polCol.A) * float32(as.Game.Pollution[i].Durability) / float32(config.PolluteMaxDurability))
		rl.DrawCircleV(as.Game.Pollution[i].Position, config.PolluteRadius, polCol)
	}
}

func (as *AquariumScreen) drawFood() {
	for i := range as.Game.Foods {
		rl.DrawCircleV(as.Game.Foods[i].Position, config.FoodRadius, config.FoodColor)
		rl.DrawCircleLinesV(as.Game.Foods[i].Position, config.FoodRadius, config.FoodBorderColor)
	}
}

func (as *AquariumScreen) HandleInput() {
	if input.MousePlayFieldClick() {
		as.Game.ClickInPlayField(input.GetMouseVector())
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
}
func (as *AquariumScreen) HandleMenuBtnClick() {
	as.Game.State = config.StateMenu
}
