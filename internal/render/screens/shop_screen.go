package screens

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/game/entities"
	"ShrimpSanctuary/internal/input"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	SSBtnFontSize    = 55
	SSMoneyFontSize  = 30
	SSBuyBtnFontSize = 40

	SSMenuButtonsY  = 111
	SSButtonsWidth  = 427
	SSButtonsHeight = 50
	SSShrimpsBtnX   = 0
	SSWallpaperBtnX = SSShrimpsBtnX + SSButtonsWidth
	SSDecorBtnX     = SSWallpaperBtnX + SSButtonsWidth

	SSBackBtnX      = 1115
	SSBackBtnY      = 15
	SSBackBtnWidth  = 150
	SSBackBtnHeight = 60

	SSMenuY = 160

	SSShrimpsBtnName   = "SHRIMPS"
	SSWallpaperBtnName = "WALLPAPER"
	SSDecorBtnName     = "DECOR"
	SSBackBtnName      = "BACK"

	SSBuyBtnX          = 500
	SSBuyBtnY          = 200
	SSBuyBtnSide       = 80
	SSBuyBtnHorOffset  = 660
	SSBuyBtnVertOffset = 120

	SSBuyBtnName = "BUY"

	SSShrimpsItemColumnCount = 4
	SSShrimpsItemRowCount    = 2
)

type ShrimpItem struct {
	BuyButton  *input.Button
	Cost       int
	ShrimpType config.ShrimpType
}

func NewShrimpItem(btn *input.Button, cost int, tp config.ShrimpType) *ShrimpItem {
	si := new(ShrimpItem)
	si.BuyButton = btn
	si.Cost = cost
	si.ShrimpType = tp
	return si
}

type ShopScreen struct {
	Game        *game.Game
	MenuButtons []*input.Button
	ShrimpItems []*ShrimpItem
	ts          *config.TextureStorage
	State       config.ShopState
}

func NewShopScreen(game *game.Game, ts *config.TextureStorage) *ShopScreen {
	ss := new(ShopScreen)
	ss.Game = game
	ss.ts = ts
	ss.State = config.ShopStateShrimps

	ss.MenuButtons = []*input.Button{
		input.NewButton(
			rl.NewRectangle(
				SSShrimpsBtnX,
				SSMenuButtonsY,
				SSButtonsWidth,
				SSButtonsHeight,
			),
			SSShrimpsBtnName,
			ss.HandleShrimpsBtnClick,
			SSBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				SSWallpaperBtnX,
				SSMenuButtonsY,
				SSButtonsWidth,
				SSButtonsHeight,
			),
			SSWallpaperBtnName,
			ss.HandleWallpaperBtnClick,
			SSBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				SSDecorBtnX,
				SSMenuButtonsY,
				SSButtonsWidth,
				SSButtonsHeight,
			),
			SSDecorBtnName,
			ss.HandleDecorBtnClick,
			SSBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				SSBackBtnX,
				SSBackBtnY,
				SSBackBtnWidth,
				SSBackBtnHeight,
			),
			SSBackBtnName,
			ss.HandleBackBtnClick,
			SSBtnFontSize,
		),
		input.NewButton(
			rl.NewRectangle(
				610,
				200,
				74,
				74,
			),
			"S",
			ss.HandleSBtnClick,
			SSBtnFontSize,
		),
	}

	ss.ShrimpItems = make([]*ShrimpItem, 0)
	for i := 0; i < SSShrimpsItemColumnCount; i++ {
		for j := 0; j < SSShrimpsItemRowCount; j++ {
			btn := input.NewButton(
				rl.NewRectangle(
					float32(SSBuyBtnX+SSBuyBtnHorOffset*j),
					float32(SSBuyBtnY+SSBuyBtnVertOffset*i),
					SSBuyBtnSide,
					SSBuyBtnSide,
				),
				SSBuyBtnName,
				ss.HandleBuyBtnClick,
				SSBuyBtnFontSize,
			)
			ss.ShrimpItems = append(ss.ShrimpItems,
				NewShrimpItem(btn, config.ShrimpCost[config.CherryShrimp], config.CherryShrimp))
		}
	}

	return ss
}

