package screens

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/game/entities"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	SSBtnFontSize    = 55 * config.ScreenCoeff
	SSMoneyFontSize  = 30 * config.ScreenCoeff
	SSBuyBtnFontSize = 40 * config.ScreenCoeff

	SSMenuButtonsY  = 111 * config.ScreenCoeff
	SSButtonsWidth  = 640 * config.ScreenCoeff
	SSButtonsHeight = 50 * config.ScreenCoeff
	SSShrimpsBtnX   = 0
	SSWallpaperBtnX = SSShrimpsBtnX + SSButtonsWidth

	SSBackBtnX      = 1115 * config.ScreenCoeff
	SSBackBtnY      = 15 * config.ScreenCoeff
	SSBackBtnWidth  = 150 * config.ScreenCoeff
	SSBackBtnHeight = 60 * config.ScreenCoeff

	SSMenuY = 160 * config.ScreenCoeff

	SSShrimpsBtnName   = "SHRIMPS"
	SSWallpaperBtnName = "WALLPAPER"
	SSBackBtnName      = "BACK"

	SSShrimpsBuyBtnSide      = 80 * config.ScreenCoeff
	SSShrimpBuyBtnX          = 500 * config.ScreenCoeff
	SSShrimpBuyBtnY          = 40*config.ScreenCoeff + SSMenuY
	SSShrimpBuyBtnHorOffset  = 660 * config.ScreenCoeff
	SSShrimpBuyBtnVertOffset = 120 * config.ScreenCoeff
	SSShrimpsItemColumnCount = 2
	SSShrimpsItemRowCount    = 4

	SSWallpaperBuyBtnSide       = 100 * config.ScreenCoeff
	SSWallpaperBuyBtnX          = 510 * config.ScreenCoeff
	SSWallpaperBuyBtnY          = 100*config.ScreenCoeff + SSMenuY
	SSWallpaperBuyBtnHorOffset  = 620 * config.ScreenCoeff
	SSWallpaperBuyBtnVertOffset = 275 * config.ScreenCoeff
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
	as             *assets.AssetStorage
	sb             *sound_bar.SoundBar
	State          config.ShopState
}

func NewShopScreen(game *game.Game, sb *sound_bar.SoundBar, as *assets.AssetStorage) *ShopScreen {
	ss := new(ShopScreen)
	ss.Game = game
	ss.sb = sb
	ss.as = as
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
	for i := 0; i < SSShrimpsItemRowCount; i++ {
		for j := 0; j < SSShrimpsItemColumnCount; j++ {
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
				NewShrimpItem(btn, SSShrimpsItemColumnCount*i+j))
		}
	}

	ss.WallpaperItems = make([]*WallpaperItem, 0)
	for i := 0; i < SSWallpapersItemRowCount; i++ {
		for j := 0; j < SSWallpapersItemColumnCount; j++ {
			btn := input.NewButton(
				rl.NewRectangle(
					SSWallpaperBuyBtnX+SSWallpaperBuyBtnHorOffset*float32(j),
					SSWallpaperBuyBtnY+SSWallpaperBuyBtnVertOffset*float32(i),
					SSWallpaperBuyBtnSide,
					SSWallpaperBuyBtnSide,
				),
				SSBuyBtnName,
				ss.HandleBuyBtnClick,
				SSBuyBtnFontSize,
				as,
			)
			wi := NewWallpaperItem(btn, SSWallpapersItemColumnCount*i+j)
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
	rl.DrawTexture(ss.as.ShopScreen, 0, 0, rl.White)
	ss.drawMoney()
	ss.drawButtons()

	switch ss.State {
	case config.ShopStateShrimps:
		ss.drawShrimpsScreen()
	case config.ShopStateWallpaper:
		ss.drawWallpaperScreen()
	}
}

func (ss *ShopScreen) drawShrimpsScreen() {
	rl.DrawTexture(ss.as.ShopShrimps, 0, SSMenuY, rl.White)
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
			ss.as.Coin,
			int32(btn.Bounds.X+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2),
			int32(btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-config.StandardSquareSpriteSide)/2),
			rl.White,
		)
		rl.DrawTextEx(
			ss.as.MolotFont,
			strconv.Itoa(si.Cost),
			rl.NewVector2(
				btn.Bounds.X+config.StandardSquareSpriteSide+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2,
				btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-SSMoneyFontSize)/2),
			SSMoneyFontSize,
			2,
			rl.Black,
		)
	}
}

func (ss *ShopScreen) drawWallpaperScreen() {
	rl.DrawTexture(ss.as.ShopWallpaper, 0, SSMenuY, rl.White)
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
				ss.as.Coin,
				int32(btn.Bounds.X+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2),
				int32(btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-config.StandardSquareSpriteSide)/2),
				rl.White,
			)
			rl.DrawTextEx(
				ss.as.MolotFont,
				strconv.Itoa(wi.Cost),
				rl.NewVector2(
					btn.Bounds.X+config.StandardSquareSpriteSide+(btn.Bounds.Width-config.StandardSquareSpriteSide-costTextWidth)/2,
					btn.Bounds.Y+btn.Bounds.Height/2+(btn.Bounds.Height/2-SSMoneyFontSize)/2),
				SSMoneyFontSize,
				2,
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
	rl.DrawTextureV(ss.as.Coin, rl.NewVector2(config.MoneyX, config.MoneyY), rl.White)
	rl.DrawTextEx(ss.as.MolotFont,
		strconv.Itoa(ss.Game.Money),
		rl.NewVector2(
			config.MoneyX+config.StandardSquareSpriteSide+config.BorderOffset,
			config.MoneyY+(config.StandardSquareSpriteSide-ASMoneyFontSize)/2),
		ASMoneyFontSize,
		2,
		rl.White,
	)
}

func (ss *ShopScreen) HandleShrimpsBtnClick() {
	ss.State = config.ShopStateShrimps
}

func (ss *ShopScreen) HandleWallpaperBtnClick() {
	ss.State = config.ShopStateWallpaper
}

func (ss *ShopScreen) HandleBackBtnClick() {
	ss.Game.State = config.StateAquarium
	ss.sb.PlayAquariumSound()
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
					ss.Game.Statistics.WallpapersCount++
					if ss.Game.Statistics.WallpapersCount == config.CollectorGoal {
						ss.Game.Statistics.Achievements[config.Collector] = true
					}
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
