package input

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func MouseButtonCollide(btn string) string {
	mousePos := rl.GetMousePosition()
	btnBounds := rl.NewRectangle(config.ButtonsX[btn], config.ButtonY, config.ButtonWidth, config.ButtonHeight)

	if utils.InBounds(mousePos, btnBounds) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			return config.ClickedBtnStatus
		}
		return config.HoveredBtnStatus
	}
	return config.WaitingBtnStatus
}

func MousePlayFieldClick() bool {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mousePos := rl.GetMousePosition()
		if utils.InBounds(mousePos, config.PlayFieldBounds) {
			return true
		}
	}
	return false
}

func GetMouseXY() (int32, int32) {
	return rl.GetMouseX(), rl.GetMouseY()
}

func GetMouseVector() rl.Vector2 {
	return rl.GetMousePosition()
}
