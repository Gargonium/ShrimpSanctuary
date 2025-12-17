package screens

import (
	"ShrimpSanctuary/assets"
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

	SSShrimpsBuyBtnSide      = 80
	SSShrimpBuyBtnX          = 500
	SSShrimpBuyBtnY          = 40 + SSMenuY
	SSShrimpBuyBtnHorOffset  = 660
	SSShrimpBuyBtnVertOffset = 120
	SSShrimpsItemColumnCount = 4
	SSShrimpsItemRowCount    = 2

	SSWallpaperBuyBtnSide       = 100
	SSWallpaperBuyBtnX          = 510
	SSWallpaperBuyBtnY          = 100 + SSMenuY
	SSWallpaperBuyBtnHorOffset  = 620
	SSWallpaperBuyBtnVertOffset = 275
	SSWallpapersItemColumnCount = 2
	SSWallpapersItemRowCount    = 2

	SSBuyBtnName   = "BUY"
	SSApplyBtnName = "APPLY"
)

type ShopScreen struct {
	Game           *game.Game
	MenuButtons    []*input.Button
	ShrimpItems    []*ShrimpItem
	WallpaperItems []*WallpaperItem
	ts             *assets.AssetStorage
	State          config.ShopState
}

func NewShopScreen(game *game.Game, as *assets.AssetStorage) *ShopScreen {
	ss := new(ShopScreen)
	ss.Game = game
	ss.ts = as
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
			as,
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
			as,
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
			as,
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
			as,
		),
	}

	ss.ShrimpItems = make([]*ShrimpItem, 0)
	for i := 0; i < SSShrimpsItemColumnCount; i++ {
		for j := 0; j < SSShrimpsItemRowCount; j++ {
			btn := input.NewButton(
				rl.NewRectangle(
					float32(SSShrimpBuyBtnX+SSShrimpBuyBtnHorOffset*j),
					float32(SSShrimpBuyBtnY+SSShrimpBuyBtnVertOffset*i),
					SSShrimpsBuyBtnSide,
					SSShrimpsBuyBtnSide,
				),
				SSBuyBtnName,
				ss.HandleBuyBtnClick,
				SSBuyBtnFontSize,
				as,
			)
			ss.ShrimpItems = append(ss.ShrimpItems,
				NewShrimpItem(btn, SSShrimpsItemRowCount*i+j))
		}
	}

	ss.WallpaperItems = make([]*WallpaperItem, 0)
	for i := 0; i < SSWallpapersItemColumnCount; i++ {
		for j := 0; j < SSWallpapersItemRowCount; j++ {
			btn := input.NewButton(
				rl.NewRectangle(
					float32(SSWallpaperBuyBtnX+SSWallpaperBuyBtnHorOffset*j),
					float32(SSWallpaperBuyBtnY+SSWallpaperBuyBtnVertOffset*i),
					SSWallpaperBuyBtnSide,
					SSWallpaperBuyBtnSide,
				),
				SSBuyBtnName,
				ss.HandleBuyBtnClick,
				SSBuyBtnFontSize,
				as,
			)
			wi := NewWallpaperItem(btn, SSWallpapersItemRowCount*i+j)
			ss.WallpaperItems = append(ss.WallpaperItems, wi)
			wi.SetBoughtAndActive(ss.wallpaperUnlocked(wi.Type), ss.wallpaperActive(wi.Type))
			if wi.IsBought {
				btn.Text = SSApplyBtnName
			}
			if wi.IsActive {
				btn.Color = config.ButtonColorFromStatus[config.ClickedBtnStatus]
			}
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
		ss.handleInputWallpaperScreen()
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

func (ss *ShopScreen) handleInputWallpaperScreen() {
	for _, wi := range ss.WallpaperItems {
		btn := wi.BuyButton
		btn.Status = input.MouseButtonCollide(btn)
		if !wi.IsActive {
			btn.Color = config.ButtonColorFromStatus[btn.Status]
		}
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

func (ss *ShopScreen) findClickedWallpaperItem() *WallpaperItem {
	for _, wi := range ss.WallpaperItems {
		if wi.BuyButton.Status == config.ClickedBtnStatus {
			return wi
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
		ss.drawWallpaperScreen()
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
				X: btn.Bounds.X + (btn.Bounds.Width-textVector.X)/2,
				Y: btn.Bounds.Y - (btn.Bounds.Height/2-textVector.Y)/2 + config.Offset5},
			btn.FontSize,
			2,
			btn.Color)
		costTextWidth := float32(rl.MeasureText(strconv.Itoa(si.Cost), SSMoneyFontSize))
		rl.DrawTexture(
			ss.ts.Coin,
			int32(btn.Bounds.X+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2),
			int32(btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-config.StandardSquareSpriteSide)/2),
			rl.White,
		)
		rl.DrawText(
			strconv.Itoa(si.Cost),
			int32(btn.Bounds.X+config.StandardSquareSpriteSide+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2),
			int32(btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-SSMoneyFontSize)/2),
			SSMoneyFontSize,
			rl.Black,
		)
	}
}

func (ss *ShopScreen) drawWallpaperScreen() {
	rl.DrawTexture(ss.ts.ShopWallpaper, 0, SSMenuY, rl.White)
	for _, wi := range ss.WallpaperItems {
		btn := wi.BuyButton
		if btn.Text == SSBuyBtnName {
			btnTextVector := rl.MeasureTextEx(btn.Font, btn.Text, btn.FontSize, 2)
			rl.DrawTextEx(
				btn.Font,
				btn.Text,
				rl.Vector2{
					X: btn.Bounds.X + (btn.Bounds.Width-btnTextVector.X)/2,
					Y: btn.Bounds.Y + (btn.Bounds.Height/2-btnTextVector.Y)/2},
				btn.FontSize,
				2,
				btn.Color,
			)
			costTextWidth := float32(rl.MeasureText(strconv.Itoa(wi.Cost), SSMoneyFontSize))
			rl.DrawTexture(
				ss.ts.Coin,
				int32(btn.Bounds.X+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2),
				int32(btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-config.StandardSquareSpriteSide)/2),
				rl.White,
			)
			rl.DrawText(
				strconv.Itoa(wi.Cost),
				int32(btn.Bounds.X+config.StandardSquareSpriteSide+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2),
				int32(btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-SSMoneyFontSize)/2),
				SSMoneyFontSize,
				rl.Black,
			)
		} else {
			btnTextVector := rl.MeasureTextEx(btn.Font, btn.Text, btn.FontSize, 2)
			rl.DrawTextEx(
				btn.Font,
				btn.Text,
				rl.Vector2{
					X: btn.Bounds.X + (btn.Bounds.Width-btnTextVector.X)/2,
					Y: btn.Bounds.Y + (btn.Bounds.Height-btnTextVector.Y)/2},
				btn.FontSize,
				2,
				btn.Color,
			)
		}

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
	switch ss.State {
	case config.ShopStateShrimps:
		si := ss.findClickedShrimpItem()
		if si != nil {
			if ss.Game.Money >= si.Cost {
				ss.Game.AddShrimpInstance(entities.NewShrimp(si.ShrimpType))
				ss.Game.Money -= si.Cost
			}
		}
	case config.ShopStateWallpaper:
		wi := ss.findClickedWallpaperItem()
		if wi != nil {
			if !wi.IsBought {
				if ss.Game.Money >= wi.Cost {
					ss.Game.Money -= wi.Cost
					wi.IsBought = true
					wi.BuyButton.Text = SSApplyBtnName
					ss.Game.UnlockedWallpaper = append(ss.Game.UnlockedWallpaper, wi.Type)
				}
			} else {
				if wi.IsActive {
					ss.Game.WallpaperState = config.DefaultWallpaperState
				} else {
					ss.Game.WallpaperState = wi.Type
					for _, wi2 := range ss.WallpaperItems {
						if wi2 != wi {
							wi2.IsActive = false
						}
					}
				}
				wi.IsActive = !wi.IsActive
			}
		}
	case config.ShopStateDecor:
	}
}

func (ss *ShopScreen) wallpaperUnlocked(wt config.WallpaperState) bool {
	for _, w := range ss.Game.UnlockedWallpaper {
		if w == wt {
			return true
		}
	}
	return false
}

func (ss *ShopScreen) wallpaperActive(wt config.WallpaperState) bool {
	if ss.Game.WallpaperState == wt {
		return true
	}
	return false
}
