package input

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/pkg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func MouseButtonCollide(btn Button) string {
	mousePos := rl.GetMousePosition()

	if utils.InBounds(mousePos, btn.Bounds) {
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

func GetMouseVector() rl.Vector2 {
	return rl.GetMousePosition()
}
