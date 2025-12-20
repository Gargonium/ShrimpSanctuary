package main

import (
	"ShrimpSanctuary/assets"
	"ShrimpSanctuary/internal/config"
	"ShrimpSanctuary/internal/game"
	"ShrimpSanctuary/internal/render"
	"ShrimpSanctuary/internal/sound_bar"
	"ShrimpSanctuary/pkg"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Shrimp Sanctuary")
	rl.InitAudioDevice()
	rl.SetTargetFPS(config.FPS)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	iconImage := rl.LoadImage("assets/assets/sprites/Shrimps/CherryShrimp.png")
	rl.SetWindowIcon(*iconImage)
	rl.UnloadImage(iconImage)

	assetMgr := assets.NewAssetManager()
	defer assetMgr.Cleanup()
	ts := assets.NewTextureStorage(assetMgr)

	sv := pkg.NewSaveManager()
	sb := sound_bar.NewSoundBar(ts)
	g := game.NewGame(sb)

	// TODO Убрать комментарий
	//if sv.SaveExists() {
	//	err := sv.LoadGame(g)
	//	if err != nil {
	//		fmt.Println("Load Error:", err)
	//	}
	//}

	r := render.NewRender(g, sb, ts)

	defer func() {
		fmt.Println("Saving game...")
		if err := sv.SaveGame(g); err != nil {
			fmt.Println("Save Error:", err)
		}
	}()

	sb.PlayBgMusic()

	var beforeAutoSave = config.AutoSaveDelay

	for !rl.WindowShouldClose() && g.State != config.StateQuit {
		r.Update()

		rl.BeginDrawing()
		r.Draw()
		rl.EndDrawing()

		if beforeAutoSave == 0 {
			err := sv.AutoSave(g)
			if err != nil {
				fmt.Println("Auto Save Error:", err)
			}
			beforeAutoSave = config.AutoSaveDelay
		}
		beforeAutoSave--
	}

	sb.StopBgMusic()
	sb.UnloadAll()

	rl.CloseWindow()
	rl.CloseAudioDevice()
}
