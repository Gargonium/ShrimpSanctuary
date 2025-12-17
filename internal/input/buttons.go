package input

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Bounds   rl.Rectangle
	Text     string
	Action   func()
	Font     rl.Font
	FontSize float32
	Color    rl.Color
	Status   config.ButtonStatus
}

func NewButton(bounds rl.Rectangle, text string, action func(), fontSize float32, as *assets.AssetStorage) *Button {
	b := new(Button)
	b.Bounds = bounds
	b.Text = text
	b.Action = action
	b.Font = as.WinterFont
	b.FontSize = fontSize
	b.Color = config.ButtonColorFromStatus["waiting"]
	b.Status = config.WaitingBtnStatus
	return b
}