func (ss *ShopScreen) HandleInput() {
	for _, btn := range ss.MenuButtons {
		btn.Status = input.MouseButtonCollide(btn)
		btn.Color = config.ButtonColorFromStatus[btn.Status]
		if btn.Status == config.ClickedBtnStatus {
			btn.Action()
		}
	}

	switch ss.State {
	case config.ShopStateShrimps:
		ss.handleInputShrimpsScreen()
	case config.ShopStateWallpaper:

	case config.ShopStateDecor:

	}
}

func (ss *ShopScreen) handleInputShrimpsScreen() {
	for _, si := range ss.ShrimpItems {
		btn := si.BuyButton
		btn.Status = input.MouseButtonCollide(btn)
		btn.Color = config.ButtonColorFromStatus[btn.Status]
		if btn.Status == config.ClickedBtnStatus {
			btn.Action()
		}
	}
}

func (ss *ShopScreen) findClickedShrimpItem() *ShrimpItem {
	for _, si := range ss.ShrimpItems {
		if si.BuyButton.Status == config.ClickedBtnStatus {
			return si
		}
	}
	return nil
}

func (ss *ShopScreen) Draw() {
	rl.DrawTexture(ss.ts.ShopScreen, 0, 0, rl.White)
	ss.drawMoney()
	ss.drawButtons()

	switch ss.State {
	case config.ShopStateShrimps:
		ss.drawShrimpsScreen()
	case config.ShopStateWallpaper:
		rl.DrawTexture(ss.ts.ShopWallpaper, 0, SSMenuY, rl.White)
	case config.ShopStateDecor:
		rl.DrawTexture(ss.ts.ShopDecor, 0, SSMenuY, rl.White)
	}
}

func (ss *ShopScreen) drawShrimpsScreen() {
	rl.DrawTexture(ss.ts.ShopShrimps, 0, SSMenuY, rl.White)
	for _, si := range ss.ShrimpItems {
		btn := si.BuyButton
		textVector := rl.MeasureTextEx(btn.Font, btn.Text, btn.FontSize, 2)
		rl.DrawTextEx(
			btn.Font,
			btn.Text,
			rl.Vector2{
				X: btn.Bounds.X + (btn.Bounds.Width-textVector.X)/2 - config.BorderOffset,
				Y: btn.Bounds.Y - (btn.Bounds.Height/2-textVector.Y)/2 + config.BorderOffset},
			btn.FontSize,
			2,
			btn.Color)
	}
}

func (ss *ShopScreen) drawButtons() {
	for _, btn := range ss.MenuButtons {
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
}

func (ss *ShopScreen) drawMoney() {
	rl.DrawTexture(ss.ts.Coin, config.MoneyX, config.MoneyY, rl.White)
	rl.DrawText(strconv.Itoa(ss.Game.Money),
		config.MoneyX+config.StandardSquareSpriteSide+config.BorderOffset,
		config.MoneyY+(config.StandardSquareSpriteSide-SSMoneyFontSize)/2,
		SSMoneyFontSize, rl.White)
}

func (ss *ShopScreen) HandleShrimpsBtnClick() {
	ss.State = config.ShopStateShrimps
}

func (ss *ShopScreen) HandleWallpaperBtnClick() {
	ss.State = config.ShopStateWallpaper
}

func (ss *ShopScreen) HandleDecorBtnClick() {
	ss.State = config.ShopStateDecor
}

func (ss *ShopScreen) HandleBackBtnClick() {
	ss.Game.State = config.StateAquarium
}

func (ss *ShopScreen) HandleBuyBtnClick() {
	si := ss.findClickedShrimpItem()
	if si != nil {
		if ss.Game.Money >= si.Cost {
			ss.Game.AddShrimpInstance(entities.NewShrimp(si.ShrimpType))
			ss.Game.Money -= si.Cost
		}
	}
}

func (ss *ShopScreen) HandleSBtnClick() {
	for i := 0; i < 10000; i++ {
		ss.Game.AddShrimpInstance(entities.NewShrimp(config.CherryShrimp))
	}
}
