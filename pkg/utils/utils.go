package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ClampAndBounce(pos, min, max, vel float32) (float32, float32) {
	if pos < min {
		return min, -vel
	}
	if pos > max {
		return max, -vel
	}
	return pos, vel
}

func Clamp(pos, min, max float32) float32 {
	if pos < min {
		return min
	}
	if pos > max {
		return max
	}
	return pos
}

func InBounds(target rl.Vector2, bounds rl.Rectangle) bool {
	if target.X > bounds.X && target.X < bounds.X+bounds.Width &&
		target.Y > bounds.Y && target.Y < bounds.Y+bounds.Height {
		return true
	}
	return false
}

func CollideCircleRect(circleCenter rl.Vector2, radius float32, X, Y, Width, Height float32) bool {
	rec := rl.NewRectangle(X, Y, Width, Height)
	return rl.CheckCollisionCircleRec(circleCenter, radius, rec)
}
