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
	var mouseX, mouseY int32
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mouseX, mouseY = rl.GetMouseX(), rl.GetMouseY()
		if utils.InBounds(rl.Vector2{X: float32(mouseX), Y: float32(mouseY)}, config.PlayFieldBounds) {
			return true
		}
	}
	return false
}

func GetMouseXY() (int32, int32) {
	return rl.GetMouseX(), rl.GetMouseY()
}
