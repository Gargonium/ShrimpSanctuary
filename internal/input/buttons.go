package input

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Bounds     rl.Rectangle
	Text       string
	Action     func()
	Font       rl.Font
	FontSize   float32
	Color      rl.Color
	HoverColor rl.Color
	ClickColor rl.Color
}

func NewButton(bounds rl.Rectangle, text string, action func(), fontSize float32) *Button {
	b := new(Button)
	b.Bounds = bounds
	b.Text = text
	b.Action = action
	b.Font = utils.LoadFont(config.WinterFont)
	b.FontSize = fontSize
	b.Color = config.ButtonColorFromStatus["waiting"]
	b.HoverColor = config.ButtonColorFromStatus["hovered"]
	b.ClickColor = config.ButtonColorFromStatus["clicked"]
	return b
}
