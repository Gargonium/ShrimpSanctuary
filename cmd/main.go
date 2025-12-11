package main

import (
	"ShrimpSanctuary/internal/colors"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/render"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Shrimp Sanctuary")
	rl.SetTargetFPS(config.FPS)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	g := game.NewGame()
	r := render.NewRender()

	for !rl.WindowShouldClose() {
		g.Update()

		rl.BeginDrawing()
		rl.ClearBackground(colors.BackgroundColor)

		r.Draw(g)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

/*
Типы данных:
	rl.Vector2
	rl.Vector3
	rl.Rectangle

Функции:
	rl.InitWindow
	rl.SetTargetFPS
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.WindowShouldClose
	rl.BeginDrawing
	rl.ClearBackground(rl.NewColor(20, 25, 35, 255))
	rl.EndDrawing()
	rl.CloseWindow()
	rl.NewVector2(0, 0)
	rl.NewRectangle
	rl.White / Yellow, Green etc.
	rl.GetFrameTime()
	rl.PushMatrix()
	rl.Translatef
	rl.Scalef
	rl.PopMatrix()
	rl.DrawRectangle
	rl.DrawRectangleLinesEx
	rl.DrawRectangleLines
	rl.DrawCircle
	rl.DrawCircleLines
	rl.DrawLine
	rl.DrawLineEx
	rl.DrawCircleGradient
	rl.GetMousePosition()
	rl.GetMouseWheelMove()
	rl.IsMouseButtonDown(rl.MouseButtonRight)
	rl.IsKeyPressed(rl.KeyR)
	rl.DrawText
	rl.Rotatef
	rl.DrawTriangleFan
	rl.DrawEllipse
	rl.Vector2Subtract
	rl.Vector2Add
	rl.Vector2Distance
	rl.DrawRectangleRounded(b.Bounds, 0.3, 8, color)
	rl.DrawRectangleRoundedLines
	rl.MeasureText
	rl.CheckCollisionPointRec
*/
