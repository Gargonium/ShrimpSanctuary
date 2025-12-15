package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	_ "ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AquariumScreen struct {
	Game                  *game.Game
	Buttons               []*input.Button
	bgTexture             rl.Texture2D
	shrimpTexture         rl.Texture2D
	shrimpTextureReversed rl.Texture2D
}

func NewAquariumScreen(game *game.Game) AquariumScreen {
	as := AquariumScreen{}

	as.Game = game

	as.Buttons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				config.ButtonsX[config.FeedBtnName],
				config.ButtonY,
				config.ButtonWidth,
				config.ButtonHeight,
			),
			config.FeedBtnName,
			as.HandleFeedBtnClick,
			config.BtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				config.ButtonsX[config.CleanBtnName],
				config.ButtonY,
				config.ButtonWidth,
				config.ButtonHeight,
			),
			config.CleanBtnName,
			as.HandleCleanBtnClick,
			config.BtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				config.ButtonsX[config.ShopBtnName],
				config.ButtonY,
				config.ButtonWidth,
				config.ButtonHeight,
			),
			config.ShopBtnName,
			as.HandleShopBtnClick,
			config.BtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				config.ButtonsX[config.ExitBtnName],
				config.ButtonY,
				config.ButtonWidth,
				config.ButtonHeight,
			),
			config.ExitBtnName,
			as.HandleExitBtnClick,
			config.BtnFontSize,
		),
	}

	as.bgTexture = utils.SpriteToTexture(config.AquariumBgSprite)

	as.shrimpTexture = utils.SpriteToTexture(config.CherryShrimpSprite)
	as.shrimpTextureReversed = utils.SpriteToTexture(config.CherryShrimpReversedSprite)

	return as
}

func (as *AquariumScreen) Draw() {
	rl.DrawTexture(as.bgTexture, 0, 0, rl.White)
	as.drawButtons()
	as.DrawShrimps()
	as.DrawFood()
	as.DrawPollute()
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

func (as *AquariumScreen) DrawShrimps() {
	for i := range as.Game.Shrimps {
		if as.Game.Shrimps[i].Vx < 0 {
			rl.DrawTextureV(as.shrimpTextureReversed, as.Game.Shrimps[i].Position, rl.White)
		} else {
			rl.DrawTextureV(as.shrimpTexture, as.Game.Shrimps[i].Position, rl.White)
		}
	}
}

func (as *AquariumScreen) DrawPollute() {
	for i := range as.Game.Pollution {
		polCol := config.PolluteColor
		polCol.A = uint8(float32(polCol.A) * float32(as.Game.Pollution[i].Durability) / float32(config.PolluteMaxDurability))
		rl.DrawCircleV(as.Game.Pollution[i].Position, config.PolluteRadius, polCol)
	}
}

func (as *AquariumScreen) DrawFood() {
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
		btnStatus := input.MouseButtonCollide(*btn)

		allowColorUpdate := true

		switch btn.Text {
		case config.FeedBtnName:
			if as.Game.IsFeeding && !as.Game.IsCleaning {
				allowColorUpdate = false
			}
		case config.CleanBtnName:
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

}
func (as *AquariumScreen) HandleExitBtnClick() {
	as.Game.State = config.StateQuit
}
