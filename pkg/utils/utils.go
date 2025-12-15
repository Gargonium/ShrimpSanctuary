package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func ClampAndBounce(pos, min, max, vel float32) (float32, float32) {
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

func CollideCircleRect(circleCenter rl.Vector2, radius float32, X2, Y2, Width2, Height2 float32) bool {
	rec := rl.NewRectangle(X2, Y2, Width2, Height2)
	return rl.CheckCollisionCircleRec(circleCenter, radius, rec)
}

func LoadFont(fontPath string) rl.Font {
	fontTtf := rl.LoadFont(fontPath) //rl.LoadFontEx(config.WinterFont, 8, 0, 250)
	return fontTtf
}

func SpriteToTexture(spritePath string) rl.Texture2D {
	image := rl.LoadImage(spritePath)
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	return texture
}
