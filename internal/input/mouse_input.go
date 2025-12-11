package input

import rl "github.com/gen2brain/raylib-go/raylib"

func HandleMouseInput() (int32, int32) {
	return rl.GetMouseX(), rl.GetMouseY()
}
