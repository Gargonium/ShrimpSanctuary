package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func ClampAndBounce(pos, min, max, vel int32) (int32, int32) {
	if pos < min {
		return min, -vel
	}
	if pos > max {
		return max, -vel
	}
	return pos, vel
}

func InBounds(target rl.Vector2, bounds rl.Rectangle) bool {
	if target.X > bounds.X && target.X < bounds.X+bounds.Width &&
		target.Y > bounds.Y && target.Y < bounds.Y+bounds.Height {
		return true
	}
	return false
}
