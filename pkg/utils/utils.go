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

func CollideCircleRect(X1, Y1, radius int32, X2, Y2, Width2, Height2 int32) bool {
	circleCenter := rl.NewVector2(float32(X1), float32(Y1))
	rec := rl.NewRectangle(float32(X2), float32(Y2), float32(Width2), float32(Height2))
	return rl.CheckCollisionCircleRec(circleCenter, float32(radius), rec)
}
