package main

import (
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/render"
	"ShrimpSanctuary/internal/sound_bar"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Shrimp Sanctuary")
	rl.InitAudioDevice()
	rl.SetTargetFPS(config.FPS)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	sb := sound_bar.NewSoundBar()
	g := game.NewGame()
	r := render.NewRender(g, sb)

	sb.PlayBgMusic()

	for !rl.WindowShouldClose() && g.State != config.StateQuit {
		r.Update()

		rl.BeginDrawing()
		r.Draw()
		rl.EndDrawing()
	}

	sb.StopBgMusic()

	rl.CloseWindow()
	rl.CloseAudioDevice()
}
