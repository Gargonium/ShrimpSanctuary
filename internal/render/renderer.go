package render

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/input"
	"ShrimpSanctuary/internal/sound_bar"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Render struct {
	game                  game.Game
	sb                    sound_bar.SoundBar
	bgTexture             rl.Texture2D
	shrimpTexture         rl.Texture2D
	shrimpTextureReversed rl.Texture2D
	btnFont               rl.Font
	buttonsColor          map[string]rl.Color
}

func NewRender(g game.Game, sb sound_bar.SoundBar) *Render {
	r := Render{}
	r.bgTexture = spriteToTexture(config.BgSprite)
	r.shrimpTexture = spriteToTexture(config.CherryShrimpSprite)
	r.shrimpTextureReversed = spriteToTexture(config.CherryShrimpReversedSprite)
	r.btnFont = loadFont(config.WinterFont)
	r.buttonsColor = make(map[string]rl.Color)
	r.buttonsColor[config.FeedBtnName] = rl.Black
	r.buttonsColor[config.CleanBtnName] = rl.Black
	r.buttonsColor[config.ShopBtnName] = rl.Black
	r.buttonsColor[config.ExitBtnName] = rl.Black

	r.game = g
	r.sb = sb

	return &r
}

func loadFont(fontPath string) rl.Font {
	fontTtf := rl.LoadFont(fontPath) //rl.LoadFontEx(config.WinterFont, 8, 0, 250)
	return fontTtf
}

func spriteToTexture(spritePath string) rl.Texture2D {
	image := rl.LoadImage(spritePath)
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	return texture
}

func (r *Render) HandleInput() {

	if input.MousePlayFieldClick() {
		r.game.ClickInPlayField(input.GetMouseXY())
	}

	feedBtnStatus := input.MouseButtonCollide(config.FeedBtnName)
	cleanBtnStatus := input.MouseButtonCollide(config.CleanBtnName)
	shopBtnStatus := input.MouseButtonCollide(config.ShopBtnName)
	exitBtnStatus := input.MouseButtonCollide(config.ExitBtnName)

	if !r.game.IsFeeding || r.game.IsCleaning {
		r.buttonsColor[config.FeedBtnName] = config.ButtonColorFromStatus[feedBtnStatus]
	}
	if !r.game.IsCleaning || r.game.IsFeeding {
		r.buttonsColor[config.CleanBtnName] = config.ButtonColorFromStatus[cleanBtnStatus]
	}
	r.buttonsColor[config.ShopBtnName] = config.ButtonColorFromStatus[shopBtnStatus]
	r.buttonsColor[config.ExitBtnName] = config.ButtonColorFromStatus[exitBtnStatus]

	if feedBtnStatus == config.ClickedBtnStatus {
		r.HandleFeedBtnClick()
	}
	if cleanBtnStatus == config.ClickedBtnStatus {
		r.HandleCleanBtnClick()
	}
	if shopBtnStatus == config.ClickedBtnStatus {
		r.HandleShopBtnClick()
	}
	if exitBtnStatus == config.ClickedBtnStatus {
		r.HandleExitBtnClick()
	}
}

func (r *Render) Draw() {
	r.DrawBackground()
	r.DrawShrimps()
	r.DrawFood()
}

func (r *Render) Update() {
	r.HandleInput()
	r.game.Update()
	r.sb.Update()
}

func (r *Render) HandleFeedBtnClick() {
	r.game.IsFeeding = !r.game.IsFeeding
	r.game.IsCleaning = false
}

func (r *Render) HandleCleanBtnClick() {
	r.game.IsCleaning = !r.game.IsCleaning
	r.game.IsFeeding = false
}

func (r *Render) HandleShopBtnClick() {

}

func (r *Render) HandleExitBtnClick() {
	r.game.IsRunning = false
}

func (r *Render) GameIsRunning() bool {
	return r.game.IsRunning
}
